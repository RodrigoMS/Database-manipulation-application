package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/RodrigoMS/app/cmd/models"
	"github.com/RodrigoMS/app/cmd/views"
	"github.com/RodrigoMS/app/pkg/utils"
)


func GetUser(w http.ResponseWriter, r *http.Request) {
	// Lógica do controlador aqui
	//user := models.GetUser()

	//fmt.Println(user)
	// Renderizar a view com os dados do usuário

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		views.HandleNotFound(w, nil)
		return
	}

	user, err := models.ReadUser(parts[2])

	if err != nil {
		fmt.Println("Erro em userModel.go: \n", err)
		// Configurar  e executar uma função que grava os logs de erro
		// em um arquivo destro da pasta logs na raiz da aplicação
		//log.Printf("Erro em userModel.go: %v", err)

		//view.Error(w, r)
		views.HandleNotFound(w, nil)
		return
	}

	if user == nil {
		views.HandleNotFound(w, nil)
		return
	}

	views.HandleSuccess(w, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.ReadAllUsers()
	if err != nil {
		fmt.Println("Erro em userModel.go")
		return
	}

	views.HandleSuccess(w, users)
}


func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	user, err := utils.ReadJSON[models.User](r.Body)
	if err != nil {
			views.HandleNotFound(w, nil)
			return
	}

	// Lógica de validação dos dados
	// ...

	user, err = models.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		views.HandleInternalServerError(w, "Não foi possível concluir o cadastro. Tente novamente mais tarde.")
		return
	}

	views.HandleResourceCreated(w, user)
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	user, err := utils.ReadJSON[models.User](r.Body)
	if err != nil {
			views.HandleNotFound(w, nil)
			return
	}

	// Lógica de validação dos dados
	// ...

	user, err = models.UpdateUser(user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		views.HandleInternalServerError(w, "Não foi possível atualizar o cadastro. Tente novamente mais tarde.")
		return
	}

	views.HandleSuccess(w, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	user, err := utils.ReadJSON[models.User](r.Body)
	if err != nil {
			views.HandleNotFound(w, nil)
			return
	}

	// Lógica de validação dos dados
	// ...

	err = models.DeleteUser(user.ID)

	if err != nil {
		views.HandleInternalServerError(w, "Erro ao excluir o usuário. Verifique se ele existe ou tente novamente em instantes.")
		return
	}

	views.HandleNoContent(w)
}