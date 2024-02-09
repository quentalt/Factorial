package Factorial_test

import (
	"errors"
	Factorial "factorial"
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {
	var ErrNegatif = "negative number"
	var ErrOverflow = errors.New("integer overflow")

	tests := []struct {
		n    int
		want uint64
		err  error
	}{
		{0, 1, nil},
		{1, 1, nil},
		{2, 2, nil},
		{5, 120, nil},
		{10, 3628800, nil},
		{1, 0, fmt.Errorf(ErrNegatif)},
		{-5, 0, fmt.Errorf(ErrNegatif)},
		{20, 0, fmt.Errorf(ErrOverflow.Error())},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			got, err := Factorial.Factorial(tt.n)
			if got != tt.want {
				t.Errorf("Factorial(%d) got = %d, want %d", tt.n, got, tt.want)
			}
			if fmt.Sprintf("%v", err) != fmt.Sprintf("%v", tt.err) {
				t.Errorf("Factorial(%d) err = %v, want %v", tt.n, err, tt.err)
			}
		})
	}
}
