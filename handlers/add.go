package handlers

import (
	"net/http"

	"github.com/abrshDev/data"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle add prodcuts")
	prod := r.Context().Value(keyProduct{}).(*data.Product)

	data.AddProduct(prod)
}
