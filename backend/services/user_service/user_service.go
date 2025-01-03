package user_service

import (
	"context"
	"v-lit-backend/models"

	"github.com/kruily/gofastcrud/core/crud"
	"github.com/kruily/gofastcrud/pkg/errors"
)

type UserService struct {
	repo crud.IRepository[models.User]
}

func NewUserService(repo crud.IRepository[models.User]) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) RegisterUserModel(ctx context.Context, request *models.UserRegisterRequest) (*models.User, error) {
	user := models.User{
		BaseUUIDEntity: &crud.BaseUUIDEntity{},
	}
	if request.Username != "" {
		user.Username = request.Username
	}
	if request.Phone != "" {
		user.Phone = request.Phone
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Password != "" {
		user.Password = request.Password
	}
	if err := u.repo.Create(ctx, &user); err != nil {
		return nil, errors.New(errors.ErrDatabase, err.Error())
	}
	return &user, nil
}
