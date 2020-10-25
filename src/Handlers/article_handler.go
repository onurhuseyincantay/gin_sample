package Handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin_sample/src/Models"
)

func GetAllArticles() gin.HandlerFunc {
	return func(c *gin.Context) {
		articles := Models.GetAllArticles()
		c.JSON(http.StatusOK, articles)
	}
}

func GetArticleByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			ErrorHandler(c, http.StatusBadRequest, err)
		} else {
			article, err := Models.GetArticleByID(id)
			if err != nil {
				ErrorHandler(c, http.StatusBadRequest, err)
			} else {
				c.JSON(http.StatusOK, article)
			}
		}
	}
}
