package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	sum := 0

	for i:= 0; i<5 ; i++ {
		sum += i 
	}
	fmt.Println(sum)

	n :=1
	for n <5{
	n *= 2
	}
	for {
	fmt.Println("Infinite Loop ...")
	}
	fmt.Println(n)
}
