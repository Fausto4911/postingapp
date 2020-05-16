package utils

import (
	"encoding/json"
)

func ToJson(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
