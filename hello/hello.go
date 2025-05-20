package main

import (
	"fmt"

	"example/greetings"
)

func main() {

	message, err := greetings.Hello("")

	if err != nil {

		fmt.Println(err.Error())
	}

	fmt.Println(message)

}
