package register

import (
	"v-lit-backend/controllers/user_controller"
	"v-lit-backend/models/model_story"

	"github.com/kruily/gofastcrud/core/crud"
	"github.com/kruily/gofastcrud/core/crud/module"
	"github.com/kruily/gofastcrud/core/server"
	"gorm.io/gorm"
)

func Register(srv *server.Server) {
	factory := module.CRUD_MODULE.GetService(module.FactoryService).(*crud.ControllerFactory)

	factory.RegisterBatch(srv, model_story.Story{}, model_story.Graph{},
		model_story.Node{}, model_story.Volume{}, model_story.Chapter{})

	factory.RegisterBatchCustomMap(srv, map[string]func(*gorm.DB) crud.ICrudController[crud.ICrudEntity]{
		"user": user_controller.NewUserController,
	})
}

// story/:story_id/volumes/:volume_id/chapters/:chapter_id
