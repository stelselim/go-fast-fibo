package go_fib_module

import (
	"testing"
)

var (
	fiboValues = map[int]int{
		// Edge Cases
		0: 0,
		1: 1,
		2: 1,
		// General Test Cases
		3:  2,
		4:  3,
		5:  5,
		6:  8,
		7:  13,
		8:  21,
		9:  34,
		10: 55,
		// Performance Testing
		21: 10946,
		22: 17711,
		23: 28657,
		24: 46368,
		25: 75025,
	}
)

const (
	internalErrorMessage   = "An error occured calculating"
	wrongValueErrorMessage = "Fibonacci value miscalculated"
	expectedValueMessage   = "Fibonacci value must be"
)

func TestFastFibo(t *testing.T) {
	for i := 0; i < 11; i++ {
		case0 := i
		fibo_0, error := FastFibo(case0)
		if error != nil {
			t.Fatal(internalErrorMessage)
		}
		if fibo_0 != fiboValues[case0] {
			t.Log(case0, "->", fibo_0)
			t.Log(expectedValueMessage, case0, "->", fiboValues[i])
			t.Error(wrongValueErrorMessage)
		}
	}

}
