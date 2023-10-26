package main

import "fmt"

func main() {
	week := []bool{true, false, false, false, true}
	for _, elem := range week {
		switch elem {
		case true:
			fmt.Println("тру")
		default:
			fmt.Println("фолс")
		}
	}
}
