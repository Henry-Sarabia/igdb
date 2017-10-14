package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

const getReleaseDateResp = `
[{
	"id": 1073,
	"game": 475,
	"created_at": 1303935024000,
	"updated_at": 1339423937521,
	"category": 0,
	"platform": 39,
	"date": 1221523200000,
	"y": 2008,
	"m": 9,
	"human": "2008-Sep-16"
}]
`

const getReleaseDatesResp = `
[{
	"id": 62566,
	"game": 26408,
	"created_at": 1481527043783,
	"updated_at": 1481527043783,
	"category": 2,
	"platform": 27,
	"date": 536371200000,
	"region": 1,
	"y": 1986,
	"m": 12,
	"human": "1986"
},
{
	"id": 32350,
	"game": 11676,
	"created_at": 1439613484873,
	"updated_at": 1439613484873,
	"category": 2,
	"platform": 13,
	"date": 978220800000,
	"y": 2000,
	"m": 12,
	"human": "2000"
},
{
	"id": 1077,
	"game": 137,
	"created_at": 1303935281000,
	"updated_at": 1339423937557,
	"category": 0,
	"platform": 9,
	"date": 1288051200000,
	"y": 2010,
	"m": 10,
	"human": "2010-Oct-26"
}]
`

func TestReleaseDateTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	rd := ReleaseDate{}
	typ := reflect.ValueOf(rd).Type()

	err := c.validateStruct(typ, ReleaseDateEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetReleaseDate(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getReleaseDateResp)
	defer ts.Close()

	rd, err := c.GetReleaseDate(1073)
	if err != nil {
		t.Error(err)
	}

	eID := 1073
	aID := rd.ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	eg := 475
	ag := rd.Game
	if ag != eg {
		t.Errorf("Expected Game ID %d, got %d", eg, ag)
	}

	ec := 1303935024000
	ac := rd.CreatedAt
	if ac != ec {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
	}
}

func TestGetReleaseDates(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getReleaseDatesResp)
	defer ts.Close()

	ids := []int{62566, 32350, 1077}
	rd, err := c.GetReleaseDates(ids)
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(rd)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	ec := DateCategory(2)
	ac := rd[0].Category
	if ac != ec {
		t.Errorf("Expected date category %d, got %d", ec, ac)
	}

	ep := 27
	ap := rd[0].Platform
	if ap != ep {
		t.Errorf("Expected platform %d, got %d", ep, ap)
	}

	ed := 978220800000
	ad := rd[1].Date
	if ad != ed {
		t.Errorf("Expected Unix time in milliseconds %d, got %d", ed, ad)
	}

	ey := 2000
	ay := rd[1].Year
	if ay != ey {
		t.Errorf("Expected year %d, got %d", ey, ay)
	}

	em := 10
	am := rd[2].Month
	if am != em {
		t.Errorf("Expected month %d, got %d", em, am)
	}

	eh := "2010-Oct-26"
	ah := rd[2].Human
	if ah != eh {
		t.Errorf("Expected date '%s', got '%s'", eh, ah)
	}

}
