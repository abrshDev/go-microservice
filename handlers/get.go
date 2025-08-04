package handlers

import (
	"net/http"

	"github.com/abrshDev/data"
)

// swagger:route GET /products listproduct `getProduct
// get product.
//
// Responses:
//
//	201: noContent

// getproduct from the database
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle get products")
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}

}
