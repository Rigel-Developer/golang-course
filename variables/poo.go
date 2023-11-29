package variables

import "fmt"

// Car es una estructura publica al igual que sus atributos
type Car struct {
	Brand string
	Year  int
}

func Structures() {

	//Modificadores de acceso en Go funciones, variables, estructuras
	// publico: se inicia con mayuscula
	// privado: se inicia con minuscula

	myCar := Car{Brand: "Ford", Year: 2020}
	fmt.Println(myCar)
}
