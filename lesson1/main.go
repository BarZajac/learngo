package main

import (
	"fmt"
	"time"
)

func main() {
	rb := time.Date(1978, 2, 3, 0, 0, 0, 0, time.UTC)
	bb := time.Date(2006, 12, 28, 0, 0, 0, 0, time.UTC)

	diff := rb.Sub(bb)
	fmt.Printf("%f\n", diff.Minutes())
}
