//Package staytus provides methods for working with https://github.com/adamcooke/staytus API
package staytus

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

//New creates new instance of Staytus API
//
//It requires a base url of your Staytus instance, Secret and Token API strings
func New(baseURL, token, secret string) (*Staytus, error) {
	if !(strings.HasPrefix(baseURL, "https://") || strings.HasPrefix(baseURL, "http://")) {
		baseURL = "http://" + baseURL
	}
	ur, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	baseURL = strings.Trim(fmt.Sprintf("%s://%s%s", ur.Scheme, ur.Host, path.Clean(ur.Path)), "./")
	return &Staytus{
		BaseURL: baseURL,
		Token:   token,
		Secret:  secret,
	}, nil
}

func (api *Staytus) post(path string, body []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", api.BaseURL, path), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("X-Auth-Token", api.Token)
	req.Header.Add("X-Auth-Secret", api.Secret)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(bodyBuf, &data)
	if err != nil {
		return nil, err
	}
	status, ok := data["status"].(string)
	if !ok {
		return nil, errors.New("unknown error")
	}
	if status == "success" {
		value := gjson.Get(string(bodyBuf), "data")
		if value.Raw == "" {
			return nil, errors.New("no data")
		}
		return []byte(value.Raw), nil
	}
	if status == "validation-error" {
		fields := gjson.Get(string(bodyBuf), "data.errors").Value().(map[string]interface{})
		var fieldErrors []string
		for field, error := range fields {
			fieldErrors = append(fieldErrors, fmt.Sprintf("%s - %s", field, error))
		}
		return nil, errors.New(fmt.Sprintf("%s: %s", status, strings.Join(fieldErrors, ",")))
	}
	errorMessage := gjson.Get(string(bodyBuf), "data.message")
	return nil, errors.New(fmt.Sprintf("%s: %s", status, errorMessage.String()))
}
