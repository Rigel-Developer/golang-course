package variables

import "fmt"

func Mapas() {
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pedro"] = 20
	m["Maria"] = 30
	m["Ana"] = 40
	m["Luis"] = 50
	fmt.Println(m)
	//Recorrer el mapa
	for key, value := range m {
		fmt.Println(key, value)
	}
	//Obtener un valor
	v, ok := m["Joses"]
	fmt.Println(v, ok)
}
