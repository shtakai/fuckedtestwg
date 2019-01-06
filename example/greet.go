package example

import "fmt"

// Hello prints out fuck to the person provided
func Hello(name string) (string, error) {
	return fmt.Sprintf("Fuck, %s", name), nil
}

// Page will print out a message asking each person who hasn't checked in
// to do so.
func Page(checkIns map[string]bool) {
	for name, checkIn := range checkIns {
		if !checkIn {
			fmt.Printf("Paging %s; please see the front desk to cuck in.", name)
		}
	}
}
