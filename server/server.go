package server

import (
	"fmt"
	"net/http"
	"strings"
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

/* Params */
type Params struct {
	params []string
}

func (p *Params) GetParams() []string {
	return p.params
}

func (p *Params) ValidParams(r *http.Request) error {
	p.params = strings.Split(r.URL.Path[len("/exchange/"):], "/")
	if len(p.params) != 4 {
		return ErrorMissParams
	}
	if !IsAllDigit(p.params[0]) || !IsAllDigit(p.params[3]) {
		return ErrorNoNumeric
	}
	keys := ValidConversion(p.params[1])
	if keys == nil {
		return ErrorNotFoundCurrency
	}
	if err := ValidExitCurrency(p.params[2], keys); err != nil {
		return err
	}
	return nil
}

/* ResponseData */
type ResponseData struct {
	value float64 
	symbol string
}

func (r *ResponseData) SetData(outValue float64, sign string) {
	r.value = outValue
	r.symbol = sign
}

func (r *ResponseData) ResponseAction(p Params) ResponseData {
	sign := CatSign(p.params[2])
	outValue := Calculate(p.params[0], p.params[3])
	return ResponseData{outValue, sign}
}

type StoreConvertHistory interface {
	ResponseAction(p Params) ResponseData
}

/* Server */
type CurrencyServer struct {
	store StoreConvertHistory
}

func (s *CurrencyServer) SetCurrencyServer(store StoreConvertHistory) {
	s.store = store
}

/* Start */
func (s *CurrencyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := Params{}
	err := p.ValidParams(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w, s.store.ResponseAction(p))
}
