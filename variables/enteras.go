package variables

import "fmt"

func MuestraEnteros() {
	fmt.Println("Enteros")
	//Declaracion
	var number int8 = -128
	// Asignacion
	otherNumber := 23

	fmt.Printf("Number: %d, OtherNumber: %d\n", number, otherNumber)
	fmt.Printf("Number: %T, OtherNumber: %T\n", number, otherNumber)
}

func Suma(a, b int) int {
	return a + b
}

func Ciclos() {
	//For tradicional
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	//For while
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	//For forever
	// i := 0
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }

}

func KeyWords() {
	//defer -> se ejecuta al final de la funcion
	//continuo -> salta a la siguiente iteracion
	//break -> sale del ciclo
	fmt.Println("Hola")
	defer fmt.Println("Mundo")
	fmt.Println("Hola de nuevo")

}
