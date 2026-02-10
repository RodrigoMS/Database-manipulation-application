package handlers

import (
	"net/http"

	"github.com/RodrigoMS/app/cmd/internal/web"
)

// Carrega o template html do login dos alunos.
func StudentLogin(w http.ResponseWriter, r *http.Request) {
	web.RenderTemplate(w, "StudentLogin", nil)
}

// Carrega a painel do aluno.
func StudentDashboard(w http.ResponseWriter, r *http.Request) {
	web.RenderTemplate(w, "StudentDashboard", nil)
}