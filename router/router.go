package router

import (
	"blog/api/service"
	"blog/controllers"

	"github.com/gin-gonic/gin"
)

// Init initializes router
func Init() {

	router := gin.Default()
	dependency := service.Init()

	articleController := controllers.NewArticle(dependency)

	router.POST("/articles", articleController.Create)
	router.GET("/articles", articleController.List)

	router.Run(":3000")
}
