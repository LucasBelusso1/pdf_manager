package webserver

import (
	"github.com/LucasBelusso1/pdf_manager/internal/api/webserver/v1/routes"
	"github.com/gin-gonic/gin"
)

func Start() {
	engine := gin.Default()

	router := routes.NewRouter(engine)
	router.RegisterRoutes()

	engine.Run(":8080")
}
