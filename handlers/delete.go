package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abrshDev/data"
	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// delete product.
//
// Responses:
//
//	201: noContent

// deleteproduct from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle get products")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	p.l.Println("got id:", id)

	err = data.DeleteProducts(id)
	if err == data.Errorproductnotfound {
		fmt.Println("error in delete prodcut:", err)
		http.Error(rw, "product not found", http.StatusNotFound)
	}

}
