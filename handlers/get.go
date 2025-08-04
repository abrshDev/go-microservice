package handlers

import (
	"net/http"

	"github.com/abrshDev/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products.
//
// Responses:
//
//	200: productsResponse
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle get products")
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}

}
