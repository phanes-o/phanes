package postgres

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	log "hello/collector/logger"
	"hello/config"
	"hello/errors"
	"hello/model"
	"hello/model/entity"
)

var Manager = &manager{}

type manager struct{}

func init() {
	Register(Manager)
}

func (a *manager) Init() {
	if config.Conf.AutoMigrate {
		p := &entity.Manager{}
		if db.Migrator().HasTable(p) {
			log.Debug("table already exist: ", zap.String("table", p.TableName()))
			return
		}
		if err := db.AutoMigrate(p); err != nil {
			log.Error("filed to create table please check config or manually create", zap.String("table", p.TableName()), zap.String("err", err.Error()))
		} else {
			log.Info("create table successfully", zap.String("table", p.TableName()))
		}
	}
}

// Create
func (a *manager) Create(ctx context.Context, m *entity.Manager) (int64, error) {
	err := GetDB(ctx).Create(m).Error
	return m.Id, err
}

// Find detail
func (a *manager) Find(ctx context.Context, in *model.ManagerInfoRequest) (*entity.Manager, error) {
	e := &entity.Manager{}

	q := GetDB(ctx).Model(&entity.Manager{})

	if in.Id == 0 {
		return e, errors.New("condition illegal")
	}
	err := q.First(&e).Error
	return e, err
}

// Update
func (a *manager) Update(ctx context.Context, id int64, dict map[string]interface{}) error {
	return GetDB(ctx).Model(&entity.Manager{}).Where("id = ?", id).Updates(dict).Error
}

// Delete
func (a *manager) Delete(ctx context.Context, id int64) error {
	return GetDB(ctx).Delete(&entity.Manager{}, id).Error
}

// List query list
func (a *manager) List(ctx context.Context, in *model.ManagerListRequest) (int, []*entity.Manager, error) {
	var (
		q        = GetDB(ctx).Model(&entity.Manager{})
		err      error
		total    int64
		managers []*entity.Manager
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
	if err = q.Limit(in.Size).Offset((in.Index - 1) * in.Size).Find(&managers).Error; err != nil {
		return 0, nil, err
	}
	return int(total), managers, nil
}

// ExecTransaction execute database transaction
func (a *manager) ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error {
	return GetDB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ContextTxKey, tx)
		return callback(ctx)
	})
}
