package lib

import (
	"fmt"
	"log"
	"net/smtp"
	"time"

	"example.com/m/v2/util"
	"github.com/go-redis/redis"
	"github.com/jordan-wright/email"
)

var RDB *redis.Client

func InitRdb() {
	var rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", SERVICE_CONFIG.Redis_host, SERVICE_CONFIG.Redis_port),
		Password: "",
		DB:       0,
	})
	RDB = rdb
}

func SendEmail(em string) {
	e := email.NewEmail()
	var code = util.GetCode()
	e.From = fmt.Sprintf("lelings <%s>", SERVICE_CONFIG.Email_name)
	e.To = []string{em}
	e.Subject = "验证码发送"
	e.Text = []byte("您的验证码为：")
	e.HTML = []byte(fmt.Sprintf("<p1>%s<p1>", code))
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", SERVICE_CONFIG.Email_name, SERVICE_CONFIG.Email_password, "smtp.qq.com"))
	if err != nil {
		log.Fatal("error")
	}
	RDB.Set("code", code, time.Minute*10)
}

func ParseCode(code string) bool {
	rcode := RDB.Get("code")
	fmt.Println(rcode.String())
	return rcode.String()[10:] == code
}
