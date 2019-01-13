package compare

import (
	"reflect"
	"testing"
)

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

	t.Logf("deepEql %v", reflect.DeepEqual(nikuo, nikuo2))
	if !reflect.DeepEqual(nikuo, nikuo2) {
		t.Errorf("nikuo != nikuo2")
	}
}

func TestDogWithFn(t *testing.T) {
	daibutsu := &DogWithFn{
		Name: "大仏",
		Age:  94,
	}
	daibutsu2 := &DogWithFn{
		Name: "大仏",
		Age:  94,
	}

	t.Logf("daibutsu=%p, daibutsu2=%p", &daibutsu, &daibutsu2)
	if !reflect.DeepEqual(daibutsu, daibutsu2) {
		t.Errorf("daibutsu != daibutsu1")
	}

}
