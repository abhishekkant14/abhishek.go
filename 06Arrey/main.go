package main

import "fmt"

func main() {
	var fruitlist [100]string
	fruitlist[0]="Mango"
	fruitlist[1]="Graps"
    fruitlist[10]="Lichi"
    fruitlist[25]="Guvawa"
    fruitlist[50]="Orang"
    fruitlist[75]="banana"
    fruitlist[99]="Jackfruit"
fmt.Println(len(fruitlist))

}
