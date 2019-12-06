package epFulfillment

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

//To help with consistency between this SDK and EasyPost's GoLang client for their shipping API (https://github.com/EasyPost/easypost-go)
//I have taken most of the main functionality and used it here. You can reference that work here https://github.com/EasyPost/easypost-go.
const defaultUserAgent = "epFulfillment Go Client V0.0.2"

var apiBaseURL = &url.URL{
	Scheme: "https", Host: "api.easypost.com", Path: "fulfillment/vendor/v2/",
}

// A Client provides an HTTP client for EasyPost API operations.
type Client struct {
	// BaseURL specifies the location of the API. It is used with
	// ResolveReference to create request URLs. (If 'Path' is specified, it
	// should end with a trailing slash.) If nil, the default will be used.
	BaseURL *url.URL
	// Client is an HTTP client used to make API requests. If nil,
	// http.DefaultClient will be used.
	Client *http.Client
	// APIKey is the user's API key. It is required.
	// Note: Treat your API Keys as passwords—keep them secret. API Keys give
	// full read/write access to your account, so they should not be included in
	// public repositories, emails, client side code, etc.
	APIKey string
	// UserAgent is a User-Agent to be sent with API HTTP requests. If empty,
	// a default will be used.
	UserAgent string
}

// New returns a new Client with the given API key.
func New(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

func (c *Client) baseURL() *url.URL {
	if c.BaseURL != nil {
		return c.BaseURL
	}
	return apiBaseURL
}

func (c *Client) userAgent() string {
	if c.UserAgent != "" {
		return c.UserAgent
	}
	return defaultUserAgent
}

func (c *Client) client() *http.Client {
	if c.Client != nil {
		return c.Client
	}
	return http.DefaultClient
}

//OLD CODE
// //BaseURL will be used to construct other URL's
// const BaseURL = "https://api.easypost.com/fulfillment/vendor/v2/"

// //APIKey is for authentication with the EasyPost API and is required
// var APIKey string

// //SetAPIKey will allow you to set the API key on the request
// func SetAPIKey(apiKey string) {
// 	APIKey = apiKey
// }

//mainRequest is the base function for doing any call to the EP Fulfillment
//API. It can be used to facilitate any crud operation
func (c *Client) setBody(req *http.Request, in interface{}) error {
	switch in := in.(type) {
	case url.Values:
		buf := []byte(in.Encode())
		req.Body = ioutil.NopCloser(bytes.NewReader(buf))
		req.GetBody = func() (io.ReadCloser, error) {
			return ioutil.NopCloser(bytes.NewReader(buf)), nil
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		// Setting Content-Length avoids chunked encoding, which the API
		// server doesn't currently support.
		req.ContentLength = int64(len(buf))
	case nil: // nop
	default:
		buf, err := json.Marshal(in)
		if err != nil {
			return err
		}
		req.Body = ioutil.NopCloser(bytes.NewReader(buf))
		req.GetBody = func() (io.ReadCloser, error) {
			return ioutil.NopCloser(bytes.NewReader(buf)), nil
		}
		req.Header.Set("Content-Type", "application/json")
		// Setting Content-Length avoids chunked encoding, which the API
		// server doesn't currently support.
		req.ContentLength = int64(len(buf))
	}
	return nil
}

func (c *Client) mainRequest(method, path string, in, out interface{}) error {
	req := &http.Request{
		Method: method,
		URL:    c.baseURL().ResolveReference(&url.URL{Path: path}),
		Header: make(http.Header, 2),
	}
	req.Header.Set("User-Agent", c.userAgent())
	if err := c.setBody(req, in); err != nil {
		return err
	}
	req.SetBasicAuth(c.APIKey, "")

	res, err := c.client().Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		if out != nil {
			return json.NewDecoder(res.Body).Decode(out)
		}
		return nil
	}

	buf, _ := ioutil.ReadAll(res.Body)
	apiErr := &APIError{Status: res.Status, StatusCode: res.StatusCode}
	if json.Unmarshal(buf, &apiErrorResponse{Error: apiErr}) != nil {
		apiErr.Message = string(buf)
	}
	return apiErr
}

//OLD FUNC
// func mainRequest(method, path string, in, out interface{}) error {
// 	body, err := json.Marshal(in)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	req, err := http.NewRequest(method, BaseURL+path, bytes.NewBuffer(body))
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	if APIKey != "" {
// 		req.SetBasicAuth(APIKey, "")
// 	} else {
// 		req.SetBasicAuth(os.Getenv("epTest"), "")
// 	}

// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("User-Agent", "epFulfillment Go Client V0.0.1")

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
// 		if out != nil {
// 			return json.NewDecoder(resp.Body).Decode(out)
// 		}
// 		return nil
// 	}

// 	buf, _ := ioutil.ReadAll(resp.Body)
// 	apiErr := &APIError{Status: resp.Status, StatusCode: resp.StatusCode}
// 	if json.Unmarshal(buf, &apiErrorResponse{Error: apiErr}) != nil {
// 		apiErr.Message = string(buf)
// 	}

// 	return apiErr
// }
