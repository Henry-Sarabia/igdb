package igdb

// CountryCode is a code representing
// a country according to the ISO-3316-1
// standard. For the full list of codes,
// visit: https://en.wikipedia.org/wiki/ISO_3166-1_numeric
type CountryCode int

// CreditCategory represents the IGDB
// enumerated type Credit Category which
// describes the type of an entry in an
// end credits list. Use the Stringer
// interface to access the corresponding
// Credit Category values as strings.
type CreditCategory int

// DateCategory represents the IGDB
// enumerated type Date Category which
// describes the format of a release
// date. Use the Stringer interface to
// access the corresponding Date Category
// values as strings.
type DateCategory int

// ESRBCode represents the IGDB
// enumerated type ESRB Rating which
// describes the recommended age group
// for consumers of a specific game.
// Use the Stringer interface to
// access the corresponding ESRB Rating
// values as strings.
type ESRBCode int

// FeatureCategory represents the IGDB enumerated type Feature Category which
// describes a type of feature value. The Feature value can either be a
// Boolean, a string, or an empty string which means "Preorder Only". Use
// the Stringer interface to access the corresponding Feature Category
// values as strings.
type FeatureCategory int

// FeedCategory represents the IGDB
// enumerated type Feed Item Category
// which describes the type of a feed
// item in a specific feed. Use the
// Stringer interface to access the
// corresponding Feed Item Category
// values as strings.
type FeedCategory int

// GameCategory represents the IGDB
// enumerated type Game Category which
// describes a type of game content.
// Use the Stringer interface to access
// the corresponding Game Category
// values as strings.
type GameCategory int

// GameStatus represents the IGDB
// enumerated type Game Status which
// describes the release status of
// a specific game. Use the Stringer
// interface to access the corresponding
// Game Status values as strings.
type GameStatus int

// GenderCode represents the IGDB
// enumerated type Gender which describes
// the gender of a specific entity. Use
// the Stringer interface to access the
// corresponding Gender values as strings.
type GenderCode int

// PEGICode represents the IGDB
// enumerated type PEGI Rating which
// describes the recommended minimum
// age for consumers of a specific
// game. Use the Stringer interface to
// access the corresponding PEGI Rating
// values as strings.
type PEGICode int

// RegionCode represents the IGDB
// enumerated type Region which describes
// a geographic region. Use the Stringer
// interface to access the corresponding
// Region values as strings.
type RegionCode int

// SpeciesCode represents the IGDB
// enumerated type Species which describes
// the species of a specific entity. Use
// the Stringer interface to access the
// corresponding species values as strings.
type SpeciesCode int

// WebsiteCategory represents the IGDB
// enumerated type Website Category which
// simply describes the category in which
// a website or URL falls under. Use the
// Stringer interface to access the
// corresponding category values as strings.
type WebsiteCategory int

// CreditCategory implements the Stringer interface
// by matching its code with the IGDBs enumerated type
// Credit Category and returns the category as a string.
// Codes with no match will return "Undefined".
// For the list of codes, visit:
// https://igdb.github.io/api/enum-fields/credit-category/
func (c CreditCategory) String() string {
	switch c {
	case 1:
		return "voice_actor"
	case 2:
		return "language"
	case 3:
		return "company_credit"
	case 4:
		return "employee"
	case 5:
		return "misc"
	case 6:
		return "support_company"
	default:
		return "Undefined"
	}
}

// DateCategory implements the Stringer interface
// by matching its code with the IGDBs enumerated type
// Date Category and returns the category as a string.
// Codes with no match will return "Undefined".
// For the list of codes, visit:
// https://igdb.github.io/api/enum-fields/date-category/
func (d DateCategory) String() string {
	switch d {
	case 0:
		return "YYYY MM DD"
	case 1:
		return "YYYY-MMM"
	case 2:
		return "YYYY"
	case 3:
		return "YYYY-Q1"
	case 4:
		return "YYYY-Q2"
	case 5:
		return "YYYY-Q3"
	case 6:
		return "YYYY-Q4"
	case 7:
		return "TBD"
	default:
		return "Undefined"
	}
}

// ESRBCode implements the Stringer interface
// by matching its code with the IGDBs enumerated
// type ESRB Rating and returns the rating as a
// string. Codes with no match will return "Undefined".
// For the list of codes, visit:
// https://igdb.github.io/api/enum-fields/esrb-rating/
func (e ESRBCode) String() string {
	switch e {
	case 1:
		return "RP"
	case 2:
		return "EC"
	case 3:
		return "E"
	case 4:
		return "E10+"
	case 5:
		return "T"
	case 6:
		return "M"
	case 7:
		return "AO"
	default:
		return "Undefined"
	}
}

// FeatureCategory implements the Stringer interface by matching its code with
// the IGDB's enumerated type Feature Category and returned the category as a
// string. Codes with no match will return "Undefined". For the list of codes,
// visit: https://igdb.github.io/api/endpoints/versions/
func (f FeatureCategory) String() string {
	switch f {
	case 0:
		return "Boolean. String should either be “1” or “0” (yes or no)."
	case 1:
		return "String. Free text."
	case 2:
		return "Preorder only. Whether the feature is only available in preorder. Value is always an empty string."
	default:
		return "Undefined"
	}
}

