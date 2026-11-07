package main

import "fmt"

/**
 * This function intSeq returns another function,
 * which we define anonymously in the body of intSeq.
 * The returned function closes over the variable i to form a closure.
 */
func intSeq() func() (int, int) {
	m := 0
	return func() (int, int) {
		n := 0
		m++
		n++
		return m, n
	}
}

func main() {
	/**
	 * We call intSeq, assigning the result (a function) to nextInt.
	 * This function value captures its own i value, which will be updated each time we call nextInt.
	 */
	nextInt := intSeq()

	/* See the effect of the closure by calling nextInt a few times. */
	m, n := nextInt()
	fmt.Println("------nextInts------")
	fmt.Println("m:", m, "n:", n)
	m, n = nextInt()
	fmt.Println("m:", m, "n:", n)
	m, n = nextInt()
	fmt.Println("m:", m, "n:", n)

	/* To confirm that the state is unique to that particular function, create and test a new one. */
	newInts := intSeq()
	m, n = newInts()
	fmt.Println("------newInts-------")
	fmt.Println("m:", m, "n", n)
}
