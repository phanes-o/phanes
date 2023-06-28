package template

func init() {
	register(StorePostgresTemplate, postgres)
}

var postgres = `
{{$true := "true"}}
{{$string := "string"}}

package postgres

import (
	"context"
	"gorm.io/gorm"
	"{{.ProjectName}}/errors"
	"{{.ProjectName}}/model"
	"{{.ProjectName}}/model/entity"
)

var {{.TitleName}} = &{{.Name}}{}

type {{.Name}} struct{}

// Create 
func (a *{{.Name}}) Create(ctx context.Context, m *entity.{{.TitleName}}) (int64, error) {
	err := GetDB(ctx).Create(m).Error
	return m.Id, err
}

// Find detail
func (a *{{.Name}}) Find(ctx context.Context, in *model.{{.TitleName}}InfoRequest ) (*entity.{{.TitleName}}, error ){
	e := &entity.{{.TitleName}}{}

	q := GetDB(ctx).Model(&entity.{{.TitleName}}{})

	if in.Id > 0 {
		err := q.First(&e, in.Id).Error
		return e, err
	}

	count := 0 
	{{range $v := .Fields}}
		{{if eq .Parameter $true}}
			{{if ne .Required $true}}
			if in.{{.Name}} != nil {
				{{if eq $string .Type}}
					q = q.Where("{{.Json}} like ?", in.{{.Name}}) 
				{{else}}
					q = q.Where("{{.Json}} = ?", in.{{.Name}}) 
				{{end}}
				count++
			}
			{{end}}
		{{end}}
	{{end}}

	if count == 0 {
		return e, errors.New("condition illegal")
	}

	err := q.First(&e).Error
	return e, err
}

// Update 
func (a *{{.Name}}) Update(ctx context.Context, id int64, dict map[string]interface{}) error {
	return GetDB(ctx).Model(&entity.{{.TitleName}}{}).Where("id = ?", id).Updates(dict).Error
}

// Delete 
func (a *{{.Name}}) Delete(ctx context.Context,id int64) error {
	return GetDB(ctx).Delete(&entity.{{.TitleName}}{}, id).Error
}

// List query list
func (a *{{.Name}}) List(ctx context.Context,in *model.{{.TitleName}}ListRequest) (int, []*entity.{{.TitleName}}, error) {
	var (
		q        = GetDB(ctx).Model(&entity.{{.TitleName}}{})
		err      error
		total    int64
		{{.Name}}s []*entity.{{.TitleName}}
	)

	{{range $v := .Fields}}
		{{if eq .Parameter $true}}
			{{if ne .Required $true}}
			if in.{{.Name}} != nil {
				{{if eq $string .Type}}
					q = q.Where("{{.Json}} like ?", in.{{.Name}}) 
				{{else}}
					q = q.Where("{{.Json}} = ?", in.{{.Name}}) 
				{{end}}
				
			}
			{{end}}
		{{end}}
	{{end}}

	if err = q.Count(&total).Error; err != nil {
		return 0, nil, err
	}
	if err = q.Limit(in.Size).Offset((in.Index - 1) * in.Size).Find(&{{.Name}}s).Error; err != nil {
		return 0, nil, err
	}
	return int(total), {{.Name}}s, nil
}

// ExecTransaction execute database transaction
func (a *{{.Name}}) ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error {
	return GetDB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, DBCONTEXTKEY, tx)
		return callback(ctx)
	})
}
`
