package routes

import "github.com/gin-gonic/gin"

type Router struct {
	routerGroup *gin.RouterGroup
}

const version = "v1"

func NewRouter(router *gin.Engine) *Router {
	routerGroup := router.Group(version)
	return &Router{routerGroup: routerGroup}
}

func (r *Router) RegisterRoutes() {
	r.setHealthRoutes()
	r.setMergeRoutes()
}
