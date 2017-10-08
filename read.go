package matrix

import (
	"fmt"
)

/*
 * Returns the number of rows
 */
func (matrix Matrix) Rows() int {
	return int(matrix[0])
}

/*
 * Returns the number of columns
 */
func (matrix Matrix) Cols() int {
	return int(matrix[1])
}

/*
 * Returns a human readable representation of matrix, ready to print
 */
func (matrix Matrix) String() string {
	output := "\n"
	for i := 0; i < int(matrix[0]); i++ {
		output += "{\t\t"
		for j := 0; j < int(matrix[1]); j++ {
			output = fmt.Sprintf("%v%v", output, matrix.At(i, j))
			if j < int(matrix[1])-1 {
				output += "\t\t"
			}
		}
		output += "\t\t}\n"
	}

	return output
}

/*
 * Returns the value at position `row`, `col`.
 *
 * Just like an array, you're responsible to make sure
 * you don't ask for an out of range value.
 */
func (matrix Matrix) At(row, col int) float64 {
	return matrix[matrix.IndexFor(row, col)]
}

/*
 * Low level method: allow to compute the flat array index
 * for a matrix position.
 */
func (matrix Matrix) IndexFor(row, col int) int {
	return row*int(matrix[1]) + col + 2
}

/*
 * Returns the given row (0-indexed) as a []float64
 */
func (matrix Matrix) GetRow(index int) (row []float64, err error) {
	if index+1 > matrix.Rows() {
		err = fmt.Errorf("Row %d is out of matrix", row)
		return
	}

	for i := 0; i < matrix.Cols(); i++ {
		row = append(row, matrix.At(index, i))
	}

	return
}
