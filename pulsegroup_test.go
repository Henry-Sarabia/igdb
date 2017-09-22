package igdb

import (
	"reflect"
	"testing"
)

func TestPulseGroupTypeIntegrity(t *testing.T) {
	c := NewClient()

	pg := PulseGroup{}
	typ := reflect.ValueOf(pg).Type()

	err := c.validateStruct(typ, PulseGroupEndpoint)
	if err != nil {
		t.Error(err)
	}
}
