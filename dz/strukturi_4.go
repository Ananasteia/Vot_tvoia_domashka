package main

import "fmt"

type Job struct {
	Title  string
	Salary int
}
type Employee struct {
	Name string
	Job  Job
}

func main() {

	nekto := Employee{"Аноним", Job{"Безработный", 0}}
	fmt.Println(nekto.Name, nekto.Job)
}
