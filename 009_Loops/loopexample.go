package main

//for is only GO's looping construct
import (
	"fmt"
)

func main() {
	//loops by example
	i := 1
	for i <= 6 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("i loop comleterd")
	//classic example
	for j := 10; j <= 20; j++ {
		fmt.Println(j)
	}
	fmt.Println("j loop comleterd")
	// continue to the next iteration
	for k := 1; k <= 10; k++ {

		if k%2 == 0 {
			continue
		}
		fmt.Println(k)
	}
	fmt.Println("k loop comleterd")
}
