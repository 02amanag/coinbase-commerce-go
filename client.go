package coinbase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	// ENDPOINT defaults to https://api.commerce.coinbase.com
	ENDPOINT = "https://api.commerce.coinbase.com"

	// API_VERSION since version two you have to
	API_VERSION = "2018-03-22"
)

// If Endpoint or ApiVersion aren't defined the library will use the consts value
type APIClient struct {
	Key        string
	Endpoint   string
	ApiVersion string
	Charge     *ChargeService
}

func Client(api_key string) (client APIClient) {
	client.Key = api_key
	client.Charge = new(ChargeService)
	client.Charge.Api = &client
	return
}

// Fetch is wrapper for http request made to coinbase and also handel all other cases
func (a *APIClient) Fetch(method, path string, body interface{}) (interface{}, error) {
	if a.Endpoint == "" {
		a.Endpoint = ENDPOINT // default endpoint
	}
	if a.ApiVersion == "" {
		a.ApiVersion = API_VERSION // default api version
	}

	var bodyBuffered io.Reader
	if body != nil {
		var data []byte
		var err error
		switch body.(type) {
		case string:
			data = []byte(body.(string))
		default:
			data, err = json.Marshal(body)
			if err != nil {
				return nil, err
			}
		}
		bodyBuffered = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, a.Endpoint+path, bodyBuffered)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CC-Version", a.ApiVersion)
	req.Header.Set("X-CC-Api-Key", a.Key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return errors.New(string(bodyBytes)), nil
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := make(map[string]interface{})
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
