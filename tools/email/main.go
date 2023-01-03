package main

import (
	"bytes"
	"fmt"
	"gopkg.in/gomail.v2"
	"text/template"
)

func main() {
	m := gomail.NewMessage()
	d := mailConfig(m)

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

func mailConfig(m *gomail.Message) *gomail.Dialer {
	//发送人
	m.SetHeader("From", myEmail)
	//接收人
	m.SetHeader("To", otherEmail)
	//抄送人
	m.SetAddressHeader("Cc", otherEmail, "小猪佩奇")
	//主题
	m.SetHeader("Subject", "小佩奇的一封信")
	//内容
	m.SetBody("text/html", getLocalHtml())
	//附件
	//m.Attach("./main.go")

	//拿到token，并进行连接,第4个参数是填授权码
	return gomail.NewDialer("smtp.qq.com", 587, myEmail, authCode)
}

func mailBody() string {
	return "<h1>新年快乐</h1>"
}

func getLocalHtml() string {
	files, err := template.ParseFiles("D:\\code\\src\\codenote\\tools\\email\\bar.html")
	if err != nil {
		fmt.Printf("ParseFiles fail, err:%+v", err)
		return "<h1>解析失败再接再厉</h1>"
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	err = files.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    "Puneet Singh",
		Message: "This is a test message in a HTML template",
	})
	if err != nil {
		fmt.Printf("execute fail, err:%+v", err)
		return "<h1>执行失败再接再厉</h1>"
	}
	return string(body.Bytes())
}
