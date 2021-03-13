package main

import "fmt"

func main() {
	//if else examples
	if 8%2 == 0 {
		fmt.Println("8 is even")
	} else {
		fmt.Println("8 is odd")
	}

	if num := 9; num <= 0 {
		fmt.Println(num, "9 is ok")

	} else if num <= 10 {
		fmt.Println(num, "num is not ok")

	} else {
		fmt.Println(num, "num is sick")
	}
}
