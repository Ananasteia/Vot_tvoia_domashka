package main

import "fmt"

func main() {
	type Person struct {
		Name      string
		Age       int
		isMarried bool
	}
	p := Person{"Шишел Мышел", 34, true}
	fmt.Println(p.Name, p.Age, p.isMarried)
	//теперь изменяю
	p.Age = 44
	fmt.Println(p.Name, p.Age, p.isMarried)
}
