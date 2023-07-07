package generate

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/tools/go/ast/astutil"
)

type GenCmd string

const (
	CmdProject  GenCmd = "project"
	CmdGenerate GenCmd = "generate"
	CmdDir      GenCmd = "dir"
)

func ReadSource(filename string) (*Generator, error) {
	var (
		err           error
		pwd           string
		file          *ast.File
		fileSet       = token.NewFileSet()
		project       string
		imports       = make([]string, 0)
		results       = make(map[StructName]*Result)
		genTypes      = make(map[StructName][]GenType)
		templateField = make(map[StructName]*TemplateField)
	)
	if file, err = parser.ParseFile(fileSet, filename, nil, parser.ParseComments); err != nil {
		return nil, err
	}

	if pwd, err = os.Getwd(); err != nil {
		fmt.Println(color.RedString(fmt.Sprintf("ERROR: %s", err)))
		os.Exit(1)
	}

	currentStruct := StructName("")

	currentGenType := make([]GenType, 0)
	currentPath := make(map[PathName]string)
	astutil.Apply(file, nil, func(c *astutil.Cursor) bool {
		// resolve field and tag
		node := c.Node()
		switch node := node.(type) {
		case *ast.TypeSpec:
			if n, ok := node.Type.(*ast.StructType); ok {
				structName := StructName(node.Name.Name)
				if currentStruct != structName {
					currentStruct = structName
				}
				tmpl := buildTemplateField(project, structName, n)
				tmpl.Imports = imports
				templateField[structName] = tmpl

				if len(currentGenType) == 0 {
					currentGenType = defaultGenType
				}
				result := &Result{
					Path:  resolvePaths(project, pwd, currentStruct, currentPath, destinations(project, pwd)),
					Codes: codeBuild(currentGenType, node, tmpl),
				}

				genTypes[structName] = currentGenType
				results[structName] = result

				currentPath = make(map[PathName]string)
				currentGenType = make([]GenType, 0)
			}
		case *ast.ImportSpec:
			imports = append(imports, node.Path.Value)
		case *ast.Comment:
			comment := strings.TrimLeft(node.Text, "//")
			split := strings.Split(comment, ":")
			if len(split) != 2 {
				fmt.Fprint(os.Stderr, "\033[31mERROR: your command error  \033[m\n")
				os.Exit(0)
				return false
			}
			switch GenCmd(split[0]) {
			case CmdProject:
				project = split[1]
				if !checkProjectExist(path.Join(pwd, project)) {
					fmt.Println(color.RedString(fmt.Sprintf("Project [%s] does not exist", project)), "❌ ")
					fmt.Println(color.BlueString("Please run [phanes new example] create a project first"))
					os.Exit(1)
				}
			case CmdGenerate:
				currentGenType = append(currentGenType, parseCommentGenType(split[1])...)
			case CmdDir:
				name, path := parseCommentDir(split[1])
				currentPath[name] = path
			}
		}
		return true
	})

	return &Generator{
		Results:       results,
		GenTypes:      genTypes,
		TemplateField: templateField,
	}, nil
}

func processFieldType(fieldType ast.Expr) string {
	var res string
	switch fieldType := fieldType.(type) {
	case *ast.Ident:
		res = fieldType.String()
	case *ast.StarExpr:
		res = fmt.Sprintf("%s%s", "*", processFieldType(fieldType.X))
	case *ast.ArrayType:
		fmt.Fprintf(os.Stderr, "\033[31mERROR: Unsupport Array Field Type\033[m\n")
		os.Exit(1)
	case *ast.MapType:
		fmt.Fprintf(os.Stderr, "\033[31mERROR: Unsupport Map Field Type\033[m\n")
		os.Exit(1)
	case *ast.StructType:
		fmt.Fprintf(os.Stderr, "\033[31mERROR: Unsupport Struct Field Type\033[m\n")
		os.Exit(1)
	case *ast.SelectorExpr:
		//fmt.Printf("Package: %s, Type: %s\n", fieldType.X, fieldType.Sel)
		res = fmt.Sprintf("%s.%s", fieldType.X, fieldType.Sel)
	default:
		fmt.Fprintf(os.Stderr, "\033[31mERROR: Unknown Field Type\033[m\n")
		os.Exit(1)
	}
	return res
}
