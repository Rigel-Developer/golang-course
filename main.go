package main

import (
	"fmt"
	"github/rigel-developer/golang-course/goroutines"
)

func main() {
	//first example
	// variables.MuestraEnteros()

	// estado, texto := variables.ConviertoATexto(123)
	// fmt.Println(estado, texto)

	// if os := runtime.GOOS; os == "windows" {
	// 	fmt.Println("El sistema operativo es: " + os)
	// } else {
	// 	fmt.Println("El sistema operativo no es windows, es: " + os)
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("El sistema operativo es: " + os)
	// case "windows":
	// 	fmt.Println("El sistema operativo es: " + os)
	// case "darwin":
	// 	fmt.Println("El sistema operativo es: " + os)
	// default:
	// 	fmt.Println("El sistema operativo no es windows, es: " + os)
	// }

	// number, mensaje := ejercicios.Ejercicio01("1000")
	// fmt.Println(number, mensaje)

	// teclado.IngresoNumeros()

	// iteraciones.Iterar()

	// fmt.Println(ejercicios.MultiplicationTable())

	// files.GrabarTable()
	// files.SumaTabla()
	// files.LeerArchivo()

	// funciones.CalculosMatematicos()

	// funciones.Closures()

	// funciones.Exponencia(10)

	// arreglos_slices.CapacidadSlice()

	// mapas.MuestroMapas()

	// users.AltaUsuario()

	// pedro := new(modelos.Hombre)
	// pedro.Vivo = false
	// camila := new(modelos.Mujer)
	// camila.Vivo = true
	// ejer_interfaces.HumanoRespirando(pedro)

	// deferkeyword.VemosPanic()

	canal1 := make(chan bool)
	go goroutines.MiNombreLentooo("Rigel Developer", canal1)
	fmt.Println("Estoy aqui")
	defer func() {
		res := <-canal1
		fmt.Println(res)
	}()

}
