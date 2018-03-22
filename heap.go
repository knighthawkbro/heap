package main

import (
	"fmt"

	"github.com/knighthawkbro/heap/array"
)

// Heap (Public) -
type Heap interface {
	Heap(numbers []int)
	Remove() int
	Add(item int)
	Size() int
	Get() int
}

func main() {
	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning driver function as an array...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	arr := array.New()
	driver(arr)
}

func driver(h Heap) {
	numbers := []int{8, 5, 3, 9, 2, 7}
	fmt.Println("Original array\n", numbers)
	h.Heap(numbers)
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

	size := h.Size()
	for x := 0; x < size; x++ {
		h.Remove()
	}
	fmt.Println("\nAfter removing everything\n", h)
}
