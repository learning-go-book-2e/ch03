package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	d := [4]int{5, 6, 7, 8}
	y := make([]int, 2)
	copy(y, d[:])
	fmt.Println(y)
	copy(d[:], x)
	fmt.Println(d)
}
