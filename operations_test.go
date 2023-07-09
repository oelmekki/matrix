package matrix

import (
	"testing"
)

func TestScalarMultiply(t *testing.T) {
	t.Run("with valid matrix", func(t *testing.T) {
		baseMatrix, err := Build(
			Builder{
				Row{1, 1, 1},
				Row{1, 1, 1},
				Row{1, 1, 1},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix while none was expected: %v", err)
		}

		matrix, err := baseMatrix.ScalarMultiply(2)
		if err != nil {
			t.Fatalf("Got an error while multiplying matrix while none was expected: %v", err)
		}

		for i := 0; i < matrix.Rows(); i++ {
			for j := 0; j < matrix.Cols(); j++ {
				val := matrix.At(i, j)
				if val != 2 {
					t.Errorf("Improperly multiplying at (%d, %d) : got value %f instead of 2.", i, j, val)
				}
			}
		}
	})

	t.Run("with an invalid matrix", func(t *testing.T) {
		baseMatrix := Matrix([]float64{10, 10, 1})
		_, err := baseMatrix.ScalarMultiply(2)
		if err == nil {
			t.Fatalf("Got no error from invalid matrix")
		}
	})
}

func TestDotProduct(t *testing.T) {
	t.Run("performing a valid operation", func(t *testing.T) {
		matrix1, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix1 while none was expected: %v", err)
		}

		matrix2, err := Build(
			Builder{
				Row{10, 11},
				Row{12, 13},
				Row{14, 15},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix2 while none was expected: %v", err)
		}

		expected, err := Build(
			Builder{
				Row{76, 82},
				Row{184, 199},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix2 while none was expected: %v", err)
		}

		actual, err := matrix1.DotProduct(matrix2)
		if err != nil {
			t.Fatalf("Got an error while none was expected: %v", err)
		}

		if !expected.EqualTo(actual) {
			t.Errorf("Expected :%s\nGot:%s", expected, actual)
		}
	})

	t.Run("performing an invalid operation", func(t *testing.T) {
		matrix1, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix1 while none was expected: %v", err)
		}

		matrix2, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix2 while none was expected: %v", err)
		}

		_, err = matrix1.DotProduct(matrix2)
		if err == nil {
			t.Fatalf("Got no error while requesting for an undefined matrix.")
		}
	})
}

func TestVectorMultiply(t *testing.T) {
	t.Run("with valid operation", func(t *testing.T) {
		matrix1, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix1 while none was expected: %v", err)
		}

		vector := []float64{7, 8, 9}
		expected := []float64{50, 122}

		actual, err := matrix1.VectorMultiply(vector)
		if err != nil {
			t.Fatalf("Got an error while none was expected: %v", err)
		}

		if len(expected) != len(actual) {
			t.Fatalf("Expected length (%d) different from actual length (%d).", len(expected), len(actual))
		}

		for i, val := range expected {
			if actual[i] != val {
				t.Errorf("At position %d, expected %f, got %f.", i, val, actual[i])
			}
		}
	})

	t.Run("with invalid operation", func(t *testing.T) {
		matrix1, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix1 while none was expected: %v", err)
		}

		vector := []float64{7}

		_, err = matrix1.VectorMultiply(vector)
		if err == nil {
			t.Fatalf("Got no error while using an invalid operation.")
		}
	})
}

func TestSetAt(t *testing.T) {
	matrix := GenerateMatrix(2, 2)
	matrix.SetAt(1, 1, 10)
	value := matrix.At(1, 1)
	if value != 10 {
		t.Errorf("Expect 10, got %f", value)
	}
}

