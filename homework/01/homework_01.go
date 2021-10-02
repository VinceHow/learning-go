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


func FactorialArray(n int) []int {
	finalSlice := make([]int, n)
	for i := 1; i <= n ; i++ {
		if i == 1 {
			finalSlice[0] = 1
		} else {
			finalSlice[i-1] = finalSlice[i-2] * i
		}
	}
	fmt.Println(finalSlice)
	return finalSlice
}

func FibonacciArray(n int) []int {
	finalSlice := make([]int ,n)
	for i:= 1; i<=n; i++ {
		if i <= 2 {
			finalSlice[i-1] = 1
		} else {
			var sumSlice int
			sumSlice += finalSlice[i-3]
			sumSlice += finalSlice[i-2]
			finalSlice[i-1] = sumSlice
		}
	}
	fmt.Println(finalSlice)
	return finalSlice
}

func MinArray(x []int) int {
	smallestInt := x[0]
	for i := range x {
		if smallestInt > x[i] {
			smallestInt = x[i]
		}
	}
	//fmt.Println("Smallest int is", smallestInt)
	return smallestInt
}

func MaxArray(x []int) int {
	biggestInt := x[0]
	for i := range x {
		if biggestInt < x[i] {
			biggestInt = x[i]
		}
	}
	//fmt.Println("Biggest int is", biggestInt)
	return biggestInt
}

func indexOf(element int, data []int) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1    //not found.
}

func GCDArray(x []int) int {
	max := MaxArray(x)
	min := MinArray(x)
	for max != min {
		maxIndex := indexOf(max, x)
		dif := max - min
		x[maxIndex] = dif
		max = MaxArray(x)
		min = MinArray(x)
	}
	fmt.Println("GCD is", max)
	return max
}

func IsPerfect(x int) bool {
	allNumbers := makeSlice(1, x/2) // not a divisor of itself
	var divisors []int
	for _, v := range allNumbers {
		if x % v == 0 {
			divisors = append(divisors, v)
		}
	}
	var sumDivisors int = 0
	for _, v := range divisors {
		sumDivisors += v
	}
	var isPerfect bool = sumDivisors == x
	fmt.Printf("Integer %v is a perfect number: %v", x , isPerfect)
	return isPerfect
}

func NextPerfectNumber(x int) int {
	x += 1 // if x is perfect, we want to find the next one
	for IsPerfect(x) != true {
		x += 1
	}
	return x
}

func main() {
	// Task #1: create a function to compute permutations and combinations
	//Combination(50 , 5)

	// Task #2: create FactorialArray
	//FactorialArray(6)

	// Task #3: create FibonacciArray
	// FibonacciArray(20)

	// Task #4: create MinArray
	//MinArray([]int{10,20,-50})

	//Task #5: create GCDArray
	//GCDArray([]int{378,273})

	// Task #6: IsPerfect
	//IsPerfect(29)

	// Task #7: NextPerfectNumber
	NextPerfectNumber(8128)

}



