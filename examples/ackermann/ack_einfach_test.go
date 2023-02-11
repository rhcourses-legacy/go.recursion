package ackermann

import "fmt"

// ExampleAck berechnet verschiedene Werte der Ackermannfunktion mittels
// der einfachen Variante Ack.
// Dabei gibt es auch die Werte für m mit aus.
func ExampleAck() {
	for m := 0; m <= 3; m++ {
		fmt.Printf("m == %d:", m)
		for n := 0; n < 3; n++ {
			fmt.Printf(" %2d", Ack(m, n))
		}
		fmt.Println()
	}

	// Output:
	// m == 0:  1  2  3
	// m == 1:  2  3  4
	// m == 2:  3  5  7
	// m == 3:  5 13 29
}
