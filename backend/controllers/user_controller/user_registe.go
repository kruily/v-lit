package user_controller

import (
	"time"
	"v-lit-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/kruily/gofastcrud/pkg/errors"
	"github.com/kruily/gofastcrud/pkg/validator"
)

func (c *UserController) Registe(ctx *gin.Context) (interface{}, error) {
	var request models.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}
	if err := validator.Validate(request); err != nil {
		return nil, errors.New(errors.ErrInvalidParam, err.Error())
	}
	user := models.User{}
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
	if err := c.Repository.Create(ctx, &user); err != nil {
		return nil, errors.New(errors.ErrDatabase, err.Error())
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
