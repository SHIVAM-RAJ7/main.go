package main

import (
	"fmt"
	"hello/greetings"
)

func main() {
	message := greetings.Hello("Shivam")
	fmt.Println(message)
}
