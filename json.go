package item

import (
	"encoding/json"
)

func JsonToString(v interface{}) string {
	bts, _ := json.Marshal(v)
	return string(bts)
}

func JsonToStringArray(v interface{}) string {
	if v == nil {
		return "[]"
	}
	switch v.(type) {
	case string:
		if v.(string) == "" {
			return "[]"
		}
	}
	bts, _ := json.Marshal(v)
	return string(bts)
}

func JsonToStringObject(v interface{}) string {
	if v == nil {
		return "{}"
	}
	switch v.(type) {
	case string:
		if v.(string) == "" {
			return "{}"
		}
	}
	bts, _ := json.Marshal(v)
	return string(bts)
}
