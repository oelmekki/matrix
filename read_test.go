package matrix

import "testing"

func TestRows(t *testing.T) {
	matrix := GenerateMatrix(3, 2)
	rows := matrix.Rows()

	if rows != 3 {
		t.Errorf("Expected 3 rows, got %d.", rows)
	}
}

func TestCols(t *testing.T) {
	matrix := GenerateMatrix(3, 2)
	cols := matrix.Cols()

	if cols != 2 {
		t.Errorf("Expected 2 cols, got %d.", cols)
	}
}

func TestString(t *testing.T) {
	matrix := GenerateMatrix(3, 2)
	expected := `
{		0		0		}
{		0		0		}
{		0		0		}
`
	actual := matrix.String()
	if actual != expected {
		t.Errorf("Expected:\n%#v\nGot:\n%#v", expected, actual)
	}
}

func TestAt(t *testing.T) {
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

	value := matrix.At(2, 1)
	if value != 5 {
		t.Errorf("Expected 5, got %f", value)
	}
}

func TestIndexFor(t *testing.T) {
	matrix := GenerateMatrix(2, 3)
	index := matrix.IndexFor(1, 1)
	if index != 6 {
		t.Errorf("Expected 6, got %d", index)
	}
}

func TestGetRow(t *testing.T) {
	t.Run("with valid argument", func(t *testing.T) {
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

		row, err := matrix.GetRow(1)
		if err != nil {
			t.Fatalf("Got an unexpected error: %v", err)
		}

		if len(row) != 3 {
			t.Errorf("Expected a row with 3 elements, got %d", len(row))
		}

		expected := []float64{-2, -25, 12}
		for i, val := range expected {
			if row[i] != val {
				t.Errorf("At position %d, expected %f, got %f", i, row[i], val)
			}
		}
	})

	t.Run("with invalid argument", func(t *testing.T) {
		matrix, err := Build(
			Builder{
				Row{10, -5.3, 22},
				Row{-2, -25, 12},
				Row{7, 5, -12.5},
			},
		)
		if err != nil {
			t.Fatalf("Got an error while building the matrix while none was expected: %v", err)
		}

		_, err = matrix.GetRow(3)
		if err == nil {
			t.Fatalf("Got no error while requesting out of bound row")
		}
	})
}
