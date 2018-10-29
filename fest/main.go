package main

import (
	"fmt"

	"github.com/juanrgon/prism"
)

func main() {
	happy := prism.InRed("Happy")
	holidays := prism.InGreen("Holidays")
	fmt.Println(happy + " " + holidays)
}
