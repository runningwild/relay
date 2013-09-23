package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	if len(data) == 0 {
		fmt.Printf("No data.\n")
		return
	}
	err = smtp.SendMail(*dstSmtpAddrAndPort, s.auth, *sender, []string{*receiver}, data)
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
