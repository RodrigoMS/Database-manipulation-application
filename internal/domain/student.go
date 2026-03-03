package domain

import "fmt"

type Student struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Class    string `json:"class"`
}

// Simulação de validação (substituir por consulta ao banco de dados)
func ValidateStudent(name, class string) bool {
    // Exemplo: apenas alunos da turma "A" são válidos
    if name == "RMS" || name == "GMS" && class == "5A" {
        fmt.Println("Model(domain)", name, "- ", class)
        return true
    }
    return false
}