package map

import (
	"errors"
)

func ReadMapValue(m map[string]interface{}, key string) (interface{}, error) {
	if m == nil {
		return errors.New("nil map")
	}

	if value, ok := m[key]; !ok {
		return errors.New("key not exist")
	}

	return m[key], nil
}