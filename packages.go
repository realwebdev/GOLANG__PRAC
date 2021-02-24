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
