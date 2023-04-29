package main

import (
	"fmt"
	"github.com/rigel-developer/golang-course/variables"
)

func main() {
	fmt.Println("Hello, World!")
	variables.MuestraEnteros()
	// una suma
	fmt.Println("Suma:", variables.Suma(2, 3))
}
