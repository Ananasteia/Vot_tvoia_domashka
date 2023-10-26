package main

import (
	"fmt"
	"os"
)

func main() {
	er := mistake()
	if er != nil {
		fmt.Println(er)
	}
}
func mistake() error {
	_, err := os.Open("filename.net")
	return err
}
