package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
	"errors"
)

/*
readRequestFile : Helper to get the file from request
 */
 func readRequestFile(request *http.Request) ([][]string, error) {
	file, _, err := request.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	if (!isSquareMatrix(records)) {
		return nil, errors.New("matrix is not a square")
	}

	return records, nil
}

func isSquareMatrix(records [][]string) bool {
	return (len(records[0]) == len(records))
}

/*
transformMatrixToString : Helper to transform a Matrix to String-like format
*/
func transformMatrixToString(matrix [][]string) string {
	var stringMatrix string
	for _, row := range matrix {
		stringMatrix = fmt.Sprintf("%s%s\n", stringMatrix, strings.Join(row, ","))
	}
	return stringMatrix
}