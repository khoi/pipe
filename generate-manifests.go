//go:build ignore
// +build ignore

// manifest/static.go

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"sort"
)

func main() {
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, "manifest", nil, 0)
	if err != nil {
		panic(err)
	}

	var manifests []string

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.VAR {
					for _, spec := range genDecl.Specs {
						if valueSpec, ok := spec.(*ast.ValueSpec); ok {
							if val, ok := valueSpec.Values[0].(*ast.CompositeLit); ok {
								if id, ok := val.Type.(*ast.Ident); ok && id.Name == "Manifest" {
									for _, name := range valueSpec.Names {
										if name.IsExported() {
											manifests = append(manifests, name.Name)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	generateCode(manifests)
}

func generateCode(manifests []string) {
	sort.Strings(manifests)

	file, err := os.Create("manifest/static.go")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprintln(file, "package manifest")
	fmt.Fprintln(file)
	fmt.Fprintln(file, "var staticManifests = []*Manifest{")

	for _, m := range manifests {
		fmt.Fprintf(file, "\t&%s,\n", m)
	}

	fmt.Fprintln(file, "}")
	fmt.Fprintln(file)
	fmt.Fprintln(file, "func StaticManifests() []*Manifest {")
	fmt.Fprintln(file, "\treturn staticManifests")
	fmt.Fprintln(file, "}")

	fmt.Println("Updated manifest/static.go")
}
