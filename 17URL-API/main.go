package main

import (
	"fmt"
	"net/url"
)

func main() {
	MyURL := "https://www.youtube.com/watch?v=oURCNFbGR78&list=PPSV&t=339s"
	paseurl, err := url.Parse(MyURL)
	if err != nil {
		fmt.Println("can't pass URL", err)
		return
	}
	fmt.Println("Scheme:", paseurl.Scheme)
	fmt.Println("Host:", paseurl.Host)
	fmt.Println("Path:", paseurl.Path)
	fmt.Println("RawQuery:",paseurl.RawQuery)
}
