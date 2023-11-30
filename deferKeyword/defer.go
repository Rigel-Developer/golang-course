package deferkeyword

import (
	"fmt"
	"log"
)

func VemosDefer() {
	defer fmt.Println("Hola")
	fmt.Println("Mundo")
}

func VemosPanic() {
	//anonymous function that will be executed when the function ends
	defer func() {
		// recover() is a function that allows us to recover from a panic
		reco := recover()
		if reco != nil {
			// logs are util because we can see the stack trace and time of the error
			log.Fatalf("Ocurrio un error que genero panic %v", reco)
		}

	}()
	fmt.Println("Hola")
	panic("Se produjo un error")
	fmt.Println("Mundo")
}
