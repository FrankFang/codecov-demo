package router

import (
	"mangosteen/config"
	"mangosteen/internal/controller"
	"mangosteen/internal/database"

	"mangosteen/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swag
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type ConFn = func() controller.Controller

func loadControllers() []ConFn {
	return [](ConFn){
		controller.NewValidationCodeController,
		controller.NewSessionController,
	}
}

func New() *gin.Engine {
	config.LoadAppConfig()
	r := gin.Default()
	docs.SwaggerInfo.Version = "1.0"

	database.Connect()

	api := r.Group("/api")
	r.GET("/ping", controller.Ping)

	for _, cf := range loadControllers() {
		cf().RegisterRoutes(api)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
