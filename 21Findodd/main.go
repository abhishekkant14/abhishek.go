package main

import "fmt"

func main() {
	numbers:=[]int {2, 4, 6, 8, 3, 7}
	var oddnumber []int
  for _,num:= range numbers {

	if num%2 !=0 {

oddnumber=append(oddnumber, num)

	}
}
fmt.Println("Oddnumer is:",oddnumber)
}