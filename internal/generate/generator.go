package generate

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"strings"

	"github.com/fatih/color"
	templ "github.com/phanes-o/phanes/internal/generate/template"
	"github.com/phanes-o/phanes/internal/utils"
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
	CamelName   string
	ProjectName string
	Module      string
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
	Name string

	SnakeName string
	// Field type
	// support type: int string int32 int64 float64 time.Time and others type
	Type string

	// Field's tags
	Tags []*Tag

	// Field's rules
	Rule *Rule
}

func (g *Generator) Generate() error {
	var (
		ok       bool
		err      error
		genTypes []GenType
		result   *Result
		buf      *bytes.Buffer
	)

	for structName, tmplField := range g.TemplateField {
		if genTypes, ok = g.GenTypes[structName]; !ok {
			continue
		}
		tmplField.CamelName = camelCase(string(tmplField.StructName))

		result = g.Results[structName]
		for _, t := range genTypes {
			if result.Codes == nil {
				result.Codes = make(map[PathName]*bytes.Buffer)
			}
			switch t {
			case GenTypeBll:
				if buf, err = buildCodeFromTemplate(templ.BllTemplate, tmplField); err != nil {
					return err
				}
				result.Codes[BllName] = buf
			case GenTypeApiAll:
				if buf, err = buildCodeFromTemplate(templ.GrpcApiTemplate, tmplField); err != nil {
					return err
				}
				result.Codes[GrpcApiName] = buf

				if buf, err = buildCodeFromTemplate(templ.HttpApiTemplate, tmplField); err != nil {
					return err
				}
				result.Codes[HttpApiName] = buf
			case GenTypeHttpApi:
				if buf, err = buildCodeFromTemplate(templ.HttpApiTemplate, tmplField); err != nil {
					return err
				}
				result.Codes[HttpApiName] = buf
			case GenTypeGrpcApi:
				if buf, err = buildCodeFromTemplate(templ.GrpcApiTemplate, tmplField); err != nil {
					return err
				}
				result.Codes[GrpcApiName] = buf

			case GenTypeStoreMysql:
				if buf, err = buildCodeFromTemplate(templ.StoreInterfaceTemplate, tmplField); err != nil {
					return err
				}
				result.Codes[StoreInterfaceName] = buf

				if buf, err = buildCodeFromTemplate(templ.StoreMysqlTemplate, tmplField); err != nil {
					return err
				}
				result.Codes[StoreMysqlName] = buf
			case GenTypeStorePostgres:
				if buf, err = buildCodeFromTemplate(templ.StoreInterfaceTemplate, tmplField); err != nil {
					return err
				}
				result.Codes[StoreInterfaceName] = buf

				if buf, err = buildCodeFromTemplate(templ.StorePostgresTemplate, tmplField); err != nil {
					return err
				}
				result.Codes[StorePostgresName] = buf
			}
		}
		if buf, err = buildCodeFromTemplate(templ.MappingTemplate, tmplField); err != nil {
			return err
		}
		result.Codes[MappingName] = buf

		g.Results[structName] = result
	}

	g.save()

	return nil
}

func (g *Generator) save() {
	for k, v := range g.Results {
		for pathName, code := range v.Codes {
			path := v.Path[pathName]
			if !utils.FileExists(path) {
				if err := utils.WriteFile(path, code.Bytes()); err != nil {
					fmt.Println(color.RedString(fmt.Sprintf("ERROR: Failed to save [%s] code", k)), "❌  ")
					continue
				}
				fmt.Println(color.GreenString(fmt.Sprintf("[%s] code generate successfully!", path)), "✅  ")
			} else {
				fmt.Println(color.YellowString(fmt.Sprintf("Notify: [%s] code already exist", path)), "⚠️ ")
			}
		}
	}

}

func buildCodeFromTemplate(templType templ.Type, fields *TemplateField) (*bytes.Buffer, error) {
	var (
		err  error
		text []byte
	)
	tmpl := templ.Get(templType)
	if text, err = parse(tmpl, fields); err != nil {
		return nil, err
	}
	return bytes.NewBuffer(text), nil
}

func genTypeTransToPathName(t GenType) PathName {
	switch t {
	case GenTypeBll:
		return BllName
	case GenTypeModel:
		return ModelName
	case GenTypeEntity:
		return EntityName
	case GenTypeStoreMysql:
		return StoreMysqlName
	case GenTypeStorePostgres:
		return StorePostgresName
	}
	return ""
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

func parse(temp string, fields *TemplateField) ([]byte, error) {
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

	if err = p.Execute(buf, fields); err != nil {
		return nil, err
	}
	newStr := strings.Replace(buf.String(), "|| {", "{", -1)
	if src, err = format.Source([]byte(newStr)); err != nil {
		return nil, err
	}
	return src, nil
}
