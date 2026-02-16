package main 

import "fmt" 

func main() {
	points:= map[string]int {
		"a":10, 
		"b":0,
	}
	fmt.Println("a", points["a"])
	fmt.Println("b", points["b"])
	fmt.Println("c", points["c"])

	//this gives out put 10, 0, 0 and this can 
	//turn out to be confusing because b is valid 
	//and c is not yet it gives the same output 

	//hence we use value ok 

	valB, okB := points["b"]
	fmt.Println(valB, okB)
	valC, okC := points["c"]
	fmt.Println(valC, okC)

	if val,ok := points["c"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("c key is not present")
	} 

}
