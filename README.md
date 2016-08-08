# Matrix

A small lib allowing to manipulate matrices.

This is not intended to exhaustively implements all possible mathematical operations on matrices. Instead, it's just what I needed to build a neural network. That being said, it provides build blocks to easily extend it (see [`Extending` section](#Extending)). If you feel some operation should be in there, feel free to send a pull request.


## Install

```
go get oelmekki/matrix
```


## Provided types

Two types are provided : `Matrix` and `Row`.

A `Matrix` is a slice of `Row`s.

A `Row` is a slice of `float64`.

Thus, you can build a Matrix that way:

```golang
myMatrix := matrix.Matrix{
  matrix.Row{  10.0, -5.3,   22.0 },
  matrix.Row{ -2.0,  -25.0,  12.0 },
  matrix.Row{  7.0,   5.3,  -12.5 },
}
```

## Demo

Documentation details are provided after this section, but here is what the lib
allows you to do in a glance:

```golang
package main

import (
  "github.com/oelmekki/matrix"
  "fmt"
)

func main() {
  firstMatrix := matrix.Matrix{
    matrix.Row{  10, -5.3,  22   },
    matrix.Row{  -2, -25,   12   },
    matrix.Row{   7,  5,   -12.5 },
  }

  // tests

  println( firstMatrix.Valid() ) // true

  secondMatrix := matrix.Matrix{
    matrix.Row{  12, -15.5 },
    matrix.Row{ -4,  -5    },
    matrix.Row{  3,   2.5  },
  }

  println( firstMatrix.SameDimensions( &secondMatrix ) ) // false

  thirdMatrix := matrix.Matrix{
    matrix.Row{  1,   7.2,    2   },
    matrix.Row{  2,   5,      1   },
    matrix.Row{  2,  -25.3,  -2.5 },
  }

  println( firstMatrix.SameDimensions( &thirdMatrix ) ) // true
  println( firstMatrix.EqualTo( &thirdMatrix ) ) // false

  // operations

  newMatrix, _ := firstMatrix.ScalarMultiply( 10.0 )
  fmt.Printf( "%v\n", newMatrix )
  /*
   * [
   *   [  100 -53   220 ]
   *   [ -20  -250  120 ]
   *   [  70   53  -125 ]
   * ]
   */

  newMatrix, _ = firstMatrix.DotProduct( &secondMatrix )
  fmt.Printf( "%v\n", newMatrix )
  /*
   * [
   *   [ 207.2  -73.5   ]
   *   [ 112     186    ]
   *   [ 26.5   -164.75 ]
   * ]
   */

  newMatrix, _ = secondMatrix.Transpose()
  fmt.Printf( "%v\n", newMatrix )
  /*
   * [
   *   [ 12     -4   3   ]
   *   [ -15.5  -5   2.5 ]
   * ]
   */

   // and more! See doc below
}

```


## Generate and test matrices

A few helper methods are provided to generate matrices and test them.


### `func ZeroMatrixFrom( origin *Matrix ) ( matrix Matrix )`

Generate a Matrix having the same dimensions than origin matrix,
but filled with 0.0.


### `func ( matrix *Matrix ) Valid() bool`

A matrix is valid if it has rows, columns and if all its rows have the same
amount of columns.


### `func ( matrix *Matrix ) SameDimensions( otherMatrix *Matrix ) bool`

True if both matrices are valid and have the same dimensions.


### `func ( matrix *Matrix ) EqualTo( otherMatrix *Matrix ) bool`

Tells if two valid matrices have the same dimensions and same
values in each cell.


## Operations on matrices

Those are the operations currently implemented. Note that all operations
produce a new matrix and return it, current matrix and argument matrix are
never modified (they are passed by reference, though, to avoid copying possibly
huge matrices around).


### `func ( matrix *Matrix ) ScalarMultiply( scalar float64 ) ( resultMatrix Matrix, err error )`

Multiply the matrix with a scalar.

Error is returned if matrix is not valid.


### `func ( matrix *Matrix ) DotProduct( otherMatrix *Matrix ) ( resultMatrix Matrix, err error )`

Perform a mathematical standard multiplication between matrix and otherMatrix, and return
the resulting resultMatrix.

Error is returned if resultMatrix is undefined (that is, if matrix columns count is not
the same than otherMatrix rows count).


### `func ( matrix *Matrix ) Transpose() ( resultMatrix Matrix, err error )`

Switch matrix dimensions, so that, eg, a 2x3 matrix returns
a 3x2 one.

Error is returned if matrix is not valid.


### `func ( matrix *Matrix ) MultiplyCells( otherMatrix *Matrix ) ( resultMatrix Matrix, err error )`

Multiply each cell from matrix with each cell at the same coordinate in otherMatrix.

Note that *this is not* standard mathematical matrix multiplication. For that one,
use `DotProduct()`.

Error is returned if matrices do not have the same dimensions.


### `func ( matrix *Matrix ) Add( otherMatrix *Matrix ) ( resultMatrix Matrix, err error )`

Add otherMatrix to matrix and return the resulting resultMatrix.

Error is returned if matrices are not valid or do not have the same dimensions.


### `func ( matrix *Matrix ) Substract( otherMatrix *Matrix ) ( resultMatrix Matrix, err error )`

Substract otherMatrix from matrix and return the resulting resultMatrix.

Error is returned if matrices are not valid or do not have the same dimensions.


### `func ( matrix *Matrix ) Sigmoid() ( resultMatrix Matrix, err error )`

Apply sigmoid function on each cell of matrix and return resulting Matrix.

Error is returned if matrix is not valid.


### `func ( matrix *Matrix ) SigmoidDerivative() ( resultMatrix Matrix, err error )`

Compute derivative for sigmoid function on each cell of matrix and return resulting Matrix.

Error is returned if matrix is not valid.


## Extending

Two generic operations are provided that should allow you to perform any cell
by cell operation by taking one or two matrices as an input: `UnaryOperation`
and `BinaryOperation`.


### `func ( matrix *Matrix ) UnaryOperation( operation func( float64 ) float64, operationName string ) ( resultMatrix Matrix, err error )`

Produce a new matrix by applying `operation` cell by cell on matrix, so that:

```
operation( cell ) -> resultCell
```

`operationName` is a name you provide for your operation so that it's easier to
know where the error comes from (error will contain that operation name).

`operation` is a function you provide, that will receive the `float64` value of
each cell, and should return the new `float64` value for that cell.

So, for example, if you want to multiply each value by 2 and add 1:

```golang
operation := func( value float64 ) float64 {
  return value * 2 + 1
}

resultMatrix, err := myMatrix.UnaryOperation( operation, "x * 2 + 1" )
```

Error is returned is the matrix is not valid.


### `func ( matrix *Matrix ) BinaryOperation( otherMatrix *Matrix, operation func( float64, float64 ) float64, operationName string ) ( resultMatrix Matrix, err error )`

Produce a new matrix by applying `operation` cell by cell on two matrices, so that:
   operation( cell1, cell2 ) -> resultCell


`operationName` is used in error message, so that it's easier to know which
operation error'd.

So, for example, if you want to divide each cell of matrix by corresponding cell in otherMatrix:

```golang
operation := func( value float64, otherValue float64 ) float64 {
  return value / otherValue
}

resultMatrix, err := myMatrix.BinaryOperation( otherMatrix, operation, "x / y" )
```


Returns error if any matrix is invalid, or both matrices aren't of same dimensions.


## Debugging

Sometime, having the lib panic'ing instead of returning error is more useful,
so that you can see the stacktrace.

If you want that, use once:

```golang
matrix.SetDebug( true )
```
