// Package matrix provides utilities to work with mathematical matrices.
//
// This is not intended to exhaustively implements all possible mathematical
// operations on matrices. Instead, it's just what I needed to build a neural
// network. That being said, it provides build blocks to easily extend it (see
// [`Extending` section](#extending)). If you feel some operation should be in
// there, feel free to send a pull request.
//
// You can generate an initialized zero matrix providing its number of rows and
// cols:
//
//     matrix := GenerateMatrix(3, 2)
//
// Or, you can use the builder to provide matrix values in an human readable way.
//
// Two types are provided to build `Matrix` : `Builder` and `Row`.
//
//     myMatrix, err := matrix.Build(
//     	matrix.Builder{
//     		matrix.Row{10.0, -5.3, 22.0},
//     		matrix.Row{-2.0, -25.0, 12.0},
//     		matrix.Row{7.0, 5.3, -12.5},
//     	},
//     )
//     if err != nil {
//     	fmt.Println("You passed a 0x0 or 1x0 matrix.")
//     }
//
// All rows will have the same amount of columns than the first one. If subsequent
// has less columns, the remaining ones will be filled with 0.
package matrix
