package funciones

import "fmt"

func Exponencia(numero int) {
	fmt.Println(numero)
	if numero > 1000 {
		return
	}
	Exponencia(numero * 2)
}
