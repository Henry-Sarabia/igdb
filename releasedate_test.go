package igdb

import (
	"reflect"
	"testing"
)

const getReleaseDateResp = `

`

const getReleaseDatesResp = `

`

const searchReleaseDatesResp = `

`

func TestReleaseDateTypeIntegrity(t *testing.T) {
	c := NewClient()

	rd := ReleaseDate{}
	typ := reflect.ValueOf(rd).Type()

	err := c.validateStruct(typ, ReleaseDateEndpoint)
	if err != nil {
		t.Error(err)
	}
}
