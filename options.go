package igdb

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Errors returned by an OptionFunc.
var (
	ErrOptionSet       = errors.New("igdb.OptionFunc: option already set")
	ErrEmptyField      = errors.New("igdb.OptionFunc: field value empty")
	ErrEmptySlice      = errors.New("igdb.OptionFunc: slice empty")
	ErrOutOfRange      = errors.New("igdb.OptionFunc: value out of range")
	ErrExclusiveOption = errors.New("igdb.OptionFunc: multiple set options are mutually exclusive")
	ErrEmptyQuery      = errors.New("igdb.OptionFunc: query value empty")
)

// options contains a value map to store optional
// parameters for the various IGDB API calls.
// The options type is not accessed directly,
// it is manipulated using the OptionFunc type.
type options struct {
	Values url.Values
}

// OptionFunc is the first-order function returned
// by the available functional options (e.g. OptLimit
// or OptFilter). OptionFunc is used to set the
// options for an API call.
type OptionFunc func(*options) error

// newOpt returns a new options object mutated by
// the provided OptionFunc arguments. Only one of
// each type of OptionFunc can be set per API call.
// OptFilter is the only exception to this rule.
func newOpt(ofs ...OptionFunc) (*options, error) {
	opt := &options{Values: url.Values{}}

	for _, of := range ofs {
		err := of(opt)
		if err != nil {
			return nil, err
		}
	}

	return opt, nil
}

// order specifies the order in which to place
// the results from an API call. The two
// available orders are Ascending and Descending.
type order string

const (
	// OrderAscending is used as an argument in the SetOrder functional
	// option to set the results from an API call in ascending order.
	OrderAscending order = ":asc"
	// OrderDescending is used as an argument in the SetOrder functional
	// option to set the results from an API call in descending order.
	OrderDescending order = ":desc"
)

// OptOrder is a functional option used to set the
// order of the results from an API call. The default
// ordering is based on relevance.
func OptOrder(field string, ord order) OptionFunc {
	return func(o *options) error {
		if len(field) == 0 {
			return ErrEmptyField
		}
		if o.Values.Get("order") != "" {
			return ErrOptionSet
		}
		o.Values.Set("order", field+string(ord))
		return nil
	}
}

// OptLimit is a functional option used to limit
// the number of results from an API call. The
// default limit is 10. The maximum limit is 50.
func OptLimit(lim int) OptionFunc {
	return func(o *options) error {
		if lim <= 0 || lim > 50 {
			return ErrOutOfRange
		}
		if o.Values.Get("limit") != "" {
			return ErrOptionSet
		}
		o.Values.Set("limit", strconv.Itoa(lim))
		return nil
	}
}

// OptOffset is a functional option used to offset
// results from an API call. The default offset is
// 0. The maximum offset is 10,000. For results with
// more than 10,000 objects, use the Scroll option.
// OptOffset and OptScroll are mutually exclusive
// options (i.e. only one can be used per API call).
func OptOffset(off int) OptionFunc {
	return func(o *options) error {
		if off < 0 || off > 50 {
			return ErrOutOfRange
		}
		if o.Values.Get("scroll") != "" {
			return ErrExclusiveOption
		}
		if o.Values.Get("offset") != "" {
			return ErrOptionSet
		}
		o.Values.Set("offset", strconv.Itoa(off))
		return nil
	}
}

// OptFields is a functional option used to
// specify which fields of the requested IGDB
// object you wnat the API to respond with.
// Subfields are accessed with a dot operator.
// The default is set to all available fields.
// All fields can be also be accessed with a
// single asterisk.
func OptFields(fields ...string) OptionFunc {
	return func(o *options) error {
		if len(fields) == 0 {
			return ErrEmptySlice
		}
		for _, f := range fields {
			if f == "" {
				return ErrEmptyField
			}
		}
		if o.Values.Get("fields") != "" {
			return ErrOptionSet
		}
		fs := strings.Join(fields, ",")
		o.Values.Set("fields", fs)
		return nil
	}
}

