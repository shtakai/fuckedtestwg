package example

import (
	"fmt"
)

func ExampleHello_spanish() {
	greeting, err := Hello("Jon")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Fuck, Jon
}
