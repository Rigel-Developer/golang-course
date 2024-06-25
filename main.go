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

	// arreglos_slices.MuestroSlice()

	// mapas.MuestroMapas()

	// users.AltaUsuario()

	// pedro := new(modelos.Hombre)
	// pedro.Vivo = false
	// camila := new(modelos.Mujer)
	// camila.Vivo = true
	// ejer_interfaces.HumanoRespirando(pedro)

	// deferkeyword.VemosPanic()

	//*rutines and channels
	// canal1 := make(chan bool)
	// go goroutines.MiNombreLentooo("Rigel Developer", canal1)
	// fmt.Println("Estoy aqui")
	// defer func() {
	// 	res := <-canal1
	// 	fmt.Println(res)
	// }()
	goroutines.CompleteTask()
	// webserver.MiWebServer()

	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
	)

}

// SumIntsOrFloats calculates the sum of integers or floats in a map.
// The key type K must be comparable, and the value type V must be either int64 or float64.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
