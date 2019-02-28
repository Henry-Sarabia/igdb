package igdb

import (
	"strconv"

	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
)

//go:generate gomodifytags -file $GOFILE -struct ExternalGame -add-tags json -w

// ExternalGame contains the ID and other metadata for a game
// on a third party service.
// For more information visit: https://api-docs.igdb.com/#external-game
type ExternalGame struct {
	ID        int                  `json:"id"`
	Category  ExternalGameCategory `json:"category"`
	CreatedAt int                  `json:"created_at"`
	Game      int                  `json:"game"`
	Name      string               `json:"name"`
	UID       string               `json:"uid"`
	UpdatedAt int                  `json:"updated_at"`
	Url       string               `json:"url"`
	Year      int                  `json:"year"`
}

// ExternalGameCategory speficies an external game, platform, or media service.
type ExternalGameCategory int

//go:generate stringer -type=ExternalGameCategory

const (
	ExternalSteam ExternalGameCategory = iota + 1
	_
	_
	_
	ExternalGOG
	_
	_
	_
	_
	ExternalYoutube
	ExternalMicrosoft
	_
	ExternalApple
	ExternalTwitch
	ExternalAndroid
)

// ExternalGameService handles all the API calls for the IGDB ExternalGame endpoint.
// This endpoint is only available for the IGDB Pro tier or above.
type ExternalGameService service

// Get returns a single ExternalGame identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any ExternalGames, an error is returned.
func (es *ExternalGameService) Get(id int, opts ...Option) (*ExternalGame, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var ext []*ExternalGame

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := es.client.get(es.end, &ext, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get ExternalGame with ID %v", id)
	}

	return ext[0], nil
}

// List returns a list of ExternalGames identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a ExternalGame is ignored. If none of the IDs
// match a ExternalGame, an error is returned.
func (es *ExternalGameService) List(ids []int, opts ...Option) ([]*ExternalGame, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var ext []*ExternalGame

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := es.client.get(es.end, &ext, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get ExternalGames with IDs %v", ids)
	}

	return ext, nil
}

// Index returns an index of ExternalGames based solely on the provided functional
// options used to sort, filter, and paginate the results. If no ExternalGames can
// be found using the provided options, an error is returned.
func (es *ExternalGameService) Index(opts ...Option) ([]*ExternalGame, error) {
	var ext []*ExternalGame

	err := es.client.get(es.end, &ext, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of ExternalGames")
	}

	return ext, nil
}

// Count returns the number of ExternalGames available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which ExternalGames to count.
func (es *ExternalGameService) Count(opts ...Option) (int, error) {
	ct, err := es.client.getCount(es.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count ExternalGames")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB ExternalGame object.
func (es *ExternalGameService) Fields() ([]string, error) {
	f, err := es.client.getFields(es.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get ExternalGame fields")
	}

	return f, nil
}
