package igdb

import (
	"net/http"
	"testing"
)

const getPageResp = `
[{
	"id": 8,
	"name": "IGN",
	"created_at": 1488280165428,
	"updated_at": 1497514032383,
	"slug": "ign",
	"url": "https://www.igdb.com/p/ign",
	"page_follows_count": 2,
	"category": 1,
	"description": "The latest game reviews, trailers, and walkthroughs from the #1 games media company.  As the leading source for gaming news on the net, IGN.com brings you the inside scoop on what's hot in the world of videogames and entertainment.  Tune in, turn on, game out!  This is only a small taste of the videos and content available from IGN. Get more, visit http://www.ign.com  Sign up for IGN's weekly top video Email at: http://go.ign.com/VideoRound-up  Interested in a career at IGN? Check us out: http://corp.ign.com/careers",
	"youtube": "https://www.youtube.com/ign",
	"logo": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/qnvv0psk2pze4v2gh0pz.jpg",
		"cloudinary_id": "qnvv0psk2pze4v2gh0pz",
		"width": 240,
		"height": 240
	},
	"background": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/g1psn86wb0rjfkxj20qd.jpg",
		"cloudinary_id": "g1psn86wb0rjfkxj20qd",
		"width": 1920,
		"height": 1080
	}
}]
`

const getPagesResp = `
[{
	"id": 36,
	"name": "TotalBiscuit, The Cynical Brit",
	"created_at": 1488280248232,
	"updated_at": 1491696787338,
	"slug": "cynicalbrit",
	"url": "https://www.igdb.com/p/cynicalbrit",
	"page_follows_count": 6,
	"category": 1,
	"description": "YouTube's #1 PC gaming critic. Sick of game reviews? Watch lengthy first impressions gameplay with honest and informative commentary. Get an idea about what you're getting into before you spend your money with \"WTF is..\", YouTube's most popular first impressions gaming series  MEDIA/DEVELOPER & PUBLISHER INQUIRIES: pr[at]cynicalbrit.com  WARNING: ONLY SEND KEYS TO THE LISTED EMAIL ADDRESSES. SEVERAL PEOPLE HAVE ATTEMPTED TO SCAM KEYS WITH SPOOFED ADDRESSES.  The channel does not take requests.  System Specs:  CPU: Intel Core i7-5930K @ 4.0ghz. RAM: Corsair Vengeance LPX 16GB DDR4-2800 BOARD: ASUS Rampage V Extreme COOLER: Corsair H110i GPU: 2x Nvidia Geforce GTX 1080 SSDs: 3x Sandisk Extreme Pro 480gb, 1x Samsung 950 PCI-E SSD. OS: Windows 7 x64 Ultimate. DISPLAY: Asus ROG SWIFT PG278Q 1440p 144hz G-sync monitor.  Capture: DXTory/FRAPs Microphone: EV RE20 Audio interface: RME Babyface.",
	"youtube": "https://www.youtube.com/user/TotalHalibut",
	"logo": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/yaaghei5o71o485emdqk.jpg",
		"cloudinary_id": "yaaghei5o71o485emdqk",
		"width": 240,
		"height": 240
	},
	"background": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/hntp6diwfxbuqq7akrje.jpg",
		"cloudinary_id": "hntp6diwfxbuqq7akrje",
		"width": 1920,
		"height": 1080
	}
},
{
	"id": 215,
	"name": "FantasticalGamer",
	"created_at": 1488280796684,
	"updated_at": 1488287514804,
	"slug": "fantasticalgamer",
	"url": "https://www.igdb.com/p/fantasticalgamer",
	"page_follows_count": 0,
	"description": "Name ► Sam or Fantastical  Age ► 19 Bio ► I Post Gaming Videos & Am Passionate About Film! I Am Also A Member of TheObeyAlliance Next Goal ► 1,000,000 Subscribers Content ► Gaming & Vlogs Partnered By ► OmniaMedia Sponsored By ►  Plunder League, DX Racer, Turtle Beach, & Corsair",
	"youtube": "https://www.youtube.com/channel/UCeCYRv2Odpjfr41ts_PkiTw",
	"logo": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/skm9bmn2roasynazwfxv.jpg",
		"cloudinary_id": "skm9bmn2roasynazwfxv",
		"width": 240,
		"height": 240
	},
	"background": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/yomgallpepunoxnjqg1j.jpg",
		"cloudinary_id": "yomgallpepunoxnjqg1j",
		"width": 1920,
		"height": 1080
	}
}]
`

