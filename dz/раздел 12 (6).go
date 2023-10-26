package main

import (
	"fmt"
	"os"
)

type MyError struct {
	Msg  string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Название файла: %s, линия с ошибкой: %d, текст ошибки: %s", e.File, e.Line, e.Msg)
}

func main() {
	er := err()
	fmt.Println(er)
}
func err() error {
	_, err := os.Open("filename.net")
	if err != nil {
		e := err.Error()
		return &MyError{e, "filename.net", 26}
	}
	return nil
}

//это было ужасно, просто ужасно
