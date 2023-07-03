package store

import (
	"context"
	"hello/model"
	"hello/model/entity"
)

type IManager interface {
	// Create
	Create(ctx context.Context, e *entity.Manager) (int64, error)
	// Find
	Find(ctx context.Context, in *model.ManagerInfoRequest) (*entity.Manager, error)
	// Update
	Update(ctx context.Context, id int64, updates map[string]interface{}) error
	// Delete
	Delete(ctx context.Context, id int64) error
	// List
	List(ctx context.Context, in *model.ManagerListRequest) (int, []*entity.Manager, error)
	// ExecTransaction
	ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error
}
