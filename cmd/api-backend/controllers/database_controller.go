package controllers

import (
	"net/http"

	models "github.com/RodrigoMS/App_Go/cmd/api-backend/models/server"
	"github.com/RodrigoMS/App_Go/cmd/api-backend/views/web"
)

func GetDatabaseInfo(w http.ResponseWriter, r *http.Request) {

	info, err := models.GetDBInfo()
	if err != nil {
		web.RenderTemplate(w, "Error500", nil)
	}

	web.RenderTemplate(w, "Index", info)
}