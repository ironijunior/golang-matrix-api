package main

import ("strconv")

/*
	invertMatrix : Invert a given matrix transposing columns and rows
*/
func invertMatrix(records [][]string) [][]string {
	numRow := len(records)

	if numRow == 0 {
		return [][]string{}
	}

	numColumn := len(records[0])
	invertedMatrix := make([][]string, numColumn)

	for i := 0; i < numColumn; i++ {
		invertedMatrix[i] = make([]string, numRow)
		for j := 0; j < numRow; j++ {
			invertedMatrix[i][j] = records[j][i]
		}
	}
	return invertedMatrix
}

/*
	flattenMatrix : Flatten a given matrix transforming all columns in a single row.
*/
func flattenMatrix(records [][]string) []string {
	var flattened []string
	for _, row := range records {
		flattened = append(flattened, row...)
	}
	return flattened
}

/*
	sumMatrix : Sum all numbers of a given matrix .
*/
func sumMatrix(records [][]string) (int, error) {
	var sum int
	for _, row := range records {
		for _, num := range row {
			i, err := strconv.Atoi(num)
			if err != nil {
				return 0, err
			}
			sum = sum + i
		}
	}
	return sum, nil
}

/*
	multiplyMatrix : Multiply all numbers of a given matrix .
*/
func multiplyMatrix(records [][]string) (int, error) {
	if (len(records) <= 0) {return 0, nil}
	
	var mult int = 1
	for _, row := range records {
		for _, num := range row {
			i, err := strconv.Atoi(num)
			if err != nil {
				return 0, err
			}
			mult = mult * i
		}
	}
	return mult, nil
}