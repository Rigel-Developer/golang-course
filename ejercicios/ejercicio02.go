package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MultiplicationTable() string {
	scanner := bufio.NewScanner(os.Stdin)
	var numero int
	var err error
	var texto string

	for {
		fmt.Println("Ingrese numero: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Ingrese un numero valido \n")
				continue
			}
		}

		for i := 1; i <= 10; i++ {
			texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
		}
		break
	}

	return texto
	// fmt.Println("Ingrese numero: ")
	// if scanner.Scan() {
	// 	for scanner.Text() != "" {
	// 		numero1, err := strconv.Atoi(scanner.Text())
	// 		if err != nil {
	// 			fmt.Println("Ingrese numero: ")
	// 			scanner.Scan()
	// 			continue
	// 		}

	// 		for i := 1; i <= 10; i++ {
	// 			fmt.Println(numero1, "x", i, "=", i*numero1)
	// 			//end

	// 		}
	// 		break

	// 	}

	// }

}
