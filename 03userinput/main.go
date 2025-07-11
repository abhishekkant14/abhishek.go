package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey What's Your Name")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
		age := 31

	fmt.Println("Hello Mr.:", input,"And age is:", age)



}
