package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin_sample/src/Routes"
)

func Run() {
	r := gin.Default()
	routes := Routes.NewList()

	for _, route := range routes {
		Routes.ServeRoute(route, r)
	}

	r.Run()
}
