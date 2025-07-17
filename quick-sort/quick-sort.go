package main

import "fmt"

func main() {
	arr := []int{4, 2, 9, 8, 3, 6, 5, 14, 1, 12}
	fmt.Printf("Unsorted: %v \n",arr)
	quicksort(arr, 0, len(arr)-1)
	fmt.Printf("Sorted: %v", arr)
}

func quicksort(arr []int, p int, r int) {
	if p < r {
		q := partition(arr, p, r)
		quicksort(arr, p, q-1)
		quicksort(arr, q+1, r)
	}
}

func partition(arr []int, p int, r int) int {
	x := arr[r]
	i := p - 1
	j := p
	for j <= r-1 {
		if arr[j] <= x {
			i++
			temp := arr[j]
			arr[j] = arr[i]
			arr[i] = temp
		}
		j++
	}

	temp := arr[i+1]
	arr[i+1] = arr[j]
	arr[j] = temp

	return i + 1
}
