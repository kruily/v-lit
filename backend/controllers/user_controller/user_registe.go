package user_controller

import (
	"time"
	"v-lit-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/kruily/gofastcrud/errors"
	"github.com/kruily/gofastcrud/validator"
)

func (c *UserController) Registe(ctx *gin.Context) (interface{}, error) {
	var request models.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}
	if err := validator.Validate(request); err != nil {
		return nil, errors.New(errors.ErrInvalidParam, err.Error())
	}

	user, err := c.userSrv.RegisterUserModel(ctx, &request)
	if err != nil {
		return nil, err
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
