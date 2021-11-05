package base64

import (
	"encoding/base64"
	"encoding/json"
)

func EncodeObject(obj map[string]interface{}) string {
	o, _ := json.Marshal(obj)
	return base64.StdEncoding.EncodeToString(o)
}

func Encode(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}
