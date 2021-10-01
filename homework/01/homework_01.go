package main

import (
	"fmt"
)

func makeSlice(min int, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func multiplySlice(numbersSlice []int) int64 {
	var totalProduct int64
	totalProduct = 1
	for _, v := range numbersSlice {
		totalProduct = totalProduct * int64(v)
	}
	return totalProduct
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Permutation(n int, k int) int64 {
	/*
		Goal: compute the number of possible ways you order k numbers from a selection of n numbers

		C(n, k) = n!/(n − k)!
	*/
	nSlice := makeSlice(1, n)
	kSlice := makeSlice(1, n-k)
	var finalSlice []int
	for _, v := range nSlice {
		if contains(kSlice, v) {
			continue
		} else {
			finalSlice = append(finalSlice, v)
		}
	}
	fmt.Println(finalSlice)
	var permutations int64
	permutations = multiplySlice(finalSlice)
	fmt.Printf("The number of permutations is %v \n", permutations)
	return permutations
}

func Combination(n int, k int) int64 {
	/*
		Goal: Create a function that it is able to compute the number of possible ways you can pick out a combination of k numbers from a total of n numbers

		C(n, k) = n!/((n − k)! * k!)
	*/
	var permutations int64
	permutations = Permutation(n, k)
	numbersSlice := makeSlice(1, k)
	var kFactorial int64
	kFactorial = multiplySlice(numbersSlice)
	combinations := permutations/kFactorial
	fmt.Printf("The number of combinations is %v", combinations)
	return combinations
}

func main() {
	Combination(50 , 5)
}



