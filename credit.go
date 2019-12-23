package igdb

import (
	"github.com/Henry-Sarabia/sliceconv"
	"github.com/pkg/errors"
	"strconv"
)

//go:generate gomodifytags -file $GOFILE -struct Credit -add-tags json -w

// Credit represents an employee responsible for working on a particular game.
// For more information visit: https://api-docs.igdb.com/#credit
type Credit struct {
	ID                    int            `json:"id,omitempty"`
	Category              CreditCategory `json:"category,omitempty"`
	Character             int            `json:"character,omitempty"`
	CharacterCreditedName string         `json:"character_credited_name,omitempty"`
	Comment               string         `json:"comment,omitempty"`
	Company               int            `json:"company,omitempty"`
	Country               int            `json:"country,omitempty"`
	CreatedAt             int            `json:"created_at,omitempty"`
	CreditedName          string         `json:"credited_name,omitempty"`
	Game                  int            `json:"game,omitempty"`
	Person                int            `json:"person,omitempty"`
	PersonTitle           int            `json:"person_title,omitempty"`
	Position              int            `json:"position,omitempty"`
	UpdatedAt             int            `json:"updated_at,omitempty"`
}

// CreditCategory specifies a specific job or role within a company.
type CreditCategory int

//go:generate stringer -type=CreditCategory

// Expected CreditCategory enums from the IGDB.
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
