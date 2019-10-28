package utils

import (
	"net/smtp"
	"strings"
	"../config"
)

var sys_email string
var sys_email_pass string
var sys_email_host string
func init(){
	sys_email = config.GetGlobalStringValue("sys_email","")
	sys_email_pass = config.GetGlobalStringValue("sys_email_pass","")
	sys_email_host = config.GetGlobalStringValue("sys_email_host","")
}



func SendToMail(to []string, subject, body string) error {
	hp := strings.Split(sys_email_host, ":")
	auth := smtp.PlainAuth("", sys_email, sys_email_pass, hp[0])
	var content_type string
	content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	toString := strings.Join(to,";")
	msg := []byte("To: " + toString + "\r\nFrom: " + sys_email + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(sys_email_host, auth, sys_email, to, msg)
	return err
}
