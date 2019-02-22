package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Credit -add-tags json -w

type Credit struct {
	ID                    int
	Category              CreditCategory
	Character             int
	CharacterCreditedName string
	Comment               string
	Company               int
	Country               int
	CreatedAt             int
	CreditedName          string
	Game                  int
	Person                int
	PersonTitle           int
	Position              int
	UpdatedAt             int
}

type CreditCategory int

//go:generate stringer -type=CreditCategory

const (
	CreditVoiceActor CreditCategory = iota + 1
	CreditLanguage
	CreditCompany
	CreditEmployee
	CreditMisc
	CreditSupportCompany
)

// CreditService handles all the API calls for the IGDB Credit endpoint.
type CreditService service

// Get returns a single Credit identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to
// retrieve. If the ID does not match any Credits, an error is returned.
func (cs *CreditService) Get(id int, opts ...Option) (*Credit, error) {
	if id < 0 {
		return nil, ErrNegativeID
	}

	var cr []*Credit

	opts = append(opts, SetFilter("id", OpEquals, strconv.Itoa(id)))
	err := cs.client.get(cs.end, &cr, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Credit with ID %v", id)
	}

	return cr[0], nil
}

// List returns a list of Credits identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results.
// Any ID that does not match a Credit is ignored. If none of the IDs
// match a Credit, an error is returned.
func (cs *CreditService) List(ids []int, opts ...Option) ([]*Credit, error) {
	for len(ids) < 1 {
		return nil, ErrEmptyIDs
	}

	for _, id := range ids {
		if id < 0 {
			return nil, ErrNegativeID
		}
	}

	var cr []*Credit

	opts = append(opts, SetFilter("id", OpContainsAtLeast, sliceconv.Itoa(ids)...))
	err := cs.client.get(cs.end, &cr, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Credits with IDs %v", ids)
	}

	return cr, nil
}

// Index returns an index of Credits based solely on the provided functional
// options used to sort, filter, and paginate the results. If no Credits can
// be found using the provided options, an error is returned.
func (cs *CreditService) Index(opts ...Option) ([]*Credit, error) {
	var cr []*Credit

	err := cs.client.get(cs.end, &cr, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get index of Credits")
	}

	return cr, nil
}

// Search returns a list of Credits found by searching the IGDB using the provided
// query. Provide functional options to sort, filter, and paginate the results. If
// no Credits are found using the provided query, an error is returned.
func (cs *CreditService) Search(qry string, opts ...Option) ([]*Credit, error) {
	var cr []*Credit

	opts = append(opts, setSearch(qry))
	err := cs.client.get(cs.end, &cr, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get Credit with query %s", qry)
	}

	return cr, nil
}

// Count returns the number of Credits available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Credits to count.
func (cs *CreditService) Count(opts ...Option) (int, error) {
	ct, err := cs.client.getCount(cs.end, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "cannot count Credits")
	}

	return ct, nil
}

// Fields returns the up-to-date list of fields in an
// IGDB Credit object.
func (cs *CreditService) Fields() ([]string, error) {
	f, err := cs.client.getFields(cs.end)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Credit fields")
	}

	return f, nil
}
