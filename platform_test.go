package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

const getPlatformResp = `
[{
	"id": 7,
	"name": "PlayStation",
	"logo": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/ptlxti6tzdpma71s5tkm.jpg",
		"cloudinary_id": "ptlxti6tzdpma71s5tkm",
		"width": 2000,
		"height": 1522
	},
	"slug": "ps",
	"url": "https://www.igdb.com/platforms/ps",
	"created_at": 1297639288000,
	"updated_at": 1392141728537,
	"generation": 5,
	"games": [
		1185,
		1186,
		1187,
		1192,
		1195,
		1201,
		425,
		675
	],
	"versions": [
		{
			"name": "Playstation 1",
			"slug": "playstation-1",
			"url": "https://www.igdb.com/platforms/ps/version/playstation-1",
			"cpu": "MIPS R3051 @ 33,8688 MHz",
			"storage": "Playstation Memory Card",
			"sound": "Stereo",
			"media": "CD-ROM Discs, Playstation game disc",
			"output": "S-Video, Composite",
			"manufacturers": [
				{
					"company": 45
				}
			],
			"logo": {
				"url": "//images.igdb.com/igdb/image/upload/t_thumb/ptlxti6tzdpma71s5tkm.png",
				"cloudinary_id": "ptlxti6tzdpma71s5tkm",
				"width": 2000,
				"height": 1522
			},
			"release_dates": [
				{
					"date": 786412800000,
					"region": 5
				},
				{
					"date": 810604800000,
					"region": 2
				},
				{
					"date": 812332800000,
					"region": 1
				},
				{
					"date": 816393600000,
					"region": 3
				}
			]
		}
	]
}]
`

const getPlatformsResp = `
[{
	"id": 23,
	"name": "Dreamcast",
	"logo": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/pj6p7nmyctusjiabjtrg.jpg",
		"cloudinary_id": "pj6p7nmyctusjiabjtrg",
		"width": 1024,
		"height": 768
	},
	"slug": "dc",
	"url": "https://www.igdb.com/platforms/dc",
	"created_at": 1300188004000,
	"updated_at": 1392140459420,
	"generation": 6,
	"games": [
		1218,
		1258,
		1259,
		968,
		1160,
		154,
		243,
		272
	],
	"versions": [
		{
			"name": "Initial version",
			"slug": "initial-version-62da4d4a-0faa-489f-bab1-a0ec26db72d9",
			"url": "https://www.igdb.com/platforms/dc/version/initial-version-62da4d4a-0faa-489f-bab1-a0ec26db72d9",
			"cpu": "Hitachi SH4 32-bit RISC @ 200 MHz",
			"storage": "VMU",
			"memory": "16 MB",
			"online": "SegaNet, GameSpy, Dreamarena",
			"media": "CD, 1.2 GB GD-ROM",
			"connectivity": "VHF, Composite Video, S-Video, VGA",
			"manufacturers": [
				{
					"company": 112
				}
			],
			"logo": {
				"url": "//images.igdb.com/igdb/image/upload/t_thumb/pj6p7nmyctusjiabjtrg.png",
				"cloudinary_id": "pj6p7nmyctusjiabjtrg",
				"width": 1024,
				"height": 768
			},
			"release_dates": [
				{
					"date": 912124800000,
					"region": 5
				},
				{
					"date": 936835200000,
					"region": 2
				},
				{
					"date": 939859200000,
					"region": 1
				},
				{
					"date": 943920000000,
					"region": 3
				}
			]
		}
	]
},
{
	"id": 130,
	"name": "Nintendo Switch",
	"logo": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/zj5x8hozy4fatoqk37nn.jpg",
		"cloudinary_id": "zj5x8hozy4fatoqk37nn",
		"width": 413,
		"height": 413
	},
	"slug": "nintendo-switch",
	"url": "https://www.igdb.com/platforms/nintendo-switch",
	"created_at": 1465978941719,
	"updated_at": 1482756176646,
	"website": "http://www.nintendo.com/switch",
	"alternative_name": "NX",
	"generation": 8,
	"games": [
		38983,
		10232,
		885,
		11529,
		19081,
		19175,
		29525,
		36846,
		27081,
		26165
	],
	"versions": [
		{
			"name": "Initial version",
			"slug": "initial-version",
			"url": "https://www.igdb.com/platforms/nintendo-switch/version/initial-version",
			"summary": "Nintendo Switch is a hybrid console/tablet. The tablet becomes a console via a docking station connected to a tv set. Nintendo regards the concept as mainly a home gaming system rather than portable. The Switch features two wireless controllers, which may be used individually or attached to a grip to get a more traditional controller.",
			"cpu": "Custom NVIDIA Tegra",
			"storage": "32 GB NAND flash, expandable via microSD cards",
			"sound": "5.1 channel Linear PCM (HDMI), analog stereo (3.5mm headphone jack)",
			"online": "Nintendo Switch Online Service",
			"media": "Nintendo Switch Game Card, microSD Card",
			"connectivity": "802.11 a/b/g/n/ac, Bluetooth 4.0 (TV Mode Only), 2x USB 3.0, 1x USB 2.0, 1x USB-C, HDMI",
			"resolutions": "1280x720 on integrated screen, up to 1920x1080 through HDMI",
			"output": "HDMI, 3.5mm headphone jack, DisplayPort over USB-C (internal)",
			"manufacturers": [
				{
					"company": 12004
				}
			],
			"developers": [
				{
					"company": 12004
				}
			],
			"logo": {
				"url": "//images.igdb.com/igdb/image/upload/t_thumb/zj5x8hozy4fatoqk37nn.png",
				"cloudinary_id": "zj5x8hozy4fatoqk37nn",
				"width": 413,
				"height": 413
			},
			"release_dates": [
				{
					"date": 1488499200000,
					"region": 8
				}
			]
		}
	]
}]
`

