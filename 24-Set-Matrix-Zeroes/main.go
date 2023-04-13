package main

import (
	"log"
)

func main() {
	input := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	input1 := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(input)
	setZeroes(input1)
}

func setZeroes(matrix [][]int) {
	lenRow := len(matrix)
	lenCol := len(matrix[0])

	var locations [][]int
	row := 0

	for row < lenRow {
		col := 0
		for col < lenCol {
			if matrix[row][col] == 0 {
				locations = append(locations, []int{row, col})
			}
			col++
		}
		row++
	}

	for _, l := range locations {
		for i := 0; i < lenCol; i++ {
			matrix[l[0]][i] = 0
		}

		for i := 0; i < lenRow; i++ {
			matrix[i][l[1]] = 0
		}
	}

	log.Println(matrix)
}
