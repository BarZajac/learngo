package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		fmt.Printf("Enter your guess: ")

		var ui string
		if _, err := fmt.Scanln(&ui); err != nil {
			fmt.Println("System error.")
			return
		}

		guess, err := strconv.ParseInt(ui, 10, 64)
		if err != nil {
			fmt.Println("You must enter an integer between 1 and 6.")
			continue
		}

		if guess < 1 || guess > 6 {
			fmt.Println("You must enter an integer between 1 and 6.")
			continue
		}

		if int64(rand.Intn(6)+1) == guess {
			fmt.Printf("You entered %d and WON!\n", guess)
			return
		} else {
			fmt.Println("You guessed WRONG!")
		}
	}
}
