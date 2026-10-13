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
		low, high := 15, 30
		for j := 0; j < k; j++ {
			var s string
			if _, err := fmt.Scan(&s); err != nil {
				fmt.Println("Error reading sign")
				return
			}
			var temp int
			if _, err := fmt.Scan(&temp); err != nil {
				fmt.Println("Error reading temperature")
				return
			}
			switch s {
			case ">=":
				if temp > low {
					low = temp
				}
			case "<=":
				if temp < high {
					high = temp
				}
			default:
				fmt.Println("Неверная операция ограничения температуры")
				return
			}
			if low > high {
				fmt.Println(-1)
			} else {
				fmt.Println(low)
			}
		}
	}
}
