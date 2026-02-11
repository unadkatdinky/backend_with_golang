package main 

import "fmt"

func main() {
	var marks [5] int 
	marks[0] = 90
	marks[1] = 80		
	marks[2] = 70
	marks[3] = 60
	marks[4] = 50
	fmt.Println(marks)

	//array literal 
	results := [5] int {90, 80, 70, 60, 50}
	fmt.Println(results)
	fmt.Println(len(results))

	res := []string {"apple", "banana", "grape"}
	fmt.Println(res)

	res = append(res, "orange")
	fmt.Println(res)
	fmt.Println(len(res))
	fmt.Println(cap(res))

	scores := make( []int, 3,5)
	// scores[0] = 90
	// scores[1] = 80
	// scores[2] = 70
	fmt.Println(scores)
	fmt.Println(len(scores))
	fmt.Println(cap(scores))

	todos := []string {"clean the house", "do the laundry", "buy groceries"}

	more := [] string {"pay bills", "call mom"}
	todos = append(todos, more...)
	fmt.Println(todos)
}