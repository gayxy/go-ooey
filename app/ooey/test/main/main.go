package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type ParamDef struct {
	Name     string // 参数名称
	Type     string // 参数类型
	Default  string // 参数默认值
	Help     string // 参数帮助信息
	Required bool   // 参数是否必须
}

func main() {

	// 解析参数定义
	params := parseParams("/Users/zhouwenzhe/src/go-ooey/app/ooey/test/script/script.go")

	// 输出参数列表
	for _, param := range params {
		fmt.Printf("Name: %s, Type: %s, Default: %s, Help: %s, Required: %t\n",
			param.Name, param.Type, param.Default, param.Help, param.Required)
	}
}
func parseParams(scriptPath string) []*ParamDef {
	// 创建一个文件集合
	fset := token.NewFileSet()

	// 解析脚本文件
	node, err := parser.ParseFile(fset, scriptPath, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("Failed to parse script:", err)
		return nil
	}

	// 遍历所有函数定义
	params := make([]*ParamDef, 0)
	for _, decl := range node.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			// 遍历函数参数列表
			for _, field := range funcDecl.Type.Params.List {
				// 解析参数名称和类型
				name := field.Names[0].Name
				typ := field.Type.(*ast.Ident).Name

				// 解析参数默认值
				defaultValue := ""
				//if field.Default != nil {
				//	defaultValue = field.Default.(*ast.BasicLit).Value
				//}

				// 解析参数帮助信息
				help := ""
				if field.Comment != nil {
					help = field.Comment.Text()
				}

				// 判断参数是否必须
				required := true
				if field.Type.(*ast.Ident).Name == "interface{}" {
					required = false
				}

				// 创建参数定义对象
				param := &ParamDef{
					Name:     name,
					Type:     typ,
					Default:  defaultValue,
					Help:     help,
					Required: required,
				}

				// 将参数定义添加到列表中
				params = append(params, param)
			}
		}
	}

	return params
}
