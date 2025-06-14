package main

/*
The race condition occur between the goroutine of main the other goroutine of printf.

The two routines are running concurrently, so it is possible the main to end before
the child routine to end, what causes a race condition.

*/



import (
	"fmt"
)

func main() {
	var x int = 1

	go fmt.Printf("The value of x is: %d\n", x)

	x = 5
	fmt.Printf("The new value of x is: %d\n", x)
}