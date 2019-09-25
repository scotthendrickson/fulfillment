package epFulfillment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//BaseURL will be used to construct other URL's
const BaseURL = "https://api.easypost.com/fulfillment/vendor/v2/"

//APIKey is for authentication with the EasyPost API and is required
var APIKey string

//SetAPIKey will allow you to set the API key on the request
func SetAPIKey(apiKey string) {
	APIKey = apiKey
}

//mainRequest is the base function for doing any call to the EP Fulfillment
//API. It can be used to facilitate any crud operation
func mainRequest(method, path string, in, out interface{}) error {
	body, err := json.Marshal(in)

	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest(method, BaseURL+path, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}

	if APIKey != "" {
		req.SetBasicAuth(APIKey, "")
	} else {
		req.SetBasicAuth(os.Getenv("epTest"), "")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		if out != nil {
			return json.NewDecoder(resp.Body).Decode(out)
		}
		return nil
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	apiErr := &APIError{Status: resp.Status, StatusCode: resp.StatusCode}
	if json.Unmarshal(buf, &apiErrorResponse{Error: apiErr}) != nil {
		apiErr.Message = string(buf)
	}

	return apiErr
}
