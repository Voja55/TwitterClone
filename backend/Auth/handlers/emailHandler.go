package handlers

import (
	"context"
	"encoding/base64"
	"errors"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	gmail "google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func SendMail(to string, body string) (bool, error) {
	//m := gomail.NewMessage()
	//m.SetHeader("From", "teamthreemail@gmail.com")
	//m.SetHeader("To", to)
	//m.SetHeader("Subject", "CCode")
	//m.SetBody("text/plain", message)
	//d := gomail.NewDialer("smtp-relay.gmail.com", 587, "teamthreemail@gmail.com", "Team.Three3")
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	//err := d.DialAndSend(m)
	//if err != nil {
	//	return false, err
	//}
	//return true, nil
	var message gmail.Message
	emailTo := "To: " + to + "\r\n"
	subject := "Subject: " + "Test email" + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte (emailTo + subject + mime + "\n" + body)
	message.Raw = base64.URLEncoding.EncodeToString(msg)

	_, err := GmailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

var GmailService *gmail.Service

func OAuthGmailService() error {
   config := oauth2.Config{
      ClientID:     "380405157054-p75sicqi7e19vjusdf07u3od59j8tnnb.apps.googleusercontent.com",
      ClientSecret: "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",
      Endpoint:     google.Endpoint,
      RedirectURL:  "https://localhost",
   }

   token := oauth2.Token{
      AccessToken:  "ya29.a0AeTM1ifWZTODN8jCTMwEv8SfikIm9Rtyi0ZBMq3MQchpln0OVuJwyEfxXjtsjjhx8xhO4djCvl9agu75Jq_epjp3-IWVq6uB15zD04aWPkYAmOeFeGTV9J0k2jXTMeb1pktQRMTn4GtqtwvjMkj9Ww8evuM8aCgYKAdYSARASFQHWtWOmfcYR3kvdXJFx9I5pzl5zAQ0163",
      RefreshToken: "1//04oMV5DkUoVUVCgYIARAAGAQSNwF-L9Ir2GDLKq5bN8FtgaY5JoNfMJgNY-c2J5qrcqOLvyZITinKjCb0uGv3sfFMnf-5wIFe6KM",
      TokenType:    "Bearer",
      Expiry: time.Now(),
   }

   var tokenSource = config.TokenSource(context.Background(), &token)

   srv, err := gmail.NewService(context.Background(), option.WithTokenSource(tokenSource))
   if err != nil {
	  return err
   }

   GmailService = srv
   if GmailService != nil {
	return nil
   }
   return errors.New("Failed to initalize gmail service")
}
