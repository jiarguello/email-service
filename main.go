package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	gomail "gopkg.in/mail.v2"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	abc := gomail.NewMessage()
	emailAdress := goDotEnvVariable("EMAIL_ADDRESS")
	password := goDotEnvVariable("EMAIL_PASSWORD")

	fmt.Println(emailAdress)
	fmt.Println(password)

	abc.SetHeader("From", emailAdress)
	abc.SetHeader("To", "jhonatan.arguello@gmail.com")
	abc.SetHeader("Subject", "Test Email Subject")
	abc.SetBody("text/plain", "This is the Test Body")

	a := gomail.NewDialer("smtp.gmail.com", 587, emailAdress, password)

	if err := a.DialAndSend(abc); err != nil {
		fmt.Println(err)
		panic(err)
	}
}