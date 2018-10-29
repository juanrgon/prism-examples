package main

import (
	"fmt"

	"github.com/juanrgon/prism"
)

func main() {
	r := prism.InRed("Red").InBold()
	w := prism.InWhite("White").InBold()
	b := prism.InBlue("Blue").InBold()

	sentence := r + ", " +  w + ", and " + b
	fmt.Println(sentence)
}
