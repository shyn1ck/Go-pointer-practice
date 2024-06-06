package main

import "fmt"

func main() {
	Num1 := 7
	Num2 := 10

	fmt.Printf("Before Num1: %d, Num2: %d\n", Num1, Num2)

	swap(&Num1, &Num2)

	fmt.Printf("Before Num1: %d, Num2: %d\n", Num1, Num2)
}

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
