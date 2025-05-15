package main

import (
	"fmt"
	"time"
)

type rowOrColumn []int
type matrix []rowOrColumn

const matrixSize = 1000
const iterations = 5

func getColumn(matrix *matrix, colIndex int) *rowOrColumn {
	col := make(rowOrColumn, matrixSize)

	for row := 0; row < matrixSize; row++ {
		col[row] = (*matrix)[row][colIndex]
	}

	return &col
}

func matrixMultiply(matrixA *matrix, matrixB *matrix, result *matrix, barrier *Barrier) {
	for rowIndex := 0; rowIndex < matrixSize; rowIndex++ {
		go func(rowIndex int) {
			for colIndex := 0; colIndex < matrixSize; colIndex++ {
				row := &(*matrixA)[rowIndex]
				col := getColumn(matrixB, colIndex)

				(*result)[rowIndex][colIndex] = rowMultiply(row, col)
			}
			barrier.Wait()
		}(rowIndex)
	}
}

func rowMultiply(row *rowOrColumn, col *rowOrColumn) int {
	sum := 0

	for i := 0; i < matrixSize; i++ {
		sum += (*row)[i] * (*col)[i]
	}

	return sum
}

func main() {
	barrier := newBarrier(matrixSize + 1)
	start := time.Now()

	for iteration := 0; iteration < iterations; iteration++ {
		matrixA := generateMatrix(matrixSize, false)
		matrixB := generateMatrix(matrixSize, false)
		result := generateMatrix(matrixSize, true)

		go matrixMultiply(matrixA, matrixB, result, barrier)
		barrier.Wait()

		// for i := 0; i < matrixSize; i++ {
		// 	fmt.Println((*matrixA)[i], (*matrixB)[i], (*result)[i])
		// }
		// fmt.Println()
	}

	fmt.Printf("Total time taken to multiply %d %dx%d matrices: %s\n", iterations, matrixSize, matrixSize, time.Since(start))
}
