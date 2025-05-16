package main

import (
	"fmt"

	"example/greetings"
)

func main() {

	message := greetings.Hello("tristan")

	fmt.Printf(message)

}
