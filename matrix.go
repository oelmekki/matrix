package matrix

import (
	"fmt"
	"math"
)

var DEBUG bool = false

type Row []float64
type Matrix []Row

/*
 * Provide `true` if you want errors to panic
 */
func SetDebug( debug bool ) {
	DEBUG = debug
}

/*
 * Generate a Matrix having the same dimensions than origin matrix,
 * but filled with 0.0
 */
func ZeroMatrixFrom( origin *Matrix ) ( matrix Matrix ) {
	for _, row := range( *origin ) {
		var newRow Row

		for range( row ) {
			newRow = append( newRow, 0.0 )
		}

		matrix = append( matrix, newRow )
	}

	return
}

/*
 * A matrix is valid if it has rows, columns and if all its rows have the same
 * amount of columns.
 */
func ( matrix *Matrix ) Valid() bool {
	if len( *matrix ) == 0 || len( (*matrix)[0] ) == 0 {
		return false
	}

	baseCol := len( (*matrix)[0] )
	for _, row := range( *matrix ) {
		if len( row ) != baseCol {
			return false
		}
	}

	return true
}

/*
 * Tells if two matrices have the same dimensions and same
 * values in each cell.
 */
func ( matrix *Matrix ) EqualTo( otherMatrix *Matrix ) bool {
	if ! matrix.SameDimensions( otherMatrix ) {
		return false
	}

	for i := 0 ; i < len( *matrix ); i++ {
		for j := 0 ; j < len( (*matrix)[0] ); j++ {
			if (*matrix)[i][j] != (*otherMatrix)[i][j] {
				return false
			}
		}
	}

	return true
}

/*
 * True if both matrices are valid and have the same dimensions
 */
func ( matrix *Matrix ) SameDimensions( otherMatrix *Matrix ) bool {
	if ! matrix.Valid() || ! otherMatrix.Valid() {
		return false
	}

	return len( *matrix ) == len( *otherMatrix ) && len( (*matrix)[0] ) == len( (*otherMatrix)[0] )
}


/*
 * Multiply the matrix with a scalar.
 *
 * Error is returned if matrix is not valid.
 */
func ( matrix *Matrix ) ScalarMultiply( scalar float64 ) ( resultMatrix Matrix, err error ) {
	operation := func( value float64 ) float64 {
		return value * scalar
	}

	resultMatrix, err = matrix.UnaryOperation( operation, "ScalarMultiply" )

	return
}


/*
 * Perform a mathematical standard multiplication between matrix and otherMatrix, and return
 * the resulting resultMatrix.
 *
 * Error is returned if resultMatrix is undefined (that is, if matrix columns count is not
 * the same than otherMatrix rows count).
 */
func ( matrix *Matrix ) DotProduct( otherMatrix *Matrix ) ( resultMatrix Matrix, err error ) {
	if  len( (*matrix)[0] ) != len( *otherMatrix ) {
		message := fmt.Sprintf( "Can't multiply matrices: %v columns count do not match %v rows count", *matrix, *otherMatrix )
		if DEBUG {
			panic( message )
		} else {
			err = fmt.Errorf( message )
		}

		return
	}

	resultRowsCount := len( *matrix )
	resultColumnsCount := len( (*otherMatrix)[0] )

	for i := 0 ; i < resultRowsCount ; i++ {
		var row Row

		for j := 0 ; j < resultColumnsCount ; j++ {
			sum := 0.0

			for k := 0 ; k < len( (*matrix)[0] ) ; k++ {
				valueInMatrix := (*matrix)[i][k]
				valueInOtherMatrix := (*otherMatrix)[k][j]
				sum += valueInMatrix * valueInOtherMatrix
			}

			row = append( row, sum )
		}

		resultMatrix = append( resultMatrix, row )
	}

	return
}

/*
 * Produce a new matrix by applying `operation` cell by cell on two matrices, so that:
 *    operation( cell1, cell2 ) -> resultCell
 *
 * Returns error if both matrices aren't of same dimensions.
 *
 * `operationName` is used in error message, so that it's easier to know which
 * operation error'd.
 */
func ( matrix *Matrix ) BinaryOperation( otherMatrix *Matrix, operation func( float64, float64 ) float64, operationName string ) ( resultMatrix Matrix, err error ) {
	if ! matrix.SameDimensions( otherMatrix ) {
		message := fmt.Sprintf( `Can't apply operation "%s" on matrices: %v is not the same dimension than %v`, operationName, *matrix, *otherMatrix )
		if DEBUG {
			panic( message )
		} else {
			err = fmt.Errorf( message )
		}

		return
	}

	for i, row := range( *matrix ) {
		var newRow Row

		for j, value := range( row ) {
			newRow = append( newRow, operation( value, (*otherMatrix)[i][j] ) )
		}

		resultMatrix = append( resultMatrix, newRow )
	}

	return
}