const searchPlatformsResp = `
[{
	"id": 11,
	"name": "Xbox",
	"logo": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/amc434vnchq0cssuy17g.jpg",
		"cloudinary_id": "amc434vnchq0cssuy17g",
		"width": 3703,
		"height": 1411
	},
	"slug": "xbox",
	"url": "https://www.igdb.com/platforms/xbox",
	"created_at": 1297639288000,
	"updated_at": 1392162016560,
	"generation": 6,
	"games": [
		1188,
		1189,
		1193,
		1194,
		1049,
		1218,
		1048,
		10,
		1137
	],
	"versions": [
		{
			"name": "Initial version",
			"slug": "initial-version-6ac4b90f-0745-458f-866d-3672aa6ed5eb",
			"url": "https://www.igdb.com/platforms/xbox/version/initial-version-6ac4b90f-0745-458f-866d-3672aa6ed5eb",
			"os": "Custom (Based on Windows NT architecture and Windows XP",
			"cpu": "Custom 733 MHz Intel Pentium III ''Coppermine-based''",
			"storage": "8 to 10 GB Internal Storage",
			"memory": "64 MB DDR SDRAM @ 200 MHz",
			"graphics": "233 Mhz nVidia NV2A",
			"online": "Xbox Live",
			"connectivity": "100 Mbit Ethernet",
			"manufacturers": [
				{
					"company": 2348
				}
			],
			"developers": [
				{
					"company": 128
				}
			],
			"logo": {
				"url": "//images.igdb.com/igdb/image/upload/t_thumb/amc434vnchq0cssuy17g.png",
				"cloudinary_id": "amc434vnchq0cssuy17g",
				"width": 3703,
				"height": 1411
			},
			"release_dates": [
				{
					"date": 1005782400000,
					"region": 2
				},
				{
					"date": 1014336000000,
					"region": 5
				},
				{
					"date": 1013644800000,
					"region": 1
				}
			]
		}
	]
},
{
	"id": 12,
	"name": "Xbox 360",
	"slug": "xbox360",
	"url": "https://www.igdb.com/platforms/xbox360",
	"created_at": 1297639288000,
	"updated_at": 1392919583542,
	"website": "http://www.xbox.com/en-US/xbox-360",
	"summary": "Xbox 360 brings you a total games and entertainment experience. The largest library of games, including titles that get you right into the thick of it with Kinect. Plus, your whole family can watch HD movies, TV shows, live events, music, sports and more—across all your devices. Xbox 360 is the center of your games and entertainment universe.",
	"alternative_name": "xbx360",
	"generation": 7,
	"games": [
		487,
		1190,
		1191,
		1059,
		1082,
		983,
		984,
		411,
		416
	],
	"versions": [
		{
			"name": "Xbox 360 Arcade",
			"slug": "xbox-360-arcade",
			"url": "https://www.igdb.com/platforms/xbox360/version/xbox-360-arcade",
			"release_dates": [
				{
					"date": 1193097600000,
					"region": 2
				}
			]
		},
		{
			"name": "Xbox 360 Original",
			"slug": "xbox-360-original",
			"url": "https://www.igdb.com/platforms/xbox360/version/xbox-360-original",
			"cpu": "3,2 GHz PowerPC Tri-Core ''Xenon''",
			"storage": "20 GB, 60 GB, 120 GB, 250 GB",
			"memory": "512 MB GDDR3 @ 700 MHz",
			"graphics": "ATI ''Xenos'' @ 500 MHz",
			"sound": "Analog Stereo, Dolby Digital 5,1",
			"online": "Xbox Live",
			"media": "DVD, Compact Dics, Digital Download",
			"connectivity": "Ethernet",
			"resolutions": "480i, 480p, 720p, 1080i, 1080p",
			"output": "HDMI",
			"logo": {
				"url": "//images.igdb.com/igdb/image/upload/t_thumb/tvvleg0mjvesgv6ylxxc.png",
				"cloudinary_id": "tvvleg0mjvesgv6ylxxc",
				"width": 400,
				"height": 260
			},
			"release_dates": [
				{
					"date": 1132617600000,
					"region": 2
				},
				{
					"date": 1133481600000,
					"region": 1
				},
				{
					"date": 1134172800000,
					"region": 5
				}
			]
		},
		{
			"name": "Xbox 360 Elite",
			"slug": "xbox-360-elite",
			"url": "https://www.igdb.com/platforms/xbox360/version/xbox-360-elite",
			"summary": "Xbox 360 brings you a total games and entertainment experience. The largest library of games, including titles that get you right into the thick of it with Kinect. Plus, your whole family can watch HD movies, TV shows, live events, music, sports and more—across all your devices. Xbox 360 is the center of your games and entertainment universe.",
			"logo": {
				"url": "//images.igdb.com/igdb/image/upload/t_thumb/kluedsa9yjdz0s6hblrm.png",
				"cloudinary_id": "kluedsa9yjdz0s6hblrm",
				"width": 1800,
				"height": 1066
			}
		}
	]
}]	
`

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

func TestGetPlatform(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getPlatformResp)
	defer ts.Close()

	p, err := c.GetPlatform(7)
	if err != nil {
		t.Error(err)
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
		t.Errorf("Expected unix epoch of %d, got %d", ed, ad)
	}
}

func TestGetPlatforms(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, getPlatformsResp)
	defer ts.Close()

	ids := []int{23, 130}
	p, err := c.GetPlatforms(ids)
	if err != nil {
		t.Error(err)
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
}

func TestSearchPlatforms(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, searchPlatformsResp)
	defer ts.Close()

	p, err := c.SearchPlatforms("xbox")
	if err != nil {
		t.Error(err)
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
		t.Errorf("Expected unix epoch of %d, got %d", ec, ac)
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
}
