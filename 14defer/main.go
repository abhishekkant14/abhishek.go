package main

import "fmt"

func add(a,b int) int {

	return a+b
} 
func main() {

	fmt.Println("Hello Welcome To GoLang")
	data := add (3,7)
	defer fmt.Println("Data is",data)
	fmt.Println("My Frist Program is Hello word")
	defer fmt.Println("My Secound Program is variables")
	fmt.Println("My Last Program is not close Yet")
}