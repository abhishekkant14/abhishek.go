package main

import "fmt"

func main() {
	Name := "Abhishek"

	var str string
	for i := 0; i < len(Name); i++ {
		str = string(Name[i]) + str
		fmt.Println("ReverseName is", str)
	}
}
