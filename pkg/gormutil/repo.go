package gormutil

import (
	"context"
	"gorm.io/gorm"
)

type ITable interface {
	TableName() string
}

type Repo[T ITable] struct {
	model T
	Db    *gorm.DB
}

func NewRepo[T ITable](db *gorm.DB) Repo[T] {
	return Repo[T]{Db: db}
}

func (r *Repo[T]) Model(ctx context.Context) *gorm.DB {
	return r.Db.WithContext(ctx).Model(r.model)
}

func (r *Repo[T]) Create(ctx context.Context, data *T) error {
	return r.Db.WithContext(ctx).Create(data).Error
}

func (r *Repo[T]) FindAll(ctx context.Context, arg ...func(*gorm.DB)) ([]*T, error) {
	bd := r.Model(ctx)
	for _, fn := range arg {
		fn(bd)
	}
	var items []*T
	if err := bd.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *Repo[T]) FindById(ctx context.Context, id int64) (*T, error) {
	var item *T
	if err := r.Db.WithContext(ctx).First(&item, id).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (r *Repo[T]) FindByWhere(ctx context.Context, where string, args ...any) (*T, error) {
	var item *T

	if err := r.Db.WithContext(ctx).
		Where(where, args...).
		First(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (r *Repo[T]) QueryExist(ctx context.Context, where string, args ...any) (bool, error) {
	var count int64
	if err := r.Model(ctx).Select("1").Where(where, args...).Limit(1).Scan(&count).Error; err != nil {
		return false, err
	}

	return count == 1, nil
}

func (r *Repo[T]) DeleteWhere(ctx context.Context, where string, args ...interface{}) error {
	var item T
	res := r.Db.WithContext(ctx).Where(where, args...).Delete(&item)
	return res.Error
}
