// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new locations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for locations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddLocation(params *AddLocationParams, opts ...ClientOption) (*AddLocationOK, error)

	ChangeLocation(params *ChangeLocationParams, opts ...ClientOption) (*ChangeLocationOK, error)

	ListLocations(params *ListLocationsParams, opts ...ClientOption) (*ListLocationsOK, error)

	RemoveLocation(params *RemoveLocationParams, opts ...ClientOption) (*RemoveLocationOK, error)

	TestLocationConfig(params *TestLocationConfigParams, opts ...ClientOption) (*TestLocationConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddLocation adds location adds backup location
*/
func (a *Client) AddLocation(params *AddLocationParams, opts ...ClientOption) (*AddLocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddLocationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddLocation",
		Method:             "POST",
		PathPattern:        "/v1/management/backup/Locations/Add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddLocationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddLocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddLocationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ChangeLocation changes location changes backup location
*/
func (a *Client) ChangeLocation(params *ChangeLocationParams, opts ...ClientOption) (*ChangeLocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeLocationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ChangeLocation",
		Method:             "POST",
		PathPattern:        "/v1/management/backup/Locations/Change",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeLocationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeLocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeLocationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListLocations lists locations returns a list of all backup locations
*/
func (a *Client) ListLocations(params *ListLocationsParams, opts ...ClientOption) (*ListLocationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListLocationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListLocations",
		Method:             "POST",
		PathPattern:        "/v1/management/backup/Locations/List",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListLocationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListLocationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListLocationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RemoveLocation removes location removes existing backup location
*/
func (a *Client) RemoveLocation(params *RemoveLocationParams, opts ...ClientOption) (*RemoveLocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveLocationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RemoveLocation",
		Method:             "POST",
		PathPattern:        "/v1/management/backup/Locations/Remove",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveLocationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveLocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RemoveLocationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
TestLocationConfig tests location config tests backup location and credentials
*/
func (a *Client) TestLocationConfig(params *TestLocationConfigParams, opts ...ClientOption) (*TestLocationConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestLocationConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TestLocationConfig",
		Method:             "POST",
		PathPattern:        "/v1/management/backup/Locations/TestConfig",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TestLocationConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TestLocationConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TestLocationConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
