package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{1, 12, 13, 14}
var matriz [5][7]int

func MuestroArreglos() {
	// tabla[0] = 1
	// tabla[5] = 15

	tabla2 := [4]string{"Hola", "Mundo", "Como", "Estas"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}

	matriz[3][5] = 1
	fmt.Println(matriz)

}
