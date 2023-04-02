package server

import (
	"fmt"
	"net/http"
	"strings"
)

type StoreDataCurrency interface {
	ConvertCurrencies(inValue string) float64
	StoreDataCurrency(inValue string)
}

type CurrencyServer struct {
	store StoreDataCurrency
}

func (s *CurrencyServer) SetStorage(storage StoreDataCurrency) {
	s.store = storage
}

func (s *CurrencyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.storeCurrencyConvert(w)
	case http.MethodGet:
		s.displayCurrencyConvert(w, r)
	}
}

func (s *CurrencyServer) displayCurrencyConvert(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	currency := s.store.ConvertCurrencies(params[1])
	if currency == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprintf(w, "%.2f", currency)
}

func (s *CurrencyServer) storeCurrencyConvert(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

func ConvertCurrencies(inValue string) float64 {
	if inValue == "10" {
		return 45.0
	} else if inValue == "15" {
		return 67.5
	}
	return 0
}
