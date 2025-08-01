package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/abrshDev/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(rw, "invalid uri", http.StatusInternalServerError)
		}
		if len(g[0]) != 2 {
			p.l.Println("invalid uri more than one capcha group")
			http.Error(rw, "invalid uri", http.StatusInternalServerError)
		}
		idString := g[0][1]
		id, _ := strconv.Atoi(idString)
		p.l.Println("got id", id)
		p.updateProducts(id, rw, r)
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)

}
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle get products")
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}

}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle add prodcuts")
	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle update products")
	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusBadRequest)
	}
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
