package main

import (
	"fmt"
	"sync"
)

type rowOrColumn []int
type matrix []rowOrColumn

const matrixSize = 3
const iterations = 3

func getColumn(matrix *matrix, colIndex int) *rowOrColumn {
	col := make(rowOrColumn, matrixSize)

	for row := 0; row < matrixSize; row++ {
		col[row] = (*matrix)[row][colIndex]
	}

	return &col
}

func matrixMultiply(matrixA *matrix, matrixB *matrix, result *matrix, wg *sync.WaitGroup) {
	for rowIndex := 0; rowIndex < matrixSize; rowIndex++ {
		go func(rowIndex int) {
			defer wg.Done()
			for colIndex := 0; colIndex < matrixSize; colIndex++ {
				row := &(*matrixA)[rowIndex]
				col := getColumn(matrixB, colIndex)

				(*result)[rowIndex][colIndex] = rowMultiply(row, col)
			}
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
	wg := sync.WaitGroup{}

	for iteration := 0; iteration < iterations; iteration++ {
		wg.Add(matrixSize)

		matrixA := generateMatrix(matrixSize, false)
		matrixB := generateMatrix(matrixSize, false)
		result := generateMatrix(matrixSize, true)

		go matrixMultiply(matrixA, matrixB, result, &wg)
		wg.Wait()

		for i := 0; i < matrixSize; i++ {
			fmt.Println((*matrixA)[i], (*matrixB)[i], (*result)[i])
		}
		fmt.Println()
	}
}
