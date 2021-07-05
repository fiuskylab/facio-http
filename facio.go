package facio

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// facio is the main struct to execute all actions and config the client
type facio struct {
	// Stores Facio's config
	config *config

	// httpClient Go's official http package
	httpClient *http.Client

	// Error stores errHandler
	Error errHandler
}

// config store
type config struct {
	baseURL string
	Error   errHandler
}

// NewConfig return a default Config
func NewConfig(baseURL string) *config {
	base, err := parseURL(baseURL)
	return &config{
		baseURL: base,
		Error:   err,
	}
}

// NewFacio returns a Facio with a given config
func NewFacio(cfg *config) *facio {
	return &facio{
		config:     cfg,
		httpClient: &http.Client{},
	}
}

type facioRes struct {
	httpRes    *http.Response
	Error      errHandler
	statusHTTP int
}

type facioReq struct {
	httpReq *http.Request
	url     string
	Error   errHandler
	method  string
}

func (f *facio) Get(endpoint string) *facioRes {
	endpoint = parseEndpoint(endpoint)
	req := &facioReq{
		url:    fmt.Sprintf("%s%s", f.config.baseURL, endpoint),
		method: http.MethodGet,
	}

	return f.request(req)
}

// request execute the given request and return a response
func (f *facio) request(req *facioReq) *facioRes {
	var err error

	res := &facioRes{}

	req.httpReq, err = http.NewRequest(req.method, req.url, nil)

	if err != nil {
		res.Error = errHandler{
			Error: err.Error(),
			isNil: false,
		}
		return res
	}

	res.httpRes, err = f.httpClient.Do(req.httpReq)

	if err != nil {
		res.Error = errHandler{
			Error: err.Error(),
			isNil: false,
		}
		return res
	}

	return res
}

// StatusCode return HTTP status
func (f *facioRes) StatusCode() int {
	if f.statusHTTP == 0 {
		f.statusHTTP = f.httpRes.StatusCode
		return f.statusHTTP
	}
	return f.statusHTTP
}

// ToStruct unmarshal JSON response into given struct
func (f *facioRes) ToStruct(i *interface{}) errHandler {
	err := newNilError()

	dec := json.NewDecoder(f.httpRes.Body)

	errJson := dec.Decode(i)

	if errJson != nil {
		err.isNil = false
		err.Error = errJson.Error()
	}

	return err
}
