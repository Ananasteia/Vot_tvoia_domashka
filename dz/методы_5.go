package main

import "fmt"

// Я ПАНК! И ДА, Я ТАК МОГУ, А ЧТО ВЫ МНЕ СДЕЛАЕТЕ????
type JopaPidora interface {
	Minutes() Seconds
}
type Seconds int

func (f Seconds) Minutes() Seconds {
	return f * 60
}
func main() {
	s := Seconds(0)
	fmt.Scan(&s)
	j := JopaPidora(s)
	fmt.Print(j.Minutes())
}
