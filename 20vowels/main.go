package main

import (
	"fmt"
	"strings"
)
func main () {

	Input:="Abhishek"
	vowels:=""
	count:=0

for _,char:=range Input {
switch strings.ToLower(string(char)) {

case "a", "e", "i", "o", "u":
   vowels += string(char) 
	count++
}
}

fmt.Println("Vowels are:", vowels)
	fmt.Println("Count:", count)
}
