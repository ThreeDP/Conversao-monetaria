package server

import (
	"reflect"
	"testing"
)

func TestIsAllDigit(t *testing.T) {
	t.Run("Check the number 100 should be true", func(t *testing.T) {
		checkBool(t, true, IsAllDigit("100"))
	})

	t.Run("Check the number 100.57 should be true", func(t *testing.T) {
		checkBool(t, true, IsAllDigit("100.57"))
	})

	t.Run("Check the number 1.00.57 should be false", func(t *testing.T) {
		checkBool(t, false, IsAllDigit("1.00.57"))
	})

	t.Run("Check the string goku should be false", func(t *testing.T) {
		checkBool(t, false, IsAllDigit("goku"))
	})

	t.Run("Check the empty string should be false", func(t *testing.T) {
		checkBool(t, false, IsAllDigit(""))
	})

	t.Run("Check the nil string should be false", func(t *testing.T) {
		var str string
		checkBool(t, false, IsAllDigit(str))
	})
}

func TestValidConversion(t *testing.T) {
	t.Run("Check BTC conversion", func(t *testing.T) {
		expected := []string{USD, BRL}
		result := ValidConversion(BTC)
		checkCurrencyTypes(t, expected, result)
	})

	t.Run("Check BRL conversion", func(t *testing.T) {
		expected := []string{USD, EUR}
		result := ValidConversion(BRL)
		checkCurrencyTypes(t, expected, result)
	})

	t.Run("Check USD conversion", func(t *testing.T) {
		expected := []string{BRL}
		result := ValidConversion(USD)
		checkCurrencyTypes(t, expected, result)
	})

	t.Run("Check USD conversion", func(t *testing.T) {
		expected := []string{BRL}
		result := ValidConversion(EUR)
		checkCurrencyTypes(t, expected, result)
	})

	t.Run("Check USD conversion", func(t *testing.T) {
		var expected []string
		result := ValidConversion("")
		checkCurrencyTypes(t, expected, result)
	})

	t.Run("Check USD conversion", func(t *testing.T) {
		var expected []string
		result := ValidConversion("GO")
		checkCurrencyTypes(t, expected, result)
	})

	t.Run("Check USD conversion", func(t *testing.T) {
		var expected []string
		var str string
		result := ValidConversion(str)
		checkCurrencyTypes(t, expected, result)
	})

}

func TestValidExitCurrency(t *testing.T) {
	t.Run("with item in list ", func(t *testing.T) {
		var expected error
		result := ValidExitCurrency("USD", []string{"USD", "EUR"})
		checkError(t, expected, result)
	})

	t.Run("without item in list", func(t *testing.T) {
		expected := ErrorNotFoundCurrency
		result := ValidExitCurrency("USD", []string{"BRL", "EUR"})
		checkError(t, expected, result)
	})

	t.Run("nil str", func(t *testing.T) {
		var str string
		expected := ErrorNotFoundCurrency
		result := ValidExitCurrency(str, []string{"BRL", "EUR"})
		checkError(t, expected, result)
	})

	t.Run("nil currencies", func(t *testing.T) {
		var str []string
		expected := ErrorNotFoundCurrency
		result := ValidExitCurrency("BRL", str)
		checkError(t, expected, result)
	})
}

func checkError(t *testing.T, expected, result error) {
	t.Helper()
	if expected != result {
		t.Errorf("expected %s, result %s", expected, result)
	}
}

func checkCurrencyTypes(t *testing.T, expected, result []string) {
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, result %v", expected, result)
	}
}

func checkBool(t *testing.T, expected, result bool) {
	t.Helper()
	if expected != result {
		t.Error("expected a result true, but has a false!")
	}
}