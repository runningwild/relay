package main

import (
	"flag"
	"fmt"
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
	fmt.Printf("Sending...")
	err := smtp.SendMail(*dstSmtpAddrAndPort, s.auth, *sender, []string{*receiver}, []byte("thundermachine"))
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
	http.ListenAndServe(":8080", &Server{auth: auth})
}
