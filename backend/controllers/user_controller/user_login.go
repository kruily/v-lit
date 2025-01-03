package user_controller

import (
	"time"
	"v-lit-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/kruily/gofastcrud/pkg/errors"
	"github.com/kruily/gofastcrud/pkg/validator"
)

func (c *UserController) Login(ctx *gin.Context) (interface{}, error) {
	var request models.UserLoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}
	if err := validator.Validate(request); err != nil {
		return nil, errors.New(errors.ErrInvalidParam, err.Error())
	}
	var user *models.User
	var err error
	if request.Username != "" {
		user, err = c.Repository.FindOne(ctx, "username = ?", request.Username)
	}
	if request.Phone != "" {
		user, err = c.Repository.FindOne(ctx, "phone = ?", request.Phone)
	}
	if request.Email != "" {
		user, err = c.Repository.FindOne(ctx, "email = ?", request.Email)
	}
	if err != nil {
		return nil, errors.New(errors.ErrUserNotFound, err.Error())
	}
	if user.Password != request.Password {
		return nil, errors.New(errors.ErrInvalidParam, "密码错误")
	}
	token, err := c.jwtmaker.CreateToken(uint(user.ID.ID()), user.Username, time.Hour*24)
	if err != nil {
		return nil, errors.New(errors.ErrInternal, err.Error())
	}
	return models.UserLoginedResponse{
		Token: token,
		User:  user,
	}, nil
}
