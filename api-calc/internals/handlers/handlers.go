package handlers

import (
	"encoding/json"
	"net/http"
)

type Input struct {
	A int `json:"a"`
	B int `json:"b"`
}

func Add(w http.ResponseWriter, r *http.Request) {

	var input Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid inputs", http.StatusBadRequest)
		return
	}

	if input.A == 0 || input.B == 0 {
		http.Error(w, "Cant add zero", http.StatusBadRequest)
		return
	}

	result := input.A + input.B

	response := map[string]int{"result": result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func Subtract(w http.ResponseWriter, r *http.Request) {

	var input Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid inputs", http.StatusBadRequest)
		return
	}

	if input.A == 0 || input.B == 0 {
		http.Error(w, "Cant subtract zero", http.StatusBadRequest)
		return
	}

	result := input.A - input.B

	response := map[string]int{"result": result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func Multiply(w http.ResponseWriter, r *http.Request) {
	var input Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid inputs", http.StatusBadRequest)
		return
	}

	if input.A == 0 || input.B == 0 {
		http.Error(w, "Cant multiply zero", http.StatusBadRequest)
		return
	}

	result := input.A * input.B

	response := map[string]int{"result": result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var input Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid inputs", http.StatusBadRequest)
		return
	}

	if input.A == 0 || input.B == 0 {
		http.Error(w, "Cant divide zero", http.StatusBadRequest)
		return
	}

	result := input.A / input.B

	response := map[string]int{"result": result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
