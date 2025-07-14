package main

import "fmt"
func main() {

evenCount:=0
oddnumber:=0
for i:=1; i<=10; i++ {
	switch i % 2 {
	case 0:
	
		fmt.Println(i,"evenCount")
		evenCount++

	case 1:
		fmt.Println(i,"oddnumber")
		oddnumber++


	}
	fmt.Println("Number of 9 to 15 is:")
	for i:=9; i<=15 ; i++{
		fmt.Println(i,"Number")
	}

}
fmt.Println("Total EvenNumber is",evenCount)
fmt.Println("Total OddNumber is",oddnumber)
}