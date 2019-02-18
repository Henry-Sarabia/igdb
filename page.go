package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Page -add-tags json -w

// Page represents an entry in the multipurpose page system
// currently used for youtubers and media organizations.
// For more information visit: https://api-docs.igdb.com/#page
type Page struct {
	ID               int             `json:"id"`
	Background       int             `json:"background"`
	Battlenet        string          `json:"battlenet"`
	Category         PageCategory    `json:"category"`
	Color            PageColor       `json:"color"`
	Company          int             `json:"company"`
	Country          int             `json:"country"`
	CreatedAt        int             `json:"created_at"`
	Description      string          `json:"description"`
	Feed             int             `json:"feed"`
	Game             int             `json:"game"`
	Name             string          `json:"name"`
	Origin           string          `json:"origin"`
	PageFollowsCount int             `json:"page_follows_count"`
	PageLogo         int             `json:"page_logo"`
	Slug             string          `json:"slug"`
	SubCategory      PageSubCategory `json:"sub_category"`
	UpdatedAt        int             `json:"updated_at"`
	Uplay            string          `json:"uplay"`
	URL              string          `json:"url"`
	User             int             `json:"user"`
	Websites         []int           `json:"websites"`
}

//go:generate stringer -type=PageCategory,PageSubCategory,PageColor

type PageCategory int

const (
	PagePersonality PageCategory = iota + 1
	PageMediaOrganization
	PageContentCreator
	PageClanTeam
)

type PageSubCategory int

const (
	PageUser PageSubCategory = iota + 1
	PageGame
	PageCompany
	PageConsumer
	PageIndustry
	PageESports
)

type PageColor int

const (
	PageGreen PageColor = iota
	PageBlue
	PageRed
	PageOrange
	PagePink
	PageYellow
)

// PageService handles all the API calls for the IGDB Page endpoint.
type PageService service

// Get returns a single Page identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Pages, an error is returned.
func (ps *PageService) Get(id int, opts ...Option) (*Page, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var pg []*Page

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &pg, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Page with ID %v", id)
	}

	return pg[0], nil
}

// List returns a list of Pages identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Page is ignored. If none of the IDs
// match a Page, an error is returned.
func (ps *PageService) List(ids []int, opts ...Option) ([]*Page, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var pg []*Page

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ps.client.get(ps.end, &pg, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Pages with IDs %v", ids)
	}

	return pg, nil
}

// Index returns an index of Pages based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Pages can
// be found using the provided options, an error is returned.
func (ps *PageService) Index(opts ...Option) ([]*Page, error) {
	var pg []*Page

	err := ps.client.get(ps.end, &pg, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Pages")
	}

	return pg, nil
}

// Count returns the number of Pages available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Pages to count.
func (ps *PageService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Pages")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Page object.
func (ps *PageService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Page fields")
	}

	return f, nil
}
