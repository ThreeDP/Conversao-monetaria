package server

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var (
	ErrorMissParams = errors.New("parameters were not informed.")
	ErrorNoNumeric = errors.New("a non-numeric value was passed to a numeric parameter.")
)

const (
	USD = "USD"
	BRL = "BRL"
	EUR = "EUR"
	BTC = "BTC"
)

const (
	SUSD = "$"
	SBRL = "R$"
	SEUR = "€"
)

type ResponseData struct {
	value float64 
	symbol string
}

type Params struct {
	params []string
}

type CurrencyServer struct {
	
}

func (s *CurrencyServer) ResponseAction(p Params) ResponseData {
	sign := CatSign(p.params[2])
	outValue := Calculate(p.params[0], p.params[3])
	return ResponseData{outValue, sign}
}

func (s *CurrencyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := Params{}
	err := p.ValidParams(r)
	if err != nil {
		return
	}
	resp := s.ResponseAction(p)
	fmt.Fprint(w, resp)
}

func (p *Params) ValidParams(r *http.Request) error {
	p.params = strings.Split(r.URL.Path[len("/exchange/"):], "/")
	if len(p.params) != 4 {
		return ErrorMissParams
	}
	if !IsAllDigit(p.params[0]) || !IsAllDigit(p.params[3]) {
		return ErrorNoNumeric
	}
	return nil
}

func CatSign(label string) string {
	switch label {
	case USD:
		return SUSD
	case EUR:
		return SEUR
	case BRL:
		return SBRL
	}
	return ""
}

func Calculate(value, rate string) float64{
	v, _ := strconv.ParseFloat(value, 64)
	r, _ := strconv.ParseFloat(rate, 64)
	return v * r
}


