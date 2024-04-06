package email

import (
	"encoding/json"
	"fmt"
	"github.com/resend/resend-go/v2"
	"github.com/shuklarituparn/marketplace/internal/consumer"
	"github.com/shuklarituparn/marketplace/internal/models"
	"log"
	"os"
)

func SendMail(Filepath string, To string, Subject string) {

	client := resend.NewClient(os.Getenv("RESEND_API_KEY"))
	emailToUse := fmt.Sprintf("Video Conversion Service <%s>", os.Getenv("EMAIL"))
	replyMail := os.Getenv("REPLY_MAIL")

	htmlcontent, err := os.ReadFile(Filepath)
	if err != nil {

		log.Println("Error opening the html content")
	}

	params := &resend.SendEmailRequest{
		From:    emailToUse,
		To:      []string{To},
		Html:    string(htmlcontent),
		Subject: Subject,
		ReplyTo: replyMail,
	}

	_, err = client.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func WelcomeEmailConsumer() {

	c, _ := consumer.NewConsumer("localhost:9092", "email_service")
	_ = c.Subscribe("welcomeMail", nil)

	defer consumer.Close(c)
	var EmailMessage models.RegisterUserModel

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		log.Printf("Received message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		err = json.Unmarshal(msg.Value, &EmailMessage)
		filepath := WelcomeTempGenerator(EmailMessage.Email, EmailMessage.Password)

		SendMail(filepath, EmailMessage.Email, "Welcome to VK-Marketplace")

		_, commitErr := c.CommitMessage(msg)
		if commitErr != nil {
			log.Printf("Failed to commit offset: %v", commitErr)
		}
	}

}
