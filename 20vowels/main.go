package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Abhishek"
	vowels := ""
	count := 0

	for _, char := range input {

		switch strings.ToLower(string(char)) {

		case "a", "e", "i", "o", "u":

			vowels += string(char)
			count++

		}
	}
	fmt.Println("vowels are", vowels)
	fmt.Println("count is", count)
}
