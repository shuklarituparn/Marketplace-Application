package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterSuccess(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register_success.html", gin.H{})
}

func LoginSuccess(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login_success.html", gin.H{})
}
