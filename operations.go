package matrix

import (
	"fmt"
	"math"
)

// ScalarMultiply multiplies each cell of the matrix individually
// with the provided value.
//
// Error is returned if matrix is not valid.
func (matrix Matrix) ScalarMultiply(scalar float64) (resultMatrix Matrix, err error) {
	operation := func(value float64) float64 {
		return value * scalar
	}

	resultMatrix, err = matrix.UnaryOperation(operation, "ScalarMultiply")

	return
}

// DotProduct performs a mathematical standard multiplication between matrix and otherMatrix,
// and return the resulting resultMatrix.
//
// Error is returned if resultMatrix is undefined (that is, if matrix columns count is not
// the same than otherMatrix rows count).
func (matrix Matrix) DotProduct(otherMatrix Matrix) (resultMatrix Matrix, err error) {
	if matrix[1] != otherMatrix[0] {
		err = generateError(fmt.Sprintf("Can't multiply matrices: %v columns count do not match %v rows count", matrix, otherMatrix))
		return
	}

	resultMatrix = GenerateMatrix(int(matrix[0]), int(otherMatrix[1]))

	for i := 0; i < int(resultMatrix[0]); i++ {
		for j := 0; j < int(resultMatrix[1]); j++ {
			sum := 0.0

			for k := 0; k < int(matrix[1]); k++ {
				valueInMatrix := matrix.At(i, k)
				valueInOtherMatrix := otherMatrix.At(k, j)
				sum += valueInMatrix * valueInOtherMatrix
			}

			resultMatrix[resultMatrix.IndexFor(i, j)] = sum
		}
	}

	return
}

// VectorMultiply multiplies matrix by given vector, return a new []float64 vector.
//
// Error is returned is operation is not valid, that is if vector has not as many entries
// than there is matrix columns.
func (matrix Matrix) VectorMultiply(vector []float64) (resultVector []float64, err error) {
	if matrix.Cols() != len(vector) {
		err = fmt.Errorf("Vector length (%d) and matrix columns count (%d) differ", len(vector), matrix.Cols())
		return
	}

	fakeMatrix := GenerateMatrix(len(vector), 1)
	for i := 0; i < len(vector); i++ {
		fakeMatrix.SetAt(i, 0, vector[i])
	}

	resultMatrix, err := matrix.DotProduct(fakeMatrix)
	if err != nil {
		return
	}

	resultVector = make([]float64, resultMatrix.Rows())

	fmt.Printf("matrix: %v\nfakeMatrix: %v\nresultMatrix: %v\n", matrix, fakeMatrix, resultMatrix)

	for i := 0; i < resultMatrix.Rows(); i++ {
		resultVector[i] = resultMatrix.At(i, 0)
	}

	return
}

// SetAt sets value at given row and col.
//
// For backward compatibility reasons, you're responsible for making sure
// the position at row and col actually exists in the matrix.
func (matrix Matrix) SetAt(row, col int, val float64) {
	matrix[matrix.IndexFor(row, col)] = val
}

// Transpose switches matrix dimensions, so that, eg, a 2x3 matrix returns
// a 3x2 one.
//
// Error is returned if matrix is not valid.
func (matrix Matrix) Transpose() (resultMatrix Matrix, err error) {
	if !matrix.Valid() {
		err = generateError(fmt.Sprintf(`Can't transpose matrix %v: matrix is not valid`, matrix))
		return
	}

	resultMatrix = GenerateMatrix(int(matrix[1]), int(matrix[0]))

	for i := 0; i < int(matrix[1]); i++ {
		for j := 0; j < int(matrix[0]); j++ {
			resultMatrix[resultMatrix.IndexFor(i, j)] = matrix.At(j, i)
		}
	}

	return
}

