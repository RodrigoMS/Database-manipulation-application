// Para adicionar os arquivos .html no executável Go.
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
		
		tmplErr := temp.ExecuteTemplate(w, "Error500", nil)
		if tmplErr != nil {
			http.Error(w, "Erro interno no servidor", http.StatusInternalServerError)
			/*
				Pode-se usar a função de view
				HandleInternalServerError(w http.ResponseWriter, errorMessage string)
			*/
		}
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