package util

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"strconv"
)


func SendMail(mailTo []string, subject string, body string) error {



	mailConn := map[string]string{
		"user": "xxxx@gmail.com",
		"pass": "xxxx",
		"host": "smtp.gmail.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From",  m.FormatAddress(mailConn["user"], "xxxx")) //这种方式可以添加别名，即“XX官方”
	//说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}

func SendEmail(email, code string) {
	//定义收件人
	mailTo := []string{
		email,
	}
	//邮件主题为"Hello"
	subject := "a demo"
	// 邮件正文
	body := fmt.Sprintf("尊敬的客户，您好:"+"<div>"+
		"您的验证码是 %s,10分钟内有效，请勿向他人泄露验证码。",code)

	err := SendMail(mailTo, subject, body)
	if err != nil {
		log.Println("err:",err)
		fmt.Println("send fail")
		return
	}

	fmt.Println("send successfully")

}