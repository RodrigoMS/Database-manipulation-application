package view

import (
	"encoding/json"
	"net/http"
)

type output interface{}

func sendResult(w http.ResponseWriter, result output) {
	// Converta os usuários para JSON
	jsonData, err := json.Marshal(result)
	if err != nil {
		HandleInternalServerError(w, "Error converting to JSON")
		return
	}

	// Defina o cabeçalho de conteúdo para JSON e escreva os dados JSON no corpo da resposta
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		HandleInternalServerError(w, "Error writing in response")
	}
}

// Status code 400
func HandleNotFound(w http.ResponseWriter, err error) {
	
		if err == nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
	
		http.Error(w, err.Error(), http.StatusBadRequest)
}

// Status code 405
func HandleMethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// Status code 500 
func HandleInternalServerError(w http.ResponseWriter, errorMessage string) {
	http.Error(w, errorMessage, http.StatusInternalServerError)
}

// Status code 200
func HandleSuccess(w http.ResponseWriter, result output) {
	w.WriteHeader(http.StatusOK)
	sendResult(w, result)
}

// Status code 201
func HandleResourceCreated(w http.ResponseWriter, result output) {
	w.WriteHeader(http.StatusCreated)
	sendResult(w, result)
}

// Status code 204
func HandleNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// Status code 500 
func Response500(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
