package variables

import (
	"fmt"
	"strings"
)

func Arrays() {
	// //Arreglos inmutables
	// var array [4]int
	// array[1] = 1
	// fmt.Println(array)
	// //Slices mutables
	// slice := []int{1, 2, 3, 4, 5}
	// fmt.Println(slice)
	// //slicing
	// fmt.Println(slice[0:2])
	// //append
	// slice = append(slice, 6)
	// fmt.Println(slice)
	// //append nueva lista usando spread operator
	// slice = append(slice, []int{7, 8, 9}...)
	// fmt.Println(slice)

	//recorrer un slice
	slice := []int{1, 2, 3, 4, 5}
	for _, valor := range slice {
		fmt.Println(valor)
	}
	isPal := isPalindromo("anita lava la tina")
	fmt.Println(isPal)

}

func isPalindromo(text string) bool {
	text = strings.ReplaceAll(text, " ", "")
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	return text == textReverse

}
