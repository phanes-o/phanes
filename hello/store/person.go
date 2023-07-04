package store

import (
	"context"
	"hello/model"
	"hello/model/entity"
)

type IPerson interface {
	// Create
	Create(ctx context.Context, e *entity.Person) (int64, error)
	// Find
	Find(ctx context.Context, in *model.PersonInfoRequest) (*entity.Person, error)
	// Update
	Update(ctx context.Context, id int64, updates map[string]interface{}) error
	// Delete
	Delete(ctx context.Context, id int64) error
	// List
	List(ctx context.Context, in *model.PersonListRequest) (int, []*entity.Person, error)
	// ExecTransaction
	ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error
}
