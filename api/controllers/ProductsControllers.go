package controllers

import (
	"net/http"

	"github.com/dimitryduarte/honestyapi/api/responses"
	"github.com/dimitryduarte/honestyapi/models"
)

func (server *Server) GetProducts(w http.ResponseWriter, r *http.Request) {

	product := models.Product{}

	products, err := product.GetProduct(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, products)
}
