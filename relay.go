package main

import (
	"net/http"
	// "net/smtp"
	"fmt"
)

type Server struct{}

func (s *Server) ServeHTTP(response http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(response, "Fudgeall")
}

func main() {
	http.ListenAndServe(":8080")
}
