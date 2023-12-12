package main

import "fmt"

func main() {

	x := 25

	if x%5 ==0 || x < 100 {
		fmt.Printf("%d is multiple of 5\n", x)
	}
}
