package main

import "fmt"

func main() {
	Name := "Abhishek"
	var ReverseName string
	for i := 0; i < len(Name); i++ {

		ReverseName = string(Name[i]) + ReverseName

	}
	fmt.Println("ReverseName Is:", ReverseName)
}
