package generate

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"strings"
)

const (
	ApiTypeHttp = "http"
	ApiTypeGrpc = "grpc"
	ApiBoth     = "both"
)

const (
	DatabaseTypePostgres = "postgres"
	DatabaseTypeMysql    = "mysql"
)

type StructName string

type Generator struct {
	// codeBuild code type
	GenTypes map[StructName][]GenType

	// Result save code source and code save path
	Results map[StructName]*Result

	// TemplateField used to template replace, key: struct name,
	TemplateField map[StructName]*TemplateField
}

type Result struct {
	// Path is code file's path
	Path map[PathName]string

	// Codes save code source
	Codes map[PathName]*bytes.Buffer
}

type TemplateField struct {
	Fields      []*Field
	Imports     []string
	StructName  StructName
	ProjectName string
}

func (tf *TemplateField) getFieldRule(name string) *Rule {
	for _, f := range tf.Fields {
		if name == f.Name {
			return f.Rule
		}
	}
	return nil
}

func (tf *TemplateField) getFieldTags(name string) []*Tag {
	for _, f := range tf.Fields {
		if name == f.Name {
			return f.Tags
		}
	}
	return nil
}

type Field struct {
	// Field name
	Name string `yaml:"name"`

	// Field type
	// support type: int string int32 int64 float64 time.Time and others type
	Type string `yaml:"type"`

	// Field's tags
	Tags []*Tag

	// Field's rules
	Rule *Rule
}

func (g *Generator) Generate() error {
	//var (
	//	err  error
	//	text []byte
	//)

	//fileName := Camel2Case("g.TemplateField.StructName")
	//templates := temp.GetTemplate()
	//for t, tmpl := range templates {
	//	var filepath string
	//	switch t {
	//	case temp.BllTemplate:
	//		filepath = fmt.Sprintf("%s/%s.go", destinations.Bll, fileName)
	//	case temp.HttpApiTemplate:
	//		if api == ApiTypeGrpc {
	//			continue
	//		}
	//		filepath = fmt.Sprintf("%s/%s.go", destinations.HttpApi, fileName)
	//	case temp.GrpcApiTemplate:
	//		filepath = fmt.Sprintf("%s/%s.go", destinations.GrpcApi, fileName)
	//		if api == ApiTypeHttp {
	//			continue
	//		}
	//	case temp.EntityTemplate:
	//		filepath = fmt.Sprintf("%s/%s.go", destinations.Entity, fileName)
	//	case temp.ModelTemplate:
	//		filepath = fmt.Sprintf("%s/%s.go", destinations.Model, fileName)
	//	case temp.StoreMysqlTemplate:
	//		if database == DatabaseTypePostgres {
	//			continue
	//		}
	//		filepath = fmt.Sprintf("%s/%s.go", destinations.StoreMysql, fileName)
	//	case temp.StorePostgresTemplate:
	//		if database == DatabaseTypeMysql {
	//			continue
	//		}
	//		filepath = fmt.Sprintf("%s/%s.go", destinations.StorePostgres, fileName)
	//	case temp.StoreInterfaceTemplate:
	//		filepath = fmt.Sprintf("%s/%s.go", destinations.StoreInterface, fileName)
	//	}
	//
	//	if text, err = parse(tmpl, g); err != nil {
	//		return err
	//	}
	//
	//	if !fileExists(filepath) {
	//		if err = writeFile(filepath, text); err != nil {
	//			return err
	//		}
	//	}
	//}

	return nil
}

func formatParams(params ...string) (ret string) {
	for i := 0; i < len(params); i++ {
		ret = fmt.Sprintf("%v/%v", ret, params[i])
	}
	return
}

func importNotExist(strType string) bool {
	var ok bool
	if _, ok = importExistMap[strType]; ok {
		return false
	}

	importExistMap[strType] = struct{}{}
	return true
}

func parse(temp string, generator *Generator) ([]byte, error) {
	var (
		tmpl = template.New("")
		err  error
		p    *template.Template
		buf  = bytes.NewBuffer([]byte{})
		src  []byte
	)
	if p, err = tmpl.Funcs(template.FuncMap{
		"notExist": importNotExist,
		"format":   formatParams,
	}).Parse(temp); err != nil {
		return nil, err
	}

	if err = p.Execute(buf, generator); err != nil {
		return nil, err
	}
	newStr := strings.Replace(buf.String(), "|| {", "{", -1)
	if src, err = format.Source([]byte(newStr)); err != nil {
		return nil, err
	}
	return src, nil
}
