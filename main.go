package main

import (
	"fmt"
	"net/http"
	"strings"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
func echoHandler(w http.ResponseWriter, r *http.Request) {
	records, err := readRequestFile(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	fmt.Fprint(w, transformMatrixToString(records))
}

func invertHandler(w http.ResponseWriter, r *http.Request) {
	records, err := readRequestFile(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	var invertedMatrix = invertMatrix(records)
	fmt.Fprint(w, transformMatrixToString(invertedMatrix))
}

func flattenHandler(w http.ResponseWriter, r *http.Request) {
	records, err := readRequestFile(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	var flattened = flattenMatrix(records)	
	fmt.Fprint(w, strings.Join(flattened, ","))
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	records, err := readRequestFile(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	sum, err := sumMatrix(records)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	fmt.Fprint(w, sum)
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	records, err := readRequestFile(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	mult, err := multiplyMatrix(records)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	fmt.Fprint(w, mult)
}

func main() {
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/invert", invertHandler)
	http.HandleFunc("/flatten", flattenHandler)
	http.HandleFunc("/sum", sumHandler)
	http.HandleFunc("/multiply", multiplyHandler)

	http.ListenAndServe(":8080", nil)
}