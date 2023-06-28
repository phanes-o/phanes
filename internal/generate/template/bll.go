package template

func init() {
	register(BllTemplate, bll)
}

var bll = `
{{$zero := 0}}
{{$empty := ""}}
{{$nil := "nil"}}
{{$true := "true"}}
{{$or := "||"}}
{{$ID := "Id"}}
{{$CreatedAt := "created_at"}}
{{$UpdatedAt := "updated_at"}}
{{$UserId := "user_id"}}

{{$string := "string"}}
{{$int64 := "int64"}}
{{$int32 := "int32"}}
{{$int := "int"}}

{{$projectName := .ProjectName}}

package bll 

import (
	"context"
	
	"{{.ProjectName}}/event"
	"{{.ProjectName}}/model"
	"{{.ProjectName}}/model/entity"
	"{{.ProjectName}}/store"
	"{{.ProjectName}}/store/postgres"
	"time"

	{{range $value :=.Fields}}
		{{if eq $value.JsonTag "user_id" }}
			"{{$projectName}}/auth"
		{{end}}
	{{end}}
)

type {{.Name}} struct{
	i{{.TitleName}} store.I{{.TitleName}}
}

var {{.TitleName}} = &{{.Name}}{
	i{{.TitleName}}: postgres.{{.TitleName}},
}

func (a *{{.Name}}) init()     func()   {
	return func() {}
}

func (a *{{.Name}}) onEvent(*event.Data) {}

// Create
func (a *{{.Name}}) Create(ctx context.Context, in *model.{{.TitleName}}CreateRequest) error  {
	var (
		err error
	)
	
	{{range $v := .Fields}}
		{{if eq .Json $UserId}}
			in.UserId, _ = auth.ContextUserID(ctx)
		{{end}}
	{{end}}

	c := build{{.TitleName}}(in)
	_, err = a.i{{.TitleName}}.Create(ctx,c)
	return err
}

// Update 
func (a *{{.Name}}) Update(ctx context.Context, in *model.{{.TitleName}}UpdateRequest) error  {
	var (
		dict = make(map[string]interface{})
	)
	{{range $v := .Fields}}
		{{if eq .Parameter $true}}
			{{if ne .Required $true}}
			if in.{{.Name}} != nil {
				dict["{{.Json}}"] = in.{{.Name}}
			}
			{{end}}
		{{end}}
	{{end}}
	// do other update here
	updateAt := time.Now().Unix()
	in.UpdatedAt = &updateAt
	return a.i{{.TitleName}}.Update(ctx, in.Id, dict)
}

// Delete 
func (a *{{.Name}}) Delete(ctx context.Context, in *model.{{.TitleName}}DeleteRequest) error  {
	return a.i{{.TitleName}}.Delete(ctx,in.Id)
}

// List 
func (a *{{.Name}}) List(ctx context.Context, in *model.{{.TitleName}}ListRequest) (*model.{{.TitleName}}ListResponse, error)  {
	var (
		err error
		total int
		list []*entity.{{.TitleName}} 
		out = &model.{{.TitleName}}ListResponse{}
	)

	if total, list, err = a.i{{.TitleName}}.List(ctx,in); err != nil {
		return nil, err
	}
	
	out.Total = total
	out.List = model.{{.TitleName}}sEntityToDto(list)

	return out, nil
}

// Find 
func (a *{{.Name}}) Find(ctx context.Context, in *model.{{.TitleName}}InfoRequest) (*model.{{.TitleName}}Info, error)  {
	var (
		err error
		data *entity.{{.TitleName}} 
		out = &model.{{.TitleName}}Info{}
	)

	if data, err = a.i{{.TitleName}}.Find(ctx,in); err != nil {
		return nil, err
	}
	
	out = model.{{.TitleName}}EntityToDto(data)
	return out, nil
}

// build{{.TitleName}} build entity
func build{{.TitleName}}(in *model.{{.TitleName}}CreateRequest) *entity.{{.TitleName}} {
	// todo: check the entity is required
	return &entity.{{.TitleName}}{
		{{range $v :=.Fields}}
			{{if eq .Json $CreatedAt}}
				{{.Name}}:time.Now().Unix(),
			{{else if eq .Json $UpdatedAt}}
				{{.Name}}:time.Now().Unix(),
			{{else}}
				{{if ne .Name $ID}}{{.Name}}: {{if eq .Parameter $true}} {{if ne .Required $true}}in.{{.Name}},{{else}}in.{{.Name}},{{end}}{{else}}{{if eq .Type $string}}"",{{else}}0,{{end}}{{end}}{{end}}
			{{end}}
		{{end}}
	} 
}
`
