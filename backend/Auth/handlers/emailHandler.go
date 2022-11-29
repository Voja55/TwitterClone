package handlers

import (
	"crypto/tls"
	gomail "gopkg.in/mail.v2"
)

func SendMail(to string, message string) (bool, error) {
	//from := "teamthreemail@gmail.com"
	//password := "Team.Three3"
	//host := "smtp.gmail.com"
	//port := "587"
	//encodeMsg := []byte(message)
	//auth := smtp.PlainAuth("", from, password, host)
	//err := smtp.SendMail(host + ":" + port, auth, from, to, encodeMsg)
	//if err != nil {
	//	return false, err
	//}
	//return true, nil
	m := gomail.NewMessage()
	m.SetHeader("From", "teamthreemail@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "CCode")
	m.SetBody("text/plain", message)
	d := gomail.NewDialer("smtp-relay.gmail.com", 587, "teamthreemail@gmail.com", "Team.Three3")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	if err != nil {
		return false, err
	}
	return true, nil
}


