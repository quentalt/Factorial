package Factorial

import (
	"errors"
	"fmt"
	"math"
)

func Factorial(n int) (uint64, error) {
	var ErrNegative = "negative number"
	var ErrOverflow = errors.New("integer overflow")

	if n < 0 {
		return 0, fmt.Errorf(ErrNegative)
	}
	var fact uint64 = 1
	for i := uint64(1); i <= uint64(n); i++ {
		if fact > math.MaxUint64/i {
			return 0, fmt.Errorf(ErrOverflow.Error())
		}
		fact *= i
	}
	return fact, nil
}
