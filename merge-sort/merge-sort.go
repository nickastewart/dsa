package main

import "fmt"

func main() {
	arr := []int{3, 41, 52, 26, 38, 57, 9, 49}
	fmt.Println(mergesort(arr, 0, 7))
}

func mergesort(arr []int, p int, r int) []int {
	if p >= r {
		return arr
	}
	q := (p + r) / 2
	mergesort(arr, p, q)
	mergesort(arr, q+1, r)
	return merge(arr, p, q, r)
}

func merge(arr []int, p int, q int, r int) []int {
	leftLength := q - p + 1
	rightLength := r - q
	left := make([]int, leftLength)
	right := make([]int, rightLength)

	for i := range leftLength {
		left[i] = arr[p+i]
	}

	for j := range rightLength {
		right[j] = arr[q+j+1]
	}

	i := 0
	j := 0
	k := p

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[k] = left[i]
			i = i + 1
		} else {
			arr[k] = right[j]
			j = j + 1
		}
		k = k + 1
	}

	for i < len(left) {
		arr[k] = left[i]
		i = i + 1
		k = k + 1
	}

	for j < len(right) {
		arr[k] = right[j]
		j = j + 1
		k = k + 1
	}

	return arr
}
