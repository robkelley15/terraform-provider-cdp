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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
)

// NewUpdateAwsClusterParams creates a new UpdateAwsClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAwsClusterParams() *UpdateAwsClusterParams {
	return &UpdateAwsClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAwsClusterParamsWithTimeout creates a new UpdateAwsClusterParams object
// with the ability to set a timeout on a request.
func NewUpdateAwsClusterParamsWithTimeout(timeout time.Duration) *UpdateAwsClusterParams {
	return &UpdateAwsClusterParams{
		timeout: timeout,
	}
}

// NewUpdateAwsClusterParamsWithContext creates a new UpdateAwsClusterParams object
// with the ability to set a context for a request.
func NewUpdateAwsClusterParamsWithContext(ctx context.Context) *UpdateAwsClusterParams {
	return &UpdateAwsClusterParams{
		Context: ctx,
	}
}

// NewUpdateAwsClusterParamsWithHTTPClient creates a new UpdateAwsClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAwsClusterParamsWithHTTPClient(client *http.Client) *UpdateAwsClusterParams {
	return &UpdateAwsClusterParams{
		HTTPClient: client,
	}
}

/*
UpdateAwsClusterParams contains all the parameters to send to the API endpoint

	for the update aws cluster operation.

	Typically these are written to a http.Request.
*/
type UpdateAwsClusterParams struct {

	// Input.
	Input *models.UpdateAwsClusterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update aws cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAwsClusterParams) WithDefaults() *UpdateAwsClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update aws cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAwsClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update aws cluster params
func (o *UpdateAwsClusterParams) WithTimeout(timeout time.Duration) *UpdateAwsClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update aws cluster params
func (o *UpdateAwsClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update aws cluster params
func (o *UpdateAwsClusterParams) WithContext(ctx context.Context) *UpdateAwsClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update aws cluster params
func (o *UpdateAwsClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update aws cluster params
func (o *UpdateAwsClusterParams) WithHTTPClient(client *http.Client) *UpdateAwsClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update aws cluster params
func (o *UpdateAwsClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the update aws cluster params
func (o *UpdateAwsClusterParams) WithInput(input *models.UpdateAwsClusterRequest) *UpdateAwsClusterParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the update aws cluster params
func (o *UpdateAwsClusterParams) SetInput(input *models.UpdateAwsClusterRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAwsClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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