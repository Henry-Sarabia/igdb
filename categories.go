package igdb

// CountryCode is a code representing
// a country according to the ISO-3316-1
// standard. For a full list of codes,
// visit: https://en.wikipedia.org/wiki/ISO_3166-1_numeric
type CountryCode int

// CreditCategory corresponds to the IGDB
// enumerated credit type which describes
// the type of entry in an end credits list.
// CreditCategory implements the Stringer
// interface.
type CreditCategory int

// DateCategory corresponds to the IGDB
// enumerated Date Category type which
// describes the format of the value in
// the release date. DateCategory
// implements the Stringer interface.
type DateCategory int

// ESRBCode corresponds to the IGDB
// enumerated ESRB rating which
// describes the recommended age
// for consumers. ESRBCode implements
// the Stringer interface.
type ESRBCode int

// FeedCategory corresponds to the IGDB
// enumerated feed item category which
// describes the type of feed item in
// a particular feed. FeedCategory
// implements the Stringer interface.
type FeedCategory int

// GameCategory corresponds to the IGDB
// enumerated type Game Category which
// describes the type of game content.
// GameCategory implements the Stringer
// interface.
type GameCategory int

// GenderCode corresponds to the IGDB
// enumerated gender type which describes
// an entity's gender. GenderCode
// implements the Stringer interface.
type GenderCode int

// PEGICode corresponds to the IGDB
// enumerated PEGI rating which
// describes the recommended minimum
// age for consumers. EPGICode implements
// the Stringer interface.
type PEGICode int

// SpeciesCode corresponds to the IGDB
// enumerated species type which describes
// an entity's species. SpeciesCode
// implements the Stringer interface.
type SpeciesCode int

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
