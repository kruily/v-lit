package user_controller

import (
	"net/http"
	"v-lit-backend/models"
	"v-lit-backend/services/user_service"

	"github.com/kruily/gofastcrud/core/crud"
	"github.com/kruily/gofastcrud/core/crud/module"
	"github.com/kruily/gofastcrud/core/crud/types"
	"github.com/kruily/gofastcrud/fast_jwt"
	"gorm.io/gorm"
)

type UserController struct {
	*crud.CrudController[models.User]
	userSrv *user_service.UserService

	jwtmaker *fast_jwt.JWTMaker
}

func NewUserController(db *gorm.DB) crud.ICrudController[crud.ICrudEntity] {
	controller := &UserController{
		CrudController: crud.NewCrudController(db, models.User{}),

		jwtmaker: module.CRUD_MODULE.GetService(module.JwtService).(*fast_jwt.JWTMaker),
	}

	controller.userSrv = user_service.NewUserService(controller.Repository)

	controller.AddRoutes([]types.APIRoute{
		{
			Method:      http.MethodPost,
			Path:        "/register",
			Handler:     controller.Registe,
			Tags:        []string{controller.GetEntityName()},
			Summary:     "用户注册",
			Description: "用户注册",
			Request:     models.UserRegisterRequest{},
			Response:    models.UserLoginedResponse{},
		},
		{
			Method:      http.MethodPost,
			Path:        "/login",
			Handler:     controller.Login,
			Tags:        []string{controller.GetEntityName()},
			Summary:     "用户登录",
			Description: "用户登录",
			Request:     models.UserLoginRequest{},
			Response:    models.UserLoginedResponse{},
		},
	})

	return controller
}
