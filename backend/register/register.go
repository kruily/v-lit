package register

import (
	"v-lit-backend/controllers/user_controller"

	"github.com/kruily/gofastcrud/core/crud"
	"github.com/kruily/gofastcrud/core/crud/module"
	"github.com/kruily/gofastcrud/core/server"
	"gorm.io/gorm"
)

func Register(srv *server.Server) {
	factory := module.CRUD_MODULE.GetService("FACTORY").(*crud.ControllerFactory)

	factory.RegisterBatchCustomMap(srv, map[string]func(*gorm.DB) crud.ICrudController[crud.ICrudEntity]{
		"user": user_controller.NewUserController,
	})
}
