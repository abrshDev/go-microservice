package data

import (
	"testing"
)

func TestCheckValidator(t *testing.T) {
	p := &Product{
		Name:  "nic",
		Price: 1.1,
		SKU:   "abx-abc-dfg",
	}
	err := p.Validator()
	if err != nil {
		t.Fatal(err)
	}
}
