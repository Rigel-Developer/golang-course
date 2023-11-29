package funciones

import "fmt"

func CalculosMatematicos() {
	suma := func(n1, n2 int) int {
		return n1 + n2
	}

	fmt.Println(suma(10, 20))
}
