package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestConvertCurrency(t *testing.T) {
	s := &CurrencyServer{}
	t.Run("Convert 10 BRL to USD using rate of 4.50", func(t *testing.T) {
		request := newRequestConvert("10", "BRL", "USD", "4.50")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		checkReturnValue(t, response.Body.String(), fmt.Sprintf("%v", ResponseData{45, "$"}))
	})

	t.Run("Convert 15 BRL to USD using rate of 4.50", func(t *testing.T) {
		request := newRequestConvert("15", "BRL", "USD", "4.50")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		checkReturnValue(t, response.Body.String(), fmt.Sprintf("%v", ResponseData{67.50, "$"}))
	})

	t.Run("Convert 10 BRL to EUR using rate of 6.50", func(t *testing.T) {
		request := newRequestConvert("10", "BRL", "EUR", "6.50")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		checkReturnValue(t, response.Body.String(), fmt.Sprintf("%v", ResponseData{65.00, "€"}))
	})

	t.Run("Convert 10 BRL to EUR without rate", func(t *testing.T) {
		request := newRequestConvert("10", "BRL", "EUR", "")
		response := httptest.NewRecorder()
		s.ServeHTTP(response, request)
		//checkParamsError(t, ErrorMissParams, err)
	})
}

func newRequestConvert(inValue, inCurrency, outCurrency, rate string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/exchange/%s/%s/%s/%s", inValue, inCurrency, outCurrency, rate), nil)
	return request
}

func checkParamsError(t *testing.T, expected, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("expected error: %s", expected)
	}
}

func checkReturnValue(t *testing.T, result, expected string) {
	t.Helper()
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, result %v", expected, result )
	}
}