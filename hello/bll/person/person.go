package bll

import (
	"context"

	"github.com/lib/pq"

	"hello/model"
	"hello/model/entity"
	"hello/model/mapping"
	"hello/store"
	"hello/store/postgres"
	"time"
)

type person struct {
	iPerson store.IPerson
}

var Person = &person{
	iPerson: postgres.Person,
}

func init() {
	Register(Person)
}

func (a *person) init() func() {
	return func() {}
}

// Create
func (a *person) Create(ctx context.Context, in *model.PersonCreateRequest) error {
	var (
		err error
	)
	c := buildPerson(in)
	_, err = a.iPerson.Create(ctx, c)
	return err
}

// Update
func (a *person) Update(ctx context.Context, in *model.PersonUpdateRequest) error {
	var (
		dict = make(map[string]interface{})
	)

	if in.Phones != nil {
		dict["phones"] = in.Phones
	}

	if in.CreatedAt != nil {
		dict["created_at"] = in.CreatedAt
	}

	if in.OrderTime != nil {
		dict["order_time"] = in.OrderTime
	}

	if in.UpdatedAt != nil {
		dict["updated_at"] = in.UpdatedAt
	}

	// do other update here
	updateAt := time.Now().Unix()
	in.UpdatedAt = &updateAt
	return a.iPerson.Update(ctx, in.Id, dict)
}

// Delete
func (a *person) Delete(ctx context.Context, in *model.PersonDeleteRequest) error {
	return a.iPerson.Delete(ctx, in.Id)
}

// List
func (a *person) List(ctx context.Context, in *model.PersonListRequest) (*model.PersonListResponse, error) {
	var (
		err   error
		total int
		list  []*entity.Person
		out   = &model.PersonListResponse{}
	)

	if total, list, err = a.iPerson.List(ctx, in); err != nil {
		return nil, err
	}

	out.Total = total
	out.List = mapping.PersonsEntityToDto(list)

	return out, nil
}

// Find
func (a *person) Find(ctx context.Context, in *model.PersonInfoRequest) (*model.PersonInfo, error) {
	var (
		err  error
		data *entity.Person
		out  = &model.PersonInfo{}
	)

	if data, err = a.iPerson.Find(ctx, in); err != nil {
		return nil, err
	}

	out = mapping.PersonEntityToDto(data)
	return out, nil
}

// buildPerson build entity
func buildPerson(in *model.PersonCreateRequest) *entity.Person {

	now := time.Now()

	ety := &entity.Person{

		Name: in.Name,

		Age: 0,

		CreatedAt: now,

		OrderTime: time.Unix(0, 0),

		UpdatedAt: now,
	}

	if len(in.Phones) != 0 {
		ety.Phones = in.Phones
	} else {
		ety.Phones = pq.StringArray{}
	}

	return ety
}
