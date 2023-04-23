// Package demo provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package demo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// ListCurrentUserDemos request
	ListCurrentUserDemos(ctx context.Context, params *ListCurrentUserDemosParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateCurrentUserDemo request with any body
	CreateCurrentUserDemoWithBody(ctx context.Context, params *CreateCurrentUserDemoParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateCurrentUserDemo(ctx context.Context, params *CreateCurrentUserDemoParams, body CreateCurrentUserDemoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ListCurrentUserDemos(ctx context.Context, params *ListCurrentUserDemosParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListCurrentUserDemosRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateCurrentUserDemoWithBody(ctx context.Context, params *CreateCurrentUserDemoParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateCurrentUserDemoRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateCurrentUserDemo(ctx context.Context, params *CreateCurrentUserDemoParams, body CreateCurrentUserDemoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateCurrentUserDemoRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewListCurrentUserDemosRequest generates requests for ListCurrentUserDemos
func NewListCurrentUserDemosRequest(server string, params *ListCurrentUserDemosParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/demos")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "error", runtime.ParamLocationQuery, params.Error); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var headerParam0 string

	headerParam0, err = runtime.StyleParamWithLocation("simple", false, "Accept-Language", runtime.ParamLocationHeader, params.AcceptLanguage)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept-Language", headerParam0)

	return req, nil
}

// NewCreateCurrentUserDemoRequest calls the generic CreateCurrentUserDemo builder with application/json body
func NewCreateCurrentUserDemoRequest(server string, params *CreateCurrentUserDemoParams, body CreateCurrentUserDemoJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateCurrentUserDemoRequestWithBody(server, params, "application/json", bodyReader)
}

// NewCreateCurrentUserDemoRequestWithBody generates requests for CreateCurrentUserDemo with any type of body
func NewCreateCurrentUserDemoRequestWithBody(server string, params *CreateCurrentUserDemoParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/demos")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	var headerParam0 string

	headerParam0, err = runtime.StyleParamWithLocation("simple", false, "Accept-Language", runtime.ParamLocationHeader, params.AcceptLanguage)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept-Language", headerParam0)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// ListCurrentUserDemos request
	ListCurrentUserDemosWithResponse(ctx context.Context, params *ListCurrentUserDemosParams, reqEditors ...RequestEditorFn) (*ListCurrentUserDemosResponse, error)

	// CreateCurrentUserDemo request with any body
	CreateCurrentUserDemoWithBodyWithResponse(ctx context.Context, params *CreateCurrentUserDemoParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateCurrentUserDemoResponse, error)

	CreateCurrentUserDemoWithResponse(ctx context.Context, params *CreateCurrentUserDemoParams, body CreateCurrentUserDemoJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateCurrentUserDemoResponse, error)
}

type ListCurrentUserDemosResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Demo
}

// Status returns HTTPResponse.Status
func (r ListCurrentUserDemosResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListCurrentUserDemosResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateCurrentUserDemoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Demo
}

// Status returns HTTPResponse.Status
func (r CreateCurrentUserDemoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateCurrentUserDemoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListCurrentUserDemosWithResponse request returning *ListCurrentUserDemosResponse
func (c *ClientWithResponses) ListCurrentUserDemosWithResponse(ctx context.Context, params *ListCurrentUserDemosParams, reqEditors ...RequestEditorFn) (*ListCurrentUserDemosResponse, error) {
	rsp, err := c.ListCurrentUserDemos(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListCurrentUserDemosResponse(rsp)
}

// CreateCurrentUserDemoWithBodyWithResponse request with arbitrary body returning *CreateCurrentUserDemoResponse
func (c *ClientWithResponses) CreateCurrentUserDemoWithBodyWithResponse(ctx context.Context, params *CreateCurrentUserDemoParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateCurrentUserDemoResponse, error) {
	rsp, err := c.CreateCurrentUserDemoWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateCurrentUserDemoResponse(rsp)
}

func (c *ClientWithResponses) CreateCurrentUserDemoWithResponse(ctx context.Context, params *CreateCurrentUserDemoParams, body CreateCurrentUserDemoJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateCurrentUserDemoResponse, error) {
	rsp, err := c.CreateCurrentUserDemo(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateCurrentUserDemoResponse(rsp)
}

// ParseListCurrentUserDemosResponse parses an HTTP response from a ListCurrentUserDemosWithResponse call
func ParseListCurrentUserDemosResponse(rsp *http.Response) (*ListCurrentUserDemosResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListCurrentUserDemosResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Demo
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreateCurrentUserDemoResponse parses an HTTP response from a CreateCurrentUserDemoWithResponse call
func ParseCreateCurrentUserDemoResponse(rsp *http.Response) (*CreateCurrentUserDemoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateCurrentUserDemoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Demo
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
