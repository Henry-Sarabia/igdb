package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct PlatformVersionCompany -add-tags json -w

// PlatformVersionCompany represents a platform developer.
// For more information visit: https://api-docs.igdb.com/#platform-version-company
type PlatformVersionCompany struct {
	Comment      string `json:"comment"`
	Company      int    `json:"company"`
	Developer    bool   `json:"developer"`
	Manufacturer bool   `json:"manufacturer"`
}

// PlatformVersionCompanyService handles all the API calls for the IGDB PlatformVersionCompany endpoint.
type PlatformVersionCompanyService service

// Get returns a single PlatformVersionCompany identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PlatformVersionCompanys, an error is returned.
func (ps *PlatformVersionCompanyService) Get(id int, opts ...FuncOption) (*PlatformVersionCompany, error) {
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

// List returns a list of PlatformVersionCompanys identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PlatformVersionCompany is ignored. If none of the IDs
// match a PlatformVersionCompany, an error is returned.
func (ps *PlatformVersionCompanyService) List(ids []int, opts ...FuncOption) ([]*PlatformVersionCompany, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var com []*PlatformVersionCompany

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := ps.client.get(ps.end, &com, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlatformVersionCompanys with IDs %v", ids)
	}

	return com, nil
}

// Index returns an index of PlatformVersionCompanys based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PlatformVersionCompanys can
// be found using the provided options, an error is returned.
func (ps *PlatformVersionCompanyService) Index(opts ...FuncOption) ([]*PlatformVersionCompany, error) {
	var com []*PlatformVersionCompany

	err := ps.client.get(ps.end, &com, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PlatformVersionCompanys")
	}

	return com, nil
}

// Count returns the number of PlatformVersionCompanys available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PlatformVersionCompanys to count.
func (ps *PlatformVersionCompanyService) Count(opts ...FuncOption) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PlatformVersionCompanys")
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
