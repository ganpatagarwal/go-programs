package main

import (
	"fmt"
)

// *Print matrix diagonals:*
// Given a matrix of integers print out its values along the diagonals.
// For example, given this matrix:

// {{1,  2,  3,  4},
//  {5,  6,  7,  8},
//  {9, 10, 11, 12}}

// The print out should be:

// 1
// 2 5
// 3 6 9
// 4 7 10
// 8 11
// 12

// O(row*col) time | O(row*col) space
func Diagonals(matrix [][]int) [][]int {
	var diagonals [][]int

	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if visited[row][col] {
				continue
			}
			getDiagonalElements(row, col, matrix, visited, &diagonals)
		}
	}

	return diagonals
}

func getDiagonalElements(row, col int, matrix [][]int, visited [][]bool, diagonals *[][]int) {
	currDiagonals := []int{matrix[row][col]}

	visited[row][col] = true

	for row < len(matrix)-1 && col > 0 {
		row += 1
		col -= 1

		if !visited[row][col] {
			currDiagonals = append(currDiagonals, matrix[row][col])
			visited[row][col] = true
		}
	}

	*diagonals = append(*diagonals, currDiagonals)
}

func main() {
	sampleInput := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(Diagonals(sampleInput))
}
