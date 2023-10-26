package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() {
	fmt.Println(r.Width * r.Height)
}
func (s Rectangle) IsSquare() bool {
	return s.Width == s.Height
}
func (x *Rectangle) DoubleSize() { //мне плевать, что так нельзя, льзя, раз я умею
	x.Width = x.Width * 2
	x.Height = x.Height * 2
}
func main() {
	big := Rectangle{6, 5}
	small := Rectangle{4, 5}
	big.Area()
	small.Area()
	fmt.Printf("Являтся ли квадратом? %v, %v\n", big.IsSquare(), small.IsSquare())
	abcd := Rectangle{8, 8}
	abcd.DoubleSize()
	fmt.Print(abcd)
}
