package enumerable

import (
	"fmt"
	"testing"
)

func benchmarkEnumerable(calls int, makeEnumerable func() Enumerable) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			enumerable := makeEnumerable()

			// .Map(..).Filter(..).Map(..).Filter(..)...
			for i := 0; i < calls; i++ {
				if isEven(i) {
					enumerable = enumerable.Map(addOne)
				} else {
					enumerable = enumerable.Filter(isEven)
				}
			}

			enumerable.Fold(0, sum)
		}
	}
}

func benchmarkEnumerables(makeEnumerable func(int) func() Enumerable, b *testing.B) {
	var (
		maxLenght  = 1000000
		lengthStep = 2
		maxCalls   = 20
		callsStep  = 2
	)

	for calls := 1; calls < maxCalls; calls = calls + callsStep {
		for length := 10; length < maxLenght; length = length * lengthStep {
			b.Run(
				fmt.Sprintf("%d,%d", length, calls),
				benchmarkEnumerable(calls, makeEnumerable(length)),
			)
		}
	}
}

func addOne(n interface{}) interface{} {
	return n.(int) + 1
}

func isEven(n interface{}) bool {
	return n.(int)%2 == 0
}

func sum(n1 interface{}, n2 interface{}) interface{} {
	return n1.(int) + n2.(int)
}
