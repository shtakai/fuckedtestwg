package cover

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct {
		base, height, want float64
	}{
		{2, 5, 5},
		{2, 2, 2},
		{11, 4, 22},
	}

	for _, tc := range tests {
		got := Triangle(tc.base, tc.height)
		if got != tc.want {
			t.Errorf("Triangle(%f, %f) = %f; want %f", tc.base, tc.height, got, tc.want)
		}
	}
}
