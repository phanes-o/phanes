package register

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"path"
	"strconv"
	"strings"

	"errors"
)

const packagePrefix = "/server/web/v1"
const commonResourceKey = "common"
const RegisterKeyPrefix = "/phanes/register_resource"

var (
	NotEnoughParams     = errors.New("not enough parameters")
	MultipleParentError = errors.New("multiple parent in a group is not allowed")
)

type Tmpl struct {
	Fields []*Field
}
type Field struct {
	Name  string
	Key   string
	Value []byte
}

func parsePackage(pwd string, project string) ([]byte, error) {
	var (
		p      = path.Join(pwd, project, packagePrefix)
		err    error
		text   []byte
		fields = &Tmpl{Fields: make([]*Field, 0)}
	)
	// parse package

	packages, err := parser.ParseDir(token.NewFileSet(), p, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("Error parsing package:", err)
		return nil, err
	}
	var common = make([]*Resource, 0)
	// find comment by ast file
	for _, pkg := range packages {
		for _, f := range pkg.Files {
			if f.Name.String() == "init.go" {
				continue
			}
			// find Init method
			for _, decl := range f.Decls {
				if fn, ok := decl.(*ast.FuncDecl); ok && fn.Name.Name == "Init" {
					// get Init method comment
					if fn.Doc != nil {
						resources, name, err := parseRegisterComment(fn.Doc)
						if err != nil {
							if errors.Is(err, NotEnoughParams) {
								fmt.Printf("Error: package: %s, line: %v, error: %+v", pkg.Name, fn.Pos(), err)
								return nil, err
							}
							if errors.Is(err, MultipleParentError) {
								fmt.Printf("Error: package: %s, line: %v, error: %+v", pkg.Name, fn.Pos(), err)
								fmt.Println(pkg.Name)
								return nil, err
							}
						}
						bytes, _ := json.Marshal(resources)
						if name != "" {
							fields.Fields = append(fields.Fields, &Field{
								Name:  name,
								Key:   RegisterKeyPrefix + "/" + name,
								Value: bytes,
							})
						} else {
							common = append(common, resources...)
						}
					}
				}
			}
		}
	}
	if len(common) > 0 {
		bytes, _ := json.Marshal(common)
		f := &Field{
			Name:  commonResourceKey,
			Key:   RegisterKeyPrefix + "/" + commonResourceKey,
			Value: bytes,
		}
		fields.Fields = append(fields.Fields, f)
	}
	// generate code by template
	if text, err = parse(tmpl, fields); err != nil {
		return nil, err
	}
	return text, nil
}

func parseRegisterComment(doc *ast.CommentGroup) ([]*Resource, string, error) {
	//#[register("auth", "method", "v1/auth")]
	//#[register("auth.authorization", "method", "v1/auth/authorization")]
	var (
		parentName string
		resources  = make([]*Resource, 0)
	)
	for _, doc := range strings.Split(doc.Text(), "\n") {
		if len(doc) <= 0 {
			continue
		}
		if doc[0] != '#' {
			continue
		}
		str := strings.TrimRight(strings.TrimLeft(doc, "#["), "]")

		if strings.Contains(str, "register(") {
			str := strings.TrimRight(strings.TrimLeft(str, "register("), ")")
			nodes := strings.Split(str, ",")
			if len(nodes) != 3 {
				return nil, "", NotEnoughParams
			}

			resource := &Resource{
				Name: strings.Trim(strings.Trim(nodes[0], " "), "\""),
				Type: strings.Trim(strings.Trim(nodes[1], " "), "\""),
				Path: strings.Trim(strings.Trim(nodes[2], " "), "\""),
			}

			if !strings.Contains(nodes[0], ".") {
				if parentName != "" {
					return nil, "", MultipleParentError
				}
				resource.Parent = true
				parentName = strings.Trim(strings.Trim(nodes[0], " "), "\"")
			}
			resources = append(resources, resource)
		}
	}
	return resources, parentName, nil
}

func parse(temp string, fields *Tmpl) ([]byte, error) {
	var (
		err  error
		src  []byte
		tmpl = template.New("")
		p    *template.Template
		buf  = bytes.NewBuffer([]byte{})
	)
	if p, err = tmpl.Funcs(template.FuncMap{"keepBytesType": keepBytesType}).Parse(temp); err != nil {
		return nil, err
	}

	if err = p.Execute(buf, fields); err != nil {
		return nil, err
	}
	if src, err = format.Source([]byte(buf.Bytes())); err != nil {
		return nil, err
	}
	return src, nil
}

func keepBytesType(data []byte) string {
	if len(data) == 0 {
		return "[]byte{}"
	}
	var str = "[]byte{"
	for i, v := range data {
		if i == 0 {
			str += strconv.Itoa(int(v))
		} else {
			str += ", " + strconv.Itoa(int(v))
		}
	}
	return str + "}"
}
