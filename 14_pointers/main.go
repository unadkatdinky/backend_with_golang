package main 

import "fmt"

func main() {
	//pointer: stores the memory address of any val 
	// &x - address of x (makes a pointer)
	//dereference *p: go to the address and read/write

	//why? change a value inside a function without returning it 

	score := 10 
	fmt.Println("before" , score)

	addScore(&score)
	fmt.Println("after", score )
}

func addScore(score *int) {
	*score = *score + 5
}