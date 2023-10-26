package main

import "fmt"

func main() {
	week := []int{1, 2, 3, 4, 5, 6, 7}
	for _, elem := range week {
		switch elem % 2 {
		case 0:
			fmt.Println("Четное")
		case 1:
			fmt.Println("Нечетное")
		}
	}
}
