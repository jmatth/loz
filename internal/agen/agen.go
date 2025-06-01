package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"

	. "github.com/dave/jennifer/jen" //nolint: staticcheck
	"github.com/jmatth/loz"
)

func main() {
	inPath := os.Args[1]
	inType := os.Args[2]
	outType := os.Args[3]
	outPath := os.Args[4]

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
				name: outType,
				typeParams: types,
			}
		case *ast.IndexListExpr:
			if recvTypeName := t.X.(*ast.Ident).Name; recvTypeName != inType {
				return true
			}
			types := loz.IterSliceMap1[ast.Expr, string](t.Indices).
				Map(func(e ast.Expr) string {
					return e.(*ast.Ident).Name
				}).
				CollectSlice()
			recvTypedef = typeDef{
				name: outType,
				typeParams: types,
			}
		default:
			return true
		}

		methodName := funcDecl.Name.Name

		var returnTypes []funcParam
		if results := funcDecl.Type.Results; results != nil {
			returnTypes = loz.IterSliceMap1[*ast.Field, funcParam](funcDecl.Type.Results.List).
				Map(fieldToFuncParam).
				CollectSlice()
			if len(returnTypes) == 1 && returnTypes[0].typeDef.name == inType {
				returnTypes[0].typeDef.name = outType
			}
		}

		paramTypes := funcDecl.Type.Params.List

		pt := loz.IterSliceMap1[*ast.Field, funcParam](paramTypes).
			Map(fieldToFuncParam).
			CollectSlice()

		genWrapper(outFile, methodName, recvTypedef, inType, pt, returnTypes)

		return true
	})

	outFileDisk, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	_, err = outFileDisk.WriteString(outFile.GoString())
	if err != nil {
		panic(err)
	}
}

func fieldToFuncParam(f *ast.Field) funcParam {
	td := typeDef{}
	switch t := f.Type.(type) {
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
		td.typeParams = loz.IterSliceMap1[ast.Expr, string](t.Indices).
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
	default:
		td.name = "TODO"
	}
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
}

type funcParam struct {
	name    string
	typeDef typeDef
}

func (t typeDef) paramsAsCode() []Code {
	return loz.IterSliceMap1[string, Code](t.typeParams).
		Map(func(s string) Code { return Id(s) }).
		CollectSlice()
}

func genWrapper(f *File, name string, recvType typeDef, originalType string, params []funcParam, returnTypes []funcParam) {
	recvId := Id(strings.ToLower(string(recvType.name[0])))
	generatedParams := loz.IterSliceMap1[funcParam, Code](params).
		Map(paramToCode).
		CollectSlice()
	innerParams := loz.IterSliceMap1[funcParam, Code](params).
		Map(func(p funcParam) Code { return Id(p.name) }).
		CollectSlice()
	generatedReturnTypes := loz.IterSliceMap1[funcParam, Code](returnTypes).
		Map(func(p funcParam) Code {
			if p.typeDef.isArr {
				return Index().Id(p.name).Id(p.typeDef.name).Types(p.typeDef.paramsAsCode()...)
			}
			return Id(p.name).Id(p.typeDef.name).Types(p.typeDef.paramsAsCode()...)
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
	return Id(param.name).
		Id(param.typeDef.name).
		TypesFunc(func(g *Group) {
			for _, pt := range param.typeDef.typeParams {
				g.Id(pt)
			}
		})
}
