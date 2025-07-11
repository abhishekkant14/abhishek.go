package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func main() {
res,err:=http.Get("https://jsonplaceholder.typicode.com/todos/1")
if err !=nil {
fmt.Println("Err while responding", err)
return
}
defer res.Body.Close()

//reading the file
data,err:=ioutil.ReadAll(res.Body)
if err !=nil{
fmt.Println("while reading file", err)
return
}

fmt.Println("responce", string(data))
}
