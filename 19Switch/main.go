package main

import "fmt"
func main() {

	evennumber:=0
	oddnumber:=0
	for i:=1; i<=10; i++ {
    switch i % 2 {
	case 0:
		fmt.Println(i,"Even")
		evennumber++
		
	case 1:
		fmt.Println(i,"oddnumber")
		oddnumber++

	}

	}
	fmt.Println("Evennumber Is:",evennumber)
	fmt.Println("Oddnumber is:", oddnumber)
}