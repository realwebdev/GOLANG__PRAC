package main

//variadic Functions // functions with variable arguments
import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//If you already have multiple args in a slice,
	//apply them to a variadic function using func(slice...) like this
	sum(nums...)
}
