package bll

import (
	"context"

	"hello/model"
	"hello/model/entity"
	"hello/model/mapping"
	"hello/store"
	"hello/store/postgres"
	"time"
)

type manager struct {
	iManager store.IManager
}

var Manager = &manager{
	iManager: postgres.Manager,
}

func (a *manager) init() func() {
	return func() {}
}

// Create
func (a *manager) Create(ctx context.Context, in *model.ManagerCreateRequest) error {
	var (
		err error
	)
	c := buildManager(in)
	_, err = a.iManager.Create(ctx, c)
	return err
}

// Update
func (a *manager) Update(ctx context.Context, in *model.ManagerUpdateRequest) error {
	var (
		dict = make(map[string]interface{})
	)

	if in.Arm != nil {
		dict["arm"] = in.Arm
	}

	if in.UpdatedAt != nil {
		dict["updated_at"] = in.UpdatedAt
	}

	// do other update here
	updateAt := time.Now().Unix()
	in.UpdatedAt = &updateAt
	return a.iManager.Update(ctx, in.Id, dict)
}

// Delete
func (a *manager) Delete(ctx context.Context, in *model.ManagerDeleteRequest) error {
	return a.iManager.Delete(ctx, in.Id)
}

// List
func (a *manager) List(ctx context.Context, in *model.ManagerListRequest) (*model.ManagerListResponse, error) {
	var (
		err   error
		total int
		list  []*entity.Manager
		out   = &model.ManagerListResponse{}
	)

	if total, list, err = a.iManager.List(ctx, in); err != nil {
		return nil, err
	}

	out.Total = total
	out.List = mapping.ManagersEntityToDto(list)

	return out, nil
}

// Find
func (a *manager) Find(ctx context.Context, in *model.ManagerInfoRequest) (*model.ManagerInfo, error) {
	var (
		err  error
		data *entity.Manager
		out  = &model.ManagerInfo{}
	)

	if data, err = a.iManager.Find(ctx, in); err != nil {
		return nil, err
	}

	out = mapping.ManagerEntityToDto(data)
	return out, nil
}

// buildManager build entity
func buildManager(in *model.ManagerCreateRequest) *entity.Manager {
	// todo: check the entity is required
	now := time.Now()
	return &entity.Manager{

		Name: in.Name,

		Arm: *in.Arm,

		Age: 0,

		Phones: in.Phones,

		CreatedAt: now,

		UpdatedAt: now,
	}
}
