package main

//In Go, an array is a numbered sequence of elements of a specific length.
// more use of slices in golang
import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("emp", a)

	// set a value at an index using the array[index] = value syntax, and get a value with array[index].
	a[4] = 100
	fmt.Println("Set:", a)
	fmt.Println("Set:", len(a))
	// declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		fmt.Println("i val: ", i)
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
			fmt.Println("j val: ", j)
		}
	}
	fmt.Println("2D: ", twoD)
}
