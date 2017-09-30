# Matrix

A small lib allowing to manipulate matrices.

This is not intended to exhaustively implements all possible mathematical operations on matrices. Instead, it's just what I needed to build a neural network. That being said, it provides build blocks to easily extend it (see [`Extending` section](#extending)). If you feel some operation should be in there, feel free to send a pull request.


## Install

```
go get github.com/oelmekki/matrix
```


## Give me a matrix

You can generate an initialized zero matrix providing its number of rows and cols:

```
matrix := GenerateMatrix( 3, 2 )
```

Or, you can use the builder to provide matrix values in an human readable way.

Two types are provided to build `Matrix` : `Builder` and `Row`.

```go
myMatrix, err := matrix.Build(
  matrix.Builder {
    matrix.Row{  10.0, -5.3,   22.0 },
    matrix.Row{ -2.0,  -25.0,  12.0 },
    matrix.Row{  7.0,   5.3,  -12.5 },
  },
)
if err != nil { println( "You passed a 0x0 or 1x0 matrix." ) }
```

All rows will have the same amount of columns than the first one. If subsequent has less columns, the remaining ones will be filled with 0.


## Demo

Documentation details are provided after this section, but here is what the lib
allows you to do in a glance:

```go
package main

import (
  "github.com/oelmekki/matrix"
  "fmt"
)

func main() {
  firstMatrix, _ := matrix.Build(
    matrix.Builder {
      matrix.Row{  10, -5.3,  22   },
      matrix.Row{  -2, -25,   12   },
      matrix.Row{   7,  5,   -12.5 },
    },
  )

  // tests

  println( firstMatrix.Valid() ) // true
  fmt.Printf( "%v\n", firstMatrix.At( 1, 1 ) ) // -25
  fmt.Printf( "%v\n", firstMatrix.Rows() ) // 3
  fmt.Printf( "%v\n", firstMatrix.Cols() ) // 3

  secondMatrix, _ := matrix.Build(
    matrix.Builder {
      matrix.Row{  12, -15.5 },
      matrix.Row{ -4,  -5    },
      matrix.Row{  3,   2.5  },
    },
  )

  println( firstMatrix.SameDimensions( secondMatrix ) ) // false

  thirdMatrix, _ := matrix.Build(
    matrix.Builder {
      matrix.Row{  1,   7.2,    2   },
      matrix.Row{  2,   5,      1   },
      matrix.Row{  2,  -25.3,  -2.5 },
    },
  )

  println( firstMatrix.SameDimensions( thirdMatrix ) ) // true
  println( firstMatrix.EqualTo( thirdMatrix ) ) // false

  // operations

  newMatrix, _ := firstMatrix.ScalarMultiply( 10.0 )
  fmt.Println( newMatrix.String() )
  /*
   * {               100             -53             220             }
   * {               -20             -250            120             }
   * {               70              50              -125            }
   */

  newMatrix, _ = firstMatrix.DotProduct( secondMatrix )
  fmt.Println( newMatrix.String() )
  /*
   * {               207.2           -73.5           }
   * {               112             186             }
   * {               26.5            -164.75         }
   */

  newMatrix, _ = secondMatrix.Transpose()
  fmt.Println( newMatrix.String() )
  /*
   * {               12              -4              3               }
   * {               -15.5           -5              2.5             }
   */

   // and more! See doc below
}

```


## Generate matrices

A few helper methods are provided to generate matrices and test them.

### `func GenerateMatrix( rows, cols int ) ( matrix Matrix )`

Generate an zero matrix with `rows` rows and `cols` cols


### `func ZeroMatrixFrom( origin Matrix ) ( matrix Matrix )`

Generate a Matrix having the same dimensions than origin matrix,
but filled with 0.0.


### `func Build( builder Builder ) ( resultMatrix Matrix, err error )`

Generate a new matrix by passing a `Builder`.

This allows to have a human friendly looking way of initializing matrices:

```go
myMatrix, _ := matrix.Build(
  matrix.Builder {
    matrix.Row{  10, -5.3,  22   },
    matrix.Row{  -2, -25,   12   },
    matrix.Row{   7,  5,   -12.5 },
  },
)
```

It returns an error if you try to provide a builder with no rows or
rows with no cols (you can safely ignore error if you're confident
your builder is valid).


### `func RandomMatrix( rows, cols int ) Matrix`

Generate a matrix of `rows` rows and `cols` cols with randomized
values.

Values are randomized with standard normal distribution, using
`math.NormFloat64()` (don't forget to seed randomizer if you
want it to be truly random).


## Read matrices

### `func ( matrix Matrix ) Rows() int`

Returns the number of rows


### `func ( matrix Matrix ) Cols() int`

Returns the number of columns


### `func ( matrix Matrix ) String() string`

Returns a human readable representation of matrix, ready to print


### `func ( matrix Matrix ) At( row, col int ) float64`

Returns the value at position `row`, `col`.

Just like an array, you're responsible to make sure
you don't ask for an out of range value.


### Looping on a matrix

You can loop on a matrix this way:

```go
myMatrix =: matrix.RandomMatrix( 5, 5 )
for i := 0 ; i < myMatrix.Rows() ; i++ {
  for j := 0 ; j < myMatrix.Cols() ; i++ {
    printf( "%v\n", myMatrix.At( i, j ) )
  }
}
```

If it's too costly for you performance wise, see the `Low level implementation`
section at the end of this doc.


