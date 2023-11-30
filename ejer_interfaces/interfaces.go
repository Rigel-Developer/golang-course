package ejer_interfaces

import (
	"fmt"
	"github/rigel-developer/golang-course/interfaces"
)

func HumanoRespirando(h interfaces.Humano) {
	h.Respirar()
	fmt.Printf("Soy un %s,  estoy respirando  y estoy vivo (%t)\n", h.Sexo(), h.EstaVivo())
}
