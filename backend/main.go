package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	json.NewEncoder(w).Encode([]int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func HandleRequest() {
	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Server started on: http://localhost:8080")
	HandleRequest()
}
