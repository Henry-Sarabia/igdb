package igdb

import (
	"errors"
	"reflect"
	"strings"
)

// validateStruct checks if the given struct contains all of the fields
// it should according to the appropriate IGDB endpoint.
func (c *Client) validateStruct(str reflect.Type, end endpoint) error {
	f, err := c.GetEndpointFields(end)
	if err != nil {
		return err
	}

	f = removeSubfields(f)

	err = validateStructTags(str, f)
	if err != nil {
		return err
	}
	return nil
}

// validateStructTags checks if the given struct contains all of
// the struct tags according to the slice of strings representing
// the appropriate struct tags.
func validateStructTags(str reflect.Type, new []string) error {
	old, err := getStructTags(str)
	if err != nil {
		return err
	}

	found := make(map[string]bool)
	for _, v := range new {
		found[v] = false
	}

	for _, v := range old {
		if _, ok := found[v]; ok {
			found[v] = true
		}
	}

	var missing []string
	for k, v := range found {
		if !v {
			missing = append(missing, k)
		}
	}

	if missing != nil {
		return errors.New("missing struct tags: " + strings.Join(missing, ", "))
	}

	return nil
}

// getStructTags collects the struct tags of every available
// field in the given struct.
func getStructTags(str reflect.Type) ([]string, error) {
	if str.Kind() != reflect.Struct {
		return nil, errors.New("input type's kind not a struct")
	}

	var f []string
	for i := 0; i < str.NumField(); i++ {
		f = append(f, str.Field(i).Tag.Get("json"))
	}
	return f, nil
}

// removeSubfields returns a slice of strings
// representing the collection of fields
// without any fields containing a period
// character.
func removeSubfields(f []string) []string {
	var out []string
	for _, val := range f {
		if !strings.Contains(val, ".") {
			out = append(out, val)
		}
	}
	return out
}
