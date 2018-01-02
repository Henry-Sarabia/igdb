package igdb

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Errors returned by a FuncOption when setting options for an API call.
var (
	// ErrOptionSet occurs when the same option is used multiple times in a single API call.
	ErrOptionSet = errors.New("igdb.FuncOption: option already set")
	// ErrEmptyField occurs when an empty string is used as a field name.
	ErrEmptyField = errors.New("igdb.FuncOption: field empty")
	// ErrEmptySlice occurs when an empty slice is used as an argument in a variadic function.
	ErrEmptySlice = errors.New("igdb.FuncOption: slice empty")
	// ErrEmptyQuery occurs when an empty string is used as a query value.
	ErrEmptyQuery = errors.New("igdb.FuncOption: query value empty")
	// ErrEmptyFilterValue occurs when an empty string is used as a filter value.
	ErrEmptyFilterValue = errors.New("igdb.FuncOption: filter value empty")
	// ErrOutOfRange occurs when a provided number value is out of valid range.
	ErrOutOfRange = errors.New("igdb.FuncOption: value out of range")
	// ErrTooManyArgs occurs when too many arguments are provided in a variadic function.
	ErrTooManyArgs = errors.New("igdb.FuncOption: too many arguments")
)

// options contains a value map that stores the optional parameters for
// the various IGDB API calls. The options type is not accessed directly,
// but instead mutated using the functional options that return a FuncOption.
type options struct {
	Values url.Values
}

// FuncOption functions are used to set the options for an API call.
// FuncOption is the first-order function returned by the available
// functional options (e.g. SetLimit or SetFilter). This first-order
// function is then passed into a service's Get, List, Search, or
// Count function.
//
// Only one of each type of FuncOption can be set per API call.
// SetFilter is the only exception to this rule.
type FuncOption func(*options) error

// newOpt returns a new options object mutated by the provided FuncOption
// arguments. If no FuncOption's are provided, an empty options object is
// returned.
func newOpt(funcOpts ...FuncOption) (*options, error) {
	opt := &options{Values: url.Values{}}

	for _, funcOpt := range funcOpts {
		err := funcOpt(opt)
		if err != nil {
			return nil, err
		}
	}

	return opt, nil
}

