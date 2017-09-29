package igdb

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

const testErrNotFound = `
{
	"status": 404,
	"message": "status not found"
}
`

func TestCheckError(t *testing.T) {
	var errTests = []struct {
		Name string
		Code int
		Body string
		Exp  string
	}{
		{"empty response", 404, "", "unexpected end of JSON input"},
		{"404 response", 404, testErrNotFound, "Status 404 - status not found"},
	}

	for _, et := range errTests {
		t.Run(et.Name, func(t *testing.T) {
			c := NewClient()

			resp := &http.Response{StatusCode: et.Code,
				Body: ioutil.NopCloser(bytes.NewBufferString(et.Body)),
			}

			err := c.checkError(resp)
			if err.Error() != et.Exp {
				t.Errorf("Expected '%v', got '%v'", et.Exp, err.Error())
			}
		})
	}
}
