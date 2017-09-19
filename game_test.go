package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

const getGameResp = `
[{
	"id": 7346,
	"name": "The Legend of Zelda: Breath of the Wild",
	"slug": "the-legend-of-zelda-breath-of-the-wild",
	"url": "https://www.igdb.com/games/the-legend-of-zelda-breath-of-the-wild",
	"created_at": 1402423357911,
	"updated_at": 1504638741974,
	"summary": "Step into a world of discovery, exploration and adventure in The Legend of Zelda: Breath of the Wild, a boundary-breaking new game in the acclaimed series. Travel across fields, through forests and to mountain peaks as you discover what has become of the ruined kingdom of Hyrule in this stunning open-air adventure.",
	"storyline": "Link awakes in a mysterious chamber after 100 years of slumber to find that Calamity Ganon has taken over Hyrule Castle and left Hyrule to decay and be taken over by nature.",
	"collection": 106,
	"franchise": 596,
	"franchises": [
		596
	],
	"hypes": 142,
	"rating": 97.05963818994093,
	"popularity": 55,
	"aggregated_rating": 98.5,
	"aggregated_rating_count": 29,
	"total_rating": 97.77981909497047,
	"total_rating_count": 94,
	"rating_count": 65,
	"developers": [
		7902
	],
	"first_release_date": 1488499200000,
	"pulse_count": 586,
	"videos": [
		{
			"name": "Trailer",
			"video_id": "Z6BeAtdoELY"
		},
		{
			"name": "Trailer",
			"video_id": "1rPxiXXxftE"
		},
		{
			"name": "Trailer",
			"video_id": "vDFZIUdo764"
		},
		{
			"name": "Trailer",
			"video_id": "zw47_q9wbBE"
		}
	],
	"cover": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/jk9el4ksl4c7qwaex2y5.jpg",
		"cloudinary_id": "jk9el4ksl4c7qwaex2y5",
		"width": 2709,
		"height": 3816
	}
}]
`

const getGamesResp = `
[{
	"id": 1721,
	"name": "Mega Man 8",
	"slug": "mega-man-8",
	"url": "https://www.igdb.com/games/mega-man-8",
	"created_at": 1352059102919,
	"updated_at": 1502526599128,
	"popularity": 1.666666666666667,
	"developers": [
		37
	],
	"cover": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/k8tdkafvthvfrxawnod4.jpg",
		"cloudinary_id": "k8tdkafvthvfrxawnod4",
		"width": 950,
		"height": 788
	}
},
{
	"id": 2777,
	"name": "Kirby Air Ride",
	"slug": "kirby-air-ride",
	"url": "https://www.igdb.com/games/kirby-air-ride",
	"created_at": 1374907021184,
	"updated_at": 1500383884617,
	"rating": 84.4665388294525,
	"popularity": 1.666666666666667,
	"developers": [
		762
	],
	"cover": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/etjab1sgankzyq6p6qgf.jpg",
		"cloudinary_id": "etjab1sgankzyq6p6qgf",
		"width": 256,
		"height": 357
	}
}]
`

const searchGamesResp = `
[{
	"id": 1074,
	"name": "Super Mario 64",
	"slug": "super-mario-64",
	"url": "https://www.igdb.com/games/super-mario-64",
	"created_at": 1339251703388,
	"updated_at": 1503657933500,
	"rating": 89.92645562042,
	"popularity": 10.66666666666667,
	"developers": [
		421
	],
	"cover": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/scutr4p9gytl4txb2soy.jpg",
		"cloudinary_id": "scutr4p9gytl4txb2soy",
		"width": 795,
		"height": 554
	}
},
{
	"id": 2350,
	"name": "Mario Kart 8",
	"slug": "mario-kart-8",
	"url": "https://www.igdb.com/games/mario-kart-8",
	"created_at": 1372712056850,
	"updated_at": 1503173840741,
	"rating": 85.5273965373542,
	"popularity": 3.666666666666667,
	"developers": [
		1103
	],
	"videos": [
		{
			"name": "Trailer",
			"video_id": "NA6CAgv6p6g"
		}
	],
	"cover": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/ivcvfoeo23xitx8xfz8m.jpg",
		"cloudinary_id": "ivcvfoeo23xitx8xfz8m",
		"width": 904,
		"height": 1273
	}
},
{
	"id": 2327,
	"name": "Mario Party",
	"slug": "mario-party",
	"url": "https://www.igdb.com/games/mario-party",
	"created_at": 1372629718263,
	"updated_at": 1500314266411,
	"rating": 80.6722124072401,
	"popularity": 2.666666666666667,
	"developers": [
		864
	],
	"cover": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/clmh270eov5rimiggwrk.jpg",
		"cloudinary_id": "clmh270eov5rimiggwrk",
		"width": 640,
		"height": 441
	}
}]
`

// Use meta endpoint to fetch list of fields. Ignore fields with periods.
// Use reflect package to gather existing struct field names.
// Iterate through meta field names, checking to see if they exist in reflect struct fields.
func TestGameTypeIntegrity(t *testing.T) {
	c := NewClient()

	g := Game{}
	typ := reflect.ValueOf(g).Type()

	err := c.validateStruct(typ, GameEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetGame(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getGameResp)
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

	var ev = []Video{
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
	ts, c := startTestServer(http.StatusOK, getGamesResp)
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
		t.Errorf("Expected unix epoch of %d, got %d", ec, ac)
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
	ts, c := startTestServer(http.StatusOK, searchGamesResp)
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
