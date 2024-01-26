package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"swaggo-handson/routes"
)

func MainHttpHandler(ctx context.Context) {
	g := gin.Default()
	routes.InitHttpRouter(g, ctx)

	_ = g.Run(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")))
}
