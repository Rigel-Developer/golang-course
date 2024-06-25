package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLentooo(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
	canal1 <- true //channels are used to communicate between goroutines
}

func task(name string, duration time.Duration) {
	time.Sleep(duration)
	fmt.Println(name, "completed")
}

func CompleteTask() {
	// Captura el tiempo de inicio
	start := time.Now()

	task("Task 1", 2*time.Second)
	task("Task 2", 1*time.Second)
	task("Task 3", 3*time.Second)

	time.Sleep(4 * time.Second) // Espera suficiente para que todas las goroutines terminen

	// Captura el tiempo de finalización
	end := time.Now()

	// Calcula la duración
	duration := end.Sub(start)

	fmt.Println("All tasks completed")
	fmt.Printf("Total execution time: %v\n", duration)
}
