package random

import (
	"math/rand"
	"testing"
	"time"
)

func TestPick(t *testing.T) {
	seed := time.Now().UnixNano()
	//var seed int64 = 1546767309267641000
	r := rand.New(rand.NewSource(seed))

	arg := make([]int, 25)

	for i := 0; i < 25; i++ {
		arg[i] = r.Int()
	}

	got := Pick(arg)

	for _, v := range arg {
		if got == v {
			return
		}
	}
	t.Errorf("Pick(seed=%d) = %d; not in slice ", seed, got)
}
