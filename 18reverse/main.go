package main

import "fmt"

func main() {
	Name := "ABHISHEK"
	var Reverse string

	for i := 0; i < len(Name); i++ {
		Reverse = string(Name[i]) + Reverse
	}

	fmt.Println("Reverse Name Is:", Reverse)

}
