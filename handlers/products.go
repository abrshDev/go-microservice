package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/abrshDev/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle get products")
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}

}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle add prodcuts")
	prod := r.Context().Value(keyProduct{}).(*data.Product)

	data.AddProduct(prod)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle update products")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	p.l.Println("got id:", id)
	prod := r.Context().Value(keyProduct{}).(data.Product)
	err = data.UpdateProduct(id, &prod)
	if err == data.Errorproductnotfound {
		http.Error(rw, "product not found ", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "product not found ", http.StatusInternalServerError)
		return
	}

}

type keyProduct struct{}

func (p *Products) MiddleWareValidate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJson(r.Body)
		fmt.Println("prod:", prod)
		if err != nil {
			fmt.Println("error to marshal json in middleware:", err)
			http.Error(rw, "unable to marshal json", http.StatusBadRequest)
			return
		}
		//validate the context
		err = prod.Validator()
		if err != nil {
			fmt.Println("error validating product")
			http.Error(rw, fmt.Sprintf("error validating products: %s", err), http.StatusBadRequest)
			return
		}
		//add the product to the context
		ctx := context.WithValue(r.Context(), keyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
