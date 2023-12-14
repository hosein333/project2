package main

import "fmt"

func main() {
	checkdayOfWeek(2)
	checkdayOfWeek(3)
	checkdayOfWeek(9)
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
	switch day {
	case 1: 
	fmt.Println("Monday")
	case 2: 
	fmt.Println("Tuesday")
	case 3: 
	fmt.Println("Wednesday")
	case 4: 
	fmt.Println("Thursday")
	case 5: 
	fmt.Println("Friday")
	case 6: 
	fmt.Println("Saturday")
	case 7: 
	fmt.Println("Sunday")
	default: 
	fmt.Println("Invalid day")
}

}