func TestTranspose(t *testing.T) {
	t.Run("with a valid matrix", func(t *testing.T) {
		matrix, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix while none was expected: %v", err)
		}

		expected, err := Build(
			Builder{
				Row{1, 4},
				Row{2, 5},
				Row{3, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building expected matrix while none was expected: %v", err)
		}

		actual, err := matrix.Transpose()
		if err != nil {
			t.Fatalf("Got an error while none was expected: %v", err)
		}

		if !actual.EqualTo(expected) {
			t.Errorf("Expected :%s\nGot:%s", expected, actual)
		}
	})

	t.Run("with an invalid matrix", func(t *testing.T) {
		matrix := Matrix([]float64{10, 10, 1})
		_, err := matrix.Transpose()
		if err == nil {
			t.Fatalf("Got no error with an invalid matrix.")
		}
	})
}

func TestMultiplyCells(t *testing.T) {
	t.Run("with valid matrices", func(t *testing.T) {
		matrix1, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix1 while none was expected: %v", err)
		}

		matrix2, err := Build(
			Builder{
				Row{2, 3, 4},
				Row{5, 6, 7},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix2 while none was expected: %v", err)
		}

		expected, err := Build(
			Builder{
				Row{2, 6, 12},
				Row{20, 30, 42},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building expected matrix while none was expected: %v", err)
		}

		actual, err := matrix1.MultiplyCells(matrix2)
		if err != nil {
			t.Fatalf("Got an error while none was expected: %v", err)
		}

		if !actual.EqualTo(expected) {
			t.Errorf("Expected :%s\nGot:%s", expected, actual)
		}
	})

	t.Run("with invalid matrices", func(t *testing.T) {
		matrix1, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix1 while none was expected: %v", err)
		}

		matrix2, err := Build(
			Builder{
				Row{2, 3},
				Row{5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix2 while none was expected: %v", err)
		}

		_, err = matrix1.MultiplyCells(matrix2)
		if err == nil {
			t.Fatalf("Got no error with an invalid matrix.")
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("with valid matrices", func(t *testing.T) {
		matrix1, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix1 while none was expected: %v", err)
		}

		matrix2, err := Build(
			Builder{
				Row{2, 3, 4},
				Row{5, 6, 7},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix2 while none was expected: %v", err)
		}

		expected, err := Build(
			Builder{
				Row{3, 5, 7},
				Row{9, 11, 13},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building expected matrix while none was expected: %v", err)
		}

		actual, err := matrix1.Add(matrix2)
		if err != nil {
			t.Fatalf("Got an error while none was expected: %v", err)
		}

		if !actual.EqualTo(expected) {
			t.Errorf("Expected :%s\nGot:%s", expected, actual)
		}
	})

	t.Run("with invalid matrices", func(t *testing.T) {
		matrix1, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix1 while none was expected: %v", err)
		}

		matrix2, err := Build(
			Builder{
				Row{2, 3},
				Row{5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix2 while none was expected: %v", err)
		}

		_, err = matrix1.Add(matrix2)
		if err == nil {
			t.Fatalf("Got no error with an invalid matrix.")
		}
	})
}

func TestSubstract(t *testing.T) {
	t.Run("with valid matrices", func(t *testing.T) {
		matrix1, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix1 while none was expected: %v", err)
		}

		matrix2, err := Build(
			Builder{
				Row{2, 3, 4},
				Row{5, 6, 7},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix2 while none was expected: %v", err)
		}

		expected, err := Build(
			Builder{
				Row{-1, -1, -1},
				Row{-1, -1, -1},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building expected matrix while none was expected: %v", err)
		}

		actual, err := matrix1.Substract(matrix2)
		if err != nil {
			t.Fatalf("Got an error while none was expected: %v", err)
		}

		if !actual.EqualTo(expected) {
			t.Errorf("Expected :%s\nGot:%s", expected, actual)
		}
	})

	t.Run("with invalid matrices", func(t *testing.T) {
		matrix1, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix1 while none was expected: %v", err)
		}

		matrix2, err := Build(
			Builder{
				Row{2, 3},
				Row{5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix2 while none was expected: %v", err)
		}

		_, err = matrix1.Substract(matrix2)
		if err == nil {
			t.Fatalf("Got no error with an invalid matrix.")
		}
	})
}

func TestSigmoid(t *testing.T) {
	t.Run("with a valid matrix", func(t *testing.T) {
		matrix, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix while none was expected: %v", err)
		}

		expected, err := Build(
			Builder{
				Row{0.7310585786300049, 0.8807970779778823, 0.9525741268224334},
				Row{0.9820137900379085, 0.9933071490757153, 0.9975273768433653},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building expected matrix while none was expected: %v", err)
		}

		actual, err := matrix.Sigmoid()
		if err != nil {
			t.Fatalf("Got an error while none was expected: %v", err)
		}

		if !actual.EqualTo(expected) {
			t.Errorf("Expected :%s\nGot:%s", expected, actual)
		}
	})

	t.Run("with an invalid matrix", func(t *testing.T) {
		matrix := Matrix([]float64{10, 10, 1})

		_, err := matrix.Sigmoid()
		if err == nil {
			t.Fatalf("Got no error with an invalid matrix.")
		}
	})
}

func TestSigmoidDerivative(t *testing.T) {
	t.Run("with a valid matrix", func(t *testing.T) {
		matrix, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix while none was expected: %v", err)
		}

		expected, err := Build(
			Builder{
				Row{0.19661193324148185, 0.10499358540350662, 0.045176659730912},
				Row{0.017662706213291107, 0.006648056670790033, 0.002466509291359931},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building expected matrix while none was expected: %v", err)
		}

		actual, err := matrix.SigmoidDerivative()
		if err != nil {
			t.Fatalf("Got an error while none was expected: %v", err)
		}

		if !actual.EqualTo(expected) {
			t.Errorf("Expected :%s\nGot:%s", expected, actual)
		}
	})

	t.Run("with an invalid matrix", func(t *testing.T) {
		matrix := Matrix([]float64{10, 10, 1})

		_, err := matrix.SigmoidDerivative()
		if err == nil {
			t.Fatalf("Got no error with an invalid matrix.")
		}
	})
}

func TestBinaryOperation(t *testing.T) {
	t.Run("with valid matrices", func(t *testing.T) {
		matrix1, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix1 while none was expected: %v", err)
		}

		matrix2, err := Build(
			Builder{
				Row{2, 3, 4},
				Row{5, 6, 7},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix2 while none was expected: %v", err)
		}

		expected, err := Build(
			Builder{
				Row{12, 16, 22},
				Row{30, 40, 52},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building expected matrix while none was expected: %v", err)
		}

		actual, err := matrix1.BinaryOperation(matrix2, func(a float64, b float64) float64 {
			return a*b + 10
		}, "a * b + 10")
		if err != nil {
			t.Fatalf("Got an error while none was expected: %v", err)
		}

		if !actual.EqualTo(expected) {
			t.Errorf("Expected :%s\nGot:%s", expected, actual)
		}
	})

	t.Run("with invalid matrices", func(t *testing.T) {
		matrix1, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix1 while none was expected: %v", err)
		}

		matrix2, err := Build(
			Builder{
				Row{2, 3},
				Row{5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix2 while none was expected: %v", err)
		}

		_, err = matrix1.BinaryOperation(matrix2, func(a float64, b float64) float64 {
			return a*b + 10
		}, "a * b + 10")
		if err == nil {
			t.Fatalf("Got no error with an invalid matrix.")
		}
	})
}

func TestUnaryOperation(t *testing.T) {
	t.Run("with a valid matrix", func(t *testing.T) {
		matrix, err := Build(
			Builder{
				Row{1, 2, 3},
				Row{4, 5, 6},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building matrix while none was expected: %v", err)
		}

		expected, err := Build(
			Builder{
				Row{11, 14, 19},
				Row{26, 35, 46},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building expected matrix while none was expected: %v", err)
		}

		actual, err := matrix.UnaryOperation(func(a float64) float64 {
			return a*a + 10
		}, "a * a + 10")
		if err != nil {
			t.Fatalf("Got an error while none was expected: %v", err)
		}

		if !actual.EqualTo(expected) {
			t.Errorf("Expected :%s\nGot:%s", expected, actual)
		}
	})

	t.Run("with an invalid matrix", func(t *testing.T) {
		matrix := Matrix([]float64{10, 10, 1})

		_, err := matrix.UnaryOperation(func(a float64) float64 {
			return a*a + 10
		}, "a * a + 10")
		if err == nil {
			t.Fatalf("Got no error with an invalid matrix.")
		}
	})
}
