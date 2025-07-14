package main

import "fmt"
func main () {

	evenNumber:=0
	oddNumber:=0
	for i:=1; i<=10; i++ {
		switch i % 2 {
		case 0:
			fmt.Println(i,"even")
			evenNumber++

		case 1:
			fmt.Println(i,"odd")
			oddNumber++
		}
	}
	fmt.Println("Total evenumber is:",evenNumber)
	fmt.Println("Total oddNumber is:",oddNumber)
}