package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num)
}
