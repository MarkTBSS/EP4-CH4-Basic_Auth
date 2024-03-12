package main

import (
	"log"
	"net/http"
)

func main() {
	// Assume req is an HTTP request object obtained from a web server
	request, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	request.SetBasicAuth("username", "password")
	username, password, ok := request.BasicAuth()
	log.Println("auth:", username, password, ok)
}
