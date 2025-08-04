package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/abrshDev/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

type keyProduct struct{}

//swagger:response noContent
type productNoContent struct {
}

// swagger :parametrs deleteProduct
type productIdParammetrWrapper struct {
	//the id of the product to delete from the database
	// in:path
	// required : true
	ID int `json:"id"`
}

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
