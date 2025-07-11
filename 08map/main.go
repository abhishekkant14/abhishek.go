package main

import "fmt"

func main() {
    //  Declare the map
    MyHighSchoolSubject := map[string]int{
        "Math":    90,
        "Science": 85,
        "English": 88,
        "History": 75,
    }

fmt.Println("MyHighSchoolSubject", MyHighSchoolSubject)
    //Loop through the map

    for key, value := range MyHighSchoolSubject {
       fmt.Printf("For key %v, value is %v\n", key, value)
   }
}