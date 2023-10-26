package main

import "fmt"

func main() {
	a := make([]int, 3, 10)
	for i, _ := range a {
		fmt.Scan(&a[i])
		if a[i] > 0 {
			fmt.Printf(" Положительное\n")
		} else {
			fmt.Printf(" Отрицательное\n")
		}
		if a[i] > 30 {
			fmt.Printf(" Больше 30\n")
		} else if a[i] > 20 {
			fmt.Printf(" Больше 20\n")
		} else if a[i] > 10 {
			fmt.Printf(" Больше 10\n")
		} else {
			fmt.Printf(" Меньше 10\n")
		}
	}
}
