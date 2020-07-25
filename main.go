package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func main() {
	run()
}

func run() {
	fset := token.NewFileSet()
	ret, err := parser.ParseDir(fset, ".", nil, parser.Mode(0))
	if err != nil {
		return
	}

	tests := []string{}
	for _, v := range ret {
		for _, f := range v.Files {
			for _, d := range f.Decls {
				tmp, ok := d.(*ast.FuncDecl)
				if !ok {
					continue
				}
				if strings.HasPrefix(tmp.Name.String(), "Test") {
					tests = append(tests, tmp.Name.String())
				}
			}
		}
	}

	if len(tests) == 0 {
		fmt.Println("Test")
		return
	}
	fmt.Println(strings.Join(tests, " "))
}
