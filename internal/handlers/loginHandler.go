package handlers

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func LoginHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

func ProcessLoginHandler(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
		return
	}

	// Set the request body as the response body
	ctx.Data(http.StatusOK, "application/json", body)
}
