package main

import (
	"fmt"
)
func main () {

	for i:=0 ; i<2 ; i++ {
	fmt.Println("Number is",i)
}

counter:=0
for {
	fmt.Println("Infinity Number:")
	counter++ 
	if counter == 2{
break
	}
}
numbers:= []int{24,14}
for index, value:= range numbers {

fmt.Printf("Index of Number %d, and value is %d\n", index,value)

}
data:="Hello Abhishek"
for index, char := range data {
fmt.Printf("Index of Number %d, and value is %c\n", index,char) 

}
	
}

