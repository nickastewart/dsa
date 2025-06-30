package main

import "fmt"

func main() {
	arr := []int{31, 41, 59, 26, 41, 58}
	fmt.Println(insertionSortAscending(arr))
}

func insertionSortAscending(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		var j int = i - 1
		for j > -1 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}

