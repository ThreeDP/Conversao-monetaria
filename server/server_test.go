package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SketchStoreData struct {
	convertedCurrency map[string]float64
	storeDataCurrency []string
}

func (e *SketchStoreData) ConvertCurrencies(inValue string) float64 {
	convertedCurrency := e.convertedCurrency[inValue]
	return convertedCurrency
}

func (e *SketchStoreData) StoreDataCurrency(inValue string) {
	e.storeDataCurrency = append(e.storeDataCurrency, inValue)
}

func TestConvertCurrencies(t *testing.T)  {

	storage := SketchStoreData{
		map[string]float64 {
			"10": 45.0,
			"15": 67.5,
		},
		nil,
	}

	server := &CurrencyServer{&storage}
	t.Run("Returns the conversion of R$ 10 to dollars", func (t *testing.T) {
		request := newRequestConvertCurries("10")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		checkBodyrequest(t, fmt.Sprintf("%.2f", 45.0), response.Body.String())
		checkResponseCode(t, http.StatusOK, response.Code)
	})

	t.Run("Returns the conversion of R$ 15 to dollars", func (t *testing.T) {
		request := newRequestConvertCurries("15")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		checkBodyrequest(t, fmt.Sprintf("%.2f", 67.5), response.Body.String())
		checkResponseCode(t, http.StatusOK, response.Code)
	})

	t.Run("Returns 404 if the currency is not found", func(t *testing.T) {
		request := newRequestConvertCurries("13")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		checkResponseCode(t, http.StatusNotFound, response.Code)
	})
}

func TestStoreData (t *testing.T) {
	store := SketchStoreData {
		map[string]float64{},
		nil,
	}
	server := &CurrencyServer{&store}
	t.Run("returns stats 'accept' for calls with method POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/exchange/11", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		checkResponseCode(t, response.Code, http.StatusAccepted)
		if len(store.storeDataCurrency) != 1 {
			t.Errorf("expected %d, result %d", 1, len(store.storeDataCurrency))
		}
	})
}

func newRequestConvertCurries(inValue string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("exchange/%s/BRL/USD/4.50", inValue), nil)
	return request
}

func checkBodyrequest(t *testing.T, expected, result string) {
	t.Helper()
	if result != expected {
		t.Errorf("expected '%s', result '%s'", expected, result)
	}
}

func checkResponseCode(t *testing.T, expected, result int) {
	t.Helper()
	if expected != result {
		t.Errorf("expected %d, result %d", expected, result)
	}
}