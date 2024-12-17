package user_controller

import (
	"net/http"
	"v-lit-backend/models"

	"github.com/kruily/gofastcrud/core/crud"
	"github.com/kruily/gofastcrud/core/crud/types"
	"gorm.io/gorm"
)

type UserController struct {
	*crud.CrudController[models.User]
}

func NewUserController(db *gorm.DB) crud.ICrudController[crud.ICrudEntity] {
	controller := &UserController{
		CrudController: crud.NewCrudController(db, models.User{}),
	}

	controller.AddRoute(types.APIRoute{
		Method:      http.MethodPost,
		Path:        "/register",
		Handler:     controller.Registe,
		Tags:        []string{controller.GetEntityName()},
		Summary:     "用户注册",
		Description: "用户注册",
		Request:     nil,
		Response:    nil,
	})

	return controller
}
