package repository

import (
	"context"
	"errors"
	"github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/model"
	"gorm.io/gorm"
	"log"
)

type IUserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)

	GetByUsername(username string) (*model.User, error)
}

var _ IUserRepository = (*UserRepository)(nil)

type UserRepository struct {
	Repo[model.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Repo: NewRepo[model.User](db)}
}

func (u *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	if err := u.Repo.Create(ctx, user); err != nil {
		log.Printf("Не удалось создать пользователя: %s", err)
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
