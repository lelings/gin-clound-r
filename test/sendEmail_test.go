package test

import (
	"fmt"
	"log"
	"net/smtp"
	"testing"

	"example.com/m/v2/util"
	"github.com/jordan-wright/email"
)

func TestSend(t *testing.T) {
	e := email.NewEmail()
	e.From = "lelings <1132235296@qq.com>"
	e.To = []string{"2453498093@qq.com"}
	e.Subject = "验证码发送"
	e.Text = []byte("您的验证码为：")
	e.HTML = []byte(fmt.Sprintf("<p1>%s<p1>", util.GetCode()))
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "1132235296@qq.com", "qphacvqsiaiwgdid", "smtp.qq.com"))
	if err != nil {
		log.Fatal("error")
	}
}
