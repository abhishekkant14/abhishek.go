package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Create("Abhishek,txt")
	if err != nil {
		fmt.Println("While creating a file", err)
		return
	}
	defer file.Close()
	contant := "Hello @bhishek what You think about Go Lang"
	_, error := io.WriteString(file, contant+"\n")
	if error != nil {
		fmt.Println("Error while writing", error)
		return
	}
	fmt.Println("File creation sucessful done")

	buffer := make([]byte, 1024)
	for {
		n, errors := file.Read(buffer)
		if errors == io.EOF {
			break
		}
		if errors != nil {
			fmt.Println("errors reading file", errors)
			break
		}

		fmt.Println(string(buffer[:n]))
	}
	file, errrs := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if errrs != nil {
		fmt.Println("Error opening file for append:", errrs)
		return
	}
	defer file.Close()

	// Append text
	_, err = file.WriteString("\nAppended line.")
	if err != nil {
		fmt.Println("Error appending to file:", err)

		file, err := os.Open("example.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text()) // Print each line
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Scanner error:", err)
		}
	}
}
