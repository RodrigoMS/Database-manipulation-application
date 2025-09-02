package views

import (
	"net/http"

	"github.com/RodrigoMS/app/pkg/utils"
)

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not found", http.StatusNotFound)
}

func HandleMethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func HandleInternalServerError(w http.ResponseWriter, errorMessage string) {
	http.Error(w, errorMessage, http.StatusInternalServerError)
}

func HandleSuccess[T any](w http.ResponseWriter, data T) {
	description, jsonData, err := utils.WriteJson(data)

	if err != nil {
		HandleInternalServerError(w, description)
		return
	}

	w.Header().Set("Content-Type", description)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func HandleResourceCreated(w http.ResponseWriter, data any) {

	description, jsonData, err := utils.WriteJson(data)
	// Verifica se houve erro na conversão para JSON
	if err != nil {
		HandleInternalServerError(w, description)
		return
	}

	w.Header().Set("Content-Type", description)
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonData)
	// Verifica se houve erro na escrita na resposta HTTP.
	if err != nil {
		HandleInternalServerError(w, "Error writing in response")
	}
}

func HandleNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
