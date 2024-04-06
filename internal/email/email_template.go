package email

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/matcornic/hermes/v2"
)

func WelcomeTempGenerator(email, password string) string {

	link := os.Getenv("EMAIL_URL")
	extradetails := fmt.Sprintf("ваш email: %s и вот ваш пароль: %s", email, password)

	h := hermes.Hermes{

		Theme: &hermes.Default{},
		Product: hermes.Product{

			Name:      "VK Маркетплейс",
			Link:      link,
			Logo:      "https://iili.io/JOGASzF.md.png",
			Copyright: "© VK Маркетплейс. Все права защищены",
		},
	}

	emailToGenerate := hermes.Email{
		Body: hermes.Body{
			Name: email,
			Intros: []string{
				"Добро пожаловать в VK Маркетплейс, мы очень рады, что вы с нами.",
				extradetails,
			},
			Actions: []hermes.Action{
				{
					Instructions: "Чтобы начать пользоваться нашим сервисом, пожалуйста, нажмите здесь:",
					Button: hermes.Button{
						Color: "#0077FF",
						Text:  "VK Маркетплейс",
						Link:  link,
					},
				},
			},
			Outros: []string{
				"Нужна помощь или у вас есть вопросы? Просто ответьте на это письмо, мы будем рады помочь.",
			},
		},
	}

	emailBody, err := h.GenerateHTML(emailToGenerate)
	if err != nil {
		panic(err)
	}

	_, err = h.GeneratePlainText(emailToGenerate)
	if err != nil {
		panic(err)
	}

	filename := fmt.Sprintf("%s_w.html", email)
	filePath := filepath.Join("./internal/email/templates", filename)

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	completeFilename := fmt.Sprintf("./internal/email/templates" + "/" + filename)
	log.Println(completeFilename)
	err = os.WriteFile(completeFilename, []byte(emailBody), 0644)
	if err != nil {
		panic(err)
	}
	return completeFilename
}
