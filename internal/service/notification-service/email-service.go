package notificationservice

import (
	"observer/internal/logger"
	"os"

	"gopkg.in/gomail.v2"
)

func (wns *NotificationService) SendEmailNotification(text string, email string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", "just.sparkless@gmail.com")
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", "Observer")
	mail.SetBody("text/plain", text)

	d := gomail.NewDialer("smtp.gmail.com", 465, "just.sparkless@gmail.com", os.Getenv("password"))
	if err := d.DialAndSend(mail); err != nil {
		logger.Error("error sending notification via email: " + err.Error())
		return err
	}

	return nil
}
