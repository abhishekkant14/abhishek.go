package main

import "fmt"
func main() {
Name:="Abhishek"
var reverse string

for i:=0; i<len(Name); i++ {

reverse=string(Name[i]) + reverse 

}
fmt.Println("reverse IS", reverse)
}