/*
 * Produce a new matrix by applying `operation` cell by cell on matrix, so that:
 *    operation( cell ) -> resultCell
 */
func ( matrix *Matrix ) UnaryOperation( operation func( float64 ) float64, operationName string ) ( resultMatrix Matrix, err error ) {
	if ! matrix.Valid() {
		message := fmt.Sprintf( `Can't apply operation "%s" on matrix %v: matrix is not valid`, operationName, *matrix )
		if DEBUG {
			panic( message )
		} else {
			err = fmt.Errorf( message )
		}

		return
	}

	for _, row := range( *matrix ) {
		var newRow Row

		for _, value := range( row ) {
			newRow = append( newRow, operation( value ) )
		}

		resultMatrix = append( resultMatrix, newRow )
	}

	return
}

/*
 * Switch matrix dimensions, so that, eg, a 2x3 matrix returns
 * a 3x2 one.
 *
 * Error is returned if matrix is not valid.
 */
func ( matrix *Matrix ) Transpose() ( resultMatrix Matrix, err error ) {
	if ! matrix.Valid() {
		message := fmt.Sprintf( `Can't transpose matrix %v: matrix is not valid`, *matrix )
		if DEBUG {
			panic( message )
		} else {
			err = fmt.Errorf( message )
		}

		return
	}

	rowsCount := len( (*matrix)[0] )
	colsCount := len( *matrix )

	for i := 0 ; i < rowsCount ; i++ {
		var row Row

		for j := 0 ; j < colsCount ; j++ {
			row = append( row, (*matrix)[j][i] )
		}

		resultMatrix = append( resultMatrix, row )
	}

	return
}


/*
 * Multiply each cell from matrix with each cell at the same coordinate in otherMatrix.
 *
 * Note that *this is not* standard mathematical matrix multiplication. For that one,
 * use `DotProduct()`.
 */
func ( matrix *Matrix ) MultiplyCells( otherMatrix *Matrix ) ( resultMatrix Matrix, err error ) {
	operation := func( value1 float64, value2 float64 ) float64 {
		return value1 * value2
	}

	resultMatrix, err = matrix.BinaryOperation( otherMatrix, operation, "MultiplyCells" )
	return
}

/*
 * Add otherMatrix to matrix and return the resulting resultMatrix
 *
 * Error is returned if matrices are not valid or do not have the same dimensions.
 */
func ( matrix *Matrix ) Add( otherMatrix *Matrix ) ( resultMatrix Matrix, err error ) {
	operation := func( value1 float64, value2 float64 ) float64 {
		return value1 + value2
	}

	resultMatrix, err = matrix.BinaryOperation( otherMatrix, operation, "Add" )
	return
}

/*
 * Substract otherMatrix from matrix and return the resulting resultMatrix
 *
 * Error is returned if matrices are not valid or do not have the same dimensions.
 */
func ( matrix *Matrix ) Substract( otherMatrix *Matrix ) ( resultMatrix Matrix, err error ) {
	operation := func( value1 float64, value2 float64 ) float64 {
		return value1 - value2
	}

	resultMatrix, err = matrix.BinaryOperation( otherMatrix, operation, "Substract" )
	return
}

/*
 * Apply sigmoid function on each cell of matrix and return resulting Matrix
 *
 * Error is returned if matrix is not valid.
 */
func ( matrix *Matrix ) Sigmoid() ( resultMatrix Matrix, err error ) {
	operation := func( value float64 ) float64 {
		return 1.0 / ( 1.0 + math.Exp( -value ) )
	}

	resultMatrix, err = matrix.UnaryOperation( operation, "Sigmoid" )
	return
}

/*
 * Compute derivative for sigmoid function on each cell of matrix and return resulting Matrix
 *
 * Error is returned if matrix is not valid.
 */
func ( matrix *Matrix ) SigmoidDerivative() ( resultMatrix Matrix, err error ) {
	resultMatrix, err = matrix.Sigmoid()
	if err != nil { return resultMatrix, err }

	operation := func( value float64 ) float64 {
		return value * ( 1.0 - value )
	}

	resultMatrix, err = resultMatrix.UnaryOperation( operation, "SigmoidDerivative" )
	return
}