// MultiplyCells multiplies each cell from matrix with each cell at the same coordinate
// in otherMatrix.
//
// Note that *this is not* standard mathematical matrix multiplication. For that one,
// use `DotProduct()`.
func (matrix Matrix) MultiplyCells(otherMatrix Matrix) (resultMatrix Matrix, err error) {
	operation := func(value1 float64, value2 float64) float64 {
		return value1 * value2
	}

	resultMatrix, err = matrix.BinaryOperation(otherMatrix, operation, "MultiplyCells")
	return
}

// Add adds up otherMatrix to matrix and returns the resulting resultMatrix.
//
// Error is returned if matrices are not valid or do not have the same dimensions.
func (matrix Matrix) Add(otherMatrix Matrix) (resultMatrix Matrix, err error) {
	operation := func(value1 float64, value2 float64) float64 {
		return value1 + value2
	}

	resultMatrix, err = matrix.BinaryOperation(otherMatrix, operation, "Add")
	return
}

// Substract removes otherMatrix from matrix and returns the resulting resultMatrix.
//
// Error is returned if matrices are not valid or do not have the same dimensions.
func (matrix Matrix) Substract(otherMatrix Matrix) (resultMatrix Matrix, err error) {
	operation := func(value1 float64, value2 float64) float64 {
		return value1 - value2
	}

	resultMatrix, err = matrix.BinaryOperation(otherMatrix, operation, "Substract")
	return
}

// Sigmoid applies sigmoid function on each cell of matrix and returns resulting Matrix.
//
// Error is returned if matrix is not valid.
func (matrix Matrix) Sigmoid() (resultMatrix Matrix, err error) {
	operation := func(value float64) float64 {
		return 1.0 / (1.0 + math.Exp(-value))
	}

	resultMatrix, err = matrix.UnaryOperation(operation, "Sigmoid")
	return
}

// SigmoidDerivative computes derivative for sigmoid function on each cell of matrix
// and returns resulting Matrix.
//
// Error is returned if matrix is not valid.
func (matrix Matrix) SigmoidDerivative() (resultMatrix Matrix, err error) {
	resultMatrix, err = matrix.Sigmoid()
	if err != nil {
		return resultMatrix, err
	}

	operation := func(value float64) float64 {
		return value * (1.0 - value)
	}

	resultMatrix, err = resultMatrix.UnaryOperation(operation, "SigmoidDerivative")
	return
}

// BinaryOperation produces a new matrix by applying `operation` cell by cell on
// two matrices, so that:
//
//    operation(cell1, cell2) -> resultCell
//
// `operationName` is used in error message, so that it's easier to know which
// operation error'd.
//
// Returns error if both matrices aren't of same dimensions.
func (matrix Matrix) BinaryOperation(otherMatrix Matrix, operation func(float64, float64) float64, operationName string) (resultMatrix Matrix, err error) {
	if !matrix.SameDimensions(otherMatrix) {
		err = generateError(fmt.Sprintf(`Can't apply operation "%s" on matrices: %v is not the same dimension than %v`, operationName, matrix, otherMatrix))
		return
	}

	resultMatrix = ZeroMatrixFrom(matrix)

	for i := 2; i < len(matrix); i++ {
		resultMatrix[i] = operation(matrix[i], otherMatrix[i])
	}

	return
}

// UnaryOperation produces a new matrix by applying `operation` cell by cell
// on matrix, so that:
//
//    operation(cell) -> resultCell
//
// `operationName` is used in error message, so that it's easier to know which
// operation error'd.
//
// Returns error if both matrices aren't of same dimensions.
func (matrix Matrix) UnaryOperation(operation func(float64) float64, operationName string) (resultMatrix Matrix, err error) {
	if !matrix.Valid() {
		err = generateError(fmt.Sprintf(`Can't apply operation "%s" on matrix %v: matrix is not valid`, operationName, matrix))
		return
	}

	resultMatrix = ZeroMatrixFrom(matrix)

	for i := 2; i < len(matrix); i++ {
		resultMatrix[i] = operation(matrix[i])
	}

	return
}
