package main

import (
	"testing"
	"errors"
)

func TestInvertMatrix(t *testing.T) {
	tables := []struct {
		input [][]string
		expected [][]string
	}{
		{[][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			[][]string{{"1", "4", "7"}, {"2", "5", "8"}, {"3", "6", "9"}}},

		{[][]string{{"1", "1", "1"}, {"-3", "4", "5"}, {"2s", "3", "a"}},
			[][]string{{"1", "-3", "2s"}, {"1", "4", "3"}, {"1", "5", "a"}}},

		{[][]string{{"a", "-1", "999"}, {"0", "4", "5"}, {"2a", "3", "-99"}, {"4", "3", "99"}},
			[][]string{{"a", "0", "2a", "4"}, {"-1", "4", "3", "3"}, {"999", "5", "-99", "99"}}},

		{[][]string{}, [][]string{}},
		{[][]string{{"a"}}, [][]string{{"a"}}},
	}

	for _, table := range tables {
		inverted := invertMatrix(table.input)
		if len(inverted) != len(table.expected) {
			t.Errorf("Inverted matrix doesnt match the length of expected, got: %v, want: %v.",
				len(inverted), len(table.expected))
		}

		if !arrayEquals(inverted, table.expected) {
			t.Errorf("INverting was incorrect, got: %v, want: %v.", inverted, table.expected)
		}
	}
}

func TestFlattenMatrix(t *testing.T) {
	tables := []struct {
		input [][]string
		expected []string
	}{
		{[][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			[]string{"1","2","3","4","5","6","7","8","9"}},

		{[][]string{{"1", "1", "1"}, {"-3", "4", "5"}, {"2s", "3", "a"}},
			[]string{"1","1","1","-3","4","5","2s","3","a"}},

		{[][]string{{"a", "-1", "999"}, {"0", "4", "5"}, {"2a", "3", "-99"}, {"4", "3", "99"}},
			[]string{"a","-1","999","0","4","5","2a","3","-99","4","3","99"}},

		{[][]string{}, []string{}},
		{[][]string{{"a"}}, []string{"a"}},
	}

	for _, table := range tables {
		flattened := flattenMatrix(table.input)
		if len(flattened) != len(table.expected) {
			t.Errorf("Flattened matrix doesnt match the length of expected, got: %v, want: %v.",
				len(flattened), len(table.expected))
		}

		if !arrayEqualsSimple(flattened, table.expected) {
			t.Errorf("Flattening was incorrect, got: %v, want: %v.", flattened, table.expected)
		}
	}
}

func TestSumMatrix(t *testing.T) {
	tables := []struct {
		input [][]string
		expected int
	}{
		{[][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			45},

		{[][]string{{"1", "1", "1"}, {"-3", "4", "5"}, {"0", "3", "8"}},
			20},

		{[][]string{{"1", "-1", "999"}, {"0", "4", "5"}, {"2", "3", "-99"}, {"4", "3", "199"}},
			1120},

		{[][]string{}, 0},
		{[][]string{{"1"}}, 1},
	}

	for _, table := range tables {
		total, err := sumMatrix(table.input)
		if err != nil {
			t.Errorf("Error during sum operation %d.", err)
		}
		if total != table.expected {
			t.Errorf("Sum of (%s) was incorrect, got: %d, want: %d.", table.input, total, table.expected)
		}
	}
}

func TestMultiplyMatrix(t *testing.T) {
	tables := []struct {
		input [][]string
		expected int
	}{
		{[][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			362880},

		{[][]string{{"1", "1", "1"}, {"-3", "4", "5"}, {"1", "-3", "8"}},
			1440},

		{[][]string{{"1", "-1", "999"}, {"0", "4", "5"}, {"2", "3", "-99"}, {"4", "3", "199"}},
			0},

		{[][]string{}, 0},
		{[][]string{{"1"}}, 1},
	}

	for _, table := range tables {
		total, err := multiplyMatrix(table.input)
		if err != nil {
			t.Errorf("Error during multiply operation %d.", err)
		}
		if total != table.expected {
			t.Errorf("Multiply of (%s) was incorrect, got: %d, want: %d.", table.input, total, table.expected)
		}
	}
}

func TestNotSquareMatrix(t *testing.T) {
	tables := []struct {
		input [][]string
		expected error
	}{
		{[][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			errors.New("matrix is not a square")},
	}

	for _, table := range tables {
		square := isSquareMatrix(table.input)
		
		if square {
			t.Errorf("Invalid squared Matrix was considered valid.")
		}
	}
}

/*
Helper to compare Arrays
*/
func arrayEquals(arr1 [][]string, arr2 [][]string) bool {
	for i := range arr1 {
		for j := range arr1[i] {
			if arr1[i][j] != arr2[i][j] {
				return false
			}
		}
	}

	return true
}

func arrayEqualsSimple(arr1 []string, arr2 []string) bool {
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}