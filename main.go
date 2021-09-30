package main

import "fmt"

// var types
//
//unit	unsigned
//int		signed
//byte	uint8
//float32
//float64
//rune	int32, alias
//bool	T or F
//strings

// explicit declaration

func main() {
	// INITIALISATION
	var greeting string
	greeting = "Hello"
	fmt.Println(greeting)
	// IMPLICIT DECLARATION, declare and initialise
	a := 30
	fmt.Println(a)
	year := "2021"
	new_year := "Hi, it's year "
	var myFloat float32 = 5.5555
	fmt.Println(new_year + year)
	fmt.Println(myFloat)
}
