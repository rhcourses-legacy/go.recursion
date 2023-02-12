package conway

import "fmt"

// ExampleElement_Next_start1 gibt die ersten 5 Elemente der Conway-Folge für den Startwert 1 aus.
func ExampleElement_Next_start1() {
	current := Element{1}

	for i := 0; i < 5; i++ {
		fmt.Println(current)
		current = current.Next()
	}

	// Output:
	// [1]
	// [1 1]
	// [2 1]
	// [1 2 1 1]
	// [1 1 1 2 2 1]
}

// ExampleElement_Next_start2 gibt die ersten 5 Elemente der Conway-Folge für den Startwert 2 aus.
func ExampleElement_Next_start2() {
	current := Element{2}

	for i := 0; i < 5; i++ {
		fmt.Println(current)
		current = current.Next()
	}

	// Output:
	// [2]
	// [1 2]
	// [1 1 1 2]
	// [3 1 1 2]
	// [1 3 2 1 1 2]
}

// ExampleConway nutzt nun die rekursive Conway-Funktion,
// um die Elemente der Conway-Folge zu bestimmen.
func ExampleConway() {
	conway_elements := Conway(1, 5)
	for _, el := range conway_elements {
		fmt.Println(el)
	}

	// Output:
	// [1]
	// [1 1]
	// [2 1]
	// [1 2 1 1]
	// [1 1 1 2 2 1]
	// [3 1 2 2 1 1]
}
