package handlers

import (
	gomail "gopkg.in/gomail.v2"
)
//TODO Sacuvati mail i pass u env
func SendMail(to string, subject string, body string) (bool, error) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "teamthreemail@gmail.com")
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body)
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "teamthreemail@gmail.com", "tgqmnactjkjrpkts")
	err := dialer.DialAndSend(msg)
	if err != nil {
		return false, err
	}
	return true, nil
}
