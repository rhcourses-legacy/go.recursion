package lists

import "fmt"

func ExampleAddElements() {
	l1 := []int{1, 2, 3, 4, 5}
	fmt.Println(AddElements(l1))

	l2 := []int{5, 4, 3, 2, 1, 0}
	fmt.Println(AddElements(l2))

	l3 := []int{18, 20, 22}
	fmt.Println(AddElements(l3))

	l4 := []int{}
	fmt.Println(AddElements(l4))

	l5 := []int{1, -1}
	fmt.Println(AddElements(l5))

	// Output:
	// 15
	// 15
	// 60
	// 0
	// 0
}
