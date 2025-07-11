package main

import (
	"fmt"
	"sort"
)

func main() {

	var studentsinclass = []string{"OM", "NAMO", "SHIV", "RAM"}
    fmt.Println(studentsinclass)
    studentsinclass = append(studentsinclass [0:])
    fmt.Println(studentsinclass)
   studentsinclass =append(studentsinclass, "Ravan")
   fmt.Println(studentsinclass)

   fmt.Println("Total vlaue of studentsinclass", len(studentsinclass))

//2nd methode

MyHighschoolResult := make([]int, 6)
MyHighschoolResult[0]= 645
MyHighschoolResult[1]= 334
MyHighschoolResult[2]= 765
MyHighschoolResult[3]= 555
MyHighschoolResult[4]= 455
MyHighschoolResult[5]= 123
fmt.Println(MyHighschoolResult)
sort.Ints(MyHighschoolResult)
fmt.Println(MyHighschoolResult)
fmt.Println(sort.IntsAreSorted(MyHighschoolResult))


}