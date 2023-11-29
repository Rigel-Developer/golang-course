package iteraciones

import "fmt"

func Iterar() {
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 100; i += 5 {
	// 	fmt.Println(i)
	// }

	for i := 100; i > 10; i -= 5 {
		if i == 50 {
			continue
		}
		fmt.Println(i)
	}
}
