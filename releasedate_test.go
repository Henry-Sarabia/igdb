package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

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
	var releaseDateTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/get_releasedate.txt", 1073, ""},
		{"Invalid ID", "test_data/empty.txt", -1000, ErrNegativeID.Error()},
		{"Empty Response", "test_data/empty.txt", 1073, errEndOfJSON.Error()},
	}
	for _, tt := range releaseDateTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			rd, err := c.GetReleaseDate(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
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
		})
	}
}

func TestGetReleaseDates(t *testing.T) {
	var releaseDateTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/get_releasedates.txt", []int{62566, 32350, 1077}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-123}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/empty.txt", nil, nil, ErrEmptyIDs.Error()},
		{"Empty Response", "test_data/empty.txt", []int{62566, 32350, 1077}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{62566, 32350, 1077}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range releaseDateTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			rd, err := c.GetReleaseDates(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
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
		})
	}
}
