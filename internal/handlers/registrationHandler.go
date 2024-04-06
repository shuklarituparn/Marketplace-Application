package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/marketplace/internal/models"
	"github.com/shuklarituparn/marketplace/internal/producer"
	"io"
	"net/http"
)

func RegistrationHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "registration.html", gin.H{})
}

func ProcessRegistrationHandler(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, ctx.Request.Body)
	body, _ := io.ReadAll(ctx.Request.Body)
	var registermailModel models.RegisterUserModel
	err := json.Unmarshal(body, &registermailModel)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	//TODO: Create a producer with the data
	p, err := producer.NewProducer("localhost:9092")
	err = producer.ProduceNewMessage(p, "welcomeMail", string(body))
	if err != nil {
		return
	}

}

//here create a kafka request to handle the userdata
