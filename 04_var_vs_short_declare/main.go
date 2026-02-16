package main

import (
	"fmt"
)

//inference:

func main() {
	var city string
	city = "London"

	var channel = "sangam"
	//this is inferrence : it automatically infers it as string
	//it assumes the type
	//shortcut is done using :=
	subscribers := 5000 
	subscribers = subscribers + 1000
	likes, comments := 100, 30
	fmt.Println (city, channel, subscribers, likes, comments)

}