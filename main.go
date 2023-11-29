package main

import (
	"fmt"
	"github/rigel-developer/golang-course/ejercicios"
	"runtime"
)

func main() {
	//first example
	// variables.MuestraEnteros()

	// estado, texto := variables.ConviertoATexto(123)
	// fmt.Println(estado, texto)

	if os := runtime.GOOS; os == "windows" {
		fmt.Println("El sistema operativo es: " + os)
	} else {
		fmt.Println("El sistema operativo no es windows, es: " + os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("El sistema operativo es: " + os)
	case "windows":
		fmt.Println("El sistema operativo es: " + os)
	case "darwin":
		fmt.Println("El sistema operativo es: " + os)
	default:
		fmt.Println("El sistema operativo no es windows, es: " + os)
	}

	number, mensaje := ejercicios.Ejercicio01("fff")
	fmt.Println(number, mensaje)
	// Define el enrutador
	// router := http.NewServeMux()

	// // Ruta de ejemplo
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println(r)
	// 	fmt.Fprintln(w, "¡Hola desde el servidor de desarrollo!")
	// })

	// // Inicia el servidor en el puerto 8000
	// log.Println("Servidor en ejecución en http://localhost:8000")
	// log.Fatal(http.ListenAndServe(":8000", router))

}
