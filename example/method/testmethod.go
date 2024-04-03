package main

import "errors"

func SumPositiveNumber(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a and b must be positive")
	}

	return a + b, nil
}
