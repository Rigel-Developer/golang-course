package files

import (
	"bufio"
	"fmt"
	"github/rigel-developer/golang-course/ejercicios"
	"os"
)

var filename string = "./files/txt/table.txt"

func GrabarTable() {
	var table string = ejercicios.MultiplicationTable()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Hubo un error al crear el archivo")
		return
	}

	fmt.Fprintln(archivo, table)
	archivo.Close()

}

func SumaTabla() {
	var table string = ejercicios.MultiplicationTable()

	if !AppendTable(filename, table) {
		fmt.Println("Hubo un error al agregar la tabla")
	}

}

func AppendTable(filen string, texto string) bool {

	archivo, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error al leer el archivo")
		return false
	}
	_, err = archivo.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error al escribir el archivo")
		return false
	}
	archivo.Close()

	return true

}

func LeerArchivo() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Hubo un error al leer el archivo")
		return
	}
	scanner := bufio.NewScanner(archivo)
	var i int
	for scanner.Scan() {
		i++
		linea := scanner.Text()
		// fmt.Println(i)
		fmt.Println(linea)
	}
	archivo.Close()
}
