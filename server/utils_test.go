package server

import (
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

func checkBool(t *testing.T, expected, result bool) {
	t.Helper()
	if expected != result {
		t.Error("expected a result true, but has a false!")
	}
}