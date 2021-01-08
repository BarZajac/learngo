package main

import (
	"fmt"
)

func main() {
	// vars1()
	// vars2()
	print1To10()
}

func vars1() {
	// Variable declarations.
	var i int
	var f float64
	var b byte
	var s string

	// Assign values.
	i = 1
	f = 1.2
	b = 2
	s = "bartek"

	// Print variables.
	// https://golang.org/pkg/fmt/

	fmt.Printf("int     i: %d\n", i)
	fmt.Printf("float64 f: %f\n", f)
	fmt.Printf("byte    b: %d\n", b)
	fmt.Printf("string  s: %s\n", s)
	fmt.Println("---")
}

func vars2() {
	// Assign values.
	i := 1
	f := 1.2
	b := byte(2)
	s := "bartek"

	// Print variables.
	fmt.Printf("%T i: %d \n", i, i)
	fmt.Printf("%T f: %f \n", f, f)
	fmt.Printf("%T b: %d \n", b, b)
	fmt.Printf("%T s: %s \n", s, s)
	fmt.Println("---")
}

func print1To10() {
	// fmt.Println(1)
	// fmt.Println(2)
	// fmt.Println(3)
	// fmt.Println(4)
	// fmt.Println(5)
	// fmt.Println(6)
	// fmt.Println(7)
	// fmt.Println(8)
	// fmt.Println(9)
	// fmt.Println(10)

	// Print 1 to 10.
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Print even numbers from 0 to 10 inclusive.
	for i := 0; i <= 10; i += 2 {
		fmt.Println(i)
	}
}
