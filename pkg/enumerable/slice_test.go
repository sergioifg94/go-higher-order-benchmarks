package enumerable

import (
	"testing"
)

func BenchmarkSlice(b *testing.B) { benchmarkEnumerables(makeSlice, b) }

func makeSlice(length int) func() Enumerable {
	return func() Enumerable {
		result := make(SliceEnumerable, length)

		for i := 0; i < length; i++ {
			result[i] = i
		}

		return result
	}
}
