package main

import (
	"log"
	"os"
	"v-lit-backend/models"

	"github.com/kruily/gofastcrud/core/crud"
	"github.com/kruily/gofastcrud/core/crud/module"
	"github.com/kruily/gofastcrud/core/database"
	"github.com/kruily/gofastcrud/core/server"
	"github.com/kruily/gofastcrud/pkg/config"
	"github.com/kruily/gofastcrud/pkg/fast_jwt"
	"github.com/kruily/gofastcrud/pkg/logger"
	"github.com/kruily/gofastcrud/pkg/utils"

	"v-lit-backend/register"
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
	configManager := module.CRUD_MODULE.GetService(module.ConfigService).(*config.ConfigManager)
	configManager.LoadConfig()

	cfg := configManager.GetConfig()

	db := module.CRUD_MODULE.GetService(module.DatabaseService).(*database.Database)
	db.RegisterModels(
		&models.User{},
	)
	// 初始化数据库
	if err := db.Init(&cfg.Database); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	// 创建日志实例
	logService, err := logger.NewLogger(logger.Config{
		Level: logger.InfoLevel,
		FileConfig: &logger.FileConfig{
			Filename:   "logs/app.log",
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     7,
			Compress:   true,
		},
		ConsoleLevel: logger.DebugLevel,
	})
	if err != nil {
		log.Fatal("Failed to create logger: ", err)
	}
	defer logService.Close()

	// 注册jwt服务
	jwtmaker, err := fast_jwt.NewJWTMaker(cfg.JWT.SecretKey)
	if err != nil {
		log.Fatal("Failed to create jwt maker: ", err)
	}
	module.CRUD_MODULE.SetService(module.JwtService, jwtmaker)

	// 注册日志服务
	module.CRUD_MODULE.WithLogger(logService)
	// 创建服务实例
	srv := server.NewServer(cfg)
	// 注册服务
	module.CRUD_MODULE.WithServer(srv)

	srv.PublishVersion(server.V1)
	factory := crud.NewControllerFactory(db.DB())
	module.CRUD_MODULE.SetService(module.FactoryService, factory)
	// 注册控制器
	register.Register(srv)
	// 运行服务（包含优雅启停）
	if err := srv.Run(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
