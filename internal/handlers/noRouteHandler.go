package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoRouteHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.html", gin.H{})

}
