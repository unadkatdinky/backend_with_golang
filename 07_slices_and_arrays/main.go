package main 

import "fmt"
func main() {
//in go array is always of fixed size, ordered and sequence of elements 
//of the same type

var marks [3] int 
//fixed, cannot grow 
//when creating project we wouldnt prefer a defined size
marks[0] = 1 
marks[1] = 2
marks[2] = 3

fmt.Println(marks)
res := [5]int{2,3,4,5,6}

fmt.Println(len(res))

//they may ask difference between arrays and slices
//why we prefer slices

//slices are the most common collection type
//dynamic and can grow
//[]type{}
//powerful and flexible data types used to work with 
//sequences of typed data
//can grow or shrink 

results := []string{"sangam", "dinky"}
fmt.Println(results, results[0], results[len(results)-1] )
results[1] = "priya"
fmt.Println(results)
//you can access these by indexes
results = append(results, "ayushi ")


}