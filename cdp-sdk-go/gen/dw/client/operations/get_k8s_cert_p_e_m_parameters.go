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

// NewGetK8sCertPEMParams creates a new GetK8sCertPEMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetK8sCertPEMParams() *GetK8sCertPEMParams {
	return &GetK8sCertPEMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetK8sCertPEMParamsWithTimeout creates a new GetK8sCertPEMParams object
// with the ability to set a timeout on a request.
func NewGetK8sCertPEMParamsWithTimeout(timeout time.Duration) *GetK8sCertPEMParams {
	return &GetK8sCertPEMParams{
		timeout: timeout,
	}
}

// NewGetK8sCertPEMParamsWithContext creates a new GetK8sCertPEMParams object
// with the ability to set a context for a request.
func NewGetK8sCertPEMParamsWithContext(ctx context.Context) *GetK8sCertPEMParams {
	return &GetK8sCertPEMParams{
		Context: ctx,
	}
}

// NewGetK8sCertPEMParamsWithHTTPClient creates a new GetK8sCertPEMParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetK8sCertPEMParamsWithHTTPClient(client *http.Client) *GetK8sCertPEMParams {
	return &GetK8sCertPEMParams{
		HTTPClient: client,
	}
}

/*
GetK8sCertPEMParams contains all the parameters to send to the API endpoint

	for the get k8s cert p e m operation.

	Typically these are written to a http.Request.
*/
type GetK8sCertPEMParams struct {

	// Input.
	Input *models.GetK8sCertPEMRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get k8s cert p e m params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetK8sCertPEMParams) WithDefaults() *GetK8sCertPEMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get k8s cert p e m params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetK8sCertPEMParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get k8s cert p e m params
func (o *GetK8sCertPEMParams) WithTimeout(timeout time.Duration) *GetK8sCertPEMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get k8s cert p e m params
func (o *GetK8sCertPEMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get k8s cert p e m params
func (o *GetK8sCertPEMParams) WithContext(ctx context.Context) *GetK8sCertPEMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get k8s cert p e m params
func (o *GetK8sCertPEMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get k8s cert p e m params
func (o *GetK8sCertPEMParams) WithHTTPClient(client *http.Client) *GetK8sCertPEMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get k8s cert p e m params
func (o *GetK8sCertPEMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get k8s cert p e m params
func (o *GetK8sCertPEMParams) WithInput(input *models.GetK8sCertPEMRequest) *GetK8sCertPEMParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get k8s cert p e m params
func (o *GetK8sCertPEMParams) SetInput(input *models.GetK8sCertPEMRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetK8sCertPEMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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