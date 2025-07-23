package main

import "fmt"

func add(a int, b int) int {
	return a + b}
	
	func multiply(x int, y int ) (result int) {
 result =x * y 
return }

func main() {
	fmt.Println("Hello, we are learning Functions in Go Lang")
	result := add(5, 7)
	fmt.Println("Sum is:", result)

	output := multiply(7, 8)
	fmt.Println("output is",output)
}