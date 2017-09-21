package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

// Note that some of these test responses have truncated information.
const getPersonResp = `
[{
	"id": 1688,
	"name": "Shigeru Miyamoto",
	"created_at": 1384907616370,
	"updated_at": 1498837418443,
	"slug": "shigeru-miyamoto",
	"url": "https://www.igdb.com/people/shigeru-miyamoto",
	"dob": -540432000000,
	"gender": 0,
	"country": 392,
	"mug_shot": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/wuyjvwyascmcquyf4qh9.jpg",
		"cloudinary_id": "wuyjvwyascmcquyf4qh9",
		"width": 2048,
		"height": 1365
	},
	"bio": "Shigeru Miyamoto is a Japanese video game designer and producer. He is best known as the creator of some of the most critically acclaimed and best-selling video games of all time.\n\nMiyamoto originally joined Nintendo in 1977, when the company was beginning its foray into video games, and starting to abandon the playing cards it had made starting in 1889. His games have been seen on every Nintendo video game console, with his earliest work appearing on arcade machines in the late 70s. Franchises Miyamoto has helped create include the Mario, Donkey Kong, The Legend of Zelda, Star Fox, F-Zero, Pikmin, and Wii series. Noteworthy games within these include Super Mario Bros., one of the most well known video games; Super Mario 64, an early example of 3D control schemes; and The Legend of Zelda: Ocarina of Time, one of the most critically acclaimed video games of all time.",
	"games": [
		2909,
		2777,
		2476,
		2923,
		2350,
		7337,
		1073,
		1070,
		1036,
		1074,
		3365
	]
}]
`

const getPersonsResp = `
[{
	"id": 52302,
	"name": "Sean Murray",
	"created_at": 1419554477079,
	"updated_at": 1471550782876,
	"slug": "sean-murray",
	"url": "https://www.igdb.com/people/sean-murray",
	"games": [
		5628,
		2354,
		3225
	]
},
{
	"id": 84908,
	"name": "Keiji Inafune",
	"created_at": 1433023287339,
	"updated_at": 1476440662740,
	"slug": "keiji-inafune",
	"url": "https://www.igdb.com/people/keiji-inafune",
	"gender": 0,
	"games": [
		5845,
		496,
		1348,
		1913,
		1430,
		1035,
		11131,
		1723,
		6914
	]
}]
`

const searchPersonsResp = `
[{
	"id": 34056,
	"name": "Hideo Kojima",
	"created_at": 1413724312587,
	"updated_at": 1472327987812,
	"slug": "hideo-kojima",
	"url": "https://www.igdb.com/people/hideo-kojima",
	"dob": -200620800000,
	"gender": 0,
	"country": 392,
	"mug_shot": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/xp2uefbgrzgro0fecvt6.jpg",
		"cloudinary_id": "xp2uefbgrzgro0fecvt6",
		"width": 1280,
		"height": 720
	},
	"bio": "Hideo Kojima was born August 24th 1963 in Tokoy and is mostly known as the mastermind behind the Metal Gear Solid universe.",
	"twitter": "https://twitter.com/Kojima_Hideo",
	"voice_acted": [
		5328,
		1985
	],
	"games": [
		5328,
		376,
		375,
		380,
		379,
		113,
		1985,
		9647,
		6326,
		378,
		382
	],
	"characters": [
		2173
	]
},
{
	"id": 93634,
	"name": "Hideo Kohima",
	"created_at": 1441492909741,
	"updated_at": 1441492909810,
	"slug": "hideo-kohima",
	"url": "https://www.igdb.com/people/hideo-kohima",
	"gender": 0,
	"games": [
		379
	]
}]
`

func TestPersonTypeIntegrity(t *testing.T) {
	c := NewClient()

	p := Person{}
	typ := reflect.ValueOf(p).Type()

	err := c.validateStruct(typ, PersonEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPerson(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getPersonResp)
	defer ts.Close()

	p, err := c.GetPerson(2107)
	if err != nil {
		t.Error(err)
	}

	en := "Shigeru Miyamoto"
	an := p.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eDOB := -540432000000
	aDOB := p.DOB
	if aDOB != eDOB {
		t.Errorf("Expected unix epoch of %d, got %d", eDOB, aDOB)
	}

	eURL := URL("//images.igdb.com/igdb/image/upload/t_thumb/wuyjvwyascmcquyf4qh9.jpg")
	aURL := p.Mugshot.URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	egID := []int{2909, 2777, 2476, 2923, 2350, 7337, 1073, 1070, 1036, 1074, 3365}
	agID := p.Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID, agID)
		}
	}
}

func TestGetPersons(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getPersonsResp)
	defer ts.Close()

	ids := []int{52302, 84908}
	p, err := c.GetPersons(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(p)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Sean Murray"
	an := p[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	es := "sean-murray"
	as := p[0].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	eg := GenderCode(0)
	ag := p[1].Gender
	if ag != eg {
		t.Errorf("Expected Gender code %d, got %d", eg, ag)
	}

	egID := []int{5845, 496, 1348, 1913, 1430, 1035, 11131, 1723, 6914}
	agID := p[1].Games
	for i := range agID {
		if agID[i] != egID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", egID, agID)
		}
	}
}

func TestSearchPersons(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, searchPersonsResp)
	defer ts.Close()

	p, err := c.SearchPersons("hideokojima")
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(p)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	eID := 34056
	aID := p[0].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	ec := CountryCode(392)
	ac := p[0].Country
	if ac != ec {
		t.Errorf("Expected Country code %d, got %d", ec, ac)
	}

	ev := []int{5328, 1985}
	av := p[0].VoiceActed
	for i := range av {
		if av[i] != ev[i] {
			t.Errorf("Expected Game ID %d, got %d\n", ev, av)
		}
	}

	eURL := URL("https://www.igdb.com/people/hideo-kohima")
	aURL := p[1].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	es := "hideo-kohima"
	as := p[1].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	eg := GenderCode(0)
	ag := p[1].Gender
	if ag != eg {
		t.Errorf("Expected Gender code %d, got %d", eg, ag)
	}
}
