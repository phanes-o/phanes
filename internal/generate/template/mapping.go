package template

func init() {
	register(MappingTemplate, mapping)
}

var mapping = `
{{$true := true}}
{{$time := "time.Time"}}
{{$starTime := "*time.Time"}}

package mapping

import (
	"{{.ProjectName}}/model"
	"{{.ProjectName}}/model/entity"
)


// {{.StructName}}sEntityToDto entity data transfer
func {{.StructName}}sEntityToDto({{.CamelName}}s []*entity.{{.StructName}}) []*model.{{.StructName}}Info {
	out := make([]*model.{{.StructName}}Info, 0, len({{.CamelName}}s))
	for _, c := range {{.CamelName}}s  {
		out = append(out, {{.StructName}}EntityToDto(c))
	}
	return out
}

// {{.StructName}}EntityToDto entity data transfer
func {{.StructName}}EntityToDto(e *entity.{{.StructName}}) *model.{{.StructName}}Info {
	return &model.{{.StructName}}Info{
		{{range $v :=.Fields}}
			{{.Name}}: {{if or (eq .Type $time) (eq .Type $starTime)}}e.{{.Name}}.Unix(),{{else}}e.{{.Name}},{{end}}
		{{end}}
	}
}
`
