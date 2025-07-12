package main

import "fmt"

func main() {

	arr1 := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 := []int8{1, 2, 3, 5, 6, 7, 8, 9}

	xor1 := xorArray(arr1)
	xor2 := xorArray(arr2)

	fmt.Printf("  %d \n", arr1)
	fmt.Printf("  %d \n", arr2)
	fmt.Printf("  %d \n", xor1 ^ xor2)
}

func xorArray(arr []int8) int8 {
	var xor int8
	for _, x := range arr {
		xor = xor ^ x
	}
	return xor
}
