package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	//fmt.Println(currentTime)
	//fmt.Printf("Time of currentTime is %T\n", currentTime)
	format := currentTime.Format("01-02-2006 03:04:05 PM Monday")
	fmt.Println(format)


layout_str :="01-02-2006 15:04:05 Monday"
datestr := "06-07-1993 11:04:05 Monday"
formatted_time, _:=time.Parse(layout_str, datestr)
fmt.Println("formate of date :",formatted_time)

}
