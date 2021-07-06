package main

import "fmt"

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	// for index, number := range numbers{ //range는 index도 같이준다 (loop)
	// 	fmt.Println(index,number)
	// }
	// return 1

	return total
}

func main() {
	// superAdd(1, 2, 3, 4, 5, 6)
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}