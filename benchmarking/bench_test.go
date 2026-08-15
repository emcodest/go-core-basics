package benchmarking

import "testing"

func BenchmarkAddToList(b *testing.B) {
	for b.Loop() {
		AddToList()
	}

}
