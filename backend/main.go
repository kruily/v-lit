package main

import (
	"log"
	"os"
	"v-lit-backend/models"
	"v-lit-backend/models/model_story"
	"v-lit-backend/register"

	"github.com/kruily/gofastcrud/config"
	"github.com/kruily/gofastcrud/core/app"
	"github.com/kruily/gofastcrud/core/crud/module"
	"github.com/kruily/gofastcrud/core/server"
	"github.com/kruily/gofastcrud/fast_jwt"
	"github.com/kruily/gofastcrud/utils"
)

func init() {
	// 获取项目根目录
	projectRoot, err := os.Getwd()
	if err != nil {
		log.Fatalf("无法获取项目根目录: %v", err)
	}
	// 加载环境变量
	utils.LoadEnv(projectRoot)
}

func main() {

	app := app.NewDefaultGoFastCrudApp()

	app.PublishVersion(server.V1)

	app.RegisterModels(
		&models.User{},
		&model_story.Story{},
		&model_story.Graph{},
		&model_story.Node{},
		&model_story.Volume{},
		&model_story.Chapter{},
	)

	configManager := module.CRUD_MODULE.GetService(module.ConfigService).(*config.ConfigManager)
	cfg := configManager.GetConfig()
	jwtmaker, err := fast_jwt.NewJWTMaker(cfg.JWT.SecretKey)
	if err != nil {
		log.Fatal("Failed to create jwt maker: ", err)
	}
	module.CRUD_MODULE.WithJwt(jwtmaker)

	app.RegisterControllers(register.Register)

	app.Start()
}
