package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	rows, cols := 5, 5
	numbers := make(map[int]bool)
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			num := rand.Intn(rows*cols) + 1
			for numbers[num] {
				num = rand.Intn(rows*cols) + 1
			}
			numbers[num] = true
			matrix[i][j] = num
		}
	}

	fmt.Println("Сгенерированный двумерный массив:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}
