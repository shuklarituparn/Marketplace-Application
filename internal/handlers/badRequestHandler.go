package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequestHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusInternalServerError, "500.html", gin.H{})
}
