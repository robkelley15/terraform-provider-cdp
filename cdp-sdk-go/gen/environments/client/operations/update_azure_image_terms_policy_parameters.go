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

// NewUpdateAzureImageTermsPolicyParams creates a new UpdateAzureImageTermsPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAzureImageTermsPolicyParams() *UpdateAzureImageTermsPolicyParams {
	return &UpdateAzureImageTermsPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAzureImageTermsPolicyParamsWithTimeout creates a new UpdateAzureImageTermsPolicyParams object
// with the ability to set a timeout on a request.
func NewUpdateAzureImageTermsPolicyParamsWithTimeout(timeout time.Duration) *UpdateAzureImageTermsPolicyParams {
	return &UpdateAzureImageTermsPolicyParams{
		timeout: timeout,
	}
}

// NewUpdateAzureImageTermsPolicyParamsWithContext creates a new UpdateAzureImageTermsPolicyParams object
// with the ability to set a context for a request.
func NewUpdateAzureImageTermsPolicyParamsWithContext(ctx context.Context) *UpdateAzureImageTermsPolicyParams {
	return &UpdateAzureImageTermsPolicyParams{
		Context: ctx,
	}
}

// NewUpdateAzureImageTermsPolicyParamsWithHTTPClient creates a new UpdateAzureImageTermsPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAzureImageTermsPolicyParamsWithHTTPClient(client *http.Client) *UpdateAzureImageTermsPolicyParams {
	return &UpdateAzureImageTermsPolicyParams{
		HTTPClient: client,
	}
}

/*
UpdateAzureImageTermsPolicyParams contains all the parameters to send to the API endpoint

	for the update azure image terms policy operation.

	Typically these are written to a http.Request.
*/
type UpdateAzureImageTermsPolicyParams struct {

	// Input.
	Input *models.UpdateAzureImageTermsPolicyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update azure image terms policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAzureImageTermsPolicyParams) WithDefaults() *UpdateAzureImageTermsPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update azure image terms policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAzureImageTermsPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update azure image terms policy params
func (o *UpdateAzureImageTermsPolicyParams) WithTimeout(timeout time.Duration) *UpdateAzureImageTermsPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update azure image terms policy params
func (o *UpdateAzureImageTermsPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update azure image terms policy params
func (o *UpdateAzureImageTermsPolicyParams) WithContext(ctx context.Context) *UpdateAzureImageTermsPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update azure image terms policy params
func (o *UpdateAzureImageTermsPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update azure image terms policy params
func (o *UpdateAzureImageTermsPolicyParams) WithHTTPClient(client *http.Client) *UpdateAzureImageTermsPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update azure image terms policy params
func (o *UpdateAzureImageTermsPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the update azure image terms policy params
func (o *UpdateAzureImageTermsPolicyParams) WithInput(input *models.UpdateAzureImageTermsPolicyRequest) *UpdateAzureImageTermsPolicyParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the update azure image terms policy params
func (o *UpdateAzureImageTermsPolicyParams) SetInput(input *models.UpdateAzureImageTermsPolicyRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAzureImageTermsPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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