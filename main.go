package main

import (
	"fmt"
	"github/rigel-developer/golang-course/variables"
)

func main() {

	// variables.MuestraEnteros()
	estado, texto := variables.ConviertoATexto(123)
	fmt.Println(estado, texto)

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
