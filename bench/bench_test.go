package bench

import (
	"fmt"
	"testing"
)

func BenchmarkHellp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hell")
	}
}
