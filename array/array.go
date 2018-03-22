package array

import "fmt"

const defaultSize = 10

// Array (Public) - Structure that defines
type Array struct {
	size       int
	collection []int
}

// Init (Public) - initializes the array with whatever size is provided, This is what can be overrided by the user.
func (a *Array) Init(capacity int) *Array {
	if capacity < 0 {
		return nil
	}
	a.collection = make([]int, capacity)
	a.size = 0
	return a
}

// New (Public) - Returns an initialized array with default size of 10.
func New() *Array { return new(Array).Init(defaultSize) }

// Heap (Public) - Converts an array to a heap
func (a *Array) Heap(numbers []int) {
	a.collection = numbers
	a.size = len(numbers)
	parent := (a.size - 2) / 2
	for i := parent; i >= 0; i-- {
		a.siftDown(i)
	}
}

// Add (Public) -
func (a *Array) Add(item int) {
	a.ensureSpace()
	a.collection[a.size] = item
	a.size++
	a.siftUp(a.size - 1)
}

// Remove (Public) -
func (a *Array) Remove() int {
	if a.size == 0 {
		return -1
	}

	removed := a.collection[0]
	a.collection[0] = a.collection[a.size-1]
	a.collection[a.size-1] = 0
	a.size--
	a.siftDown(0)
	return removed
}

// Get (Public) -
func (a *Array) Get() int {
	if a.size == 0 {
		return -1
	}
	return a.collection[0]
}

// Size (Public) - returns the size of the Array
func (a *Array) Size() int {
	return a.size
}

// String (Public) - formats the array when fmt.Print is called.
func (a *Array) String() string {
	if a.size == 0 {
		return "[ ]"
	}
	s := "[ "
	for x := 0; x < a.size; x++ {
		s += fmt.Sprintf("%v ", a.collection[x])
	}
	//return s + "]"
	return fmt.Sprintf("%v] %v", s, a.size)
}

// ensureSpace (Private) - Sees if the size and capacity of the array are the same. If so,
// It creates a new array with double the capacity and overwrites the old array with a new
// array, then clears the new array for the GC.
func (a *Array) ensureSpace() {
	if a.size == cap(a.collection) {
		new := new(Array).Init(cap(a.collection) * 2)
		new.size = a.size
		for i := 0; i < a.size; i++ {
			new.collection[i] = a.collection[i]
		}
		*a = *new
		new = nil
	}
}

// checkIndex (Private) -
func (a *Array) checkIndex(index int) error {
	if index < 0 || index >= a.size {
		return fmt.Errorf("index outside of list")
	}
	return nil
}

// siftUp (Private) -
func (a *Array) siftUp(child int) {
	// Add item to bottom of tree
	parent := (child - 1) / 2
	for parent >= 0 {
		if a.collection[child] > a.collection[parent] {
			a.collection[child], a.collection[parent] = a.collection[parent], a.collection[child]
			child = parent
			parent = (child - 1) / 2
		} else {
			break
		}
	}
}

// siftDow (Private) -
func (a *Array) siftDown(parent int) {
	leftChild := 2*parent + 1
	rightChild := 2*parent + 2
	for leftChild > 0 && leftChild < a.size {
		var larger int
		if err := a.checkIndex(rightChild); err != nil {
			larger = leftChild
		} else {
			if a.collection[rightChild] > a.collection[leftChild] {
				larger = rightChild
			} else {
				larger = leftChild
			}
		}

		if a.collection[parent] > a.collection[larger] {
			break
		} else {
			a.collection[parent], a.collection[larger] = a.collection[larger], a.collection[parent]
		}

		parent = larger
		leftChild = 2*parent + 1
		rightChild = 2*parent + 2
	}
}
