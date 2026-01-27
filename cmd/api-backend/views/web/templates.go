package web

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates/*.html
var templateFiles embed.FS

var temp = template.Must(template.ParseFS(templateFiles, "templates/*.html"))

func RenderTemplate(w http.ResponseWriter, page string, data any) {

	err := temp.ExecuteTemplate(w, page, data)
	if err != nil {
		//log.Printf("Erro ao renderizar template %s: %v", page, err)

		tmplErr := temp.ExecuteTemplate(w, "Error500", nil)
		if tmplErr != nil {
			//log.Printf("Erro ao renderizar template de erro: %v", tmplErr)
			http.Error(w, "Erro interno no servidor", http.StatusInternalServerError)
			/*
				Pode-se usar a função de view
				HandleInternalServerError(w http.ResponseWriter, errorMessage string)
			*/
		}
	}
}