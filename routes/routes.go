package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"swaggo-handson/controller"
)

func InitHttpRouter(g *gin.Engine, ctx context.Context) {
	g.GET("health-check", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{"status": "OK"})
	})

	userController := controller.InitHttpUserController(ctx)
	user := g.Group("user")
	user.Use()
	{
		user.POST("", userController.Create)
		user.PATCH(":id", userController.Update)
		user.DELETE(":id", userController.Delete)
		user.GET("", userController.GetAll)
		user.GET(":id", userController.GetByID)
	}

}
