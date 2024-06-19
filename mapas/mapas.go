package mapas

import "fmt"

func MuestroMapas() {
	paises := make(map[string]string, 5)

	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)
	fmt.Println(paises["Mexico"])

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	fmt.Println(campeonato)

	// for equipo, puntaje := range campeonato {
	// 	fmt.Printf("El equipo %s, tiene un puntaje de: %d \n", equipo, puntaje)
	// }

	//quitar elemento
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	//buscar elemento
	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)
}
