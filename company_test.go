package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

const GetCompanyResp = `
[{
	"id": 58,
	"name": "Mojang AB",
	"created_at": 1300110282000,
	"updated_at": 1499997426290,
	"slug": "mojang-ab",
	"url": "https://www.igdb.com/companies/mojang-ab",
	"logo": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/vewa4pfeszo8e9jncwex.jpg",
		"cloudinary_id": "vewa4pfeszo8e9jncwex",
		"width": 800,
		"height": 400
	},
	"description": "Mojang AB is a Swedish independent video game developer founded in May 2009. Mojang is best known for creating the popular independent game Minecraft, a free-to-build sandbox game.",
	"start_date": 1241128800000,
	"start_date_category": 0,
	"country": 752,
	"website": "https://mojang.com",
	"parent": 1010,
	"change_date_category": 7,
	"twitter": "https://twitter.com/Mojang",
	"published": [
		2600,
		1898,
		121,
		18977,
		2999,
		8339
	],
	"developed": [
		2600,
		1898,
		121,
		18977,
		8339
	]
}]
`

const GetCompaniesResp = `
[{
	"id": 854,
	"name": "Playdead",
	"created_at": 1348518254559,
	"updated_at": 1504811027097,
	"slug": "playdead",
	"url": "https://www.igdb.com/companies/playdead",
	"logo": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/twygesjnprgnhfya3lrw.jpg",
		"cloudinary_id": "twygesjnprgnhfya3lrw",
		"width": 371,
		"height": 52
	},
	"description": "Playdead is an independent game developer and publisher based in Copenhagen, Denmark. Playdead was founded in 2006 by Arnt Jensen and Dino Patti, to develop LIMBO, which was released in 2010. Today we’re more than 25 people, creating in various aspects in the company. A small team are currently working on expanding LIMBO to new platforms and the rest are working on Arnt Jensen's new game INSIDE.",
	"start_date": 1149120000000,
	"start_date_category": 0,
	"country": 208,
	"website": "http://playdead.com/",
	"change_date_category": 7,
	"published": [
		1331,
		7342
	],
	"developed": [
		59851,
		1331,
		7342
	]
},
{
	"id": 7260,
	"name": "Night School Studio",
	"created_at": 1447881708788,
	"updated_at": 1496952641480,
	"slug": "night-school-studio",
	"url": "https://www.igdb.com/companies/night-school-studio",
	"start_date_category": 7,
	"change_date_category": 7,
	"published": [
		14587
	],
	"developed": [
		14587,
		22748,
		36858
	]
}]
`

const SearchCompaniesResp = `
[{
	"id": 6545,
	"name": "tobyfox",
	"created_at": 1442097213278,
	"updated_at": 1500415107616,
	"slug": "tobyfox",
	"url": "https://www.igdb.com/companies/tobyfox",
	"logo": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/asqsr9anzx2t22smxskt.jpg",
		"cloudinary_id": "asqsr9anzx2t22smxskt",
		"width": 700,
		"height": 335
	},
	"description": "Creator of UNDERTALE.\n\nComposes music for Homestuck.",
	"start_date": 1372024800000,
	"start_date_category": 0,
	"country": 840,
	"website": "http://tobyfox.net/",
	"parent": 3911,
	"change_date_category": 7,
	"published": [
		12517
	],
	"developed": [
		12517
	]
},
{
	"id": 5529,
	"name": "Terrible Toybox",
	"created_at": 1431324943488,
	"updated_at": 1490098043838,
	"slug": "terrible-toybox",
	"url": "https://www.igdb.com/companies/terrible-toybox",
	"logo": {
		"url": "//images.igdb.com/igdb/image/upload/t_thumb/xx9lm1jz1gldwhigj6xv.jpg",
		"cloudinary_id": "xx9lm1jz1gldwhigj6xv",
		"width": 480,
		"height": 360
	},
	"description": "Based in Seattle, Washington.",
	"start_date_category": 7,
	"country": 840,
	"change_date_category": 7,
	"published": [
		10232
	],
	"developed": [
		10232
	]
}]
`

func TestCompanyTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	com := Company{}
	typ := reflect.ValueOf(com).Type()

	err := c.validateStruct(typ, CompanyEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetCompany(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, GetCompanyResp)
	defer ts.Close()

	com, err := c.GetCompany(58)
	if err != nil {
		t.Error(err)
	}

	en := "Mojang AB"
	an := com.Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eURL := URL("https://www.igdb.com/companies/mojang-ab")
	aURL := com.URL
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	eID := []int{2600, 1898, 121, 18977, 8339}
	aID := com.Developed
	for i := range aID {
		if aID[i] != eID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", eID[i], aID[i])
		}
	}
}

func TestGetCompanies(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, GetCompaniesResp)
	defer ts.Close()

	ids := []int{854, 7260}
	com, err := c.GetCompanies(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(com)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	en := "Playdead"
	an := com[0].Name
	if an != en {
		t.Errorf("Expected name '%s', got '%s'", en, an)
	}

	eu := 1504811027097
	au := com[0].UpdatedAt
	if au != eu {
		t.Errorf("Expected unix epoch of %d, got %d", eu, au)
	}

	eURL := URL("https://www.igdb.com/companies/night-school-studio")
	aURL := com[1].URL
	if aURL != eURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	eID := []int{14587, 22748, 36858}
	aID := com[1].Developed
	for i := range aID {
		if aID[i] != eID[i] {
			t.Errorf("Expected Game ID %d, got %d\n", eID[i], aID[i])
		}
	}
}

func TestSearchCompanies(t *testing.T) {
	ts, c := startTestServer(http.StatusOK, SearchCompaniesResp)
	defer ts.Close()

	com, err := c.SearchCompanies("mario")
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(com)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	eID := 6545
	aID := com[0].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	eu := 1500415107616
	au := com[0].UpdatedAt
	if au != eu {
		t.Errorf("Expected unix epoch of %d, got %d", eu, au)
	}

	es := "terrible-toybox"
	as := com[1].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	ePub := 10232
	aPub := com[1].Published[0]
	if aPub != ePub {
		t.Errorf("Expected Game ID %d, got %d", ePub, aPub)
	}
}
