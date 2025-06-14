package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {
	
	fn := func (t float64) float64 {
		return (0.5 * a* math.Pow(t, 2)) + (vo * t) + so
	}

	return fn

}

func main() {
	var a float64
	var vo float64
	var so float64
	var t float64
	
	fmt.Printf("Input the aceleration: ")
	fmt.Scan(&a)

	fmt.Printf("Input the Initial velocity: ")
	fmt.Scan(&vo)

	fmt.Printf("Input the initial space: ")
	fmt.Scan(&so)


	fn := GenDisplaceFn(a, vo, so)

	fmt.Printf("Input the time: ")
	fmt.Scan(&t)	

	fmt.Printf("The final space is: ")
	fmt.Println(fn(t))
}