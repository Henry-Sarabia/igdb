package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

// AchievementIcon is an icon for a specific achievement.
// For more information visit: https://api-docs.igdb.com/#achievement-icon
type AchievementIcon struct {
	Image
	ID int `json:"id"`
}

// AchievementIconService handles all the API calls for the IGDB
// AchievementIcon endpoint.
type AchievementIconService service

// Get returns a single AchievementIcon identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any AchievementIcons, an error is returned.
func (as *AchievementIconService) Get(id int, opts ...Option) (*AchievementIcon, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var icon []*AchievementIcon

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := as.client.get(as.end, &icon, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get AchievementIcon with ID %v", id)
	}

	return icon[0], nil
}

// List returns a list of AchievementIcons identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a AchievementIcon is ignored. If none of the IDs
// match a AchievementIcon, an error is returned.
func (as *AchievementIconService) List(ids []int, opts ...Option) ([]*AchievementIcon, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var icon []*AchievementIcon

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := as.client.get(as.end, &icon, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get AchievementIcons with IDs %v", ids)
	}

	return icon, nil
}

// Index returns an index of AchievementIcons based solely on the provided functional
// options used to sort, filter, and paginate the results. If no AchievementIcons can
// be found using the provided options, an error is returned.
func (as *AchievementIconService) Index(opts ...Option) ([]*AchievementIcon, error) {
	var icon []*AchievementIcon

	err := as.client.get(as.end, &icon, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of AchievementIcons")
	}

	return icon, nil
}

// Count returns the number of AchievementIcons available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which AchievementIcons to count.
func (as *AchievementIconService) Count(opts ...Option) (int, error) {
	ct, err := as.client.getCount(as.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count AchievementIcons")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB AchievementIcon object.
func (as *AchievementIconService) Fields() ([]string, error) {
	f, err := as.client.getFields(as.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get AchievementIcon fields")
	}

	return f, nil
}
