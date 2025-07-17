package main

import "fmt"

func main() {
	heap := []int{4, 9, 5, 17, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	fmt.Print("Unsorted: ")
	fmt.Print(heap)
	fmt.Print("\n")
	maxHeapSort(heap, len(heap))

	fmt.Print("Heap Sorted: ")
	fmt.Print(heap)
}

func maxHeapSort(heap []int, n int) {
	length := n - 1
	buildMaxHeap(heap, n)
	for i := length; i >= 0; i-- {
		temp := heap[i]
		heap[i] = heap[0]
		heap[0] = temp
		heap := heap[0 : i]
		maxHeapify(heap, 0)
	}
}

func buildMaxHeap(heap []int, n int) {
	heap = heap[:n]
	for i := (n / 2) + 1; i >= 0; i-- {
		maxHeapify(heap, i)
	}
}

func maxHeapify(heap []int, i int) {
	l := left(i)
	r := right(i)
	var largest int
	if l < len(heap) && heap[l] > heap[i] {
		largest = l
	} else {
		largest = i
	}

	if r < len(heap) && heap[r] > heap[largest] {
		largest = r
	}

	if largest != i {
		temp := heap[i]
		heap[i] = heap[largest]
		heap[largest] = temp
		maxHeapify(heap, largest)
	}
}

func parent(i int) int {
	return i / 2
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return (2 * i) + 1
}
