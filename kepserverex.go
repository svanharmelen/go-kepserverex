//
// Copyright 2019, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package kepserverex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	apiVersionPath = "config/v1/project/"
	userAgent      = "go-kepserverex"
)

// A Client manages communication with the KEPServerEX API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	baseURL *url.URL

	// Username and password used for authentication.
	username, password string

	// Services used for talking to different parts of the KEPServerEX API.
	Channels  *ChannelService
	Devices   *DeviceService
	TagGroups *TagGroupService
	Tags      *TagService
}

// NewClient returns a new KEPServerEX API client. If a nil httpClient is
// provided, http.DefaultClient will be used. To use API methods which require
// authentication, provide a valid username and password.
func NewClient(httpClient *http.Client, host, username, password string) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		client:   httpClient,
		username: username,
		password: password,
	}

	if err := c.SetBaseURL("https://" + host); err != nil {
		return nil, err
	}

	// Create all the public services.
	c.Channels = &ChannelService{client: c}
	c.Devices = &DeviceService{client: c}
	c.TagGroups = &TagGroupService{client: c}
	c.Tags = &TagService{client: c}

	return c, nil
}

// BaseURL return a copy of the baseURL.
func (c *Client) BaseURL() *url.URL {
	u := *c.baseURL
	return &u
}

// SetBaseURL sets the base URL for API requests to a custom endpoint.
func (c *Client) SetBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(baseURL.Path, apiVersionPath) {
		baseURL.Path += apiVersionPath
	}

	// Update the base URL of the client.
	c.baseURL = baseURL

	return nil
}

// NewRequest creates an API request. A relative URL path can be provided in
// path, in which case it is resolved relative to the base URL of the Client.
// Relative URL paths should always be specified without a preceding slash. If
// specified, the value pointed to by opt is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, path string, opt interface{}) (*http.Request, error) {
	u := *c.baseURL
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	// Set the encoded path data
	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	req := &http.Request{
		Method:     method,
		URL:        &u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       u.Host,
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", userAgent)
	req.SetBasicAuth(c.username, c.password)

	switch {
	case method == "GET" && opt != nil:
		q, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		req.URL.RawQuery = q.Encode()

	case method == "POST" || method == "PUT":
		body := &bytes.Buffer{}
		encoder := json.NewEncoder(body)
		encoder.SetEscapeHTML(false)
		if err = encoder.Encode(opt); err != nil {
			return nil, err
		}
		req.Body = ioutil.NopCloser(body)
		req.GetBody = func() (io.ReadCloser, error) {
			return ioutil.NopCloser(body), nil
		}
		req.ContentLength = int64(body.Len())
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *http.Request, v interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return err
}

// An ErrorResponse reports errors caused by an API request.
type ErrorResponse struct {
	Code     int
	Response *http.Response
	Message  string
}

func (e *ErrorResponse) Error() string {
	path, _ := url.QueryUnescape(e.Response.Request.URL.Path)
	u := fmt.Sprintf("%s://%s%s", e.Response.Request.URL.Scheme, e.Response.Request.URL.Host, path)
	return fmt.Sprintf("%s %s: %d %s", e.Response.Request.Method, u, e.Response.StatusCode, e.Message)
}

// CheckResponse checks the API response for errors, and returns them if present.
func CheckResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204:
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	err := json.NewDecoder(r.Body).Decode(errorResponse)
	if err != nil {
		errorResponse.Code = 500
		errorResponse.Message = "failed to parse unknown error format"
	}

	return errorResponse
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int {
	p := new(int)
	*p = v
	return p
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	p := new(string)
	*p = v
	return p
}
