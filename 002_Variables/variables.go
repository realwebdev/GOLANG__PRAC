// Declaring single variable
package main

import (
	"fmt"
	"math"
)

func main() {
	var age int
	var name string // declaring variables without initial value

	fmt.Println("no inital declaration of value to variable of age", age)

	name = "Haseeb"
	age = 8
	fmt.Println("This is my name", name, "and this is my age", age)

	// Type inference Go will automatically infer type from assigned value to variable

	var age_2 = 29
	fmt.Println("this is my new age", age_2)

	//Multiple variable declaration

	var age3, name2 = 100, "hamza"
	fmt.Print(age3, name2)

	//varable in different formats undersame declaration
	//  var (
	// 	name1=initial value
	// 	name2=initial value
	// )
	var (
		person_Name = "Naveen"
		p_age       = 29
		Height      = 5.4
	)
	fmt.Println(" this is what it's look like", person_Name, p_age, Height)

	//short hand declaration

	count := 10
	fmt.Println(count)

	count1, count2 := 12, 13

	fmt.Println(count1, count2)
	//short hand notation should be used already assigned variable
	//don't use var keyword with short hand notation

	name12, age12 := 2, 9
	fmt.Println(name12, age12)

	//there should be atleast one new variable across multiple declarations in short hand syntax

	//     a, b := 20, 30 //a and b declared
	//     fmt.Println("a is", a, "b is", b)
	//     a, b := 40, 50 //error, no new variables

	//----->Runitme variable assignment

	a, b := 5.10, 10.3
	c := math.Min(a, b)
	fmt.Println("minimum value is", c)

	// Go is strongly typed, variables declared as belonging to one type cannot be assigned a value of another type

	// age := 29 -> age = "naveen" it gives error

}
