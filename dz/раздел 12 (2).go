package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errorss()
	if err != nil {
		fmt.Println(err)
	}
}
func errorss() error {
	var err = errors.New("Файл не удалось открыть")
	return err
}
