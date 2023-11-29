package arreglos_slices

import "fmt"

var otraTabla []int = []int{1, 12, 13, 14}
var tabla2 [10]int = [10]int{1, 12, 13, 14, 15, 16, 17, 18, 19, 20}

func MuestroSlice() {

	porcion := tabla2[0:5]
	porcion2 := tabla2[:5]
	porcion3 := tabla2[5:]
	porcion4 := tabla2[2:7]
	porcion0 := tabla2[:]

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
	fmt.Println(porcion4)
	fmt.Println(porcion0)
}

func CapacidadSlice() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
