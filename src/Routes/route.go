package Routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin_sample/src/Handlers"
)

type Route struct {
	MethodType  MethodType
	Path        string
	HandlerFunc gin.HandlerFunc
}

func NewList() []Route {
	articleRoute := Route{MethodType: GET, Path: "/article", HandlerFunc: Handlers.GetAllArticles()}
	articleByIdRoute := Route{MethodType: GET, Path: "/article/:id", HandlerFunc: Handlers.GetArticleByID()}
	routes := []Route{articleRoute, articleByIdRoute}
	return routes
}

func ServeRoute(route Route, engine *gin.Engine) {
	switch route.MethodType {
	case GET:
		engine.GET(route.Path, route.HandlerFunc)
	default:
		log.Fatal("Method not defined")
	}
}
