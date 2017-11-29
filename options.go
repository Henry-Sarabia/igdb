package igdb

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Errors returned by an OptionFunc when setting API call options.
var (
	// ErrOptionSet occurs when the same option is used multiple times in a single API call.
	ErrOptionSet = errors.New("igdb.OptionFunc: option already set")
	// ErrEmptyField occurs when an empty string is used as a field name.
	ErrEmptyField = errors.New("igdb.OptionFunc: field value empty")
	// ErrEmptySlice occurs when an empty slice is used as an argument in a variadic function.
	ErrEmptySlice = errors.New("igdb.OptionFunc: slice empty")
	// ErrOutOfRange occurs when a provided number value is out of valid range.
	ErrOutOfRange = errors.New("igdb.OptionFunc: value out of range")
	// ErrEmptyQuery occurs when an empty string is used as a query value.
	ErrEmptyQuery = errors.New("igdb.OptionFunc: query value empty")
)

// options contains a value map that stores the optional parameters for
// the various IGDB API calls. The options type is not accessed directly,
// but instead mutated using the functional options that return an OptionFunc.
type options struct {
	Values url.Values
}

// OptionFunc functions are used to set the options for an API call.
// OptionFunc is the first-order function returned by the available
// functional options (e.g. OptLimit or OptFilter). This first-order
// function is then passed into a service's Get, List, Search, or
// Count function.
//
// Only one of each type of OptionFunc can be set per API call.
// OptFilter is the only exception to this rule.
type OptionFunc func(*options) error

// newOpt returns a new options object mutated by the provided OptionFunc
// arguments. If no OptionFunc's are provided, an empty options object is
// returned.
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

// order specifies the order in which to organize the results from an API call.
// There are three orders in which results are organized: relevance, ascending,
// and descending. Relevance is only available as a default and cannot be
// explicitly specified.
type order string

// The available orders for the functional option OptOrder.
const (
	// OrderAscending is used as an argument in the OptOrder functional
	// option to organize the results from an API call in ascending order.
	OrderAscending order = ":asc"
	// OrderDescending is used as an argument in the OptOrder functional
	// option to organize the results from an API call in descending order.
	OrderDescending order = ":desc"
)

// OptOrder is a functional option used to set the order of the results from
// an API call, either ascending or descending. The default order is based on
// relevance and cannot be explicitly set.
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

// OptLimit is a functional option used to limit the number of results from
// an API call. The default limit is 10. The maximum limit is 50.
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

// OptOffset is a functional option used to offset the results from an API
// call. The default offset is 0. The maximum offset is 10,000.
func OptOffset(off int) OptionFunc {
	return func(o *options) error {
		if off < 0 || off > 50 {
			return ErrOutOfRange
		}
		if o.Values.Get("offset") != "" {
			return ErrOptionSet
		}
		o.Values.Set("offset", strconv.Itoa(off))
		return nil
	}
}

// OptFields is a functional option used to specify which fields of the
// requested IGDB object you want the API to provide. Subfields are accessed
// with a dot operator (e.g. cover.url). To select all available fields at
// once, use an asterisk character (i.e. *).
//
// The default for Get and List functions is set to all available fields.
// The default for Search functions is set to solely the ID field.
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

// operator represents the postfix operation used to filter the results from
// an API call using the provided field value. For the list of postfix
// operators, visit:
// https://igdb.github.io/api/references/filters/#available-postfixes
type operator string

// The available operators for the functional option OptFilter.
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

// OptFilter is a functional option used to filter the results from an API
// call. Most filtering operations need three different arguments: an operator
// and 2 operands. The provided field and val strings act as the operands
// for the provided operator. Optfilter is the only option allowed to be set
// multiple times in a single API call. By default, results are unfiltered.
//
// Note that the ID field cannot be used for filtering except when paired with
// the OpNotIn operator to filter away specific results. Also note that when
// filtering a field that consists of an enumerated type (e.g. Gender Code,
// Feed Category, Game Status, etc.), you must provide the number corresponding
// to the intended field value.
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

// optSearch is a functional option used to search the IGDB using the
// provided query.
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
