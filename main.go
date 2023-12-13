package main

import "fmt"

func main() {
	checkdayOfWeek(1)
}

func chechMultiple(x int) {

	if x%2 ==0 {
		fmt.Printf("%d is multiple of 2\n", x)
	}
	if x%3 ==0 {
                fmt.Printf("%d is multiple of 3\n", x)
	}
	if x%4 ==0 {
                fmt.Printf("%d is multiple of 4\n", x)
	}
	if x%5 ==0 {
                fmt.Printf("%d is multiple of 5\n", x)
	}
	if x%6 ==0 {
                fmt.Printf("%d is multiple of 6\n", x)
	}
	if x%7 ==0 {
                fmt.Printf("%d is multiple of 7\n", x)
	}
	if x%8 ==0 {
                fmt.Printf("%d is multiple of 8\n", x)
	}
	if x%9 ==0 {
                fmt.Printf("%d is multiple of 9\n", x)
	}


}



func checkdayOfWeek(day int) {
	if day == 1 {
		fmt.Println("Saturday")
	}
	if day == 2 {
                fmt.Println("Sunday")
        }
	if day == 3 {
                fmt.Println("Monday")
	}
	if day == 4 {
                fmt.Println("Thursday")
	}
	if day == 5 {
                fmt.Println("Wednesday")
	}
	if day == 6 {
                fmt.Println("Thursday")
	}
	if day == 7 {
                fmt.Println("Friday")
	}
}
