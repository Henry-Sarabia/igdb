package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct TimeToBeat -add-tags json -w

// TimeToBeat represents the average completion times for a particular game.
// For more information: https://api-docs.igdb.com/#time-to-beat
type TimeToBeat struct {
	ID         int `json:"id"`
	Completely int `json:"completely"`
	Game       int `json:"game"`
	Hastly     int `json:"hastly"`
	Normally   int `json:"normally"`
}

// TimeToBeatService handles all the API calls for the IGDB TimeToBeat endpoint.
type TimeToBeatService service

// Get returns a single TimeToBeat identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any TimeToBeats, an error is returned.
func (ts *TimeToBeatService) Get(id int, opts ...Option) (*TimeToBeat, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var time []*TimeToBeat

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := ts.client.get(ts.end, &time, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get TimeToBeat with ID %v", id)
	}

	return time[0], nil
}

// List returns a list of TimeToBeats identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a TimeToBeat is ignored. If none of the IDs
// match a TimeToBeat, an error is returned.
func (ts *TimeToBeatService) List(ids []int, opts ...Option) ([]*TimeToBeat, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var time []*TimeToBeat

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := ts.client.get(ts.end, &time, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get TimeToBeats with IDs %v", ids)
	}

	return time, nil
}

// Index returns an index of TimeToBeats based solely on the provided functional
// options used to sort, filter, and paginate the results. If no TimeToBeats can
// be found using the provided options, an error is returned.
func (ts *TimeToBeatService) Index(opts ...Option) ([]*TimeToBeat, error) {
	var time []*TimeToBeat

	err := ts.client.get(ts.end, &time, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of TimeToBeats")
	}

	return time, nil
}

// Count returns the number of TimeToBeats available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which TimeToBeats to count.
func (ts *TimeToBeatService) Count(opts ...Option) (int, error) {
	ct, err := ts.client.getCount(ts.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count TimeToBeats")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB TimeToBeat object.
func (ts *TimeToBeatService) Fields() ([]string, error) {
	f, err := ts.client.getFields(ts.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get TimeToBeat fields")
	}

	return f, nil
}
