package main

import (
	"fmt"
	"time"
)

func main() {
	// Definir las dos fechas
	startDate := time.Date(2024, time.January, 12, 10, 38, 0, 0, time.UTC)
	endDate := time.Date(2024, time.January, 12, 5, 38, 0, 0, time.UTC)

	// Restar las fechas y obtener la duración
	duration := endDate.Sub(startDate)

	// Convertir la duración a horas
	hours := duration.Hours()

	fmt.Printf("La diferencia es de %.2f horas\n", hours)
}
