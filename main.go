package main

import "fmt"

func main() {
	x := 25

	if x%5 ==0 && x != 25 {
		fmt.Printf("%d is multiple of 5\n", x)
	} else {
                fmt.Printf("%d is not multiple of 5  or equal to 25\n", x)
	}

}
