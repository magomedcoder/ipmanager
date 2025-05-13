package repository

import (
	"context"
	"github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/model"
	"github.com/magomedcoder/ipmanager/pkg/gormutil"
	"gorm.io/gorm"
	"log"
)

type IUserSessionRepository interface {
	Create(ctx context.Context, userSession *model.UserSession) (*model.UserSession, error)

	CheckExpiredToken(ctx context.Context, accessToken string) (bool, error)

	DeleteByToken(ctx context.Context, accessToken string) error
}

var _ IUserSessionRepository = (*UserSessionRepository)(nil)

type UserSessionRepository struct {
	gormutil.Repo[model.UserSession]
}

func NewUserSessionRepository(db *gorm.DB) *UserSessionRepository {
	return &UserSessionRepository{Repo: gormutil.NewRepo[model.UserSession](db)}
}

func (u *UserSessionRepository) Create(ctx context.Context, userSession *model.UserSession) (*model.UserSession, error) {
	if err := u.Repo.Create(ctx, userSession); err != nil {
		log.Printf("Не удалось создать: %s", err)
		return nil, err
	}

	return userSession, nil
}

func (u *UserSessionRepository) CheckExpiredToken(ctx context.Context, accessToken string) (bool, error) {
	exists, err := u.Repo.QueryExist(ctx, "access_token = ? AND TO_TIMESTAMP(expires_at) > NOW()", accessToken)
	if err != nil {
		return false, err
	}

	if !exists {
		if delErr := u.DeleteByToken(ctx, accessToken); delErr != nil {
			return false, delErr
		}
	}

	return exists, nil
}

func (u *UserSessionRepository) DeleteByToken(ctx context.Context, accessToken string) error {
	return u.Repo.DeleteWhere(ctx, "access_token = ?", accessToken)
}