// ComposeOptions composes multiple functional options into a single FuncOption.
// This is primarily used to conveniently create a single functional option that
// can be used repeatedly across different API calls.
func ComposeOptions(opts ...FuncOption) FuncOption {
	return func(o *options) error {
		for _, opt := range opts {
			if err := opt(o); err != nil {
				return err
			}
		}
		return nil
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
	OrderAscending order = ":asc"
	// OrderDescending is used as an argument in the SetOrder functional
	// option to organize the results from an API call in descending order.
	OrderDescending order = ":desc"
)

// subfilter specifies which type of filter to apply to the associated IGDB
// object's array field when using SetOrder's optional subfiltering
// functionality.
type subfilter string

// Available subfilters for the functional option SetOrder
const (
	// SubMax filters based on the maximum element in the array.
	SubMax subfilter = ":max"

	// SubMin filters based on the minimum element in the array.
	SubMin subfilter = ":min"

	// SubSum filters based on the sum of the elements in the array.
	SubSum subfilter = ":sum"

	// SubAverage filters based on the average of the elements in the array.
	SubAverage subfilter = ":avg"

	// SubMedian filters based on the median element in the array.
	SubMedian subfilter = ":median"
)

// SetOrder is a functional option used to set the order of the results from
// an API call, either ascending or descending. The provided field and order
// specify which field to sort by and in what order, respectively. Subfields
// are accessed with a dot operator. Note that the field string must match an
// IGDB object's JSON field tag exactly, not the Go struct field name. The
// default order is based on relevance and cannot be explicitly set.
//
// Optionally, SetOrder also allows you to provide a subfilter argument with
// which to perform array subfiltering on any of an IGBD object's array fields
// (e.g. a Game object's ReleaseDates field). In other words, you can order
// based on the max, min, sum, average, or median value of an array field's
// contents. If more than one subfilter is provided, an error is returned.
//
// For more information, visit: https://igdb.github.io/api/references/ordering/
func SetOrder(field string, ord order, sub ...subfilter) FuncOption {
	return func(o *options) error {
		if strings.TrimSpace(field) == "" {
			return ErrEmptyField
		}
		if len(sub) > 1 {
			return ErrTooManyArgs
		}
		if o.Values.Get("order") != "" {
			return ErrOptionSet
		}

		if len(sub) == 0 {
			o.Values.Set("order", field+string(ord))
			return nil
		}

		o.Values.Set("order", field+string(ord)+string(sub[0]))
		return nil
	}
}

// SetLimit is a functional option used to limit the number of results from
// an API call. The default limit is 10. The maximum limit is 50.
//
// For more information, visit: https://igdb.github.io/api/references/pagination/
func SetLimit(lim int) FuncOption {
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

// SetOffset is a functional option used to offset the results from an API
// call. The default offset is 0. The maximum offset is 10,000.
//
// For more information, visit: https://igdb.github.io/api/references/pagination/
func SetOffset(off int) FuncOption {
	return func(o *options) error {
		if off < 0 || off > 10000 {
			return ErrOutOfRange
		}
		if o.Values.Get("offset") != "" {
			return ErrOptionSet
		}
		o.Values.Set("offset", strconv.Itoa(off))
		return nil
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
// For more information, visit: https://igdb.github.io/api/references/fields/
func SetFields(fields ...string) FuncOption {
	return func(o *options) error {
		if len(fields) == 0 {
			return ErrEmptySlice
		}
		if hasBlankElem(fields) {
			return ErrEmptyField
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

// Available operators for the functional option SetFilter
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
	// OpIn checks if the field contains all of the given values. Can provide multiple values.
	OpIn operator = "in"
	// OpNotIn checks if the field does not contain any of the given values. Can provide multiple values.
	OpNotIn operator = "not_in"
	// OpAny checks if the field contains any of the given values. Can provide multiple values.
	OpAny operator = "any"
)

// SetFilter is a functional option used to filter the results from an API
// call. Most filtering operations need three different arguments: an operator
// and 2 operands. The provided field and val string act as the operands
// for the provided operator. If multiple values are provided, they will be
// concatenated into a comma separated list. If no values are provided, an
// error is returned. The only exception to this rule is if one of the provided
// operators is OpExists or OpNotExists, which do not need a provided value.

// SetFilter is the only option allowed to be set multiple times in a single
// API call. By default, results are unfiltered.
//
// Note that the ID field cannot be used for filtering except when paired with
// the OpNotIn operator to filter away specific results. Also note that when
// filtering a field that consists of an enumerated type (e.g. Gender Code,
// Feed Category, Game Status, etc.), you must provide the number corresponding
// to the intended field value.
//
// For more information, visit: https://igdb.github.io/api/references/filters/
func SetFilter(field string, op operator, val ...string) FuncOption {
	switch op {
	case OpExists, OpNotExists:
		return zeroValueFilter(field, op, val...)
	case OpIn, OpNotIn, OpAny:
		return multiValueFilter(field, op, val...)
	default:
		return singleValueFilter(field, op, val...)
	}
}

// zeroValueFilter handles the SetFilter functional option when the filter does
// not need a value. A placeholder value, "1", will be used instead.
func zeroValueFilter(field string, op operator, val ...string) FuncOption {
	return func(o *options) error {
		if field == "" {
			return ErrEmptyField
		}
		if len(val) > 0 {
			return ErrTooManyArgs
		}

		o.setFilterValue(field, op, "1") // 1 is just a placeholder for the required field value

		return nil
	}
}

// multiValueFilter handles the SetFilter functional option when the filter
// accepts multiple values. The values will be concatenated into a comma
// separated string.
func multiValueFilter(field string, op operator, val ...string) FuncOption {
	return func(o *options) error {
		if field == "" {
			return ErrEmptyField
		}
		if len(val) == 0 || hasBlankElem(val) {
			return ErrEmptyFilterValue
		}

		joined := strings.Join(val, ",")
		o.setFilterValue(field, op, joined)

		return nil
	}
}

// singleValueFilter handles the SetFilter functional option when the filter only
// accepts a single value. The value will be used as is.
func singleValueFilter(field string, op operator, val ...string) FuncOption {
	return func(o *options) error {
		if field == "" {
			return ErrEmptyField
		}
		if len(val) == 0 || hasBlankElem(val) {
			return ErrEmptyFilterValue
		}
		if len(val) > 1 {
			return ErrTooManyArgs
		}

		o.setFilterValue(field, op, val[0])

		return nil
	}
}

// setFilterValue sets the filter value in the options object.
func (o *options) setFilterValue(field string, op operator, val string) {
	s := fmt.Sprintf("filter[%s][%s]", field, string(op))
	o.Values.Set(s, val)
}

// setSearch is a functional option used to search the IGDB using the
// provided query.
func setSearch(qry string) FuncOption {
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
