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

func main() {
	output := os.Args[1]
	template := template.Must(template.New("map.tmpl.go").Funcs(template.FuncMap{
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
	}).ParseGlob("./internal/gen/*.tmpl.go"))

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
