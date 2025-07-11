package main

import "fmt"

func main() {

	//var ptr int
	//fmt.Println("The vale of ptr", ptr)

	number := 60
	
	var ptr = &number

	fmt.Println("tHe Value of Number", ptr)
    fmt.Println("tHe Value of Number", *ptr)

}
