package compare

import "testing"

// See fucked up https://golang.org/ref/spec#Comparison_operators to check more

func TestSquare(t *testing.T) {
	arg := 4
	want := 16
	got := Square(arg)
	if got != want {
		t.Errorf("Square(%d) = %d; want %d", arg, got, want)
	}
}

func TestDog(t *testing.T) {
	nikuo := Dog{
		Name: "肉夫",
		Age:  93,
	}
	nikuo2 := Dog{
		Name: "肉夫",
		Age:  93,
	}

	t.Logf("nikup=%p, nikuo2=%p", &nikuo, &nikuo2)
	if nikuo != nikuo2 {
		t.Errorf("nikuo != nikuo2")
	}

	daibutsu := Dog{
		Name: "螺髪",
		Age:  49,
	}

	if nikuo == daibutsu {
		t.Errorf("nikuo == daibutsu")
	}
}

func TestPointer(t *testing.T) {
	nikuo := &Dog{
		Name: "肉夫",
		Age:  93,
	}
	nikuo2 := &Dog{
		Name: "肉夫",
		Age:  93,
	}
	t.Logf("nikup=%p, nikuo2=%p", &nikuo, &nikuo2)
	if nikuo == nikuo2 {
		t.Errorf("nikuo == nikuo2")
	}
}
