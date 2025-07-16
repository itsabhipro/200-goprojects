package struc

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func getNode() (*ast.File, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "inspection.go", nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return node, nil
}

func ExploreFile() {
	node, err := getNode()
	if err != nil {
		fmt.Println("Error from inspection.go : ", err)
		return
	}
	ast.Inspect(node, func(n ast.Node) bool {
		genDecl, ok := n.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			return true
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			if structType, ok := typeSpec.Type.(*ast.StructType); ok {
				fmt.Printf("Struct: %s\n", typeSpec.Name.Name)
				for _, field := range structType.Fields.List {
					for _, fieldName := range field.Names {
						fmt.Printf("  Field: %s\n", fieldName.Name)
					}
				}
			}
		}
		return true
	})
}
