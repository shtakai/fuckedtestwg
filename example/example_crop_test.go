package example_test

import (
	"alyson/twg/example"
	"fmt"
	"io"

	// Needed for initialize side effect
	_ "image/png"
)

func ExampleCrop() {
	var r io.Reader

	img, err := example.Decode(r)
	if err != nil {
		panic(err)
	}

	err = example.Crop(img, 0, 0, 20, 20)
	if err != nil {
		panic(err)
	}

	var w io.Writer
	err = example.Encode(img, w)
	if err != nil {
		panic(err)
	}
	fmt.Println("See out.jpg for the cropped image.")

	// Output:
	// See out.jpg for the cropped image.
}
