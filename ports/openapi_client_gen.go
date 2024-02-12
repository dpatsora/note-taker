// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package ports

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
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
	// GetNotes request
	GetNotes(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateNoteWithBody request with any body
	CreateNoteWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateNote(ctx context.Context, body CreateNoteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteNote request
	DeleteNote(ctx context.Context, noteUUID openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetNote request
	GetNote(ctx context.Context, noteUUID openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateNoteWithBody request with any body
	UpdateNoteWithBody(ctx context.Context, noteUUID openapi_types.UUID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateNote(ctx context.Context, noteUUID openapi_types.UUID, body UpdateNoteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetNotes(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetNotesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateNoteWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateNoteRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateNote(ctx context.Context, body CreateNoteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateNoteRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteNote(ctx context.Context, noteUUID openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteNoteRequest(c.Server, noteUUID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetNote(ctx context.Context, noteUUID openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetNoteRequest(c.Server, noteUUID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateNoteWithBody(ctx context.Context, noteUUID openapi_types.UUID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateNoteRequestWithBody(c.Server, noteUUID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateNote(ctx context.Context, noteUUID openapi_types.UUID, body UpdateNoteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateNoteRequest(c.Server, noteUUID, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetNotesRequest generates requests for GetNotes
func NewGetNotesRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/notes")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateNoteRequest calls the generic CreateNote builder with application/json body
func NewCreateNoteRequest(server string, body CreateNoteJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateNoteRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateNoteRequestWithBody generates requests for CreateNote with any type of body
func NewCreateNoteRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/notes")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteNoteRequest generates requests for DeleteNote
func NewDeleteNoteRequest(server string, noteUUID openapi_types.UUID) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "noteUUID", runtime.ParamLocationPath, noteUUID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/notes/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetNoteRequest generates requests for GetNote
func NewGetNoteRequest(server string, noteUUID openapi_types.UUID) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "noteUUID", runtime.ParamLocationPath, noteUUID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/notes/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateNoteRequest calls the generic UpdateNote builder with application/json body
func NewUpdateNoteRequest(server string, noteUUID openapi_types.UUID, body UpdateNoteJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateNoteRequestWithBody(server, noteUUID, "application/json", bodyReader)
}

// NewUpdateNoteRequestWithBody generates requests for UpdateNote with any type of body
func NewUpdateNoteRequestWithBody(server string, noteUUID openapi_types.UUID, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "noteUUID", runtime.ParamLocationPath, noteUUID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/notes/%s", pathParam0)
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
	// GetNotesWithResponse request
	GetNotesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetNotesResponse, error)

	// CreateNoteWithBodyWithResponse request with any body
	CreateNoteWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateNoteResponse, error)

	CreateNoteWithResponse(ctx context.Context, body CreateNoteJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateNoteResponse, error)

	// DeleteNoteWithResponse request
	DeleteNoteWithResponse(ctx context.Context, noteUUID openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteNoteResponse, error)

	// GetNoteWithResponse request
	GetNoteWithResponse(ctx context.Context, noteUUID openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetNoteResponse, error)

	// UpdateNoteWithBodyWithResponse request with any body
	UpdateNoteWithBodyWithResponse(ctx context.Context, noteUUID openapi_types.UUID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateNoteResponse, error)

	UpdateNoteWithResponse(ctx context.Context, noteUUID openapi_types.UUID, body UpdateNoteJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateNoteResponse, error)
}

type GetNotesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Notes
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r GetNotesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetNotesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateNoteResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r CreateNoteResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateNoteResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteNoteResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r DeleteNoteResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteNoteResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetNoteResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Note
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r GetNoteResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetNoteResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateNoteResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r UpdateNoteResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateNoteResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetNotesWithResponse request returning *GetNotesResponse
func (c *ClientWithResponses) GetNotesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetNotesResponse, error) {
	rsp, err := c.GetNotes(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetNotesResponse(rsp)
}

// CreateNoteWithBodyWithResponse request with arbitrary body returning *CreateNoteResponse
func (c *ClientWithResponses) CreateNoteWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateNoteResponse, error) {
	rsp, err := c.CreateNoteWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateNoteResponse(rsp)
}

func (c *ClientWithResponses) CreateNoteWithResponse(ctx context.Context, body CreateNoteJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateNoteResponse, error) {
	rsp, err := c.CreateNote(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateNoteResponse(rsp)
}

// DeleteNoteWithResponse request returning *DeleteNoteResponse
func (c *ClientWithResponses) DeleteNoteWithResponse(ctx context.Context, noteUUID openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteNoteResponse, error) {
	rsp, err := c.DeleteNote(ctx, noteUUID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteNoteResponse(rsp)
}

// GetNoteWithResponse request returning *GetNoteResponse
func (c *ClientWithResponses) GetNoteWithResponse(ctx context.Context, noteUUID openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetNoteResponse, error) {
	rsp, err := c.GetNote(ctx, noteUUID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetNoteResponse(rsp)
}

// UpdateNoteWithBodyWithResponse request with arbitrary body returning *UpdateNoteResponse
func (c *ClientWithResponses) UpdateNoteWithBodyWithResponse(ctx context.Context, noteUUID openapi_types.UUID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateNoteResponse, error) {
	rsp, err := c.UpdateNoteWithBody(ctx, noteUUID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateNoteResponse(rsp)
}

func (c *ClientWithResponses) UpdateNoteWithResponse(ctx context.Context, noteUUID openapi_types.UUID, body UpdateNoteJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateNoteResponse, error) {
	rsp, err := c.UpdateNote(ctx, noteUUID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateNoteResponse(rsp)
}

// ParseGetNotesResponse parses an HTTP response from a GetNotesWithResponse call
func ParseGetNotesResponse(rsp *http.Response) (*GetNotesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetNotesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Notes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseCreateNoteResponse parses an HTTP response from a CreateNoteWithResponse call
func ParseCreateNoteResponse(rsp *http.Response) (*CreateNoteResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateNoteResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseDeleteNoteResponse parses an HTTP response from a DeleteNoteWithResponse call
func ParseDeleteNoteResponse(rsp *http.Response) (*DeleteNoteResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteNoteResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseGetNoteResponse parses an HTTP response from a GetNoteWithResponse call
func ParseGetNoteResponse(rsp *http.Response) (*GetNoteResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetNoteResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Note
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseUpdateNoteResponse parses an HTTP response from a UpdateNoteWithResponse call
func ParseUpdateNoteResponse(rsp *http.Response) (*UpdateNoteResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateNoteResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}
