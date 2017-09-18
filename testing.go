package igdb

import (
	"errors"
	"reflect"
	"strings"
)

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
