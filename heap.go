package main

import "fmt"

// Heap (Public) -
type Heap interface {
	Create(numbers []int)
	Remove()
	Add(int)
	Size() int
}

func main() {

}

func driver(h Heap) {
	numbers := []int{8, 5, 3, 9, 2, 7}
	fmt.Println("Original array\n", numbers)
	h.Create(numbers)
	fmt.Println("\nAfter creating a heap\n", h)

	h.Remove()
	fmt.Println("\nAfter removing the largest item\n", h)
	h.Remove()
	fmt.Println("\nAfter removing the largest item\n", h)

	h.Add(13)
	fmt.Println("\nAdding 13 to the heap\n", h)
	h.Add(1)
	fmt.Println("\nAdding 1 to the heap\n", h)
	h.Add(4)
	fmt.Println("\nAdding 4 to the heap\n", h)

	for x := 0; x < h.Size(); x++ {
		h.Remove()
	}
	fmt.Println("\nAfter removing everything\n", h)
}
