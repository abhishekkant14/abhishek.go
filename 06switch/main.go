package main

import (
	"fmt"
	"time"
)

func main() {

	i := 5
	switch i {

	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Foure")

     default :
    fmt.Println("Other")
	}

//multiple condition

switch time.Now().Weekday() {

case time.Sunday, time.Saturday:
	fmt.Println("it's weekend")
default:
	fmt.Println("workday")
}

//type switch

iam:=func(i interface{}) {

	switch t:=i.(type) {

	case int:
	fmt.Println("integer")
	case string:
		fmt.Println("String",t)
	}
}
iam("AAbhishek")
}