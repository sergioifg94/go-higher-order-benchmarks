package enumerable

import (
	"testing"
)

func BenchmarkStream(b *testing.B) { benchmarkEnumerables(makeStream, b) }

func makeStream(length int) func() Enumerable {
	return func() Enumerable {
		result := make(chan interface{})

		go func() {
			defer close(result)
			for i := 0; i < length; i++ {
				result <- i
			}
		}()

		return ChanEnumerable(result)
	}
}