const searchPagesResp = `
[{
	"id": 133,
	"name": "FRANKIEonPC",
	"created_at": 1488280541409,
	"updated_at": 1488982145551,
	"slug": "frankieonpc",
	"url": "https://www.igdb.com/p/frankieonpc",
	"country": 826,
	"page_follows_count": 1,
	"category": 1,
	"description": "Hopefully some entertaining videos!  The best way to contact me if you have any concerns / ideas / interesting things is by email (below) and I will get back to you in a few hours usually!",
	"facebook": "https://www.facebook.com/Frankieonpcin1080p",
	"twitter": "https://twitter.com/frankieonpc",
	"twitch": "https://www.twitch.tv/frankieonpcin1080p",
	"youtube": "https://www.youtube.com/frankieonpc",
	"logo": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/ywletykhceqsyaqvqkdn.jpg",
		"cloudinary_id": "ywletykhceqsyaqvqkdn",
		"width": 240,
		"height": 240
	},
	"background": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/o4b1inee4xyyq26x8jiz.jpg",
		"cloudinary_id": "o4b1inee4xyyq26x8jiz",
		"width": 1920,
		"height": 1080
	}
},
{
	"id": 550,
	"name": "PC Gamer",
	"created_at": 1492618811752,
	"updated_at": 1493056315218,
	"slug": "pc-gamer",
	"url": "https://www.igdb.com/p/pc-gamer",
	"page_follows_count": 0,
	"category": 2,
	"sub_category": 1,
	"description": "The global authority on PC games. ",
	"facebook": "https://www.facebook.com/pcgamermagazine",
	"twitter": "https://twitter.com/pcgamer",
	"youtube": "https://www.youtube.com/user/pcgamer",
	"logo": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/dmfigjz0frsofnfuxzch.jpg",
		"cloudinary_id": "dmfigjz0frsofnfuxzch",
		"width": 480,
		"height": 480
	},
	"background": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/z8fcgxj1moujzq9eal7w.jpg",
		"cloudinary_id": "z8fcgxj1moujzq9eal7w",
		"width": 2120,
		"height": 351
	}
}]
`

func TestGetPage(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getPageResp)
	defer ts.Close()

	pg, err := c.GetPage(8)
	if err != nil {
		t.Error(err)
	}

	en := "IGN"
	an := pg.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	ep := 2
	ap := pg.PageFollowCount
	if ap != ep {
		t.Errorf("Expected ID %d, got %d", ep, ap)
	}

	eyt := "https://www.youtube.com/ign"
	ayt := pg.Youtube
	if ayt != eyt {
		t.Errorf("Expected URL '%s', got '%s'", eyt, ayt)
	}

	ew := 1920
	aw := pg.Background.Width
	if aw != ew {
		t.Errorf("Expected width of %d, got %d", ew, aw)
	}
}

func TestGetPages(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getPagesResp)
	defer ts.Close()

	ids := []int{36, 215}
	pg, err := c.GetPages(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(pg)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "TotalBiscuit, The Cynical Brit"
	an := pg[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eyt := "https://www.youtube.com/user/TotalHalibut"
	ayt := pg[0].Youtube
	if ayt != eyt {
		t.Errorf("Expected URL '%s', got '%s'", eyt, ayt)
	}

	eu := 1488287514804
	au := pg[1].UpdatedAt
	if au != eu {
		t.Errorf("Expected unix epoch of %d, got %d", eu, au)
	}

	eh := 240
	ah := pg[1].Logo.Height
	if ah != eh {
		t.Errorf("Expected height of %d, got %d", eh, ah)
	}
}

func TestSearchPages(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, searchPagesResp)
	defer ts.Close()

	pg, err := c.SearchPages("PC")
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(pg)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	eID := 133
	aID := pg[0].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	ec := CountryCode(826)
	ac := pg[0].Country
	if ac != ec {
		t.Errorf("Expected country code %d, got %d", ec, ac)
	}

	efb := "https://www.facebook.com/Frankieonpcin1080p"
	afb := pg[0].Facebook
	if afb != efb {
		t.Errorf("Expected Facebook info '%s', got '%s'", efb, afb)
	}

	eCat := 2
	aCat := pg[1].Category
	if aCat != eCat {
		t.Errorf("Expected category %d, got %d", eCat, aCat)
	}

	ed := "The global authority on PC games. "
	ad := pg[1].Description
	if ad != ed {
		t.Errorf("Expected description '%s', got '%s'", ed, ad)
	}

	etw := "https://twitter.com/pcgamer"
	atw := pg[1].Twitter
	if atw != etw {
		t.Errorf("Expected Twitter info '%s', got '%s'", etw, atw)
	}
}
