package main

import (
	"alyson/twg/math"
	"fmt"
)

func main() {
	sum := math.Sum([]int{10, -2, 3})
	if sum != 11 {
		msg := fmt.Sprintf("FAIL: wanted 11 but received %d", sum)
		panic(msg)
	}

	add := math.Add(5, 10)
	if add != 15 {
		msg := fmt.Sprintf("FAIL: wanted 15 but received %d", sum)
		panic(msg)
	}
	fmt.Println("pass")
}
