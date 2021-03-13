//Keyword  for constants "const"  define fix unchangeable vals
package main

import (
	"fmt"
)

func main() {
	const a = 50
	fmt.Println("this is a const", a)

	//declaring a group of const
	const (
		name   = "Hazar"
		age    = 29
		symbol = "youth"
	)
	fmt.Println(name, age, symbol) //reassignment not allowed in const keyword

	//const val should be known at compile time hence it value is not allowed in runtime declaration
	//like this....
	// const ab = math.Sqrt(4)
	// fmt.Println(a)
	// .\constant.go:23:8: const initializer math.Sqrt(4) is not a constant

	// String Constants, Typed and Untyped Constants
	// Strin of const is not declared as string rather as n. but GO is strngly types language
	const n = "Sam"
	var varname = n
	fmt.Printf("type of n %T value %v", varname, varname)

	// var defaultName = "Sam" //allowed
	type yourString string
	// var customName myString = "Sam" //allowed
	// //customName = defaultName        //not allowed
}
