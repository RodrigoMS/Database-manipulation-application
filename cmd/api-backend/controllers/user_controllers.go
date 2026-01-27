package controllers

import (
	"net/http"

	model "github.com/RodrigoMS/App_Go/cmd/api-backend/models/user"
	"github.com/RodrigoMS/App_Go/cmd/api-backend/views"
)

type InvalidData struct {
	Name string
	Func func()
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	user := model.User{
		Name:     "Rodrigo",
		Email:    "rmsproject@localhost.com",
		Active:   true,
		Password: "*********",
	}

	/*user := InvalidData{
		Name: "Teste",
		Func: func() {},
	}*/

	views.HandleSuccess(w, user)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	views.HandleResourceCreated(w, "POST - User")
}

func PutUser(w http.ResponseWriter, r *http.Request) {

	user := model.User{
		Name:     "Rodrigo",
		Email:    "rmsproject@localhost.com",
		Active:   true,
		Password: "*********",
	}

	views.HandleSuccess(w, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	views.HandleNoContent(w)
}
