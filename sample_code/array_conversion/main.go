package main

import "fmt"

func main() {
	arrayConversions()
	arrayPointerConversions()
	panicArrayConversions()
}

func arrayConversions() {
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	xSlice[0] = 10
	fmt.Println(xSlice)
	fmt.Println(xArray)
	fmt.Println(smallArray)
}

func arrayPointerConversions() {
	xSlice := []int{1, 2, 3, 4}
	xArrayPointer := (*[4]int)(xSlice)
	xSlice[0] = 10
	xArrayPointer[1] = 20
	fmt.Println(xSlice)
	fmt.Println(xArrayPointer)
}

func panicArrayConversions() {
	xSlice := []int{1, 2, 3, 4}
	panicArray := [5]int(xSlice)
	fmt.Println(panicArray)
}
