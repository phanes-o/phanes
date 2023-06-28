package template

func init() {
	register(ModelTemplate, model)
}

var model = `

{{$ID := "Id"}}
{{$create := "CreatedAt"}}
{{$true := "true"}}
{{$false := "false"}}

{{$string := "string"}}
{{$int64 := "int64"}}
{{$int32 := "int32"}}
{{$int := "int"}}
{{$point := "Point"}}
{{$strSlice := "pq.StringArray"}}
{{$int64Slice := "pq.Int64Array"}}
{{$projectName := .ProjectName}}
{{$titleName := .TitleName}}
{{$moduleName := "model"}}

package model

import (
	"{{.ProjectName}}/model/entity"

	{{range $value :=.Fields}}
		{{if and (eq $point .Type) (notExist (format $moduleName $point $titleName))}}
			"{{$projectName}}/model/po"	
		{{end}}

		{{if or (eq $strSlice .Type) (eq $int64Slice .Type) }}
			{{if notExist (format $moduleName "pq" $titleName)}}
				"github.com/lib/pq"
			{{end}}
		{{end}}
	{{end}}
)

// {{.TitleName}}CreateRequest create data request
type {{.TitleName}}CreateRequest struct {
{{range $value :=.Fields}}
	{{if ne $ID .Name}} 
		{{if eq .Parameter $true}}
			{{if eq .Required $true}}
				{{if eq $point .Type}} 
					{{.Name}} po.{{.Type}} {{.Char}}json:"{{$value.JsonTag}}" validate:"required"{{.Char}}
				{{else}} 
					{{.Name}} {{.Type}} {{.Char}}json:"{{$value.JsonTag}}" validate:"required"{{.Char}}
				{{end}}
			{{else}}
				{{if eq $point .Type}} 
					{{.Name}} po.{{.Type}} {{.Char}}json:"{{$value.JsonTag}}"{{.Char}}
				{{else}} 
					{{.Name}} {{.Type}} {{.Char}}json:"{{$value.JsonTag}}"{{.Char}}
				{{end}}
			{{end}}
		{{end}}
	{{end}}
{{end}}
}

// {{.TitleName}}UpdateRequest update request data
type {{.TitleName}}UpdateRequest struct {
	Id int64 {{.Char}}json:"id"{{.Char}}
{{range $value :=.Fields}}
	{{if eq $create .Name}} 
		{{.Name}} {{.Type}} {{.Char}}json:"{{$value.JsonTag}}"{{.Char}}
	{{else if eq .Parameter $true}}
		{{if eq .Required $true}}
			{{if eq $point .Type}} 
				{{.Name}} *po.{{.Type}} {{.Char}}json:"{{$value.JsonTag}}" validate:"required"{{.Char}}
			{{else}} 
				{{.Name}} *{{.Type}} {{.Char}}json:"{{$value.JsonTag}}" validate:"required"{{.Char}}
			{{end}}
		{{else}}
			{{if eq $point .Type}} 
				{{.Name}} *po.{{.Type}} {{.Char}}json:"{{$value.JsonTag}}"{{.Char}}
			{{else}} 
				{{.Name}} *{{.Type}} {{.Char}}json:"{{$value.JsonTag}}"{{.Char}}
			{{end}}
			
		{{end}}
	{{end}}
{{end}}
}

// {{.TitleName}}ListRequest data list request
type {{.TitleName}}ListRequest struct {
Index int {{.Char}}json:"index"{{.Char}}
Size int {{.Char}}json:"size"{{.Char}}
{{range $value :=.Fields}}
	{{if eq $ID .Name}} 
		{{.Name}} {{.Type}} {{.Char}}json:"{{$value.JsonTag}}"{{.Char}}
	{{else if eq .Parameter $true}}
		{{if eq .Required $true}}
			{{if eq $point .Type}} 
				{{.Name}} *po.{{.Type}} {{.Char}}json:"{{$value.JsonTag}}" validate:"required"{{.Char}}
			{{else}} 
				{{.Name}} *{{.Type}} {{.Char}}json:"{{$value.JsonTag}}" validate:"required"{{.Char}}
			{{end}}
		{{else}}
			{{if eq $point .Type}} 
				{{.Name}} *po.{{.Type}} {{.Char}}json:"{{$value.JsonTag}}"{{.Char}}
			{{else}} 
				{{.Name}} *{{.Type}} {{.Char}}json:"{{$value.JsonTag}}"{{.Char}}
			{{end}}
		{{end}}
	{{end}}
{{end}}
}


// {{.TitleName}}ListResponse data list response
type {{.TitleName}}ListResponse struct {
	Total int {{.Char}}json:"total"{{.Char}}
	List []*{{.TitleName}}Info {{.Char}}json:"list"{{.Char}}
}

// {{.TitleName}}InfoRequest detail info request
type {{.TitleName}}InfoRequest struct {
{{range $value :=.Fields}}
	{{if eq $ID .Name}} 
		{{.Name}} {{.Type}} {{.Char}}json:"{{$value.JsonTag}}"{{.Char}}
	{{else if eq .Parameter $true}}
		{{if eq .Required $true}}
			{{if eq $point .Type}} 
				{{.Name}} *po.{{.Type}} {{.Char}}json:"{{$value.JsonTag}}" validate:"required"{{.Char}}
			{{else}} 
				{{.Name}} *{{.Type}} {{.Char}}json:"{{$value.JsonTag}}" validate:"required"{{.Char}}
			{{end}}
		{{else}}
			{{if eq $point .Type}} 
				{{.Name}} *po.{{.Type}} {{.Char}}json:"{{$value.JsonTag}}"{{.Char}}
			{{else}} 
				{{.Name}} *{{.Type}} {{.Char}}json:"{{$value.JsonTag}}"{{.Char}}
			{{end}}
		{{end}}
	{{end}}
{{end}}

}

// {{.TitleName}}Info detail info
type {{.TitleName}}Info struct {
{{range $value :=.Fields}}
	{{if eq $point .Type}} 
		{{.Name}} po.{{.Type}} {{.Char}}json:"{{$value.JsonTag}}"{{.Char}}
	{{else}}
		{{.Name}} {{.Type}} {{.Char}}json:"{{$value.JsonTag}}"{{.Char}}
	{{end}}
{{end}}
}

// {{.TitleName}}DeleteRequest delete struct
type {{.TitleName}}DeleteRequest struct {
{{range $value :=.Fields}}
	{{if eq $ID .Name}} 
		{{.Name}} {{.Type}} {{.Char}}json:"{{$value.JsonTag}}"{{.Char}}
	{{end}}
{{end}}
}

// {{.TitleName}}sEntityToDto entity data transfer
func {{.TitleName}}sEntityToDto({{.Name}}s []*entity.{{.TitleName}}) []*{{.TitleName}}Info {
	out := make([]*{{.TitleName}}Info, 0, len({{.Name}}s))
	for _, c := range {{.Name}}s  {
		out = append(out, {{.TitleName}}EntityToDto(c))
	}
	return out
}

// {{.TitleName}}EntityToDto entity data transfer
func {{.TitleName}}EntityToDto(e *entity.{{.TitleName}}) *{{.TitleName}}Info {
	return &{{.TitleName}}Info{
		{{range $v :=.Fields}}
			{{.Name}}: {{if eq .Time $true}}e.{{.Name}}.Unix(),{{else}}e.{{.Name}},{{end}}
		{{end}}
	}
}
`
