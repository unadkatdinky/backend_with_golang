package main 

import "fmt" 

func main() {
	prices := map[string]int {
		"xyz" : 500, 
		"def" : 1800,
	}

	total := 0 
	for item, price := range prices {
		fmt.Println(item, price)
		total = total + price
	}

	fmt.Println(total)
}
