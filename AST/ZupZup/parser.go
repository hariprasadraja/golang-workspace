// Sample program for golang ast parsing
//
// Source: https://zupzup.org/go-ast-traversal/

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "AST/ZupZup/sample.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// TEMP: Snippet for debugging. remove it before commit
	spew.Print(strings.Repeat("-", 15) + "	Debug Begin: node" + strings.Repeat("-", 15) + "\n")
	spew.Dump(node)
	spew.Print(strings.Repeat("-", 15) + "	Debug End: node	" + strings.Repeat("-", 15) + "\n")

	fmt.Println("Imports:")
	for _, i := range node.Imports {
		fmt.Println(i.Path.Value)
	}

	fmt.Println("Comments:")
	for _, c := range node.Comments {
		fmt.Print(c.Text())
	}

	fmt.Println("Functions:")
	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}
		fmt.Println(fn.Name.Name)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		// Find Return Statements
		ret, ok := n.(*ast.ReturnStmt)
		if ok {
			fmt.Printf("return statement found on line %d:\n\t", fset.Position(ret.Pos()).Line)
			printer.Fprint(os.Stdout, fset, ret)
		}

		// Find Functions
		fn, ok := n.(*ast.FuncDecl)
		if ok {
			var exported string
			if fn.Name.IsExported() {
				exported = "exported "
			}
			fmt.Printf("%sfunction declaration found on line %d: \n\t%s\n", exported, fset.Position(fn.Pos()).Line, fn.Name.Name)
		}

		return true
	})
}