// operator is a postfix operator used in
// the filter option to filter the results
// based on the provided field values. For
// the list of postfix operators, visit:
// https://igdb.github.io/api/references/filters/#available-postfixes
type operator string

const (
	// OpEquals checks for equality. Must match exactly.
	OpEquals operator = "eq"
	// OpNotEquals checks for inequality. Any non-exact match.
	OpNotEquals operator = "not_eq"
	// OpGreaterThan checks if a field value is greater than a given value. Only works on numbers.
	OpGreaterThan operator = "gt"
	// OpGreaterThanEqual checks if a field value is greater than or equal to a given value. Only works on numbers.
	OpGreaterThanEqual operator = "gte"
	// OpLessThan checks if a field value is less than a given value. Only works on numbers.
	OpLessThan operator = "lt"
	// OpLessThanEqual checks if a field value is less than or equal to a given value. Only works on numbers.
	OpLessThanEqual operator = "lte"
	// OpPrefix checks if a field value contains the given prefix value. Only works on strings.
	OpPrefix operator = "prefix"
	// OpExists checks if a field value is a non-null value. Does not need a provided value.
	OpExists operator = "exists"
	// OpNotExists checks if a field value is a null value. Does not need a provided value.
	OpNotExists operator = "not_exists"
	// OpIn checks if the field contains all of the given comma separated values.
	OpIn operator = "in"
	// OpNotIn checks if the field does not contain any of the given comma separated values.
	OpNotIn operator = "not_in"
	// OpAny checks if the field contains any of the given comma separated values.
	OpAny operator = "any"
)

// OptFilter is a functional option used to filter the results from
// an API call. The provided field name specifies which field to
// operate on with the provided operator. The given string value
// represents the value of the field being operated on. OptFilter
// is the only option allowed to be set multiple times in a single
// API call. Note, the ID field cannot be used in the filter option
// except when paired with the OpNotIn operator to filter away
// specific results.
func OptFilter(field string, op operator, val string) OptionFunc {
	return func(o *options) error {
		if op == OpExists || op == OpNotExists {
			val = "1"
		}
		if field == "" || val == "" {
			return ErrEmptyField
		}
		s := fmt.Sprintf("filter[%s][%s]", field, string(op))
		o.Values.Set(s, val)
		return nil
	}
}

// OptScroll is a functional option used
// to paginate the results of an API call
// using IGDB's Scroll API. The given
// integer denotes which page of results
// to retrieve from the API call. When not
// included in an API call, this option has
// no default value. This option and OptScroll
// are mutually exclusive options; only one can
// be used per API call.
// WORK IN PROGRESS
func OptScroll(page int) OptionFunc {
	return func(o *options) error {
		if page < 1 {
			return ErrOutOfRange
		}
		if o.Values.Get("offset") != "" {
			return ErrExclusiveOption
		}
		if o.Values.Get("scroll") != "" {
			return ErrOptionSet
		}
		o.Values.Set("scroll", strconv.Itoa(page))
		return nil
	}
}

// OptScrollNew is a functional option used to
// paginate the results of an API call using
// the Scroll API. If OptScrollNew is set,
// the API call will return two additional
// headers that will be stored in the Client.
// One header contains the number of results
// found for the API call and the second header
// contains a special Scroll path that can be
// queried to get the next page of results.
// This path can be repeatedly queried to
// iterate through pages because it does not
// change. The path will expire after 3 minutes
// of not being queried.
func OptScrollNew() OptionFunc {
	return func(o *options) error {
		if o.Values.Get("offset") != "" {
			return ErrExclusiveOption
		}
		if o.Values.Get("scroll") != "" {
			return ErrOptionSet
		}

		o.Values.Set("scroll", "1")
		return nil
	}
}

// optSearch is an unexported functional
// option used to search the IGDB for the
// given query.
func optSearch(qry string) OptionFunc {
	return func(o *options) error {
		if qry == "" {
			return ErrEmptyQuery
		}
		if o.Values.Get("search") != "" {
			return ErrOptionSet
		}
		o.Values.Set("search", qry)
		return nil
	}
}
