package main

import (
	"math/rand"
)

func generateMatrix(size int, empty bool) *matrix {
	matrix := make(matrix, size)
	for i := range matrix {
		matrix[i] = make(rowOrColumn, size)
	}

	if empty {
		return &matrix
	}

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			// inserts a random number between -20 and 19
			matrix[row][col] = rand.Intn(40) - 20
		}
	}

	return &matrix
}
