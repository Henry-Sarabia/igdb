package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct PlatformVersionCompany -add-tags json -w

// PlatformVersionCompany represents a platform developer.
// For more information visit: https://api-docs.igdb.com/#platform-version-company
type PlatformVersionCompany struct {
	ID           int    `json:"id,omitempty"`
	Comment      string `json:"comment,omitempty"`
	Company      int    `json:"company,omitempty"`
	Developer    bool   `json:"developer,omitempty"`
	Manufacturer bool   `json:"manufacturer,omitempty"`
}

// PlatformVersionCompanyService handles all the API calls for the IGDB PlatformVersionCompany endpoint.
type PlatformVersionCompanyService service

// Get returns a single PlatformVersionCompany identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PlatformVersionCompanies, an error is returned.
func (ps *PlatformVersionCompanyService) Get(id int, opts ...Option) (*PlatformVersionCompany, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var com []*PlatformVersionCompany

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &com, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlatformVersionCompany with ID %v", id)
	}

	return com[0], nil
}

// List returns a list of PlatformVersionCompanies identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PlatformVersionCompany is ignored. If none of the IDs
// match a PlatformVersionCompany, an error is returned.
func (ps *PlatformVersionCompanyService) List(ids []int, opts ...Option) ([]*PlatformVersionCompany, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var com []*PlatformVersionCompany

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ps.client.get(ps.end, &com, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlatformVersionCompanies with IDs %v", ids)
	}

	return com, nil
}

// Index returns an index of PlatformVersionCompanies based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PlatformVersionCompanies can
// be found using the provided options, an error is returned.
func (ps *PlatformVersionCompanyService) Index(opts ...Option) ([]*PlatformVersionCompany, error) {
	var com []*PlatformVersionCompany

	err := ps.client.get(ps.end, &com, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PlatformVersionCompanies")
	}

	return com, nil
}

// Count returns the number of PlatformVersionCompanies available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PlatformVersionCompanies to count.
func (ps *PlatformVersionCompanyService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PlatformVersionCompanies")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB PlatformVersionCompany object.
func (ps *PlatformVersionCompanyService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get PlatformVersionCompany fields")
	}

	return f, nil
}
