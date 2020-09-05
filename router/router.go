package router

import (
	"blog/api/service"
	"blog/controllers"

	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
)

// Init initializes router
func Init(db orm.Ormer) {

	router := gin.Default()
	dependency := service.Init(db)

	articleController := controllers.NewArticle(dependency)
	// routs of apis
	router.POST("/articles", articleController.Create)
	router.GET("/articles", articleController.List)
	router.GET("/articles/:id", articleController.Get)
	router.PATCH("/articles/:id", articleController.Update)

	router.Run(":3000")
}
