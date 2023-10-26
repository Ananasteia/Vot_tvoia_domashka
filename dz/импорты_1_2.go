package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// 1 ЗАДАЧА
	x := 23.0
	fmt.Printf("%.1f\n", math.Sqrt(x))
	// 2 ЗАДАЧА
	y := time.Now().Second()
	z := float64(y)
	fmt.Printf("%.1f", math.Sqrt(z))
}
