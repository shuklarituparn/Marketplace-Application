package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UnAuthHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusUnauthorized, "bez_auth.html", gin.H{})

}
