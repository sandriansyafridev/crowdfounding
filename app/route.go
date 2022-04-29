package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	r := gin.New()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "data")
	})

	return r
}
