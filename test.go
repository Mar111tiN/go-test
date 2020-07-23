package main

import (
	"fmt"

	"github.com/Mar111tiN/go-test/testpackage"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Print(testpackage.ReverseRunes("Hello."))
}
