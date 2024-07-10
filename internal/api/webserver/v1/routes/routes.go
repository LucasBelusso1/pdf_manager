package routes

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	routerGroup *gin.RouterGroup
}

const version = "/v1"

func NewRouter(routerGroup *gin.RouterGroup) *Router {
	return &Router{routerGroup: routerGroup.Group(version)}
}

func (r *Router) RegisterRoutes() {
	r.setHealthRoutes()
	r.setMergeRoutes()
	r.setCountPagesRoutes()
}
