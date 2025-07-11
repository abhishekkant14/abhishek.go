package main

import "fmt"

func main() {
// Declaring variables using var
var name string ="Abhishek"
var Age int =31
var isStudents bool = true
var height float64 =23.56767

fmt.Println("Name is: ",name)
fmt.Println("Age is: ",Age)
fmt.Println("isStudents is: ",isStudents)
fmt.Println("height is: ",height)
fmt.Printf("Name is:%s, Age is:%d, height is:%.3f\n",name, Age, height)

	var a int
	var b string
	var c float64
	var d bool

	fmt.Println("Default int:", a)
	fmt.Println("Default string:", b)
	fmt.Println("Default float:", c)
	fmt.Println("Default bool:", d)
}

