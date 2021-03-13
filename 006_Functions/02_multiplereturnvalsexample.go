package main

import "fmt"

func vals() (int, int, int) {
	return 3, 7, 9
}
func main() {
	//Go has built in support for multiple return values
	//return both results and error values from functions

	a, b, c := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	_, b, d := vals()
	fmt.Println(d)
}
