package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"regexp"
	"strings"

	. "github.com/dave/jennifer/jen" //nolint: staticcheck
	"github.com/jmatth/loz"
)

func main() {
	appendPtr := flag.Bool("a", false, "append to file instead of overwriting")
	flag.Parse()
	inPath := flag.Arg(0)
	inType := flag.Arg(1)
	outType := flag.Arg(2)
	outPath := flag.Arg(3)

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, inPath+".go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	outFile := NewFile("loz")

	ast.Inspect(file, func(n ast.Node) bool {
		funcDecl, ok := n.(*ast.FuncDecl)
		if !ok || funcDecl.Recv == nil || len(funcDecl.Recv.List) != 1 {
			return true
		}
		recv := funcDecl.Recv.List[0]
		var recvTypedef typeDef
		switch t := recv.Type.(type) {
		case *ast.IndexExpr:
			if recvTypeName := t.X.(*ast.Ident).Name; recvTypeName != inType {
				return true
			}
			types := []string{
				t.Index.(*ast.Ident).Name,
			}
			recvTypedef = typeDef{
				name:       outType,
				typeParams: types,
			}
		case *ast.IndexListExpr:
			if recvTypeName := t.X.(*ast.Ident).Name; recvTypeName != inType {
				return true
			}
			types := loz.Map1[ast.Expr, string](loz.IterSlice(t.Indices)).
				Map(func(e ast.Expr) string {
					return e.(*ast.Ident).Name
				}).
				CollectSlice()
			recvTypedef = typeDef{
				name:       outType,
				typeParams: types,
			}
		default:
			return true
		}

		methodName := funcDecl.Name.Name

		var returnTypes []funcParam
		if results := funcDecl.Type.Results; results != nil {
			returnTypes = loz.Map1[*ast.Field, funcParam](loz.IterSlice(funcDecl.Type.Results.List)).
				Map(fieldToFuncParam).
				CollectSlice()
			if len(returnTypes) == 1 && returnTypes[0].typeDef.name == inType {
				returnTypes[0].typeDef.name = outType
			}
		}

		paramTypes := funcDecl.Type.Params.List

		pt := loz.Map1[*ast.Field, funcParam](loz.IterSlice(paramTypes)).
			Map(fieldToFuncParam).
			CollectSlice()

		genWrapper(outFile, methodName, recvTypedef, inType, pt, returnTypes)

		return true
	})

	var outFileDisk *os.File
	fileFlags := os.O_WRONLY
	if *appendPtr {
		fileFlags |= os.O_APPEND
	} else {
		fileFlags |= os.O_TRUNC | os.O_CREATE
	}
	outFileDisk, err = os.OpenFile(outPath, fileFlags, 0o644)
	if err != nil {
		panic(err)
	}
	var outWriter io.Writer = outFileDisk
	if *appendPtr {
		outWriter = packageSkipper{
			writer: outFileDisk,
		}
	}
	err = outFile.Render(outWriter)
	if err != nil {
		panic(err)
	}
	err = outFileDisk.Close()
	if err != nil {
		panic(err)
	}
}

func typeToTypedef(t ast.Expr) typeDef {
	td := typeDef{}
	switch t := t.(type) {
	case *ast.ArrayType:
		td.name = t.Elt.(*ast.Ident).Name
		td.isArr = true
	case *ast.Ident:
		td.name = t.Name
	case *ast.IndexExpr:
		td.name = t.X.(*ast.Ident).Name
		td.typeParams = []string{t.Index.(*ast.Ident).Name}
	case *ast.IndexListExpr:
		td.name = t.X.(*ast.Ident).Name
		td.typeParams = loz.Map1[ast.Expr, string](loz.IterSlice(t.Indices)).
			Map(func(e ast.Expr) string {
				switch t := e.(type) {
				case *ast.Ident:
					return t.Name
				case *ast.IndexExpr:
					// Dirty hack until the go compiler is fixed.
					outerType := t.X.(*ast.Ident).Name
					innerType := t.Index.(*ast.Ident).Name
					return fmt.Sprintf("%s[%s]", outerType, innerType)
				default:
					panic(t)
				}
			}).
			CollectSlice()
	case *ast.StarExpr:
		td = typeToTypedef(t.X)
		td.isStar = true
	default:
		td.name = "TODO"
	}
	return td
}

