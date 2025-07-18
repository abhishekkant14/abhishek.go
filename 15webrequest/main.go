package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
url:= "https://jsonplaceholder.typicode.com/posts/1"

res,err:=http.Get(url)
if err!=nil {
fmt.Println("Reading while err", err)
return

}
defer res.Body.Close()

data,err:=ioutil.ReadAll(res.Body)
if err!= nil {
fmt.Println("Err while reading time",err)
return
}
fmt.Println("Responce")
fmt.Println("Reading data",string(data))
}