package repository

import (
	"context"
	"errors"
	"github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/model"
	"github.com/magomedcoder/ipmanager/pkg/gormutil"
	"gorm.io/gorm"
	"log"
)

type IUserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)

	GetUsers(ctx context.Context, arg ...func(*gorm.DB)) ([]*model.User, error)

	GetById(ctx context.Context, id int64) (*model.User, error)

	GetByUsername(username string) (*model.User, error)

	UpdatePasswordById(ctx context.Context, id int64, password string) (int64, error)
}

var _ IUserRepository = (*UserRepository)(nil)

type UserRepository struct {
	gormutil.Repo[model.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Repo: gormutil.NewRepo[model.User](db)}
}

func (u *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	if err := u.Repo.Create(ctx, user); err != nil {
		log.Printf("Не удалось создать пользователя: %s", err)
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUsers(ctx context.Context, arg ...func(*gorm.DB)) ([]*model.User, error) {
	users, err := u.FindAll(ctx, arg...)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) GetById(ctx context.Context, id int64) (*model.User, error) {
	user, err := u.Repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetByUsername(username string) (*model.User, error) {
	user, err := u.Repo.FindByWhere(context.TODO(), "username = ?", username)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Не удалось получить пользователя: %s", err)
		}
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) UpdatePasswordById(ctx context.Context, id int64, password string) (int64, error) {
	_id, err := u.Repo.UpdateById(ctx, id, map[string]interface{}{
		"password": password,
	})
	if err != nil {
		return _id, nil
	}

	return _id, nil
}
