package main

import (
	"errors"
	"fmt"
)

func main() {
	// vars1()
	// vars2()
	// print1To10()
	// printOdd1ToN(200)

	// to := 10
	// sum, err := sumOdd1toN(to)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// fmt.Printf("sum of odd numbers between 1 and %d: %d\n", to, sum)

	c := concatStrings("rafal", "bartek")
	fmt.Println(c)
}

func vars1() {
	// Variable declarations.
	var i int
	var f float64
	var b byte
	var s string

	// Assign values.
	i = 1
	f = 1.234567
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
	// Print 1 to 10.
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func printEven1to10() {
	// Print even numbers from 0 to 10 inclusive.
	for i := 0; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

// printOdd1ToN prints odd numbers between 0 and 100.
func printOdd1ToN(n int) {
	for i := 1; i <= n; i += 2 {
		fmt.Println(i)
	}
}

// sumOdd1toN returns sum of odd numbers between 1 and n.
func sumOdd1toN(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("only numbers greater then 1 are allowed")
	}

	var sum int
	for i := 1; i <= n; i += 2 {
		sum += i // sum = sum + i
	}
	return sum, nil
}

func concatStrings(s1, s2 string) string {
	return s1 + " - " + s2
}
