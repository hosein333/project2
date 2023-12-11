package main

import ( 
     "fmt"
     "time"
)
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

	cnt := 0

	for {
	cnt++
	fmt.Printf("%d Infinite Loop ..., cnt")
	time.Sleep(1 * time.Second)
	}

	fmt.Println("After Loops")
}
