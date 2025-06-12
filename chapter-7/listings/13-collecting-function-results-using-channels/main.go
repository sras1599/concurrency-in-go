package main

import (
	"fmt"
)

func findFactors(number int) []int {
	fmt.Printf("Finding factors of %d\n", number)
	ret := make([]int, 0)

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			ret = append(ret, i)
		}
	}

	return ret
}

func main() {
	result := make(chan []int)

	go func() {
		result <- findFactors(18271812821)
	}()

	fmt.Println(findFactors(717271781))
	fmt.Println(<-result)
}
