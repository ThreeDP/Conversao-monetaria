package main

import (
	"log"
	"net/http"
	"server"
)

const Port = ":8007"

func main() {
	server := &server.CurrencyServer{}
	if err := http.ListenAndServe(Port, server); err != nil {
		log.Fatalf("unable to listen on port %s %v", Port, err)
	}
}
