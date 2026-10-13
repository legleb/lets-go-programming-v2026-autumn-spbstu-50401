package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Error reading departments count")
		return
	}
	for i := 0; i < n; i++ {
		var k int
		if _, err := fmt.Scan(&k); err != nil {
			fmt.Println("Error reading employee count")
			return
		}
	}
}
