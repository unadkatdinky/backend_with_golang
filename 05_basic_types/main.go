package main

import(
	"fmt"
	"strings"
)

func main() {
	firstName := "dinky"
	secondName := "unadkat"
	fullName := firstName + " " + secondName 
	fmt.Println(fullName)
	fmt.Println(strings.ToUpper(fullName))
	views1 := 1000
	views2:= 200
	totalViews := views1 + views2
	//you can also do /, %, ++ and so on
	//same works for float
	fmt.Println(totalViews)
	//for booleans we have bool

	const appName = "Go Basics"
	//typed constants 
	const maxUpload int = 25
	const discountPrice float64 = 10.3
	



	 
}