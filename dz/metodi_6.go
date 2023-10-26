package main

import "fmt"

type Person struct {
	Name string
	Adress
}
type Adress struct {
	Street  string
	City    string
	Country string
}

func (q Person) FullAddress() string {
	return q.Street + q.City + q.Country
}
func main() {
	nonGrata := Person{"Ricardo", Adress{"Pushkina", "MuchoSransk", "Atlantida"}}
	w := nonGrata.FullAddress()
	fmt.Print(w)
}
