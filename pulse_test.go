package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

const getPulseResp = `
[{
	"id": 145346,
	"category": 2,
	"title": "Nintendo announces new Mario, Zelda amiibo",
	"summary": "Seven new amiibo are coming, and they look great",
	"image": "https://cdn.vox-cdn.com/thumbor/XyUL3xlst_v2qfQf3X0lWoJhV5k=/0x0:1280x720/1600x900/cdn.vox-cdn.com/uploads/chorus_image/image/55234879/vlcsnap_2017_06_13_12h50m33s654.0.png",
	"url": "https://www.polygon.com/e3/2017/6/13/15792800/nintendo-e3-2017-mario-zelda-amiibo",
	"uid": "0cf56e50dac7df79b22a7917d092259f",
	"author": "Ben Kuchera",
	"created_at": 1497375274087,
	"updated_at": 1497396776349,
	"pulse_source": 2,
	"published_at": 1497374540000,
	"pulse_image": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/i1fti435exzyu1ydftu4.jpg",
		"cloudinary_id": "i1fti435exzyu1ydftu4",
		"width": 1600,
		"height": 900
	},
	"tags": [
		1,
		17,
		38,
		268435468,
		268435487,
		536871422,
		536872221
	]
}]
`

const getPulsesResp = `
[{
	"id": 132354,
	"category": 10,
	"title": "Battleborn: a great game with a fatal flaw",
	"summary": "Things could have been different, . One year ago today, Battleborn released to the world. And sure, we have fun with Randy Pitchford's ridiculous tweet about it, but the game holds a secret not many know: actually, it is very good. It could have been popular. Heck, I think it...",
	"image": "https://www.destructoid.com//ul/434247-Rendain.jpg",
	"url": "http://feedproxy.google.com/~r/Destructoid/~3/Icdv9HZL2V8/battleborn-a-great-game-with-a-fatal-flaw-434247.phtml",
	"uid": "d98ff53330026bd45afdb0d28949fab7",
	"author": "Darren Nakamura",
	"created_at": 1493846597080,
	"updated_at": 1493846597080,
	"pulse_source": 10,
	"published_at": 1493827200000,
	"pulse_image": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/fmtbomzsqfc5uvgguked.jpg",
		"cloudinary_id": "fmtbomzsqfc5uvgguked",
		"width": 1920,
		"height": 1080
	},
	"tags": [
		1,
		17,
		18,
		27,
		268435461,
		268435468,
		536871378,
		536871458,
		536872244,
		536873085,
		536873286,
		536873287,
		805314055,
		1073741825
	]
},
{
	"id": 257394,
	"category": 18,
	"title": "Cliff Bleszinski Reckons His Personality May have Something To Do With LawBreaker’s Lukewarm Reception",
	"summary": "But he also believes that this is a marathon, not a sprint.",
	"image": "http://gamingbolt.com/wp-content/uploads/2016/03/LawBreakers.jpg",
	"url": "http://gamingbolt.com/cliff-bleszinski-reckons-his-personality-may-have-something-to-do-with-lawbreakers-lukewarm-reception",
	"uid": "5fc0c5269a2aa7887d1a2c13a27c5bd2",
	"author": "Pramath",
	"created_at": 1502757342455,
	"updated_at": 1502757342455,
	"pulse_source": 18,
	"published_at": 1502756274000,
	"pulse_image": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/ibmrifg0uxp8w3y6hfzo.jpg",
		"cloudinary_id": "ibmrifg0uxp8w3y6hfzo",
		"width": 620,
		"height": 349
	},
	"tags": [
		1,
		18,
		268435461,
		536871049,
		536871378,
		536871436,
		536871882,
		536871937,
		536872363,
		805318165,
		1073741825
	]
},
{
	"id": 109415,
	"category": 8,
	"title": "Overwatch's 'Winter Wonderland' Event Now Live",
	"summary": "Blizzard is kicking off the holidays in Overwatch with winter-themed skins, maps, and more.",
	"url": "http://www.ign.com/articles/2016/12/13/overwatch-winter-wonderland-event-now-live",
	"uid": "585051a4e4b079863aa40471",
	"author": "Brandin Tyrrel",
	"created_at": 1481664038655,
	"updated_at": 1481664038655,
	"pulse_source": 8,
	"published_at": 1481659863000,
	"tags": [
		1,
		18,
		27,
		268435461,
		268435468,
		268435471,
		536871198
	]
}]
`

