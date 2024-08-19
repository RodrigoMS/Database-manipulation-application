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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Defina o cabeçalho de conteúdo para JSON e escreva os dados JSON no corpo da resposta
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func Response200(w http.ResponseWriter, result output) {
	w.WriteHeader(http.StatusOK)
	sendResult(w, result)
}

func Response201(w http.ResponseWriter, result output) {
	w.WriteHeader(http.StatusCreated)
	sendResult(w, result)
}

func Response204(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func Response400(w http.ResponseWriter, err error) {
	// Status code 400 - Página não encontrada.
	if err == nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	http.Error(w, err.Error(), http.StatusBadRequest)
}

func Response500(w http.ResponseWriter, err error) {
	// Status code 500 - Erro no servidor.
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
