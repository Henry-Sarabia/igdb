package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestPlatformTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	p := Platform{}
	typ := reflect.ValueOf(p).Type()

	err := c.validateStruct(typ, PlatformEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestPlatformsGet(t *testing.T) {
	var platformTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/platforms_get.txt", 7, ""},
		{"Invalid ID", "test_data/empty.txt", -10, ErrNegativeID.Error()},
		{"Empty response", "test_data/empty.txt", 7, errEndOfJSON.Error()},
		{"No results", "test_data/empty_array.txt", 0, ErrNoResults.Error()},
	}
	for _, tt := range platformTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			p, err := c.Platforms.Get(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			en := "PlayStation"
			an := p.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			egID := []int{1185, 1186, 1187, 1192, 1195, 1201, 425, 675}
			agID := p.Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}

			eCPU := "MIPS R3051 @ 33,8688 MHz"
			aCPU := p.Versions[0].CPU
			if aCPU != eCPU {
				t.Errorf("Expected CPU '%s', got '%s'", eCPU, aCPU)
			}

			el := 4
			al := len(p.Versions[0].ReleaseDates)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			ed := 810604800000
			ad := p.Versions[0].ReleaseDates[1].Date
			if ad != ed {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", ed, ad)
			}
		})
	}
}

func TestPlatformsList(t *testing.T) {
	var platformTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/platforms_list.txt", []int{23, 130}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-123}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/platforms_list.txt", nil, nil, ""},
		{"Empty response", "test_data/empty.txt", []int{23, 130}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{23, 130}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", []int{0, 9999999}, nil, ErrNoResults.Error()},
	}
	for _, tt := range platformTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			p, err := c.Platforms.List(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(p)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			en := "Dreamcast"
			an := p[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			ew := 1024
			aw := p[0].Logo.Width
			if aw != ew {
				t.Errorf("Expected width of %d, got %d", ew, aw)
			}

			evn := "Initial version"
			avn := p[0].Versions[0].Name
			if avn != evn {
				t.Errorf("Expected name '%s', got '%s'", evn, avn)
			}

			eWeb := "http://www.nintendo.com/switch"
			aWeb := p[1].Website
			if aWeb != eWeb {
				t.Errorf("Expected website '%s', got '%s'", eWeb, aWeb)
			}

			eh := 413
			ah := p[1].Versions[0].Logo.Height
			if ah != eh {
				t.Errorf("Expected height of %d, got %d", eh, ah)
			}

			egID := []int{38983, 10232, 885, 11529, 19081, 19175, 29525, 36846, 27081, 26165}
			agID := p[1].Games
			for i := range agID {
				if agID[i] != egID[i] {
					t.Errorf("Expected Game ID %d, got %d\n", egID[i], agID[i])
				}
			}
		})
	}
}

func TestPlatformsSearch(t *testing.T) {
	var platformTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/platforms_search.txt", "xbox", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/empty.txt", "", []OptionFunc{OptLimit(50)}, ErrEmptyQuery.Error()},
		{"Empty response", "test_data/empty.txt", "xbox", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "xbox", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
		{"No results", "test_data/empty_array.txt", "non-existant entry", nil, ErrNoResults.Error()},
	}
	for _, tt := range platformTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			p, err := c.Platforms.Search(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(p)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eID := 11
			aID := p[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			ec := 1297639288000
			ac := p[0].CreatedAt
			if ac != ec {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
			}

			eg := 6
			ag := p[0].Generation
			if ag != eg {
				t.Errorf("Expected generation %d, got %d", eg, ag)
			}

			evl := 3
			avl := len(p[1].Versions)
			if avl != evl {
				t.Errorf("Expected length of %d, got %d", evl, avl)
			}

			er := "480i, 480p, 720p, 1080i, 1080p"
			ar := p[1].Versions[1].Resolutions
			if ar != er {
				t.Errorf("Expected resolutions '%s', got '%s'", er, ar)

			}

			es := "xbox-360-elite"
			as := p[1].Versions[2].Slug
			if as != es {
				t.Errorf("Expected slug '%s', got '%s'", es, as)
			}
		})
	}
}

func TestPlatformsCount(t *testing.T) {
	var countTests = []struct {
		Name     string
		Resp     string
		Opts     []OptionFunc
		ExpCount int
		ExpErr   string
	}{
		{"Happy path", `{"count": 100}`, []OptionFunc{OptFilter("popularity", OpGreaterThan, "75")}, 100, ""},
		{"Empty response", "", nil, 0, errEndOfJSON.Error()},
		{"Invalid option", "", []OptionFunc{OptLimit(100)}, 0, ErrOutOfRange.Error()},
		{"No results", "[]", nil, 0, ErrNoResults.Error()},
	}

	for _, tt := range countTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, tt.Resp)
			defer ts.Close()

			count, err := c.Platforms.Count(tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if count != tt.ExpCount {
				t.Fatalf("Expected count %d, got %d", tt.ExpCount, count)
			}
		})
	}
}

func TestPlatformsListFields(t *testing.T) {
	var fieldTests = []struct {
		Name      string
		Resp      string
		ExpFields []string
		ExpErr    string
	}{
		{"Happy path", `["name", "slug", "url"]`, []string{"url", "slug", "name"}, ""},
		{"Dot operator", `["logo.url", "background.id"]`, []string{"background.id", "logo.url"}, ""},
		{"Asterisk", `["*"]`, []string{"*"}, ""},
		{"Empty response", "", nil, errEndOfJSON.Error()},
		{"No results", "[]", nil, ""},
	}

	for _, tt := range fieldTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c := testServerString(http.StatusOK, tt.Resp)
			defer ts.Close()

			fields, err := c.Platforms.ListFields()
			assertError(t, err, tt.ExpErr)

			ok, err := equalSlice(fields, tt.ExpFields)
			if err != nil {
				t.Fatal(err)
			}
			if !ok {
				t.Fatalf("Expected fields '%v', got '%v'", tt.ExpFields, fields)
			}
		})
	}
}
