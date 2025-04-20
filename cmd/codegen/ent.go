package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func getEntField(entityName string, path string) []*tableColumn {
	f, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
	if err != nil {
		panic(err)
	}

	columns := make([]*tableColumn, 0)

	ast.Inspect(f, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.TypeSpec:
			if n.Name.Name == entityName {
				for _, field := range n.Type.(*ast.StructType).Fields.List {
					// 構造体のフィールドで埋め込みフィールドとフィールド名CreateAtとUpdateAtは必要ない
					if len(field.Names) != 0 &&
						field.Names[0].Name != "CreatedAt" &&
						field.Names[0].Name != "UpdatedAt" &&
						field.Names[0].Name != "Edges" &&
						field.Names[0].Name != "selectValues" {
						var fieldType string
						switch t := field.Type.(type) {
						case *ast.Ident:
							// 基本型の場合（stringなど）
							fieldType = t.Name
						case *ast.SelectorExpr:
							// パッケージ名.型の場合（uuid.UUIDなど）
							fieldType = t.X.(*ast.Ident).Name + "." + t.Sel.Name
						case *ast.StarExpr:
							// ポインタ型の場合（*uuid.UUIDなど）
							fieldType = "*" + t.X.(*ast.SelectorExpr).X.(*ast.Ident).Name + "." + t.X.(*ast.SelectorExpr).Sel.Name
						default:
							panic("invalid ent structure")
						}

						columns = append(columns, &tableColumn{
							name:      field.Names[0].Name,
							fieldType: fieldType,
						})
					}
				}
			}
		}

		return true
	})

	return columns
}
