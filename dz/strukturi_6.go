package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {
	p := Book{"Таинственный сад", "Фрэнсис Элизы Бёрнетт", 320}
	b := Book{"Глаз волка", "Даниэль Пеннак", 96}
	c := bolsheStranits(p, b)
	fmt.Println(c)
}
func bolsheStranits(p Book, b Book) Book {
	if p.Pages > b.Pages {
		return p
	} else {
		return b
	}
}
