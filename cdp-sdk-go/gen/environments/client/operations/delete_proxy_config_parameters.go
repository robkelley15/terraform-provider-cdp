// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// NewDeleteProxyConfigParams creates a new DeleteProxyConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteProxyConfigParams() *DeleteProxyConfigParams {
	return &DeleteProxyConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProxyConfigParamsWithTimeout creates a new DeleteProxyConfigParams object
// with the ability to set a timeout on a request.
func NewDeleteProxyConfigParamsWithTimeout(timeout time.Duration) *DeleteProxyConfigParams {
	return &DeleteProxyConfigParams{
		timeout: timeout,
	}
}

// NewDeleteProxyConfigParamsWithContext creates a new DeleteProxyConfigParams object
// with the ability to set a context for a request.
func NewDeleteProxyConfigParamsWithContext(ctx context.Context) *DeleteProxyConfigParams {
	return &DeleteProxyConfigParams{
		Context: ctx,
	}
}

// NewDeleteProxyConfigParamsWithHTTPClient creates a new DeleteProxyConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteProxyConfigParamsWithHTTPClient(client *http.Client) *DeleteProxyConfigParams {
	return &DeleteProxyConfigParams{
		HTTPClient: client,
	}
}

/*
DeleteProxyConfigParams contains all the parameters to send to the API endpoint

	for the delete proxy config operation.

	Typically these are written to a http.Request.
*/
type DeleteProxyConfigParams struct {

	// Input.
	Input *models.DeleteProxyConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete proxy config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProxyConfigParams) WithDefaults() *DeleteProxyConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete proxy config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProxyConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete proxy config params
func (o *DeleteProxyConfigParams) WithTimeout(timeout time.Duration) *DeleteProxyConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete proxy config params
func (o *DeleteProxyConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete proxy config params
func (o *DeleteProxyConfigParams) WithContext(ctx context.Context) *DeleteProxyConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete proxy config params
func (o *DeleteProxyConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete proxy config params
func (o *DeleteProxyConfigParams) WithHTTPClient(client *http.Client) *DeleteProxyConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete proxy config params
func (o *DeleteProxyConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the delete proxy config params
func (o *DeleteProxyConfigParams) WithInput(input *models.DeleteProxyConfigRequest) *DeleteProxyConfigParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the delete proxy config params
func (o *DeleteProxyConfigParams) SetInput(input *models.DeleteProxyConfigRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProxyConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Input != nil {
		if err := r.SetBodyParam(o.Input); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}