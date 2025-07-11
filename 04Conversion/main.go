package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {

	fmt.Println("Hello Welcome to Our Pizza Shop")
	fmt.Println("Please share rating beetween 1 to 5")
	reader :=bufio.NewReader(os.Stdin)
	input,_ :=reader.ReadString('\n')
    fmt.Println("THANKS FOR RATING", input)

}
