package controllers

import (
	"net/http"

	model "github.com/RodrigoMS/App_Go/cmd/api-backend/models/product"
	"github.com/RodrigoMS/App_Go/cmd/api-backend/views"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	product := model.Product{
		Name:     "Teclado",
		Barcode:  "7898506450077",
		Price:    21.99,
		Quantity: 7,
	}

	views.HandleSuccess(w, product)
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	views.HandleResourceCreated(w, "POST - Product")
}

func PutProduct(w http.ResponseWriter, r *http.Request) {

	views.HandleSuccess(w, "PUT - Product")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	views.HandleNoContent(w)
}
