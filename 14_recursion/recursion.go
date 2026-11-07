package main

import "fmt"

/* This fact function calls itself until it reaches the base case of fact(0). */
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	/**
	 * Anonymous functions can also be recursive,
	 * but this requires explicitly declaring a variable with var to store the function before it’s defined.
	 */
	var flib func(n int) int
	flib = func(n int) int {
		if n < 2 {
			return n
		}

		/* Since fib was previously declared in main, Go knows which function to call with fib here */
		return flib(n-1) + flib(n-2)
	}

	fmt.Println(flib(7))
}
