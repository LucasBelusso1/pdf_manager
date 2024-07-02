package routes

import "github.com/LucasBelusso1/pdf_manager/internal/api/webserver/v1/handlers"

func (r *Router) setHealthRoutes() {
	r.routerGroup.POST("/health", handlers.Health)
}
