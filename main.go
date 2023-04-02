package main

import (
	"log"
	"net/http"
	"server"
)

type StoreData struct {}

func (s *StoreData) ConvertCurrencies(inValue string) float64 {
	return 123.0
}

func (s *StoreData) StoreDataCurrency(inValue string) {}

func main() {
	port := ":5000"
	server := &server.CurrencyServer{}
	server.SetStorage(&StoreData{})
	if err := http.ListenAndServe(port, server); err != nil {
		log.Fatalf("unable to listen on port %s %v", port[1:], err)
	}
}