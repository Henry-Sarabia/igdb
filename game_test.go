package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGameTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	g := Game{}
	typ := reflect.ValueOf(g).Type()

	err := c.validateStruct(typ, GameEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetGame(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_game.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	g, err := c.GetGame(7346)
	if err != nil {
		t.Error(err)
	}

	eID := 7346
	aID := g.ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	en := "The Legend of Zelda: Breath of the Wild"
	an := g.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'\n", en, an)
	}

	er := 98.5
	ar := g.AggregateRating
	if ar != er {
		t.Errorf("Expected rating of %f, got %f", er, ar)
	}

	ew := 2709
	aw := g.Covers.Width
	if aw != ew {
		t.Errorf("Expected width of %d, got %d\n", ew, aw)
	}

	var ev = []YoutubeVideo{
		{"Trailer", "Z6BeAtdoELY"},
		{"Trailer", "1rPxiXXxftE"},
		{"Trailer", "vDFZIUdo764"},
		{"Trailer", "zw47_q9wbBE"}}
	av := g.Videos
	for i := range av {
		if av[i] != ev[i] {
			t.Errorf("Expected video %v, got video %v\n", ev[i], av[i])
		}
	}
}

func TestGetGames(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/get_games.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	ids := []int{1721, 2777}
	g, err := c.GetGames(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(g)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	eURL := URL("https://www.igdb.com/games/mega-man-8")
	aURL := g[0].URL
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	ec := 1352059102919
	ac := g[0].CreatedAt
	if ac != ec {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
	}

	ed := 762
	ad := g[1].Developers[0]
	if ad != ed {
		t.Errorf("Expected developer ID %d, got %d", ed, ad)
	}

	eID := "etjab1sgankzyq6p6qgf"
	aID := g[1].Covers.ID
	if aID != eID {
		t.Errorf("Expected cloudinary ID '%s', got '%s'", eID, aID)
	}
}

func TestSearchGames(t *testing.T) {
	ts, c, err := testServerFile(http.StatusOK, "test_data/search_games.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()

	g, err := c.SearchGames("mario")
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(g)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Super Mario 64"
	an := g[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	es := "super-mario-64"
	as := g[0].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	er := 85.5273965373542
	ar := g[1].Rating
	if ar != er {
		t.Errorf("Expected rating of %f, got %f", er, ar)
	}

	ep := 3.666666666666667
	ap := g[1].Popularity
	if ap != ep {
		t.Errorf("Expected popularity of %f, got %f", ep, ap)
	}

	ed := 864
	ad := g[2].Developers[0]
	if ad != ed {
		t.Errorf("Expected developer ID %d, got %d", ed, ad)
	}

	eURL := URL("//images.igdb.com/igdb/image/upload/t_thumb/clmh270eov5rimiggwrk.jpg")
	aURL := g[2].Covers.URL
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}
}
