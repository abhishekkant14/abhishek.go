package main

import (
	"fmt"
	"strings"
)
func main() {
input:="Abhishek"
vowels:=""
count:=0

for _,char:=range input {
switch strings.ToLower(string(char)) {
case "a", "e", "i", "o", "u":
	vowels+=string(char)
	count++

}

}
fmt.Println("Vowels are:", vowels)
fmt.Println("Count:", count)
}
