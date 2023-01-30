// +build ignore

package igdb

import "testing"

func TestCreditCategory(t *testing.T) {
	var creditTests = []struct {
		Name      string
		Credit    CreditCategory
		ExpString string
	}{
		{"Voice actor", CreditCategory(1), "voice_actor"},
		{"Language", CreditCategory(2), "language"},
		{"Company credit", CreditCategory(3), "company_credit"},
		{"Employee", CreditCategory(4), "employee"},
		{"Miscellaneous", CreditCategory(5), "misc"},
		{"Support company", CreditCategory(6), "support_company"},
		{"Undefined default", CreditCategory(100), "Undefined"},
	}
	for _, tt := range creditTests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.Credit.String() != tt.ExpString {
				t.Errorf("Expected Credit Category '%s', got '%s'", tt.ExpString, tt.Credit.String())
			}
		})
	}
}

func TestDateCategory(t *testing.T) {
	var dateTests = []struct {
		Name      string
		Date      DateCategory
		ExpString string
	}{
		{"Year Month Day", DateCategory(0), "YYYY MM DD"},
		{"Year-Month", DateCategory(1), "YYYY-MMM"},
		{"Year", DateCategory(2), "YYYY"},
		{"Year Quarter 1", DateCategory(3), "YYYY-Q1"},
		{"Year Quarter 2", DateCategory(4), "YYYY-Q2"},
		{"Year Quarter 3", DateCategory(5), "YYYY-Q3"},
		{"Year Quarter 4", DateCategory(6), "YYYY-Q4"},
		{"To Be Determined", DateCategory(7), "TBD"},
		{"Undefined default", DateCategory(100), "Undefined"},
	}
	for _, tt := range dateTests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.Date.String() != tt.ExpString {
				t.Errorf("Expected Date Category '%s', got '%s'", tt.ExpString, tt.Date.String())
			}
		})
	}
}

func TestESRBCode(t *testing.T) {
	var ESRBTests = []struct {
		Name      string
		ESRB      ESRBCode
		ExpString string
	}{
		{"Rating Pending", ESRBCode(1), "RP"},
		{"Early Childhood", ESRBCode(2), "EC"},
		{"Everyone", ESRBCode(3), "E"},
		{"Everyone 10+", ESRBCode(4), "E10+"},
		{"Teen", ESRBCode(5), "T"},
		{"Mature", ESRBCode(6), "M"},
		{"Adult Only", ESRBCode(7), "AO"},
		{"Undefined default", ESRBCode(100), "Undefined"},
	}
	for _, tt := range ESRBTests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.ESRB.String() != tt.ExpString {
				t.Errorf("Expected ESRB Code '%s', got '%s'", tt.ExpString, tt.ESRB.String())
			}
		})
	}
}

func TestFeatureCategory(t *testing.T) {
	var featureTests = []struct {
		Name      string
		Feature   FeatureCategory
		ExpString string
	}{
		{"Boolean", FeatureCategory(0), "Boolean. String should either be “1” or “0” (yes or no)."},
		{"String", FeatureCategory(1), "String. Free text."},
		{"Empty String", FeatureCategory(2), "Preorder only. Whether the feature is only available in preorder. Value is always an empty string."},
		{"Undefined default", FeatureCategory(100), "Undefined"},
	}
	for _, tt := range featureTests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.Feature.String() != tt.ExpString {
				t.Errorf("Expected Feature Category '%s', got '%s'", tt.ExpString, tt.Feature.String())
			}
		})
	}
}

func TestFeedCategory(t *testing.T) {
	var feedTests = []struct {
		Name      string
		Feed      FeedCategory
		ExpString string
	}{
		{"Rating pending", FeedCategory(1), "Pulse Article"},
		{"Coming soon", FeedCategory(2), "Coming Soon"},
		{"New trailer", FeedCategory(3), "New Trailer"},
		{"User contributed item", FeedCategory(5), "User Contributed Item"},
		{"User contributions item", FeedCategory(6), "User Contributions Item"},
		{"Page contributed item", FeedCategory(7), "Page Contributed Item"},
		{"Undefined default", FeedCategory(100), "Undefined"},
	}
	for _, tt := range feedTests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.Feed.String() != tt.ExpString {
				t.Errorf("Expected Feed Category '%s', got '%s'", tt.ExpString, tt.Feed.String())
			}
		})
	}
}

func TestGameCategory(t *testing.T) {
	var gameTests = []struct {
		Name      string
		Game      GameCategory
		ExpString string
	}{
		{"Main game", GameCategory(0), "Main Game"},
		{"DLC/Addon", GameCategory(1), "DLC / Addon"},
		{"Expansion", GameCategory(2), "Expansion"},
		{"Bundle", GameCategory(3), "Bundle"},
		{"Standalone expansion", GameCategory(4), "Standalone Expansion"},
		{"Mod", GameCategory(5), "Mod"},
		{"Episode", GameCategory(6), "Episode"},
		{"Season", GameCategory(7), "Season"},
		{"Remake", GameCategory(8), "Remake"},
		{"Remaster", GameCategory(9), "Remaster"},
		{"ExpandedGame", GameCategory(10), "ExpandedGame"},
		{"Port", GameCategory(11), "Port"},
		{"Fork", GameCategory(12), "Fork"},
		{"Pack", GameCategory(13), "Pack"},
		{"Update", GameCategory(14), "Update"},
		{"Undefined default", GameCategory(100), "Undefined"},
	}
	for _, tt := range gameTests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.Game.String() != tt.ExpString {
				t.Errorf("Expected Game Category '%s', got '%s'", tt.ExpString, tt.Game.String())
			}
		})
	}
}

