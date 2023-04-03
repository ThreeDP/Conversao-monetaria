package server

import (
	"unicode"
	"strconv"
	"errors"
)

var (
	ErrorMissParams = errors.New("parameters were not informed.")
	ErrorNoNumeric = errors.New("a non-numeric value was passed to a numeric parameter.")
	ErrorNotFoundCurrency = errors.New("currency not registered")
)

func IsAllDigit(value string) bool {
	dot := false
	if len(value) <= 0 {
		return false
	}
	for i := 0; i < len(value); i++ {
		if rune(value[i]) == '.' && !dot {
			dot = true
			continue
		}
		if !unicode.IsDigit(rune(value[i])) {
			return false
		}
	}
	return true
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

func ValidConversion(str string) []string {
	switch str {
	case BTC:
		return []string {USD, BRL}
	case BRL:
		return []string {USD, EUR}
	case USD:
		return []string {BRL}
	case EUR:
		return []string {BRL}
	}
	return nil
}

func ValidExitCurrency(str string, currencies []string) error {
	for i := 0; i < len(currencies); i++ {
		if str == currencies[i] {
			return nil
		}
	}
	return ErrorNotFoundCurrency
}