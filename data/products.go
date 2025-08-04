package data

import (
	/* "encoding/json"
	"fmt"
	"io" */
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// product defines the structure of product
//
//swagger:model
type Product struct {
	//the id for this user
	//required :true
	//min:1
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}
type products []*Product

func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
func (p *products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func (p *Product) Validator() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSku)
	return validate.Struct(p)
}
func validateSku(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]`)
	matches := reg.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true

}
func GetProducts() products {
	return productList
}
func AddProduct(p *Product) {
	p.ID = GetNextId()
	productList = append(productList, p)
}
func GetNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findproduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

var Errorproductnotfound = fmt.Errorf("product not found")

func findproduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, Errorproductnotfound
}
func DeleteProducts(id int) error {
	for i, pro := range productList {
		if pro.ID == id {
			productList = append(productList[:i], productList[i+1:]...)
			return nil
		}
	}
	return Errorproductnotfound
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
