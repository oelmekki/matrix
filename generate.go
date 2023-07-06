package matrix

import (
	"fmt"
	"math/rand"
)

// GenerateMatrix creates a zero matrix with `rows` rows and `cols` cols.
func GenerateMatrix(rows, cols int) (matrix Matrix) {
	matrix = make(Matrix, rows*cols+2)
	matrix[0] = float64(rows)
	matrix[1] = float64(cols)

	return
}

// Build generates a new matrix by passing from `Builder`.
//
// This allows to have a human friendly looking way of initializing matrices:
//
// myMatrix, _ := matrix.Build(
// 	matrix.Builder {
// 		matrix.Row{10, -5.3, 22},
// 		matrix.Row{-2, -25, 12},
// 		matrix.Row{ 7, 5, -12.5},
// 	},
// )
//
// It returns an error if you try to provide a builder with no rows or
// rows with no cols (you can safely ignore error if you're confident
// your builder is valid).
func Build(builder Builder) (resultMatrix Matrix, err error) {
	if len(builder) == 0 || len(builder[0]) == 0 {
		err = generateError(fmt.Sprintf("Can't build empty matrix. If you want to generate a zero matrix, use GenerateMatrix()"))
		return
	}

	resultMatrix = GenerateMatrix(len(builder), len(builder[0]))
	for i, row := range builder {
		for j, value := range row {
			resultMatrix[resultMatrix.IndexFor(i, j)] = value
		}
	}

	return
}

// ZeroMatrixFrom generates a Matrix having the same dimensions than origin matrix,
// but filled with 0.0
func ZeroMatrixFrom(origin Matrix) Matrix {
	return GenerateMatrix(int(origin[0]), int(origin[1]))
}

// RandomMatrix generates a matrix of `rows` rows and `cols` cols with randomized
// values.
//
// Values are randomized with standard normal distribution, using
// `math.NormFloat64()` (don't forget to seed randomizer if you
// want it to be truly random).
func RandomMatrix(rows, cols int) Matrix {
	matrix := GenerateMatrix(rows, cols)
	for i := 2; i < len(matrix); i++ {
		matrix[i] = rand.NormFloat64()
	}

	return matrix
}
