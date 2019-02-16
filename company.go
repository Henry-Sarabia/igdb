package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Company -add-tags json -w

// Company represents a video game company.
// This includes both publishers and developers.
// For more information visit: https://api-docs.igdb.com/#company
type Company struct {
	ID                 int          `json:"id"`
	ChangeDate         int          `json:"change_date"`
	ChangeDateCategory DateCategory `json:"change_date_category"`
	ChangedCompanyID   int          `json:"changed_company_id"`
	Country            int          `json:"country"`
	CreatedAt          int          `json:"created_at"`
	Description        string       `json:"description"`
	Developed          []int        `json:"developed"`
	Logo               int          `json:"logo"`
	Name               string       `json:"name"`
	Parent             int          `json:"parent"`
	Published          []int        `json:"published"`
	Slug               string       `json:"slug"`
	StartDate          int          `json:"start_date"`
	StartDateCategory  DateCategory `json:"start_date_category"`
	UpdatedAt          int          `json:"updated_at"`
	URL                string       `json:"url"`
	Websites           []int        `json:"websites"`
}

// CompanyService handles all the API calls for the IGDB Company endpoint.
type CompanyService service

// Get returns a single Company identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Companies, an error is returned.
func (cs *CompanyService) Get(id int, opts ...Option) (*Company, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var comp []*Company

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.get(cs.end, &comp, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Company with ID %v", id)
	}

	return comp[0], nil
}

// List returns a list of Companies identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Company is ignored. If none of the IDs
// match a Company, an error is returned.
func (cs *CompanyService) List(ids []int, opts ...Option) ([]*Company, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var comp []*Company

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.get(cs.end, &comp, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Companies with IDs %v", ids)
	}

	return comp, nil
}

// Index returns an index of Companies based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Companies can
// be found using the provided options, an error is returned.
func (cs *CompanyService) Index(opts ...Option) ([]*Company, error) {
	var comp []*Company

	err := cs.client.get(cs.end, &comp, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Companies")
	}

	return comp, nil
}

// Count returns the number of Companies available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Companies to count.
func (cs *CompanyService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Companies")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Company object.
func (cs *CompanyService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Company fields")
	}

	return f, nil
}
