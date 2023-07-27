package template

func init() {
	register(BllTemplate, bll)
}

var bll = `
{{$zero := 0}}
{{$empty := ""}}
{{$nil := "nil"}}
{{$true := true}}
{{$ID := "Id"}}

{{$time := "time.Time"}}
{{$starTime := "*time.Time"}}
{{$string := "string"}}
{{$int64 := "int64"}}
{{$int32 := "int32"}}
{{$int := "int"}}

{{$pqStringArray := "pq.StringArray"}}
{{$pqFloat32Array := "pq.Float32Array"}}
{{$pqFloat64Array := "pq.Float64Array"}}
{{$pqInt32Array := "pq.Int32Array"}}
{{$pqInt64Array := "pq.Int64Array"}}

{{$required := "Required"}}
{{$projectName := .ProjectName}}
{{ $break := false }}
package bll 

import (
	"context"
	{{range $v :=.Fields}}
		{{if $break}}
			{{break}}
		{{end}}
		{{if eq .Rule.Parameter $true}}
			{{if ne .Rule.Required $true}}
				{{if or (eq .Type $pqStringArray) (eq .Type $pqFloat32Array) (eq .Type $pqFloat64Array) (eq .Type $pqInt32Array) (eq .Type $pqInt64Array)}}
				"github.com/lib/pq"
				{{ $break = true }}
				{{end}}
			{{end}}
		{{end}}
	{{end}}
	{{ $break = false }}
	"{{.Module}}/model"
	"{{.Module}}/model/entity"
	"{{.Module}}/model/mapping"
	"{{.Module}}/store"
	"{{.Module}}/store/postgres"
	"time"
)

type {{.CamelName}} struct{
	i{{.StructName}} store.I{{.StructName}}
}

var {{.StructName}} = &{{.CamelName}}{
	i{{.StructName}}: postgres.{{.StructName}},
}

func init() {
	Register({{.StructName}})
}

func (a *{{.CamelName}}) init()     func()   {
	return func() {}
}

// Create
func (a *{{.CamelName}}) Create(ctx context.Context, in *model.{{.StructName}}CreateRequest) error  {
	var (
		err error
	)
	c := build{{.StructName}}(in)
	_, err = a.i{{.StructName}}.Create(ctx,c)
	return err
}

// Update 
func (a *{{.CamelName}}) Update(ctx context.Context, in *model.{{.StructName}}UpdateRequest) error  {
	var (
		dict = make(map[string]interface{})
	)
	{{range $v := .Fields}}
		{{if eq .Rule.Parameter $true}}
			{{if ne .Rule.Required $true}}
			if in.{{.Name}} != nil {
				dict["{{.SnakeName}}"] = in.{{.Name}}
			}
			{{end}}
		{{end}}
	{{end}}
	// do other update here
	updateAt := time.Now().Unix()
	in.UpdatedAt = &updateAt
	return a.i{{.StructName}}.Update(ctx, in.Id, dict)
}

// Delete 
func (a *{{.CamelName}}) Delete(ctx context.Context, in *model.{{.StructName}}DeleteRequest) error  {
	return a.i{{.StructName}}.Delete(ctx,in.Id)
}

// List 
func (a *{{.CamelName}}) List(ctx context.Context, in *model.{{.StructName}}ListRequest) (*model.{{.StructName}}ListResponse, error)  {
	var (
		err error
		total int
		list []*entity.{{.StructName}} 
		out = &model.{{.StructName}}ListResponse{}
	)

	if total, list, err = a.i{{.StructName}}.List(ctx,in); err != nil {
		return nil, err
	}
	
	out.Total = total
	out.List = mapping.{{.StructName}}sEntityToDto(list)

	return out, nil
}

// Find 
func (a *{{.CamelName}}) Find(ctx context.Context, in *model.{{.StructName}}InfoRequest) (*model.{{.StructName}}Info, error)  {
	var (
		err error
		data *entity.{{.StructName}} 
		out = &model.{{.StructName}}Info{}
	)

	if data, err = a.i{{.StructName}}.Find(ctx,in); err != nil {
		return nil, err
	}
	
	out = mapping.{{.StructName}}EntityToDto(data)
	return out, nil
}

// build{{.StructName}} build entity
func build{{.StructName}}(in *model.{{.StructName}}CreateRequest) *entity.{{.StructName}} {
	{{range $v :=.Fields}}
		{{if eq $break $true}}
			{{break}}
		{{end}}
		{{if eq .Type $time}}
			{{if eq .Rule.AutoFill $true}}
				now := time.Now()
				{{ $break = true }}
			{{end}}
		{{end}}
	{{end}}

	ety := &entity.{{.StructName}}{
		{{range $v :=.Fields}}
			{{if eq .Type $time}}
				{{if eq .Rule.AutoFill $true}}	
					{{.Name}}: now,
				{{else}}
					{{.Name}}: time.Unix(0, 0),
				{{end}}
			{{else if eq .Type $starTime}}
				{{if eq .Rule.AutoFill $true}}	
					{{.Name}}: &now,
				{{else}}
					{{.Name}}: time.Unix(0, 0),
				{{end}}
			{{else}}
				{{if ne .Name $ID}}
					{{if eq .Rule.Parameter $true}} 
						{{if eq .Rule.Required $true}}
							{{.Name}}: in.{{.Name}},
						{{end}}
					{{else}}
						{{if eq .Type $string}}
							{{.Name}}: "",
						{{else if eq .Type $pqStringArray}}
							{{.Name}}: pq.StringArray{},
						{{else if eq .Type $pqFloat32Array}}
							{{.Name}}: pq.Float32Array{},
						{{else if eq .Type $pqFloat64Array}}
							{{.Name}}: pq.Float32Array{},
						{{else if eq .Type $pqInt32Array}}
							{{.Name}}: pq.Int32Array{},
						{{else if eq .Type $pqInt32Array}}
							{{.Name}}: pq.Int64Array{},
						{{else}}
							{{.Name}}: 0,
						{{end}}
					{{end}}
				{{end}}
			{{end}}
		{{end}}
	} 
	{{range $v :=.Fields}}
		{{if eq .Rule.Parameter $true}}
			{{if ne .Rule.Required $true}}
				{{if or (eq .Type $string) (eq .Type $int) (eq .Type $int32) (eq .Type $int64)}}
					if in.{{.Name}} != nil {
						ety.{{.Name}} = *in.{{.Name}}
					}
				{{else if eq .Type $pqStringArray}}
					if len(in.{{.Name}}) != 0 {
						ety.{{.Name}} = in.{{.Name}}
					} else {
						ety.{{.Name}} = pq.StringArray{}
					}
				{{else if eq .Type $pqFloat32Array}}
					if len(in.{{.Name}}) != 0 {
						ety.{{.Name}} = *in.{{.Name}}
					}else {
						ety.{{.Name}} = pq.Float32Array{}
					}
				{{else if eq .Type $pqFloat64Array}}
					if len(in.{{.Name}}) != 0 {
						ety.{{.Name}} = *in.{{.Name}}
					}else {
						ety.{{.Name}} = pq.Float64Array{}
					}
				{{else if eq .Type $pqInt32Array}}
					if len(in.{{.Name}}) != 0 {
						ety.{{.Name}} = *in.{{.Name}}
					}else {
						ety.{{.Name}} = pq.Int32Array{}
					}
				{{else if eq .Type $pqInt64Array}}
					if len(in.{{.Name}}) != 0 {
						ety.{{.Name}} = *in.{{.Name}}
					}else {
						ety.{{.Name}} = pq.Int64Array{}
					}
				{{end}}
			{{end}}
		{{end}}
	{{end}}
	return ety
}
`
