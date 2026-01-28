package handlers

import (
	"fmt"
	"net/http"

	"github.com/RodrigoMS/app/cmd/internal/domain"
	"github.com/RodrigoMS/app/cmd/internal/web"
)

func GetDatabaseInfo(w http.ResponseWriter, r *http.Request) {

	info, err := domain.GetDBInfo()
	if err != nil {
		fmt.Printf("opa")
		web.RenderTemplate(w, "Error500", nil)
	}

	web.RenderTemplate(w, "Index", info)
}