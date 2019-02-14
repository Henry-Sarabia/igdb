package igdb

import (
	"fmt"
	"github.com/Henry-Sarabia/apicalypse"
	"github.com/Henry-Sarabia/whitespace"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

// Errors returned by a FuncOption when setting options for an API call.
var (
	// ErrEmptyQuery occurs when an empty string is used as a query value.
	ErrEmptyQuery = errors.New("igdb.FuncOption: provided query value is empty")
	// ErrEmptyFields occurs when an empty string is used as a field value.
	ErrEmptyFields = errors.New("igdb.FuncOption: one or more provided field values are empty")
	// ErrEmptyFilterValues occurs when an empty string is used as a filter value.
	ErrEmptyFilterValues = errors.New("igdb.FuncOption: one or more provided filter values are empty")
	// ErrOutOfRange occurs when a provided number value is out of valid range.
	ErrOutOfRange = errors.New("igdb.FuncOption: provided value is out of range")
)

// FuncOption functions are used to set the options for an API call.
// FuncOption is the first-order function returned by the available
// functional options (e.g. SetLimit or SetFilter). This first-order
// function is then passed into a service's Get, List, Index, Search, or
// Count function.
type FuncOption func() (apicalypse.Option, error)

func unwrapOptions(opts ...FuncOption) ([]apicalypse.Option, error) {
	unwrapped := make([]apicalypse.Option, len(opts))
	for i, opt := range opts {
		var err error
		if unwrapped[i], err = opt(); err != nil {
			return nil, errors.Wrap(err, "cannot unwrap invalid option")
		}
	}

	return unwrapped, nil
}

// ComposeOptions composes multiple functional options into a single FuncOption.
// This is primarily used to create a single functional option that can be used
// repeatedly across multiple queries.
func ComposeOptions(funcOpts ...FuncOption) FuncOption {
	return func() (apicalypse.Option, error) {
		unwrapped, err := unwrapOptions(funcOpts...)
		if err != nil {
			return nil, errors.Wrap(err, "cannot compose invalid functional options")
		}

		return apicalypse.ComposeOptions(unwrapped...), nil
	}
}

// order specifies the order in which to organize the results from an API call.
// There are three orders in which results are organized: relevance, ascending,
// and descending. Relevance is only available as a default and cannot be
// explicitly specified.
type order string

// Available orders for the functional option SetOrder
const (
	// OrderAscending is used as an argument in the SetOrder functional
	// option to organize the results from an API call in ascending order.
	OrderAscending order = "asc"
	// OrderDescending is used as an argument in the SetOrder functional
	// option to organize the results from an API call in descending order.
	OrderDescending order = "desc"
)

func SetOrder(field string, order order) FuncOption {
	return func() (apicalypse.Option, error) {
		if whitespace.IsBlank(field) {
			return nil, ErrEmptyFields
		}
		return apicalypse.Sort(field, string(order)), nil
	}
}

// SetLimit is a functional option used to limit the number of results from
// an API call. The default limit is 10. The maximum limit is 50.
//
// For more information, visit: https://api-docs.igdb.com/#pagination
func SetLimit(lim int) FuncOption {
	return func() (apicalypse.Option, error) {
		if lim <= 0 || lim > 50 {
			return nil, ErrOutOfRange
		}

		return apicalypse.Limit(lim), nil
	}
}

// SetOffset is a functional option used to offset the results from an API
// call. The default offset is 0. The maximum offset is 10,000.
//
// For more information, visit: https://api-docs.igdb.com/#pagination
func SetOffset(off int) FuncOption {
	return func() (apicalypse.Option, error) {
		if off < 0 || off > 10000 {
			return nil, ErrOutOfRange
		}

		return apicalypse.Offset(off), nil
	}
}

// SetFields is a functional option used to specify which fields of the
// requested IGDB object you want the API to provide. Subfields are accessed
// with a dot operator (e.g. cover.url). To select all available fields at
// once, use an asterisk character (i.e. *). Note that the field string must
// match an IGDB object's JSON field tag exactly, not the Go struct field
// name.
//
// The default for Get and List functions is set to all available fields.
// The default for Search functions is set to solely the ID field.
//
// For more information, visit: https://api-docs.igdb.com/#fields
func SetFields(fields ...string) FuncOption {
	return func() (apicalypse.Option, error) {
		if len(fields) <= 0 {
			return nil, ErrEmptyFields
		}
		for _, f := range fields {
			if whitespace.IsBlank(f) {
				return nil, ErrEmptyFields
			}
		}
		return apicalypse.Fields(fields...), nil
	}
}

// operator represents the postfix operation used to filter the results from
// an API call using the provided field value. For the list of postfix
// operators, visit: https://api-docs.igdb.com/#filters
type operator string

// Available operators for the functional option SetFilter
const (
	// OpEquals checks for equality. Must match exactly.
	OpEquals operator = "%s = %s"
	// OpNotEquals checks for inequality. Any non-exact match.
	OpNotEquals operator = "%s != %s"
	// OpGreaterThan checks if a field value is greater than a given value. Only works on numbers.
	OpGreaterThan operator = "%s > %s"
	// OpGreaterThanEqual checks if a field value is greater than or equal to a given value. Only works on numbers.
	OpGreaterThanEqual operator = "%s >= %s"
	// OpLessThan checks if a field value is less than a given value. Only works on numbers.
	OpLessThan operator = "%s < %s"
	// OpLessThanEqual checks if a field value is less than or equal to a given value. Only works on numbers.
	OpLessThanEqual operator = "%s <= %s"
	// OpContainsAll checks if the given value exists in within the array.
	OpContainsAll operator = "%s = [%s]"
	// OpNotContainsAll checks if the given value does not exist in within the array.
	OpNotContainsAll operator = "%s != [%s]"
	// OpContainsAtLeast checks if any of the given values exist within the array.
	OpContainsAtLeast operator = "%s = (%s)"
	// OpNotContainsAtLeast checks if any of the given values do not exist within the array.
	OpNotContainsAtLeast operator = "%s != (%s)"
	// OpContainsExactly checks if the the given values exactly match the array.
	OpContainsExactly operator = "%s = {%s}"
)

// SetFilter is a functional option used to filter the results from an API
// call. Filtering operations need three different arguments: an operator
// and 2 operands, the field and its value. The provided field and val string
// act as the operands for the provided operator. If multiple values are provided,
//they will be concatenated into a comma separated list. If no values are
//provided, an error is returned.
//
//SetFilter is the only option allowed to be set multiple times in a single
//API call. By default, results are unfiltered.
//
//Note that when filtering a field that consists of an enumerated type (e.g. Gender Code,
// Feed Category, Game Status, etc.), you must provide the number corresponding
// to the intended field value.
//
// For more information, visit: https://api-docs.igdb.com/#filters
func SetFilter(field string, op operator, val ...string) FuncOption {
	return func() (apicalypse.Option, error) {
		if isBlank(field) {
			return nil, ErrEmptyFields
		}
		if len(val) <= 0 || hasBlankElem(val) {
			return nil, ErrEmptyFilterValues
		}

		j := strings.Join(val, ",")
		return apicalypse.Where(fmt.Sprintf(string(op), field, j)), nil
	}
}

// setSearch is a functional option used to search the IGDB using the
// provided query.
func setSearch(qry string) FuncOption {
	return func() (apicalypse.Option, error) {
		if isBlank(qry) {
			return nil, ErrEmptyQuery
		}
		return apicalypse.Search("", qry), nil
	}
}

// hasBlankElem returns true if the slice of strings contains a blank string
// element, either an empty string or a string of space characters. Otherwise,
// return false.
func hasBlankElem(s []string) bool {
	for _, v := range s {
		if strings.TrimSpace(v) == "" {
			return true
		}
	}
	return false
}

// isBlank returns true if the provided string is empty or only consists of whitespace.
// Returns false otherwise.
func isBlank(s string) bool {
	if removeWhitespace(s) == "" {
		return true
	}

	return false
}

// removeWhitespace returns the provided string with all of the whitespace removed.
// This includes spaces, tabs, newlines, returns, and form feeds.
func removeWhitespace(s string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(s, "")
}
