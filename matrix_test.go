package matrix

import "testing"

func TestSetDebug(t *testing.T) {
	SetDebug(true)
	if !DEBUG {
		t.Errorf("Did not properly set debug mode.")
	}

	SetDebug(false)
}

func TestValid(t *testing.T) {
	t.Run("with valid matrix", func(t *testing.T) {
		matrix := GenerateMatrix(3, 2)
		if !matrix.Valid() {
			t.Errorf("Valid is returning false with a matrix straight from GenerateMatrix.")
		}
	})

	t.Run("with invalid matrix", func(t *testing.T) {
		matrix := Matrix([]float64{2, 2, 0, 1})
		if matrix.Valid() {
			t.Errorf("Valid is returning true with garbage matrix.")
		}
	})
}

func TestEqualTo(t *testing.T) {
	t.Run("with equal matrices", func(t *testing.T) {
		matrix1 := GenerateMatrix(2, 3)
		matrix2 := GenerateMatrix(2, 3)

		if !matrix1.EqualTo(matrix2) {
			t.Errorf("EqualTo returns false while it should be true.")
		}
	})

	t.Run("with inequal matrices", func(t *testing.T) {
		matrix1 := RandomMatrix(2, 3)
		matrix2 := ZeroMatrixFrom(matrix1)

		if matrix1.EqualTo(matrix2) {
			t.Errorf("EqualTo returns true when matrices has different values.")
		}

		matrix1 = GenerateMatrix(2, 3)
		matrix2 = GenerateMatrix(3, 2)

		if matrix1.EqualTo(matrix2) {
			t.Errorf("EqualTo returns true when matrices has different dimensions.")
		}

		matrix1 = GenerateMatrix(2, 3)
		matrix2 = GenerateMatrix(20, 30)

		if matrix1.EqualTo(matrix2) {
			t.Errorf("EqualTo returns true when matrices has different dimensions.")
		}
	})
}

func TestSameDimensions(t *testing.T) {
	t.Run("with matrices of the same dimensions", func(t *testing.T) {
		matrix1 := GenerateMatrix(2, 3)
		matrix2 := GenerateMatrix(2, 3)

		if !matrix1.SameDimensions(matrix2) {
			t.Errorf("SameDimensions returns false while it should be true.")
		}
	})

	t.Run("with matrices of different dimensions", func(t *testing.T) {
		matrix1 := GenerateMatrix(2, 3)
		matrix2 := GenerateMatrix(3, 2)

		if matrix1.SameDimensions(matrix2) {
			t.Errorf("SameDimensions returns true while it should be false.")
		}
	})

	t.Run("with invalid matrices", func(t *testing.T) {
		matrix1 := Matrix([]float64{10, 10, 1})
		matrix2 := Matrix([]float64{10, 10, 1})

		if matrix1.SameDimensions(matrix2) {
			t.Errorf("SameDimensions does not issue error on invalid matrices.")
		}
	})
}
