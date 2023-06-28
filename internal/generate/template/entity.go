package template

func init() {
	register(EntityTemplate, entity)
}

var entity = `
{{$ID := "Id"}}
{{$true := "true"}}
{{$string := "string"}}
{{$int64 := "int64"}}
{{$int32 := "int32"}}
{{$int := "int"}}
{{$text := "text"}}
{{$point := "Point"}}
{{$json := "json"}}
{{$jsonb := "jsonb"}}
{{$time := "time.Time"}}
{{$strSlice := "pq.StringArray"}}
{{$int64Slice := "pq.Int64Array"}}
{{$projectName := .ProjectName}}
{{$titleName := .TitleName}}
{{$moduleName := "entity"}}
{{$postgres := "postgres"}}

package entity

import (
	{{if eq $postgres .DbType}}
		"github.com/lib/pq"
	{{end}}
)

type {{.TitleName}} struct {
{{range $value :=.Fields}}
	{{if eq $ID .Name}} 
		{{if eq $true $value.EntityTag.Auto}}
			{{.Name}} {{.Type}} {{.Char}}gorm:"column:{{$value.Name}};type:BIGINT;primary_key" json:"{{$value.Name}}"{{.Char}}
		{{else}}
			{{.Name}} {{.Type}} {{.Char}} {{range $v := $value.EntityTag.Tags}}$v.Name:"$v.Value" {{end}}{{.Char}}
		{{end}}
	{{else if eq $time .Type}}
		{{if eq $true $value.EntityTag.Auto}}
			{{.Name}} {{.Type}} {{.Char}}gorm:"column:{{$value.Name}};type:TIMESTAMP" json:"{{$value.Name}}"{{.Char}}
		{{else}}
			{{.Name}} {{.Type}} {{.Char}} {{range $v := $value.EntityTag.Tags}}$v.Name:"$v.Value" {{end}}{{.Char}}
		{{end}}
	{{else if eq $int64 .Type}} 
		{{if eq $true $value.EntityTag.Auto}}
			{{.Name}} {{.Type}} {{.Char}}gorm:"column:{{$value.Name}};type:BIGINT" json:"{{$value.Name}}"{{.Char}}
		{{else}}
			{{.Name}} {{.Type}} {{.Char}} {{range $v := $value.EntityTag.Tags}}$v.Name:"$v.Value" {{end}}{{.Char}}
		{{end}}
	{{else if eq $string .Type}} 
		{{if eq $true $value.EntityTag.Auto}}
			{{.Name}} {{.Type}} {{.Char}}gorm:"column:{{$value.Name}};type:VARCHAR(255)" json:"{{$value.Name}}"{{.Char}}
		{{else}}
			{{.Name}} {{.Type}} {{.Char}} {{range $v := $value.EntityTag.Tags}}$v.Name:"$v.Value" {{end}}{{.Char}}
		{{end}}
	{{else if eq $int32 .Type}} 
		{{if eq $true $value.EntityTag.Auto}}
			{{.Name}} {{.Type}} {{.Char}}gorm:"column:{{$value.Name}};type:TINYINT" json:"{{$value.Name}}"{{.Char}}
		{{else}}
			{{.Name}} {{.Type}} {{.Char}} {{range $v := $value.EntityTag.Tags}}$v.Name:"$v.Value" {{end}}{{.Char}}
		{{end}}
	{{else if eq $int .Type}} 
		{{if eq $true $value.EntityTag.Auto}}
			{{.Name}} {{.Type}} {{.Char}}gorm:"column:{{$value.Name}};type:TINYINT" json:"{{$value.Name}}"{{.Char}}
		{{else}}
			{{.Name}} {{.Type}} {{.Char}} {{range $v := $value.EntityTag.Tags}}$v.Name:"$v.Value" {{end}}{{.Char}}
		{{end}}
	{{else if eq $text .Type}} 
		{{if eq $true $value.EntityTag.Auto}}
			{{.Name}} {{.Type}} {{.Char}}gorm:"column:{{$value.Name}};type:TEXT" json:"{{$value.Name}}"{{.Char}}
		{{else}}
			{{.Name}} {{.Type}} {{.Char}} {{range $v := $value.EntityTag.Tags}}$v.Name:"$v.Value" {{end}}{{.Char}}
		{{end}}
	{{else if eq $point .Type}} 
		{{if eq $true $value.EntityTag.Auto}}
			{{.Name}} po.{{.Type}} {{.Char}}gorm:"column:{{$value.Name}};type:POINT" json:"{{$value.Name}}"{{.Char}}
		{{else}}
			{{.Name}} {{.Type}} {{.Char}} {{range $v := $value.EntityTag.Tags}}$v.Name:"$v.Value" {{end}}{{.Char}}
		{{end}}
	{{else if eq $strSlice .Type}} 
		{{if eq $true $value.EntityTag.Auto}}
			{{.Name}} {{.Type}} {{.Char}}gorm:"column:{{$value.Name}};type:VARCHAR[]" json:"{{$value.Name}}"{{.Char}}
		{{else}}
			{{.Name}} {{.Type}} {{.Char}} {{range $v := $value.EntityTag.Tags}}$v.Name:"$v.Value" {{end}}{{.Char}}
		{{end}}
	{{else if eq $jsonb .Type}} 
		{{if eq $true $value.EntityTag.Auto}}
			{{.Name}} {{.Type}} {{.Char}}gorm:"column:{{$value.Name}};type:JSONB" json:"{{$value.Name}}"{{.Char}}
		{{else}}
			{{.Name}} {{.Type}} {{.Char}} {{range $v := $value.EntityTag.Tags}}$v.Name:"$v.Value" {{end}}{{.Char}}
		{{end}}
	{{else if eq $json .Type}} 
		{{if eq $true $value.EntityTag.Auto}}
			{{.Name}} {{.Type}} {{.Char}}gorm:"column:{{$value.Name}};type:JSON" json:"{{$value.Name}}"{{.Char}}
		{{else}}
			{{.Name}} {{.Type}} {{.Char}} {{range $v := $value.EntityTag.Tags}}$v.Name:"$v.Value" {{end}}{{.Char}}
		{{end}}
	{{end}}
{{end}}
}

func (a *{{.TitleName}}) TableName() string {
	return "{{.FileName}}s"
}
`
