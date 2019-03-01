package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct PersonWebsite -add-tags json -w

// PersonWebsite represents a website associated
// with a person in the video game industry.
// For more information visit: https://api-docs.igdb.com/#person-website
type PersonWebsite struct {
	ID       int             `json:"id"`
	Category WebsiteCategory `json:"category"`
	Trusted  bool            `json:"trusted"`
	URL      string          `json:"url"`
}

// PersonWebsiteService handles all the API calls for the IGDB PersonWebsite endpoint.
type PersonWebsiteService service

// Get returns a single PersonWebsite identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PersonWebsites, an error is returned.
func (ps *PersonWebsiteService) Get(id int, opts ...Option) (*PersonWebsite, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var web []*PersonWebsite

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &web, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PersonWebsite with ID %v", id)
	}

	return web[0], nil
}

// List returns a list of PersonWebsites identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PersonWebsite is ignored. If none of the IDs
// match a PersonWebsite, an error is returned.
func (ps *PersonWebsiteService) List(ids []int, opts ...Option) ([]*PersonWebsite, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var web []*PersonWebsite

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ps.client.get(ps.end, &web, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PersonWebsites with IDs %v", ids)
	}

	return web, nil
}

// Index returns an index of PersonWebsites based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PersonWebsites can
// be found using the provided options, an error is returned.
func (ps *PersonWebsiteService) Index(opts ...Option) ([]*PersonWebsite, error) {
	var web []*PersonWebsite

	err := ps.client.get(ps.end, &web, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PersonWebsites")
	}

	return web, nil
}

// Count returns the number of PersonWebsites available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PersonWebsites to count.
func (ps *PersonWebsiteService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PersonWebsites")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB PersonWebsite object.
func (ps *PersonWebsiteService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get PersonWebsite fields")
	}

	return f, nil
}
