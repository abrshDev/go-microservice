package main

import (
	"testing"

	"github.com/abrshDev/sdk/client"
	"github.com/abrshDev/sdk/client/products"
)

func TestMain(t *testing.T) {
	c := client.NewHTTPClientWithConfig(nil, client.DefaultTransportConfig().WithHost("localhost:1991"))
	params := products.NewDeleteProductParams()

	c.Products.DeleteProduct(params, nil)

	t.Log("Main function executed successfully")
}
