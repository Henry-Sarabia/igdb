package igdb

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

const checkErrorResp = `
{
	"status": 404,
	"message": "status not found"
}
`

func TestCheckError(t *testing.T) {
	ts, c := startTestServer(http.StatusNotFound, checkErrorResp)
	defer ts.Close()

	_, err := c.GetGame(1022)

	if err == nil {
		t.Error("Expected error, got nil")
	}

	expErr := "Status 404 - status not found"
	actErr := err.Error()
	if actErr != expErr {
		t.Errorf("Expected error '%s', got '%s'", expErr, actErr)
	}
}

func TestCheckErrorResp(t *testing.T) {
	var errTests = []struct {
		Name   string
		Status string
		Code   int
		Body   string
		Exp    string
	}{
		{"empty response", "404 Not Found", 404, "", "unexpected end of JSON input"},
	}

	for _, et := range errTests {
		t.Run(et.Name, func(t *testing.T) {
			c := NewClient()

			resp := &http.Response{Status: et.Status,
				StatusCode: et.Code,
				Body:       ioutil.NopCloser(bytes.NewBufferString(et.Body)),
			}
			err := c.checkError(resp)
			if err.Error() != et.Exp {
				t.Errorf("Expected '%v', got '%v'", et.Exp, err.Error())
			}
		})
	}
}
