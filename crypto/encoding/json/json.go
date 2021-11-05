package json

import (
	"encoding/json"
	"logmontire/common/errors"
	"logmontire/crypto/encoding"
)

/**
From Interface converts interface to json
*/
func Fi(obj interface{}) []byte {
	jsonResult, _ := json.Marshal(obj)
	return jsonResult
}

func FromInterface(obj interface{}) []byte {
	jsonResult, _ := json.Marshal(obj)
	return jsonResult
}

/**
To Interface
*/
func Ti(encodedJson []byte) (map[string]interface{}, *errors.Error) {
	var inter map[string]interface{}

	err := json.Unmarshal(encodedJson, &inter)
	if err != nil {
		return nil, encoding.JSONParseError(string(encodedJson))
	}

	return inter, nil
}

func ToInterface(encodedJson []byte) (interface{}, *errors.Error) {
	var inter interface{}

	err := json.Unmarshal(encodedJson, &inter)
	if err != nil {
		return nil, encoding.JSONParseError(string(encodedJson))
	}

	return inter, nil
}
