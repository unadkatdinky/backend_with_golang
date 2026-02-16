package main

import (
	"fmt" 
	
)

func main() {
	score := 72 

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 75 {
		fmt.Println("B")
	} else if score >= 45 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

	// if with short statement
	items := 3;
	pricePerItem := 49 
	if total := items * pricePerItem; total >= 100 {
	//if with short statement
	//we are creating total and comparing the if statement in. the same line
	fmt.Println("Not eligible") } else {
		fmt.Println("Not Eligible for shipping")
	}
	// the closing bracket } and else should be in the same line

	for i := 1; i<=5; i++ {
		fmt.Println(i)
	}

	N := 10 
	sum := 0

	for i := 0; i < N; i++ {
		sum++
	}
	fmt.Println(sum)

	day := 3

	switch day {
	case 1: 
		fmt.Println("monday")
	case 2: 
		fmt.Println("tuesday")
	case 3: 
		fmt.Println("wednesday")
	default:
		fmt.Println("other")
	}


}
