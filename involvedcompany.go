package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct InvolvedCompany -add-tags json -w

// InvolvedCompany represents a company involved in the development
// of a particular video game.
// For more information visit: https://api-docs.igdb.com/#involved-company
type InvolvedCompany struct {
	ID         int  `json:"id"`
	Company    int  `json:"company"`
	CreatedAt  int  `json:"created_at"`
	Developer  bool `json:"developer"`
	Game       int  `json:"game"`
	Porting    bool `json:"porting"`
	Publisher  bool `json:"publisher"`
	Supporting bool `json:"supporting"`
	UpdatedAt  int  `json:"updated_at"`
}

// InvolvedCompanyService handles all the API calls for the IGDB InvolvedCompany endpoint.
type InvolvedCompanyService service

// Get returns a single InvolvedCompany identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any InvolvedCompanies, an error is returned.
func (is *InvolvedCompanyService) Get(id int, opts ...Option) (*InvolvedCompany, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var com []*InvolvedCompany

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := is.client.post(is.end, &com, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get InvolvedCompany with ID %v", id)
	}

	return com[0], nil
}

// List returns a list of InvolvedCompanies identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a InvolvedCompany is ignored. If none of the IDs
// match a InvolvedCompany, an error is returned.
func (is *InvolvedCompanyService) List(ids []int, opts ...Option) ([]*InvolvedCompany, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var com []*InvolvedCompany

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := is.client.post(is.end, &com, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get InvolvedCompanies with IDs %v", ids)
	}

	return com, nil
}

// Index returns an index of InvolvedCompanies based solely on the provided functional
// options used to sort, filter, and paginate the results. If no InvolvedCompanies can
// be found using the provided options, an error is returned.
func (is *InvolvedCompanyService) Index(opts ...Option) ([]*InvolvedCompany, error) {
	var com []*InvolvedCompany

	err := is.client.post(is.end, &com, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of InvolvedCompanies")
	}

	return com, nil
}

// Count returns the number of InvolvedCompanies available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which InvolvedCompanies to count.
func (is *InvolvedCompanyService) Count(opts ...Option) (int, error) {
	ct, err := is.client.getCount(is.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count InvolvedCompanies")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB InvolvedCompany object.
func (is *InvolvedCompanyService) Fields() ([]string, error) {
	f, err := is.client.getFields(is.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get InvolvedCompany fields")
	}

	return f, nil
}
