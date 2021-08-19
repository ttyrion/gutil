package maputil

import (
	"errors"
)

func ReadMapValue(m map[string]interface{}, key string) (interface{}, error) {
	if m == nil {
		return nil, errors.New("nil map")
	}

	if _, ok := m[key]; !ok {
		return nil, errors.New("key not exist")
	}

	return m[key], nil
}