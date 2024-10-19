package data

import (
	"testing"
)

func TestStruct(t *testing.T) {
	p := &Product{}

	err := p.Validator()

	if err != nil {
		t.Fatal(err)
	}
}
