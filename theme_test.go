package igdb

import (
	"reflect"
	"testing"
)

func TestThemeTypeIntegrity(t *testing.T) {
	c := NewClient()

	r := Theme{}
	typ := reflect.ValueOf(r).Type()

	err := c.validateStruct(typ, ThemeEndpoint)
	if err != nil {
		t.Error(err)
	}
}
