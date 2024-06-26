// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListStorageClassesParams creates a new ListStorageClassesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListStorageClassesParams() *ListStorageClassesParams {
	return &ListStorageClassesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListStorageClassesParamsWithTimeout creates a new ListStorageClassesParams object
// with the ability to set a timeout on a request.
func NewListStorageClassesParamsWithTimeout(timeout time.Duration) *ListStorageClassesParams {
	return &ListStorageClassesParams{
		timeout: timeout,
	}
}

// NewListStorageClassesParamsWithContext creates a new ListStorageClassesParams object
// with the ability to set a context for a request.
func NewListStorageClassesParamsWithContext(ctx context.Context) *ListStorageClassesParams {
	return &ListStorageClassesParams{
		Context: ctx,
	}
}

// NewListStorageClassesParamsWithHTTPClient creates a new ListStorageClassesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListStorageClassesParamsWithHTTPClient(client *http.Client) *ListStorageClassesParams {
	return &ListStorageClassesParams{
		HTTPClient: client,
	}
}

/*
ListStorageClassesParams contains all the parameters to send to the API endpoint

	for the list storage classes operation.

	Typically these are written to a http.Request.
*/
type ListStorageClassesParams struct {
	// Body.
	Body ListStorageClassesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list storage classes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListStorageClassesParams) WithDefaults() *ListStorageClassesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list storage classes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListStorageClassesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list storage classes params
func (o *ListStorageClassesParams) WithTimeout(timeout time.Duration) *ListStorageClassesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list storage classes params
func (o *ListStorageClassesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list storage classes params
func (o *ListStorageClassesParams) WithContext(ctx context.Context) *ListStorageClassesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list storage classes params
func (o *ListStorageClassesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list storage classes params
func (o *ListStorageClassesParams) WithHTTPClient(client *http.Client) *ListStorageClassesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list storage classes params
func (o *ListStorageClassesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list storage classes params
func (o *ListStorageClassesParams) WithBody(body ListStorageClassesBody) *ListStorageClassesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list storage classes params
func (o *ListStorageClassesParams) SetBody(body ListStorageClassesBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListStorageClassesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
