package template

func init() {
	register(StoreMysqlTemplate, mysql)
}

var mysql = `
{{$string := "string"}}
{{$true := true}}

package mysql

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	logger "{{.Module}}/collector/logger"
	"{{.Module}}/config"
	"{{.Module}}/errors"
	"{{.Module}}/model"
	"{{.Module}}/model/entity"
)

var {{.StructName}} = &{{.CamelName}}{}

type {{.CamelName}} struct{}

func init() {
	Register({{.StructName}})
}


func (a *{{.CamelName}}) Init() {
	if config.Conf.AutoMigrate {
		p := &entity.{{.StructName}}{}
		if db.Migrator().HasTable(p) {
			logger.Debug("table already exist: ", zap.String("table", p.TableName()))
			return
		}
		if err := db.AutoMigrate(p); err != nil {
			logger.Error("filed to create table please check config or manually create", zap.String("table", p.TableName()), zap.String("err", err.Error()))
		} else {
			logger.Info("create table successfully", zap.String("table", p.TableName()))
		}
	}
}


// Create 
func (a *{{.CamelName}}) Create(ctx context.Context, m *entity.{{.StructName}}) (int64, error) {
	err := GetDB(ctx).Create(m).Error
	return m.Id, err
}

// Find detail
func (a *{{.CamelName}}) Find(ctx context.Context, in *model.{{.StructName}}InfoRequest ) (*entity.{{.StructName}}, error ){
	e := &entity.{{.StructName}}{}

	q := GetDB(ctx).Model(&entity.{{.StructName}}{})

	if in.Id == 0 {
		return e, errors.New("condition illegal")
	}
	err := q.First(&e).Error
	return e, err
}

// Update 
func (a *{{.CamelName}}) Update(ctx context.Context, id int64, dict map[string]interface{}) error {
	return GetDB(ctx).Model(&entity.{{.StructName}}{}).Where("id = ?", id).Updates(dict).Error
}

// Delete 
func (a *{{.CamelName}}) Delete(ctx context.Context,id int64) error {
	return GetDB(ctx).Delete(&entity.{{.StructName}}{}, id).Error
}

// List query list
func (a *{{.CamelName}}) List(ctx context.Context,in *model.{{.StructName}}ListRequest) (int, []*entity.{{.StructName}}, error) {
	var (
		q        = GetDB(ctx).Model(&entity.{{.StructName}}{})
		err      error
		total    int64
		{{.CamelName}}s []*entity.{{.StructName}}
	)

	{{range $v := .Fields}}
		{{if eq $v.Rule.Parameter $true}}
			{{if ne $v.Rule.Required $true}}
			if in.{{.Name}} != nil {
				{{if eq $string .Type}}
					q = q.Where("{{.SnakeName}} like ?", in.{{.Name}}) 
				{{else}}
					q = q.Where("{{.SnakeName}} = ?", in.{{.Name}}) 
				{{end}}
			}
			{{end}}
		{{end}}
	{{end}}

	if err = q.Count(&total).Error; err != nil {
		return 0, nil, err
	}
	if err = q.Limit(in.Size).Offset((in.Index - 1) * in.Size).Find(&{{.CamelName}}s).Error; err != nil {
		return 0, nil, err
	}
	return int(total), {{.CamelName}}s, nil
}

// ExecTransaction execute database transaction
func (a *{{.CamelName}}) ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error {
	return GetDB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ContextTxKey, tx)
		return callback(ctx)
	})
}
`