func TestGameStatus(t *testing.T) {
	var gameTests = []struct {
		Name      string
		Game      GameStatus
		ExpString string
	}{
		{"Released", GameStatus(0), "Released"},
		{"Alpha", GameStatus(2), "Alpha"},
		{"Beta", GameStatus(3), "Beta"},
		{"Early access", GameStatus(4), "Early Access"},
		{"Offline", GameStatus(5), "Offline"},
		{"Cancelled", GameStatus(6), "Cancelled"},
		{"Undefined default", GameStatus(100), "Undefined"},
	}
	for _, tt := range gameTests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.Game.String() != tt.ExpString {
				t.Errorf("Expected Game Status '%s', got '%s'", tt.ExpString, tt.Game.String())
			}
		})
	}
}

func TestGenderCode(t *testing.T) {
	var genderTests = []struct {
		Name      string
		Gender    GenderCode
		ExpString string
	}{
		{"Male", GenderCode(0), "Male"},
		{"Female", GenderCode(1), "Female"},
		{"Unknown", GenderCode(2), "Unknown"},
		{"Undefined default", GenderCode(100), "Undefined"},
	}
	for _, tt := range genderTests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.Gender.String() != tt.ExpString {
				t.Errorf("Expected Gender '%s', got '%s'", tt.ExpString, tt.Gender.String())
			}
		})
	}
}

func TestPEGICode(t *testing.T) {
	var PEGITests = []struct {
		Name      string
		PEGI      PEGICode
		ExpString string
	}{
		{"Age 3", PEGICode(1), "3"},
		{"Age 7", PEGICode(2), "7"},
		{"Age 12", PEGICode(3), "12"},
		{"Age 16", PEGICode(4), "16"},
		{"Age 18", PEGICode(5), "18"},
		{"Undefined default", PEGICode(100), "Undefined"},
	}
	for _, tt := range PEGITests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.PEGI.String() != tt.ExpString {
				t.Errorf("Expected PEGI Code '%s', got '%s'", tt.ExpString, tt.PEGI.String())
			}
		})
	}
}

func TestRegionCode(t *testing.T) {
	var regionTests = []struct {
		Name      string
		Region    RegionCode
		ExpString string
	}{
		{"Europe", RegionCode(1), "Europe (EU)"},
		{"North America", RegionCode(2), "North America (NA)"},
		{"Australia", RegionCode(3), "Australia (AU)"},
		{"New Zealand", RegionCode(4), "New Zealand (NZ)"},
		{"Japan", RegionCode(5), "Japan (JP)"},
		{"China", RegionCode(6), "China (CH)"},
		{"Asia", RegionCode(7), "Asia (AS)"},
		{"Worldwide", RegionCode(8), "Worldwide"},
		{"Undefined default", RegionCode(100), "Undefined"},
	}
	for _, tt := range regionTests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.Region.String() != tt.ExpString {
				t.Errorf("Expected Region Code '%s', got '%s'", tt.ExpString, tt.Region.String())
			}
		})
	}
}

func TestSpeciesCode(t *testing.T) {
	var speciesTests = []struct {
		Name      string
		Species   SpeciesCode
		ExpString string
	}{
		{"Human", SpeciesCode(1), "Human"},
		{"Alien", SpeciesCode(2), "Alien"},
		{"Animal", SpeciesCode(3), "Animal"},
		{"Android", SpeciesCode(4), "Android"},
		{"Unknown", SpeciesCode(5), "Unknown"},
		{"Undefined default", SpeciesCode(100), "Undefined"},
	}
	for _, tt := range speciesTests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.Species.String() != tt.ExpString {
				t.Errorf("Expected Species Code '%s', got '%s'", tt.ExpString, tt.Species.String())
			}
		})
	}
}

func TestWebsiteCategory(t *testing.T) {
	var websiteTests = []struct {
		Name      string
		Website   WebsiteCategory
		ExpString string
	}{
		{"Official", WebsiteCategory(1), "official"},
		{"Wikia", WebsiteCategory(2), "wikia"},
		{"Wikipedia", WebsiteCategory(3), "wikipedia"},
		{"Facebook", WebsiteCategory(4), "facebook"},
		{"Twitter", WebsiteCategory(5), "twitter"},
		{"Twitch", WebsiteCategory(6), "twitch"},
		{"Instagram", WebsiteCategory(8), "instagram"},
		{"Youtube", WebsiteCategory(9), "youtube"},
		{"iPhone", WebsiteCategory(10), "iphone"},
		{"iPad", WebsiteCategory(11), "ipad"},
		{"Android", WebsiteCategory(12), "android"},
		{"Steam", WebsiteCategory(13), "steam"},
		{"Undefined default", WebsiteCategory(100), "Undefined"},
	}
	for _, tt := range websiteTests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.Website.String() != tt.ExpString {
				t.Errorf("Expected Website Category '%s', got '%s'", tt.ExpString, tt.Website.String())
			}
		})
	}
}
