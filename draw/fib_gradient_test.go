package draw_test

import (
	twgdraw "alyson/twg/draw"
	"fmt"
	"image/draw"
	"testing"
)

func TestFibGradient(t *testing.T) {
	var im draw.Image // = image.NewRGBA(image.Rect(0, 0, 20, 20))
	twgdraw.FibGradient(im)
}

func TestFibFunc(t *testing.T) {
	fmt.Println(twgdraw.A)
	got := twgdraw.Fib(2)
	if got != 1 {
		t.Errorf("Fib(2)=%d, want 1", got)
	}
}

//func TestInfo(t *testing.T) {
//	//d := twgdraw.Dog{
//	//	name: "udon",
//	//	age:  3,
//	//}
//	twgdraw.Info(d)
//}
