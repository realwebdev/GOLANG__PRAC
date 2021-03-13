package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//types in Go
	// bool
	// string
	// numeric
	/*int8, int16, int32, int64, int
	  uint8, uint16, uint32, uint64, uint
	  float32, float64
	  complex64, complex128
	  byte
	  rune*/

	// fmt.Println("This is bool declaration", a)

	//further example
	//var a bool = true
	// b := false
	// fmt.Println("a:", a, "b:", b)
	// c := a && b
	// fmt.Println("c:", c)
	// d := a || b
	// fmt.Println("d:", d)

	//%T is the format specifier to print the type and %d is used to print the size.
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))   //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) //type and size of b

	//floating numbers
	q, w := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", q, w)
	sum := q + w
	diff := q - w
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)
	//complex number 64 and 128 bit contain real float and imaginary part

	// syntax complex(r,i, FloatType)ComplexType // natural decalartion is 128 bit for complex with out inference
	var x complex128 = complex(1, 2)
	c2 := 7 + 2i
	cadd := x + c2
	cmul := x * c2
	fmt.Println("this is addition", cadd, "this is multiplication", cmul)

	//STring types
	fname := "Haseeb"
	lname := "Zafar"
	fullname := fname + "    " + lname
	fmt.Println("this is my fullname", fullname)

	//conversion code
	i := 55            //int
	k := 67.8          //float64
	sum1 := i + int(k) //j is converted to int
	fmt.Println(sum1)

	l := 10
	var j float64 = float64(l) //this statement will not work without explicit conversion
	fmt.Println("j", j)
}
