package igdb

import (
	"reflect"
	"testing"
)

func TestTitleTypeIntegrity(t *testing.T) {
	c := NewClient()

	r := Title{}
	typ := reflect.ValueOf(r).Type()

	err := c.validateStruct(typ, TitleEndpoint)
	if err != nil {
		t.Error(err)
	}
}
