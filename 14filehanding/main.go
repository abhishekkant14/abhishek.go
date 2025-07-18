package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("Hello Abhishek")
	if err != nil {
	defer	fmt.Println("Err While Creating", err)

		return
	}
	
	 fmt.Println("file creation done")

	//while writing file
	data := "Hello Abhishek Welocome to the Go Lang"
	_,eror := io.WriteString(file,data + "\n")
	if eror != nil {

		fmt.Println("Err while write ")
	}

}
