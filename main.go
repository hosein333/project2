package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	sum := 0

	for i:= 0; i<5 ; i++ {
		sum += i 
	}
	fmt.Println(sum)
}
