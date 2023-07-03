package postgres

import (
	"context"
	"gorm.io/gorm"
	"hello/errors"
	"hello/model"
	"hello/model/entity"
)

var Person = &person{}

type person struct{}

// Create
func (a *person) Create(ctx context.Context, m *entity.Person) (int64, error) {
	err := GetDB(ctx).Create(m).Error
	return m.Id, err
}

// Find detail
func (a *person) Find(ctx context.Context, in *model.PersonInfoRequest) (*entity.Person, error) {
	e := &entity.Person{}

	q := GetDB(ctx).Model(&entity.Person{})

	if in.Id == 0 {
		return e, errors.New("condition illegal")
	}
	err := q.First(&e).Error
	return e, err
}

// Update
func (a *person) Update(ctx context.Context, id int64, dict map[string]interface{}) error {
	return GetDB(ctx).Model(&entity.Person{}).Where("id = ?", id).Updates(dict).Error
}

// Delete
func (a *person) Delete(ctx context.Context, id int64) error {
	return GetDB(ctx).Delete(&entity.Person{}, id).Error
}

// List query list
func (a *person) List(ctx context.Context, in *model.PersonListRequest) (int, []*entity.Person, error) {
	var (
		q       = GetDB(ctx).Model(&entity.Person{})
		err     error
		total   int64
		persons []*entity.Person
	)

	if in.Arm != nil {

		q = q.Where("arm like ?", in.Arm)

	}

	if in.UpdatedAt != nil {

		q = q.Where("updated_at = ?", in.UpdatedAt)

	}

	if err = q.Count(&total).Error; err != nil {
		return 0, nil, err
	}
	if err = q.Limit(in.Size).Offset((in.Index - 1) * in.Size).Find(&persons).Error; err != nil {
		return 0, nil, err
	}
	return int(total), persons, nil
}

// ExecTransaction execute database transaction
func (a *person) ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error {
	return GetDB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ContextTxKey, tx)
		return callback(ctx)
	})
}
