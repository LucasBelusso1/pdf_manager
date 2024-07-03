package webserver

import (
	"github.com/LucasBelusso1/pdf_manager/internal/api/webserver/v1/routes"
	"github.com/gin-gonic/gin"
)

const apiPrefix = "/api"

const multipartMemoryLimit_8MB = 8 << 10

func Start() {
	engine := gin.Default()
	engine.MaxMultipartMemory = multipartMemoryLimit_8MB
	defaultRouterGroup := engine.Group(apiPrefix)

	router := routes.NewRouter(defaultRouterGroup)
	router.RegisterRoutes()

	engine.Run(":8080")
}
