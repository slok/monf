package main

import (
	"fmt"

	"github.com/slok/monf/configuration"
)

func main() {
	fmt.Println("Hello world!")
	configuration.LoadSettings("")
}
