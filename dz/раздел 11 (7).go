package main

import "fmt"

// Тут я чуть передалала задание, чтобы провал был только когда значение четное и больше пяти
func main() {
	week := []int{1, 2, 3, 4, 5, 6, 7}
	for _, elem := range week {
		switch elem % 2 {
		case 0:
			fmt.Println("четное значение")
			if elem < 5 {
				continue
			}
			fallthrough
		default: // такого быть не может, так что будем брать для провала
			fmt.Println("еще и больше пяти")
		case 1:
			fmt.Println("нечетное значение")
		}
	}
}
