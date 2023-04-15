package main

import (
	"fmt"
)

func main() {
	sampleInput := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(sampleInput))
}

//Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//Output: [1,2,3,4,8,12,11,10,9,5,6,7]

// time - O(n)
// space - O(1)
func spiralOrder(matrix [][]int) []int {
	var result []int

	leftCol, rightCol := 0, len(matrix[0])
	topRow, bottomRow := 0, len(matrix)

	for leftCol < rightCol && topRow < bottomRow {
		// top row, left to right
		for i := leftCol; i < rightCol; i++ {
			result = append(result, matrix[topRow][i])
		}
		topRow += 1

		// last col, top to bottom
		for i := topRow; i < bottomRow; i++ {
			result = append(result, matrix[i][rightCol-1])
		}
		rightCol -= 1

		// case for single row OR single column matrix
		if !(leftCol < rightCol && topRow < bottomRow) {
			break
		}

		// bottom row, right to left
		for i := rightCol - 1; i >= leftCol; i-- {
			result = append(result, matrix[bottomRow-1][i])
		}
		bottomRow -= 1

		// first col, bottom to top
		for i := bottomRow - 1; i >= topRow; i-- {
			result = append(result, matrix[i][leftCol])
		}
		leftCol += 1
	}

	return result
}
