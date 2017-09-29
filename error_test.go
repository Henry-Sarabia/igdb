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
		{"empty response", http.StatusBadRequest, "", "unexpected end of JSON input"},
		{"404 response", http.StatusNotFound, testErrNotFound, "Status 404 - status not found"},
		{"200 response", http.StatusOK, "", ""},
	}

	for _, et := range errTests {
		t.Run(et.Name, func(t *testing.T) {
			c := NewClient()

			resp := &http.Response{StatusCode: et.Code,
				Body: ioutil.NopCloser(bytes.NewBufferString(et.Body)),
			}

			err := c.checkError(resp)
			if resp.StatusCode == http.StatusOK {
				if err != nil {
					t.Fatalf("Expected nil err, got '%v'", err)
				}
				return
			}
			if err.Error() != et.Exp {
				t.Fatalf("Expected '%v', got '%v'", et.Exp, err.Error())
			}
		})
	}
}
