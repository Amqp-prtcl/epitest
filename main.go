package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")

	var a **string = new(*string)

	fmt.Println(*a)
}