func fieldToFuncParam(f *ast.Field) funcParam {
	td := typeToTypedef(f.Type)
	name := ""
	if len(f.Names) > 0 {
		name = f.Names[0].Name
	}
	return funcParam{
		name:    name,
		typeDef: td,
	}
}

type typeDef struct {
	name       string
	typeParams []string
	isArr      bool
	isStar     bool
}

type funcParam struct {
	name    string
	typeDef typeDef
}

func (t typeDef) paramsAsCode() []Code {
	return loz.Map1[string, Code](loz.IterSlice(t.typeParams)).
		Map(func(s string) Code { return Id(s) }).
		CollectSlice()
}

func genWrapper(f *File, name string, recvType typeDef, originalType string, params []funcParam, returnTypes []funcParam) {
	recvId := Id(strings.ToLower(string(recvType.name[0])))
	generatedParams := loz.Map1[funcParam, Code](loz.IterSlice(params)).
		Map(paramToCode).
		CollectSlice()
	innerParams := loz.Map1[funcParam, Code](loz.IterSlice(params)).
		Map(func(p funcParam) Code { return Id(p.name) }).
		CollectSlice()
	generatedReturnTypes := loz.Map1[funcParam, Code](loz.IterSlice(returnTypes)).
		Map(func(p funcParam) Code {
			result := Id(p.typeDef.name).Types(p.typeDef.paramsAsCode()...)
			if p.typeDef.isArr {
				result = Index().Add(result)
			}
			if p.typeDef.isStar {
				result = Op("*").Add(result)
			}
			return result
		}).
		CollectSlice()

	f.Func().
		Params(recvId.Clone().Id(recvType.name).Types(recvType.paramsAsCode()...)).
		Id(name).
		Params(generatedParams...).
		Params(generatedReturnTypes...).
		BlockFunc(func(g *Group) {
			innerCall := Id(originalType).
				Types(recvType.paramsAsCode()...).
				Call(recvId).
				Dot(name).
				Call(innerParams...)
			if len(returnTypes) == 1 && returnTypes[0].typeDef.name == recvType.name {
				// The method allows chaining, cast back to our subtype
				innerCall = Id(recvType.name).Types(recvType.paramsAsCode()...).Call(innerCall)
			}
			if len(generatedReturnTypes) < 1 {
				// No return parameters so return without rendering the return keyword
				g.Custom(Options{}, innerCall)
				return
			}
			g.Return(innerCall)
		})
}

func paramToCode(param funcParam) Code {
	typeResult :=
		Id(param.typeDef.name).
			TypesFunc(func(g *Group) {
				for _, pt := range param.typeDef.typeParams {
					g.Id(pt)
				}
			})
	if param.typeDef.isArr {
		typeResult = Index().Add(typeResult)
	}
	if param.typeDef.isStar {
		typeResult = Op("*").Add(typeResult)
	}
	return Id(param.name).Add(typeResult)
}

type packageSkipper struct {
	skipped bool
	builder strings.Builder
	writer  io.Writer
}

var packageMatcher = regexp.MustCompile(`^package \w+`)

func (s packageSkipper) Write(b []byte) (int, error) {
	if s.skipped {
		return s.writer.Write(b)
	}
	written, err := s.builder.Write(b)
	if err != nil {
		return written, err
	}
	stringSoFar := s.builder.String()
	matchIndexes := packageMatcher.FindStringIndex(stringSoFar)
	if matchIndexes == nil {
		return written, err
	}
	substrIndex := matchIndexes[1] + 1
	if substrIndex >= len(stringSoFar) {
		return written, nil
	}
	_, err = s.writer.Write([]byte(stringSoFar[substrIndex:]))
	s.builder.Reset()
	s.skipped = true
	return written, err
}
