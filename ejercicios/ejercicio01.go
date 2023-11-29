package ejercicios

import "strconv"

func Ejercicio01(param string) (int, string) {
	number, err := strconv.Atoi(param)
	if err != nil {
		// handle the error if the conversion fails
		return 0, err.Error()
	}
	if number > 100 {
		return number, "Es mayor a 100"
	} else {
		return number, "Es menor a 100"
	}

	// continue with your logic using the converted number

}
