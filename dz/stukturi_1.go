package main

import "fmt"

func main() {
	type Book struct {
		Title  string
		Author string
		Pages  int
	}
	p := Book{"Название", "Писака", 45}
	fmt.Println(p.Title, p.Author, p.Pages)
}