## Test matrices

### `func ( matrix Matrix ) Valid() bool`

A matrix is valid if it has rows, columns and if all its rows have the same
amount of columns.


### `func ( matrix Matrix ) SameDimensions( otherMatrix Matrix ) bool`

True if both matrices are valid and have the same dimensions.


### `func ( matrix Matrix ) EqualTo( otherMatrix Matrix ) bool`

Tells if two valid matrices have the same dimensions and same
values in each cell.


## Operations on matrices

Those are the operations currently implemented. Note that all operations
produce a new matrix and return it, current matrix and argument matrix are
never modified.


### `func ( matrix Matrix ) ScalarMultiply( scalar float64 ) ( resultMatrix Matrix, err error )`

Multiply the matrix with a scalar.

Error is returned if matrix is not valid.


### `func ( matrix Matrix ) DotProduct( otherMatrix Matrix ) ( resultMatrix Matrix, err error )`

Perform a mathematical standard multiplication between matrix and otherMatrix, and return
the resulting resultMatrix.

Error is returned if resultMatrix is undefined (that is, if matrix columns count is not
the same than otherMatrix rows count).


### `func ( matrix Matrix ) Transpose() ( resultMatrix Matrix, err error )`

Switch matrix dimensions, so that, eg, a 2x3 matrix returns
a 3x2 one.

Error is returned if matrix is not valid.


### `func ( matrix Matrix ) MultiplyCells( otherMatrix Matrix ) ( resultMatrix Matrix, err error )`

Multiply each cell from matrix with each cell at the same coordinate in otherMatrix.

Note that *this is not* standard mathematical matrix multiplication. For that one,
use `DotProduct()`.

Error is returned if matrices do not have the same dimensions.


### `func ( matrix Matrix ) Add( otherMatrix Matrix ) ( resultMatrix Matrix, err error )`

Add otherMatrix to matrix and return the resulting resultMatrix.

Error is returned if matrices are not valid or do not have the same dimensions.


### `func ( matrix Matrix ) Substract( otherMatrix Matrix ) ( resultMatrix Matrix, err error )`

Substract otherMatrix from matrix and return the resulting resultMatrix.

Error is returned if matrices are not valid or do not have the same dimensions.


### `func ( matrix Matrix ) Sigmoid() ( resultMatrix Matrix, err error )`

Apply sigmoid function on each cell of matrix and return resulting Matrix.

Error is returned if matrix is not valid.


### `func ( matrix Matrix ) SigmoidDerivative() ( resultMatrix Matrix, err error )`

Compute derivative for sigmoid function on each cell of matrix and return resulting Matrix.

Error is returned if matrix is not valid.


## Extending

Two generic operations are provided that should allow you to perform any cell
by cell operation by taking one or two matrices as an input: `UnaryOperation`
and `BinaryOperation`.


### `func ( matrix Matrix ) UnaryOperation( operation func( float64 ) float64, operationName string ) ( resultMatrix Matrix, err error )`

Produce a new matrix by applying `operation` cell by cell on matrix, so that:

```
operation( cell ) -> resultCell
```

`operationName` is a name you provide for your operation so that it's easier to
know where the error comes from (error will contain that operation name).

`operation` is a function you provide, that will receive the `float64` value of
each cell, and should return the new `float64` value for that cell.

So, for example, if you want to multiply each value by 2 and add 1:

```go
operation := func( value float64 ) float64 {
  return value * 2 + 1
}

resultMatrix, err := myMatrix.UnaryOperation( operation, "x * 2 + 1" )
```

Error is returned is the matrix is not valid.


### `func ( matrix Matrix ) BinaryOperation( otherMatrix Matrix, operation func( float64, float64 ) float64, operationName string ) ( resultMatrix Matrix, err error )`

Produce a new matrix by applying `operation` cell by cell on two matrices, so that:
   operation( cell1, cell2 ) -> resultCell


`operationName` is used in error message, so that it's easier to know which
operation error'd.

So, for example, if you want to divide each cell of matrix by corresponding cell in otherMatrix:

```go
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

```go
matrix.SetDebug( true )
```

## Low level implementation

Under the hood, a Matrix is a `[]float64`. First entry is the number of rows,
second entry is the number of cols.

To retrieve the index of a value for a matrix position in that array, you can
use `IndexFor( row, col int ) float64`:

```go
myMatrix, _ := matrix.Build(
  matrix.Builder {
    matrix.Row{  10.0, -5.3,   22.0 },
    matrix.Row{ -2.0,  -25.0,  12.0 },
    matrix.Row{  7.0,   5.3,  -12.5 },
  },
)
println( myMatrix.IndexFor( 1, 2 ) ) // 7
fmt.Printf( "%v\n", myMatrix[7] ) // 12
```

This implementation was chosen because it's 50% faster than using a
`[][]float64` to map the matrix. It's also slightly faster than using a
`struct{ Rows int, Cols int, Values []float64 }` and prevent having to pass the
Matrix by reference everywhere: since it's a slice, it's always a reference.

if you need to iterate directly on matrix values for some performance critical
operation, you can do it this way:

```go
for i := 2 ; i < len( myMatrix ); i++ {
 // ...
}
```

Or, to iterate on rows:

```go
for i := 2 ; i < len( myMatrix ); i += myMatrix[1] {
 // `i` is the start of a row
}
```
