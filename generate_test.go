package matrix

import "testing"

func TestGenerateMatrix(t *testing.T) {
	matrix := GenerateMatrix(3, 10)

	if matrix.Rows() != 3 {
		t.Errorf("Expected 3 rows, got %d", matrix.Rows())
	}

	if matrix.Cols() != 10 {
		t.Errorf("Expected 10 cols, got %d", matrix.Cols())
	}

	for i := 0; i < matrix.Rows(); i++ {
		for j := 0; j < matrix.Cols(); j++ {
			val := matrix.At(i, j)
			if val != 0 {
				t.Errorf("Improperly initialized matrix at %d, %d : got value %f instead of 0.", i, j, val)
			}
		}
	}
}

func TestBuild(t *testing.T) {
	t.Run("with valid matrix", func(t *testing.T) {
		matrix, err := Build(
			Builder{
				Row{10, -5.3, 22},
				Row{-2, -25, 12},
				Row{7, 5, -12.5},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while none was expected: %v", err)
		}

		if matrix.Rows() != 3 {
			t.Errorf("Expected 3 rows, got %d", matrix.Rows())
		}

		if matrix.Cols() != 3 {
			t.Errorf("Expected 10 cols, got %d", matrix.Cols())
		}

		if matrix.At(0, 0) != 10 {
			t.Errorf("Expected 10, at (0, 0), got %f", matrix.At(0, 0))
		}

		if matrix.At(2, 2) != -12.5 {
			t.Errorf("Expected -12.5, at (2, 2), got %f", matrix.At(2, 2))
		}
	})

	t.Run("with invalid matrix", func(t *testing.T) {
		_, err := Build(
			Builder{},
		)
		if err == nil {
			t.Errorf("Got no error with undefined matrix: %v", err)
		}

		_, err = Build(
			Builder{
				Row{},
			},
		)
		if err == nil {
			t.Errorf("Got no error with undefined row: %v", err)
		}
	})
}

func TestZeroMatrixFrom(t *testing.T) {
	otherMatrix, err := Build(
		Builder{
			Row{10, -5.3, 22},
			Row{-2, -25, 12},
			Row{7, 5, -12.5},
		},
	)
	if err != nil {
		t.Fatalf("Got an error while none was expected: %v", err)
	}

	matrix := ZeroMatrixFrom(otherMatrix)

	if matrix.Rows() != 3 {
		t.Errorf("Expected 3 rows, got %d", matrix.Rows())
	}

	if matrix.Cols() != 3 {
		t.Errorf("Expected 3 cols, got %d", matrix.Cols())
	}

	for i := 0; i < matrix.Rows(); i++ {
		for j := 0; j < matrix.Cols(); j++ {
			val := matrix.At(i, j)
			if val != 0 {
				t.Errorf("Improperly initialized matrix at %d, %d : got value %f instead of 0.", i, j, val)
			}
		}
	}
}

func TestRandomMatrix(t *testing.T) {
	matrix := RandomMatrix(3, 4)

	if matrix.Rows() != 3 {
		t.Errorf("Expected 3 rows, got %d", matrix.Rows())
	}

	if matrix.Cols() != 4 {
		t.Errorf("Expected 4 cols, got %d", matrix.Cols())
	}

	for i := 0; i < matrix.Rows(); i++ {
		for j := 0; j < matrix.Cols(); j++ {
			if matrix.At(i, j) == 0 {
				t.Errorf("Improperly initialized matrix at %d, %d : expected random value, got 0.", i, j)
			}
		}
	}
}
