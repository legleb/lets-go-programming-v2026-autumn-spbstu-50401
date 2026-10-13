package main

import "fmt"

func main() {
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println("Error reading departments count")

		return
	}

	for range departments {
		var count int
		if _, err := fmt.Scan(&count); err != nil {
			fmt.Println("Error reading employee count")

			return
		}

		low, high := 15, 30

		for range count {
			var sign string
			var temp int

			if _, err := fmt.Scanf("%c= %d", &sign, &temp); err != nil {
				fmt.Println("Error reading temperature")

				return
			}

			switch sign {
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
