package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Println(div(x, y))
}
func div(x int, y int) (int, error) {
	if x == 0 || y == 0 {
		return 0, fmt.Errorf("Деление нулей запрещено")
	}
	return x / y, nil
}
