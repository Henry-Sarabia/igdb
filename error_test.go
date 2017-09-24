package igdb

import "testing"
import "net/http"

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
