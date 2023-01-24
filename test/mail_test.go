package test

import (
	"CloudStorage/core/define"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = define.EmailFrom
	e.To = []string{"1536941014@qq.com"}
	e.Subject = "Verification Code"
	e.HTML = []byte("Your Verification Code is below<h1>123456</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("1", define.EmailAccount, define.EmailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
