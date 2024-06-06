package main

import "fmt"

func main() {

	value := 10
	ptr := &value
	fmt.Printf("After incrementing: %d\n", value)
	increaseValue(ptr, 5)
	fmt.Printf("Before incrementing: %d\n", value)
}

func increaseValue(ptr *int, amount int) {
	*ptr += amount
}
