package main

import "fmt"

func readSignAndTemp() (string, int, error) {
	var sign string
	if _, err := fmt.Scan(&sign); err != nil {
		return "", 0, fmt.Errorf("scan sign: %w", err)
	}

	var temp int
	if _, err := fmt.Scan(&temp); err != nil {
		return "", 0, fmt.Errorf("scan temp: %w", err)
	}

	return sign, temp, nil
}

func main() {
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println(err)

		return
	}

	for range departments {
		var count int
		if _, err := fmt.Scan(&count); err != nil {
			fmt.Println(err)

			return
		}

		low, high := 15, 30

		for range count {
			sign, temp, err := readSignAndTemp()
			if err != nil {
				fmt.Println(err)

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
				fmt.Println("Wrong sign")

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
