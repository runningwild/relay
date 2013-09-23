package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/smtp"
)

var srcUser = flag.String("src-user", "mistcakethegame", "username of sender")
var srcPassword = flag.String("src-password", "", "password of sender")
var srcSmtpAddr = flag.String("src-addr", "smtp.gmail.com", "smtp server of sender")
var dstSmtpAddrAndPort = flag.String("dst-smtp", "smtp.gmail.com:587", "smtp server and port of destination")
var sender = flag.String("sender", "mistcakethegame@gmail.com", "sender email address")
var receiver = flag.String("receiver", "runningwild@gmail.com", "recipient email address")

type Server struct {
	auth smtp.Auth
}

func (s *Server) ServeHTTP(response http.ResponseWriter, req *http.Request) {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "To: %s\r\nSubject: BugReport\r\n\r\n", *receiver)
	_, err := io.Copy(&buf, req.Body)
	err = smtp.SendMail(*dstSmtpAddrAndPort, s.auth, *sender, []string{*receiver}, buf.Bytes())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func main() {
	flag.Parse()
	auth := smtp.PlainAuth("", *srcUser, *srcPassword, *srcSmtpAddr)
	if auth == nil {
		fmt.Printf("Error: Failed to authorize/\n")
		return
	}
	http.Handle("/relay", &Server{auth: auth})
	http.ListenAndServe(":8080", nil)
}
