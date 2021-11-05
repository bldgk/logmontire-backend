package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"logmontire/crypto/encoding/json"
	"net/http"
	"net/url"
	"strings"
)

func Post(uri string, params map[string]interface{}, body map[string]interface{}) (interface{}, error) {
	uri += stringify(params)

	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(json.FromInterface(body)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	response, _ := json.ToInterface(respBody)

	return response, nil
}

func PostFormUrl(uri string, params map[string]interface{}, body url.Values) (interface{}, error) {
	uri += stringify(params)

	client := &http.Client{}

	req, err := http.NewRequest("POST", uri, strings.NewReader(body.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	response, _ := json.ToInterface(respBody)

	return response, nil
}

func Get(uri string, params map[string]interface{}) (interface{}, error) {
	uri += stringify(params)

	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	resp, _ := json.ToInterface(body)
	return resp, nil
}

func stringify(params map[string]interface{}) string {
	if len(params) > 0 && params != nil {
		urlParams := "?"

		for key, value := range params {
			urlParams += fmt.Sprintf("%s=%s&", key, value)
		}
		return urlParams
	}
	return ""
}
