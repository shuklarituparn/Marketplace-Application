package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegistrationHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "registration.html", gin.H{})
}

func ProcessRegistrationHandler(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, ctx.Request.Body)
}
