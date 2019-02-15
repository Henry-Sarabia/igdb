package igdb

import (
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct PlatformVersion -add-tags json -w

// PlatformVersion represents a particular version of a platform.
// For more information visit: https://api-docs.igdb.com/#platform-version
type PlatformVersion struct {
	ID                          int    `json:"id"`
	Companies                   []int  `json:"companies"`
	Connectivity                string `json:"connectivity"`
	CPU                         string `json:"cpu"`
	Graphics                    string `json:"graphics"`
	MainManufacturer            int    `json:"main_manufacturer"`
	Media                       string `json:"media"`
	Memory                      string `json:"memory"`
	Name                        string `json:"name"`
	OS                          string `json:"os"`
	Output                      string `json:"output"`
	PlatformLogo                int    `json:"platform_logo"`
	PlatformVersionReleaseDates []int  `json:"platform_version_release_dates"`
	Resolutions                 string `json:"resolutions"`
	Slug                        string `json:"slug"`
	Sound                       string `json:"sound"`
	Storage                     string `json:"storage"`
	Summary                     string `json:"summary"`
	URL                         string `json:"url"`
}

// PlatformVersionService handles all the API calls for the IGDB PlatformVersion endpoint.
type PlatformVersionService service

// Get returns a single PlatformVersion identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any PlatformVersions, an error is returned.
func (ps *PlatformVersionService) Get(id int, opts ...Option) (*PlatformVersion, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var ver []*PlatformVersion

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ps.client.get(ps.end, &ver, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlatformVersion with ID %v", id)
	}

	return ver[0], nil
}

// List returns a list of PlatformVersions identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a PlatformVersion is ignored. If none of the IDs
// match a PlatformVersion, an error is returned.
func (ps *PlatformVersionService) List(ids []int, opts ...Option) ([]*PlatformVersion, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var ver []*PlatformVersion

	opts = append(opts, SetFilter("id", OpContainsAtLeast, intsToStrings(ids)...))
	err := ps.client.get(ps.end, &ver, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get PlatformVersions with IDs %v", ids)
	}

	return ver, nil
}

// Index returns an index of PlatformVersions based solely on the provided functional
// options used to sort, filter, and paginate the results. If no PlatformVersions can
// be found using the provided options, an error is returned.
func (ps *PlatformVersionService) Index(opts ...Option) ([]*PlatformVersion, error) {
	var ver []*PlatformVersion

	err := ps.client.get(ps.end, &ver, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of PlatformVersions")
	}

	return ver, nil
}

// Count returns the number of PlatformVersions available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which PlatformVersions to count.
func (ps *PlatformVersionService) Count(opts ...Option) (int, error) {
	ct, err := ps.client.getCount(ps.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count PlatformVersions")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB PlatformVersion object.
func (ps *PlatformVersionService) Fields() ([]string, error) {
	f, err := ps.client.getFields(ps.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get PlatformVersion fields")
	}

	return f, nil
}
