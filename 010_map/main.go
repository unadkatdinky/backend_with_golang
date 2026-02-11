package main 

import "fmt" 

func main() {
	ages := map[string]int{
		"dinky" : 23,
		"minky" : 25,
		"pinky" : 27,
	}
	fmt.Println(ages["minky"])

	var siblings map[string]int
	fmt.Println(siblings, siblings["dinky"])
	delete(ages, "minky")
	delete(ages, "rinky")
	fmt.Println(ages)
}