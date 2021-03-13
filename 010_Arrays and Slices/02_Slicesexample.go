//some slices concept
/*Slices are a key data type in Go,
giving a more powerful interface to sequences than arrays*/
// They are refrence to the existing arrays

package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:2]
	fmt.Println(b)
	//the output to this code is last index not include
	s := make([]string, 3)

	fmt.Println("emp: ", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len of s", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy", c)

	l := s[2:5]
	fmt.Println("sl1", l)

	l = s[:5]
	fmt.Println("sl2", l)

	l = s[2:]
	fmt.Println("sl3", l)

	// declare and initialize a variable for slice in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	twoD2 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD2[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD2[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD2)

}
