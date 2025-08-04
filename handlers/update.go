package handlers

import (
	"net/http"
	"strconv"

	"github.com/abrshDev/data"
	"github.com/gorilla/mux"
)

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle update products")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	p.l.Println("got id:", id)
	prod := r.Context().Value(keyProduct{}).(*data.Product)
	err = data.UpdateProduct(id, prod)
	if err == data.Errorproductnotfound {
		http.Error(rw, "product not found ", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "product not found ", http.StatusInternalServerError)
		return
	}

}