const searchPulsesResp = `
[{
	"id": 255433,
	"category": 10,
	"title": "Review: Mega Man Legacy Collection 2",
	"summary": "8 gets its due, . JUMP JUMP!SLIDE SLIDE!",
	"image": "https://www.destructoid.com//ul/452158-AAA1.jpg",
	"url": "http://feedproxy.google.com/~r/Destructoid/~3/Q5Gll61WQPM/review-mega-man-legacy-collection-2-452158.phtml",
	"uid": "4ba55b81080f244a9b405d33b5228877",
	"author": "Chris Carter",
	"created_at": 1502176226691,
	"updated_at": 1502177553457,
	"pulse_source": 10,
	"published_at": 1502157660000,
	"videos": [
		{
			"category": 1,
			"video_id": "iamiA3ki4-o"
		}
	],
	"pulse_image": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/pz3baoetyfezjx2j0nw0.jpg",
		"cloudinary_id": "pz3baoetyfezjx2j0nw0",
		"width": 1258,
		"height": 499
	}
},
{
	"id": 260033,
	"category": 10,
	"title": "Retroid talks Mega Man tonight",
	"summary": "Get equipped with Question Beam, . Retroid is recording our sixth episode tonight, and for this one we'll be taking a look back at the Mega Man/Rockman franchise. There have been nearly 100 games tied to Mega Man, so there's a lot to cover! Fortunately, I'll be joined by som...",
	"url": "http://feedproxy.google.com/~r/Destructoid/~3/ChzHjztL10M/retroid-talks-mega-man-tonight-456282.phtml",
	"uid": "b08c5a4959be5c3deb6da77bc8cb07a9",
	"author": "Kevin McClusky",
	"created_at": 1503433253964,
	"updated_at": 1503435195108,
	"pulse_source": 10,
	"published_at": 1503414000000
},
{
	"id": 105878,
	"category": 1,
	"title": "Mega Man Is Very Sad",
	"summary": "<img src=\"https://i.kinja-img.com/gawker-media/image/upload/s--P-wPzHZq--/c_fit,fl_progressive,q_80,w_636/mg1trrbab9lyzriu6vgu.png\"><p>Mega Man fans are sad, this we know. But I had no idea that Mega Man himself was “imbued with sadness” by his creator.<br></p><p><a href=\"http://kotaku.com/mega-man-is-very-sad-1789064690\">Read more...</a></p>",
	"image": "https://i.kinja-img.com/gawker-media/image/upload/s--P-wPzHZq--/c_fit,fl_progressive,q_80,w_636/mg1trrbab9lyzriu6vgu.png",
	"url": "http://kotaku.com/mega-man-is-very-sad-1789064690",
	"uid": "1789064690",
	"created_at": 1479378087141,
	"updated_at": 1479383085195,
	"pulse_source": 1,
	"published_at": 1479339000000,
	"pulse_image": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/y4wwcqqbkuyeteoq4l2n.jpg",
		"cloudinary_id": "y4wwcqqbkuyeteoq4l2n",
		"width": 636,
		"height": 365
	}
}]
`

func TestPulseTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	p := Pulse{}
	typ := reflect.ValueOf(p).Type()

	err := c.validateStruct(typ, PulseEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPulse(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getPulseResp)
	defer ts.Close()

	p, err := c.GetPulse(145346)
	if err != nil {
		t.Error(err)
	}

	et := "Nintendo announces new Mario, Zelda amiibo"
	at := p.Title
	if at != et {
		t.Errorf("Expected title '%s', got '%s'", et, at)
	}

	ep := 2
	ap := p.PulseSource
	if ap != ep {
		t.Errorf("Expected Pulse Source %d, got %d", ep, ap)
	}

	eURL := URL("//images.igdb.com/igdb/image/upload/t_thumb/i1fti435exzyu1ydftu4.jpg")
	aURL := p.PulseImage.URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	etID := []Tag{1, 17, 38, 268435468, 268435487, 536871422, 536872221}
	atID := p.Tags
	for i := range atID {
		if atID[i] != etID[i] {
			t.Errorf("Expected Tag ID %d, got %d\n", etID[i], atID[i])
		}
	}
}

func TestGetPulses(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getPulsesResp)
	defer ts.Close()

	ids := []int{132354, 257394, 109415}
	p, err := c.GetPulses(ids)
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(p)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	et := "Battleborn: a great game with a fatal flaw"
	at := p[0].Title
	if at != et {
		t.Errorf("Expected title '%s', got '%s'", et, at)
	}

	ea := "Darren Nakamura"
	aa := p[0].Author
	if aa != ea {
		t.Errorf("Expected slug '%s', got '%s'", ea, aa)
	}

	eUID := "5fc0c5269a2aa7887d1a2c13a27c5bd2"
	aUID := p[1].UID
	if aUID != eUID {
		t.Errorf("Expected ID '%s', got '%s'", eUID, aUID)
	}

	eURL := URL("//images.igdb.com/igdb/image/upload/t_thumb/ibmrifg0uxp8w3y6hfzo.jpg")
	aURL := p[1].PulseImage.URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	ep := 8
	ap := p[2].PulseSource
	if ap != ep {
		t.Errorf("Expected Pulse Source %d, got %d", ep, ap)
	}

	etID := []Tag{1, 18, 27, 268435461, 268435468, 268435471, 536871198}
	atID := p[2].Tags
	for i := range atID {
		if atID[i] != etID[i] {
			t.Errorf("Expected Tag ID %d, got %d\n", etID[i], atID[i])
		}
	}
}

func TestSearchPulses(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, searchPulsesResp)
	defer ts.Close()

	p, err := c.SearchPulses("megaman")
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(p)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	eCat := 10
	aCat := p[0].Category
	if aCat != eCat {
		t.Errorf("Expected category %d, got %d", eCat, aCat)
	}

	ec := 1502176226691
	ac := p[0].CreatedAt
	if ac != ec {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
	}

	et := "Retroid talks Mega Man tonight"
	at := p[1].Title
	if at != et {
		t.Errorf("Expected title '%s', got '%s'", et, at)
	}

	eURL := URL("http://feedproxy.google.com/~r/Destructoid/~3/ChzHjztL10M/retroid-talks-mega-man-tonight-456282.phtml")
	aURL := p[1].URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	ep := 1479339000000
	ap := p[2].PublishedAt
	if ap != ep {
		t.Errorf("Expected Unix time in milliseconds of %d, got %d", ep, ap)
	}

	eID := "y4wwcqqbkuyeteoq4l2n"
	aID := p[2].PulseImage.ID
	if aID != eID {
		t.Errorf("Expected ID '%s', got '%s'", eID, aID)
	}
}
