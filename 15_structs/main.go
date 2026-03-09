package main

import "fmt"

type User struct {
	//struct groups related fields to one type
	//struct fields are mutable
	//you can also make partial structs where all fields are not defined
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	u1 := User{
		ID:    1,
		Name:  "Dinky",
		Email: "dinkyunadkat02@gmail.com",
		Age:   23,
	}

	fmt.Println(u1)
	fmt.Println(u1.Age)

}