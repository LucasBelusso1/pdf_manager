package routes

import "github.com/LucasBelusso1/pdf_manager/internal/api/webserver/v1/handlers"

func (r *Router) setMergeRoutes() {
	r.routerGroup.POST("/merge", handlers.Merge)
}
