package main

import (
	"fmt"
	"io"
	"os"
)
func main() {
//while creating file
	file,err:=os.Create("Abhishek.text")
	if err !=nil {
		fmt.Println("err while creatin file", err)
		return
	
	}
	defer file.Close()

  fmt.Println("file creation sucessful done")

//while Writing file
contant := "Hello @abhishek welocome to Go Lang"
_,errors:=io.WriteString(file,contant+"\n")
if errors !=nil {
fmt.Println("Errors while writing file",errors)
}
}