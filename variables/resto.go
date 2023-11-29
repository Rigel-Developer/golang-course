package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	// asignacion de variables
	Nombre = "Rigel"
	Estado = true
	Sueldo = 5000.54
	Fecha = time.Now()

	//Impresion de variables
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha.Local().Format("02-01-2006 15:04:05"))

	// MuestraEnteros()
}

func ConviertoATexto(numero int) (bool, string) {
	// texto := fmt.Sprintf("%d", numero)
	texto := strconv.Itoa(numero)
	return true, texto
}
