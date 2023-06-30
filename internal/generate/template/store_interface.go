package template

func init() {
	register(StoreInterfaceTemplate, storeInterface)
}

var storeInterface = `
package store

import (
	"context"
	"{{.ProjectName}}/model"
	"{{.ProjectName}}/model/entity"
)

type I{{.StructName}} interface {
	// Create 
	Create(ctx context.Context, e *entity.{{.StructName}}) (int64, error)
	// Find 
	Find(ctx context.Context, in *model.{{.StructName}}InfoRequest) (*entity.{{.StructName}}, error)
	// Update 
	Update(ctx context.Context, id int64, updates map[string]interface{}) (error)
	// Delete 
	Delete(ctx context.Context, id int64) (error)
	// List 
	List(ctx context.Context, in *model.{{.StructName}}ListRequest) (int, []*entity.{{.StructName}}, error)
	// ExecTransaction 
	ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error 
}
`
