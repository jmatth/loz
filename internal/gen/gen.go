//go:build ignore

package main

import (
	"fmt"
	"iter"
	"os"
	"os/exec"
	"text/template"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type countLoopVal struct {
	I       int
	IsFirst bool
	IsLast  bool
}

type subTmplArgs struct {
	Index int
	BaseName string
}

func main() {
	output := os.Args[1]
	template := template.Must(template.New("map.go.tmpl").Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"skip": func(toSkip int, s iter.Seq[countLoopVal]) iter.Seq[countLoopVal] {
			return func(yield func(countLoopVal) bool) {
				skipped := 0
				for v := range s {
					if skipped < toSkip {
						skipped++
						continue
					}
					if !yield(v) {
						break
					}
				}
			}
		},
		"numsTo": func(n int) iter.Seq[countLoopVal] {
			return func(yield func(countLoopVal) bool) {
				for i := 1; i <= n; i++ {
					if !yield(countLoopVal{
						I:       i,
						IsFirst: i == 1,
						IsLast:  i == n,
					}) {
						break
					}
				}
			}
		},
		"subTmplArgs": func(baseName string, index int) subTmplArgs {
			return subTmplArgs{
				Index: index,
				BaseName: baseName,
			}
		},
	}).ParseGlob("./internal/gen/*.go.tmpl"))

	outFile, err := os.Create(fmt.Sprintf("%v.go", output))
	panicIfErr(err)

	err = template.Execute(outFile, map[string]any{
		"package": "loz",
		"levels":  9,
	})
	panicIfErr(err)

	err = exec.Command("go", "fmt", outFile.Name()).Run()
	panicIfErr(err)
}
