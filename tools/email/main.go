package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
)

func main() {
	m := gomail.NewMessage()
	d, err := mailConfig(m)
	if err != nil {
		fmt.Printf("send mail fail.err :%+v\n", err)
		return
	}

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		panic(err)
	}
	fmt.Printf("send mail success.\n dial :%+v\n", d)
}

const (
	myEmail    = "7997@qq.com" // TODO fix it when use it
	otherEmail = "4957@qq.com" // TODO fix it when use it
	authCode   = "vrbprtlnubhpbbad"
)

func mailConfig(m *gomail.Message) (*gomail.Dialer, error) {
	body, err := getLocalHtml()
	if err != nil {
		return nil, err
	}
	//发送人
	m.SetHeader("From", myEmail)
	//接收人
	m.SetHeader("To", otherEmail)
	//抄送人
	m.SetAddressHeader("Cc", otherEmail, "小猪佩奇")
	//主题
	m.SetHeader("Subject", "小佩奇的一封信")
	//内容
	m.SetBody("text/html", body)
	//附件
	//m.Attach("./main.go")

	//拿到token，并进行连接,第4个参数是填授权码
	return gomail.NewDialer("smtp.qq.com", 587, myEmail, authCode), nil
}

func mailBody() string {
	return "<h1>新年快乐</h1>"
}

func getLocalHtml() (string, error) {
	b, err := os.ReadFile("D:\\code\\src\\codenote\\tools\\email\\bar.html")
	if err != nil {
		fmt.Printf("ReadFile fail, err:%+v", err)
		return "", err
	}
	return string(b), nil
}
