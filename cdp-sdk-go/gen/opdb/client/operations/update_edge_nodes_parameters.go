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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
)

// NewUpdateEdgeNodesParams creates a new UpdateEdgeNodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateEdgeNodesParams() *UpdateEdgeNodesParams {
	return &UpdateEdgeNodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEdgeNodesParamsWithTimeout creates a new UpdateEdgeNodesParams object
// with the ability to set a timeout on a request.
func NewUpdateEdgeNodesParamsWithTimeout(timeout time.Duration) *UpdateEdgeNodesParams {
	return &UpdateEdgeNodesParams{
		timeout: timeout,
	}
}

// NewUpdateEdgeNodesParamsWithContext creates a new UpdateEdgeNodesParams object
// with the ability to set a context for a request.
func NewUpdateEdgeNodesParamsWithContext(ctx context.Context) *UpdateEdgeNodesParams {
	return &UpdateEdgeNodesParams{
		Context: ctx,
	}
}

// NewUpdateEdgeNodesParamsWithHTTPClient creates a new UpdateEdgeNodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateEdgeNodesParamsWithHTTPClient(client *http.Client) *UpdateEdgeNodesParams {
	return &UpdateEdgeNodesParams{
		HTTPClient: client,
	}
}

/*
UpdateEdgeNodesParams contains all the parameters to send to the API endpoint

	for the update edge nodes operation.

	Typically these are written to a http.Request.
*/
type UpdateEdgeNodesParams struct {

	// Input.
	Input *models.UpdateEdgeNodesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update edge nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateEdgeNodesParams) WithDefaults() *UpdateEdgeNodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update edge nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateEdgeNodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update edge nodes params
func (o *UpdateEdgeNodesParams) WithTimeout(timeout time.Duration) *UpdateEdgeNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update edge nodes params
func (o *UpdateEdgeNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update edge nodes params
func (o *UpdateEdgeNodesParams) WithContext(ctx context.Context) *UpdateEdgeNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update edge nodes params
func (o *UpdateEdgeNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update edge nodes params
func (o *UpdateEdgeNodesParams) WithHTTPClient(client *http.Client) *UpdateEdgeNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update edge nodes params
func (o *UpdateEdgeNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the update edge nodes params
func (o *UpdateEdgeNodesParams) WithInput(input *models.UpdateEdgeNodesRequest) *UpdateEdgeNodesParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the update edge nodes params
func (o *UpdateEdgeNodesParams) SetInput(input *models.UpdateEdgeNodesRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEdgeNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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