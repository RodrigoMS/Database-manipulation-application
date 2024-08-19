package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/RodrigoMS/app/cmd/models"
	"github.com/RodrigoMS/app/cmd/view"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Lógica do controlador aqui
	//user := models.GetUser()

	//fmt.Println(user)
	// Renderizar a view com os dados do usuário

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		view.Response400(w, nil)
		return
	}

	user, err := models.ReadUser(parts[2])

	if err != nil {
		fmt.Println("Erro em userModel.go: \n", err)
		// Configurar  e executar uma função que grava os logs de erro
		// em um arquivo destro da pasta logs na raiz da aplicação
		//log.Printf("Erro em userModel.go: %v", err)

		//view.Error(w, r)
		view.Response400(w, err)
		return
	}

	if user == nil {
		view.Response400(w, nil)
		return
	}

	view.Response200(w, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.ReadAllUsers()
	if err != nil {
		fmt.Println("Erro em userModel.go")
		return
	}

	view.Response200(w, users)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		view.Response400(w, err)
		return
	}

	// Lógica de validação dos dados
	// ...

	user, err = models.CreateUser(user.Email, user.Password)
	if err != nil {
		view.Response500(w, err)
		return
	}

	view.Response200(w, user)
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		view.Response400(w, err)
		return
	}

	// Lógica de validação dos dados
	// ...

	user, err = models.UpdateUser(user.Email, user.Password, user.ID)
	if err != nil {
		view.Response500(w, err)
		return
	}

	view.Response200(w, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		view.Response400(w, err)
		return
	}

	// Lógica de validação dos dados
	// ...

	err = models.DeleteUser(user.ID)
	fmt.Println(err)
	if err != nil {
		view.Response500(w, err)
		return
	}

	view.Response204(w)
}

// gravar logs (ver como se faz)
/*
logFile, err := os.OpenFile("meuLog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    log.Fatal(err)
}
log.SetOutput(logFile)
*/
