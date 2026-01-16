// Para adicionar os arquivos .html no executável Go.
package web

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates/*.html
var templateFiles embed.FS

// Carregamos o template a partir do sistema de arquivos embutido
var temp = template.Must(template.ParseFS(templateFiles, "templates/*.html"))

func InformationDatabase(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "Index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
// Para caminho relativo do .html
package web

import (
	"fmt"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("internal/web/templates/*.html"))

func InformationDatabase(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chegou!!!")
	temp.ExecuteTemplate(w, "Index", nil)
}

*/