// FeedCategory implements the Stringer interface
// by matching its code with the IGDBs enumerated
// type Feed Item Category and returns the category
// as a string. Codes with no match will return
// "Undefined". For the list of codes, visit:
// https://igdb.github.io/api/enum-fields/feed-item-category/
func (f FeedCategory) String() string {
	switch f {
	case 1:
		return "Pulse Article"
	case 2:
		return "Coming Soon"
	case 3:
		return "New Trailer"
	case 5:
		return "User Contributed Item"
	case 6:
		return "User Contributions Item"
	case 7:
		return "Page Contributed Item"
	default:
		return "Undefined"
	}
}

// GameCategory implements the stringer interface
// by matching its code with the IGDBs enumerated type
// Game Category and returns the category as a string.
// Codes with no match will return "Undefined".
// For the list of codes, visit:
// https://igdb.github.io/api/enum-fields/game-category/
func (g GameCategory) String() string {
	switch g {
	case 0:
		return "Main Game"
	case 1:
		return "DLC / Addon"
	case 2:
		return "Expansion"
	case 3:
		return "Bundle"
	case 4:
		return "Standalone Expansion"
	default:
		return "Undefined"
	}
}

// GameStatus implements the Stringer interface
// by matching its code with the IGDBs enumerated
// Game Status type and returns the status as a
// string. Codes with no match return "Undefined".
// For the list of codes, visit:
// https://igdb.github.io/api/enum-fields/game-status/
func (g GameStatus) String() string {
	switch g {
	case 0:
		return "Released"
	case 2:
		return "Alpha"
	case 3:
		return "Beta"
	case 4:
		return "Early Access"
	case 5:
		return "Offline"
	case 6:
		return "Cancelled"
	default:
		return "Undefined"
	}
}

// GenderCode implements the Stringer interface
// by matching its code with the IGDBs enumerated
// type Gender and returns the value as a string.
// Codes with no match will return "Undefined".
// For the list of codes, visit:
// https://igdb.github.io/api/enum-fields/gender/
func (g GenderCode) String() string {
	switch g {
	case 0:
		return "Male"
	case 1:
		return "Female"
	case 2:
		return "Unknown"
	default:
		return "Undefined"
	}
}

// PEGICode implements the Stringer interface
// by matching its code with the IGDBs enumerated
// type PEGI Rating and returns the value as a string.
// Codes with no match will return "Undefined".
// For the list of codes, visit:
// https://igdb.github.io/api/enum-fields/pegi-rating/
func (p PEGICode) String() string {
	switch p {
	case 1:
		return "3"
	case 2:
		return "7"
	case 3:
		return "12"
	case 4:
		return "16"
	case 5:
		return "18"
	default:
		return "Undefined"
	}
}

// RegionCode implements the Stringer interface
// by matching its code with the IGDBs enumerated
// type Region and returns the value as a string.
// Codes with no match return "Undefined".
// For the list of codes, visit:
// https://igdb.github.io/api/enum-fields/region/
func (r RegionCode) String() string {
	switch r {
	case 1:
		return "Europe (EU)"
	case 2:
		return "North America (NA)"
	case 3:
		return "Australia (AU)"
	case 4:
		return "New Zealand (NZ)"
	case 5:
		return "Japan (JP)"
	case 6:
		return "China (CH)"
	case 7:
		return "Asia (AS)"
	case 8:
		return "Worldwide"
	default:
		return "Undefined"
	}
}

// String will return the enumerated type as a string
// corresponding to its IGDB code. For more information,
// SpeciesCode implements the Stringer interface
// by matching its code with the IGDBs enumerated
// type Species and returns the value as a string.
// For the list of codes, visit:
// https://igdb.github.io/api/enum-fields/species/
func (s SpeciesCode) String() string {
	switch s {
	case 1:
		return "Human"
	case 2:
		return "Alien"
	case 3:
		return "Animal"
	case 4:
		return "Android"
	case 5:
		return "Unknown"
	default:
		return "Undefined"
	}
}

// WebsiteCategory implements the Stringer interface
// by matching its code with the IGDBs enumerated
// Website Category and returns the value as a
// string. Codes with no match return "Undefined".
// For the list of codes, visit:
// https://igdb.github.io/api/enum-fields/website-category/
func (w WebsiteCategory) String() string {
	switch w {
	case 1:
		return "official"
	case 2:
		return "wikia"
	case 3:
		return "wikipedia"
	case 4:
		return "facebook"
	case 5:
		return "twitter"
	case 6:
		return "twitch"
	case 8:
		return "instagram"
	case 9:
		return "youtube"
	case 10:
		return "iphone"
	case 11:
		return "ipad"
	case 12:
		return "android"
	case 13:
		return "steam"
	default:
		return "Undefined"
	}
}
