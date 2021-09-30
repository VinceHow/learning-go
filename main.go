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

/*
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
*/

//pointers, holding the memory location of another variable

//& has ONE meaning = the address of a variable that it is next to
//a := 25
//b := &a
//B is now the value of the memory address of A, instead o A itself

//* has TWO meanings
//1. when it is next to a variable: get the value of the variable that this pointer is pointing to (DEREFERENCING)
// a := 25
// b := &a
// c := *b
// c is now 25, which is the value stored at address b
//
//
//2. when * is next to a TYPE: this means the var being created is a POINTER, and the address is holding a TYPE following the *
// var *string myName
//myName is variable that holds the memory address of a string variable

/*
func main(){
	a := 25
	b := &a
	fmt.Println(b)
	var c *(*int) = &b
	fmt.Println(c)
}
*/

func zero_my_int(x *int) {
	*x = 0
}
func main(){
	x := 5
	zero_my_int(&x)
	if x == 0 {
		fmt.Println("Successfully zeroed")
	} else {
		fmt.Println("x is still 5")
	}
}



