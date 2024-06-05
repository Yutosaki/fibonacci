package main

import (
	"math/big"
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		n        int
		expected *big.Int
	}{
		{0, big.NewInt(0)},
		{1, big.NewInt(1)},
		{2, big.NewInt(1)},
		{3, big.NewInt(2)},
		{4, big.NewInt(3)},
		{5, big.NewInt(5)},
		{10, big.NewInt(55)},
		{20, big.NewInt(6765)},
		{30, big.NewInt(832040)},
	}

	for _, test := range tests {
		result := fibonacci(test.n)
		if result.Cmp(test.expected) != 0 {
			t.Errorf("fibonacci(%d) = %s, expected %s", test.n, result.String(), test.expected.String())
		}
	}
}
