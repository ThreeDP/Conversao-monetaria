package server

import (
	"encoding/json"
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

func (p *Params) ValidParams(path string) error {
	p.params = strings.Split(path, "/")
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
	Value float64 `json:"valorConvertido"`
	Symbol string `json:"simboloMoeda"`
}

func (r *ResponseData) SetData(outValue float64, sign string) {
	r.Value = outValue
	r.Symbol = sign
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
	http.Handler
}

func (s *CurrencyServer) SetCurrencyServer(store StoreConvertHistory) {
	s.store = store
}

/* Start */
func NewCurrencyServer(store StoreConvertHistory) *CurrencyServer{
	s := new(CurrencyServer)
	s.store = store
	router := http.NewServeMux()
	router.Handle("/consult", http.HandlerFunc(s.showLogs))
	router.Handle("/exchange/", http.HandlerFunc(s.handleConvert))
	s.Handler = router
	return s
}

func (s *CurrencyServer) showLogs(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *CurrencyServer) handleConvert(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/exchange/"):]
	p := Params{}
	err := p.ValidParams(path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	jsonResp, err := json.Marshal(s.store.ResponseAction(p))
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
        return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
