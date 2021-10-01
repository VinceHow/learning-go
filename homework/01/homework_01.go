package main

import (
	"fmt"
)

func makeSlice(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func multiplySlice(numbersSlice []int) int {
	var totalProduct int
	totalProduct = 1
	for _, v := range numbersSlice {
		totalProduct = totalProduct * v
	}
	return totalProduct
}

func Combination(n int, k int) {
	/*
		Goal: Create a function that it is able to compute the number of possible ways you can pick out a combination of k numbers from a total of n numbers

		C(n, k) = n!/((n − k)! * k!)
	*/
	numbersSlice := makeSlice(k, n, 1)

}
func main() {
	numbersSlice := []int{1,2,3}
	fmt.Println(multiplySlice(numbersSlice))
}






func Permutation(n int, k int) {
	/*
		Goal: compute the number of possible ways you order k numbers from a selection of n numbers

		C(n, k) = n!/(n − k)!
	*/
}
