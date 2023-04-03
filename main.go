package main

import (
	"log"
	"net/http"
	"server"
)

const Port = ":5000"

type MainStore struct {
	Value float64 
	Symbol string
}

func (s *MainStore) ResponseAction(p server.Params) server.ResponseData {
	sign := server.CatSign(p.GetParams()[2])
	outValue := server.Calculate(p.GetParams()[0], p.GetParams()[3])
	rd := server.ResponseData{}
	rd.SetData(outValue, sign)
	return rd
}

func main() {
	server := server.NewCurrencyServer(&MainStore{})
	server.SetCurrencyServer(&MainStore{})
	if err := http.ListenAndServe(Port, server); err != nil {
		log.Fatalf("unable to listen on port %s %v", Port, err)
	}
}
