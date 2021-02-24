package main

import "fmt"

func main() {
	fmt.Println("It is awesome to learn go programming language.")
	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()

}
func foo() {
	fmt.Println(" I'm In foo")

}
func bar() {
	fmt.Println("and then we exited")
}

//control flow 3 types ( how a computer  reads or execute in order a code)
//(1) sequenctial
//(2) loop; iterative
//(3) conditional
