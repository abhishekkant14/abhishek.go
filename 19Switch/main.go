package main

import "fmt"

func main() {
	sum := 0

	for i := 1; i <= 10; i++ {
		switch i {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
			sum += i
		default:
		}
	}

	fmt.Println("Sum using switch case from 1 to 10 is:", sum)
}