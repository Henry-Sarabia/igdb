package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

const getPulseSourceResp = `
[{
	"id": 1,
	"name": "Kotaku",
	"page": 501
}]
`

const getPulseSourcesResp = `
[{
	"id": 6,
	"name": "Escapist",
	"page": 553
},
{
	"id": 10,
	"name": "Destructoid",
	"page": 552
}]
`

const searchPulseSourcesResp = `
[{
	"id": 3,
	"name": "GameInformer",
	"page": 261
},
{
	"id": 16,
	"name": "GameIndustry.biz",
	"page": 563
},
{
	"id": 17,
	"name": "GameSpot",
	"page": 564
}]
`

func TestPulseSourceTypeIntegrity(t *testing.T) {
	c := NewClient()

	ps := PulseSource{}
	typ := reflect.ValueOf(ps).Type()

	err := c.validateStruct(typ, PulseSourceEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPulseSource(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getPulseSourceResp)
	defer ts.Close()

	ps, err := c.GetPulseSource(4943)
	if err != nil {
		t.Error(err)
	}

	en := "Kotaku"
	an := ps.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eID := 1
	aID := ps.ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	ep := 501
	ap := ps.Page
	if ap != ep {
		t.Errorf("Expected Page ID %d, got %d", ep, ap)
	}
}

func TestGetPulseSources(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getPulseSourcesResp)
	defer ts.Close()

	ids := []int{6, 10}
	ps, err := c.GetPulseSources(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(ps)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Escapist"
	an := ps[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eID := 6
	aID := ps[0].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	ep := 552
	ap := ps[1].Page
	if ap != ep {
		t.Errorf("Expected Page ID %d, got %d", ep, ap)
	}
}

func TestSearchPulseSources(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, searchPulseSourcesResp)
	defer ts.Close()

	ps, err := c.SearchPulseSources("game")
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(ps)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	eID := 3
	aID := ps[0].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	en := "GameIndustry.biz"
	an := ps[1].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	ep := 564
	ap := ps[2].Page
	if ap != ep {
		t.Errorf("Expected Page ID %d, got %d", ep, ap)
	}
}
