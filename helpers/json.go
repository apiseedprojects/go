package helpers

import (
	"encoding/json"
	"fmt"
)

func ToJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("error encoding to JSON: %s", err.Error())
	}
	return string(b)
}
