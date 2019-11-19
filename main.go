package main

import (
    "fmt"
	"math"
	/* "math/rand"	var rr int = rand.Intn(100)*/
)

func main() {
	var a string = "hello world"
	var b complex128 = -5 +12i
	var c [2]int
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(math.Cbrt(6))
	numbers := make([]int,5,10)
	numbers = append(numbers, 1, 2, 3, 4)
	number2 := make([]int, 15)
	copy(number2, numbers)
	fmt.Println(numbers)

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	}else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}