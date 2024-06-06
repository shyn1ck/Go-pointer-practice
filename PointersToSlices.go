package main

import "fmt"

func main() {

	number := []int{1, 2, 3, 4, 5}
	fmt.Println("slice after change:", number)
	ptrNums := &number

	doubleSlice(ptrNums)

	fmt.Println("slice before change:", number)
}

func doubleSlice(slice *[]int) {
	for i := range *slice {
		(*slice)[i] *= 2
	}
}
