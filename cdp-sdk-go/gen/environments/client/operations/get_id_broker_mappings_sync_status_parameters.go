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

// NewGetIDBrokerMappingsSyncStatusParams creates a new GetIDBrokerMappingsSyncStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIDBrokerMappingsSyncStatusParams() *GetIDBrokerMappingsSyncStatusParams {
	return &GetIDBrokerMappingsSyncStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIDBrokerMappingsSyncStatusParamsWithTimeout creates a new GetIDBrokerMappingsSyncStatusParams object
// with the ability to set a timeout on a request.
func NewGetIDBrokerMappingsSyncStatusParamsWithTimeout(timeout time.Duration) *GetIDBrokerMappingsSyncStatusParams {
	return &GetIDBrokerMappingsSyncStatusParams{
		timeout: timeout,
	}
}

// NewGetIDBrokerMappingsSyncStatusParamsWithContext creates a new GetIDBrokerMappingsSyncStatusParams object
// with the ability to set a context for a request.
func NewGetIDBrokerMappingsSyncStatusParamsWithContext(ctx context.Context) *GetIDBrokerMappingsSyncStatusParams {
	return &GetIDBrokerMappingsSyncStatusParams{
		Context: ctx,
	}
}

// NewGetIDBrokerMappingsSyncStatusParamsWithHTTPClient creates a new GetIDBrokerMappingsSyncStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIDBrokerMappingsSyncStatusParamsWithHTTPClient(client *http.Client) *GetIDBrokerMappingsSyncStatusParams {
	return &GetIDBrokerMappingsSyncStatusParams{
		HTTPClient: client,
	}
}

/*
GetIDBrokerMappingsSyncStatusParams contains all the parameters to send to the API endpoint

	for the get Id broker mappings sync status operation.

	Typically these are written to a http.Request.
*/
type GetIDBrokerMappingsSyncStatusParams struct {

	// Input.
	Input *models.GetIDBrokerMappingsSyncStatusRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Id broker mappings sync status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIDBrokerMappingsSyncStatusParams) WithDefaults() *GetIDBrokerMappingsSyncStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Id broker mappings sync status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIDBrokerMappingsSyncStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Id broker mappings sync status params
func (o *GetIDBrokerMappingsSyncStatusParams) WithTimeout(timeout time.Duration) *GetIDBrokerMappingsSyncStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Id broker mappings sync status params
func (o *GetIDBrokerMappingsSyncStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Id broker mappings sync status params
func (o *GetIDBrokerMappingsSyncStatusParams) WithContext(ctx context.Context) *GetIDBrokerMappingsSyncStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Id broker mappings sync status params
func (o *GetIDBrokerMappingsSyncStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Id broker mappings sync status params
func (o *GetIDBrokerMappingsSyncStatusParams) WithHTTPClient(client *http.Client) *GetIDBrokerMappingsSyncStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Id broker mappings sync status params
func (o *GetIDBrokerMappingsSyncStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get Id broker mappings sync status params
func (o *GetIDBrokerMappingsSyncStatusParams) WithInput(input *models.GetIDBrokerMappingsSyncStatusRequest) *GetIDBrokerMappingsSyncStatusParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get Id broker mappings sync status params
func (o *GetIDBrokerMappingsSyncStatusParams) SetInput(input *models.GetIDBrokerMappingsSyncStatusRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetIDBrokerMappingsSyncStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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