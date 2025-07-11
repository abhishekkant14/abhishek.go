package main
import "fmt"

func main() {

	fmt.Println("structs in Go lang")
	abhishek:= User {"abhishek", "Abhishekkanmt@gmail.com", 29, true}
fmt.Printf("abhishek details is: %+v\n", abhishek)
fmt.Printf("Name is %v and email %v.",abhishek.Name, abhishek.Email)
}

type User struct {
 Name string
 Email string
 age int
 status bool
}