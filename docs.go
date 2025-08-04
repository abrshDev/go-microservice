// Package classification of Product API
//
// The documentation for the Product API.
//
// Schemes: http
// BasePath: /
// Version: 2.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

//swagger:response noContent
type productNoContent struct {
}

// swagger:parameters parametrs  deleteProduct
type productIdParammetrWrapper struct {
	//the id of the product to delete from the database
	//
	// in:path
	// required : true
	ID int `json:"id"`
}
