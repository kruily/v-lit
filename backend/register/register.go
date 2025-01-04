package register

import (
	"v-lit-backend/controllers/user_controller"
	"v-lit-backend/models/model_story"

	"github.com/kruily/gofastcrud/core/crud"
	"github.com/kruily/gofastcrud/core/server"
	"gorm.io/gorm"
)

func Register(factory *crud.ControllerFactory, srv *server.Server) {
	story := factory.Register(srv, model_story.Story{})

	graph := factory.RegisterWithFather(srv, story, model_story.Graph{})

	factory.RegisterWithFather(srv, graph, model_story.Node{})

	volume := factory.RegisterWithFather(srv, story, model_story.Volume{})

	factory.RegisterWithFather(srv, volume, model_story.Chapter{})

	factory.RegisterWithFather(srv, story, model_story.Card{})

	factory.RegisterBatchCustomMap(srv, map[string]func(*gorm.DB) crud.ICrudController[crud.ICrudEntity]{
		"user": user_controller.NewUserController,
	})
}

// story/:story_id/volumes/:volume_id/chapters/:chapter_id
