package main

import "fmt"

func main() {

	a := [][]int{{1, 2}, {4, 5}}

	b := [][]int{{1, 2}, {4, 5}}

	c := [][]int{{0, 0}, {0, 0}}

	matrixMultiply(a, b, c, 2)
	fmt.Println(c)
}

func matrixMultiply(a [][]int, b [][]int, c [][]int, n int) {

	a1 := [][]*int {{&a[0][0]}}
	a2 := [][]*int {{&a[0][1]}}
	a3 := [][]*int {{&a[1][0]}}
	a4 := [][]*int {{&a[1][1]}}

	b1 := [][]*int {{&b[0][0]}}
	b2 := [][]*int {{&b[0][1]}}
	b3 := [][]*int {{&b[1][0]}}
	b4 := [][]*int {{&b[1][1]}}

	c1 := [][]*int {{&c[0][0]}}
	c2 := [][]*int {{&c[0][1]}}
	c3 := [][]*int {{&c[1][0]}}
	c4 := [][]*int {{&c[1][1]}}

	matrixMultiplyPointer(a1, b1, c1, n/2)
	matrixMultiplyPointer(a1, b2, c2, n/2)
	matrixMultiplyPointer(a3, b1, c3, n/2)
	matrixMultiplyPointer(a3, b2, c4, n/2)
	matrixMultiplyPointer(a2, b3, c1, n/2)
	matrixMultiplyPointer(a2, b4, c2, n/2)
	matrixMultiplyPointer(a4, b3, c3, n/2)
	matrixMultiplyPointer(a4, b4, c4, n/2)
}

func matrixMultiplyPointer(a [][]*int, b[][]*int, c[][]*int, n int){

	if n == 1 {
		*c[0][0] = *c[0][0] + *a[0][0] * *b[0][0]
	}

}
