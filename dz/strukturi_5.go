package main

import "fmt"

type Cup struct {
	color string
	drink string
	ml    int
}

func main() {
	one := Cup{"blue", "tea", 450}
	two := Cup{"red", "tea", 200}
	fmt.Println(one.color == two.color)
	fmt.Println(one.drink == two.drink)
	fmt.Println(one.ml == two.ml)
	fmt.Println(one == two)
}
