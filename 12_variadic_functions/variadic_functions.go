package main

import "fmt"

/* Here’s a function that will take an arbitrary number of ints as arguments. */
func sum(nums ...int) {
	fmt.Println("nums:", nums)
	total := 0

	/**
	 * Within the function, the type of nums is equivalent to []int.
	 * We can call len(nums), iterate over it with range, etc.
	 */
	for index, num := range nums {
		fmt.Println("num:", num)
		fmt.Println("index:", index)
		total += num
	}
	fmt.Println("total:", total)
}

func main() {
	/* Variadic functions can be called in the usual way with individual arguments. */
	sum(1, 2)
	sum(1, 2, 3)

	/* If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this. */
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
