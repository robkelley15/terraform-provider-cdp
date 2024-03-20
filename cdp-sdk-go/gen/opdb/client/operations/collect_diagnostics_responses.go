// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
)

// CollectDiagnosticsReader is a Reader for the CollectDiagnostics structure.
type CollectDiagnosticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollectDiagnosticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCollectDiagnosticsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCollectDiagnosticsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCollectDiagnosticsOK creates a CollectDiagnosticsOK with default headers values
func NewCollectDiagnosticsOK() *CollectDiagnosticsOK {
	return &CollectDiagnosticsOK{}
}

/*
CollectDiagnosticsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type CollectDiagnosticsOK struct {
	Payload *models.CollectDiagnosticsResponse
}

// IsSuccess returns true when this collect diagnostics o k response has a 2xx status code
func (o *CollectDiagnosticsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this collect diagnostics o k response has a 3xx status code
func (o *CollectDiagnosticsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this collect diagnostics o k response has a 4xx status code
func (o *CollectDiagnosticsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this collect diagnostics o k response has a 5xx status code
func (o *CollectDiagnosticsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this collect diagnostics o k response a status code equal to that given
func (o *CollectDiagnosticsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the collect diagnostics o k response
func (o *CollectDiagnosticsOK) Code() int {
	return 200
}

func (o *CollectDiagnosticsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/opdb/collectDiagnostics][%d] collectDiagnosticsOK  %+v", 200, o.Payload)
}

func (o *CollectDiagnosticsOK) String() string {
	return fmt.Sprintf("[POST /api/v1/opdb/collectDiagnostics][%d] collectDiagnosticsOK  %+v", 200, o.Payload)
}

func (o *CollectDiagnosticsOK) GetPayload() *models.CollectDiagnosticsResponse {
	return o.Payload
}

func (o *CollectDiagnosticsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CollectDiagnosticsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCollectDiagnosticsDefault creates a CollectDiagnosticsDefault with default headers values
func NewCollectDiagnosticsDefault(code int) *CollectDiagnosticsDefault {
	return &CollectDiagnosticsDefault{
		_statusCode: code,
	}
}

/*
CollectDiagnosticsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type CollectDiagnosticsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this collect diagnostics default response has a 2xx status code
func (o *CollectDiagnosticsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this collect diagnostics default response has a 3xx status code
func (o *CollectDiagnosticsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this collect diagnostics default response has a 4xx status code
func (o *CollectDiagnosticsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this collect diagnostics default response has a 5xx status code
func (o *CollectDiagnosticsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this collect diagnostics default response a status code equal to that given
func (o *CollectDiagnosticsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the collect diagnostics default response
func (o *CollectDiagnosticsDefault) Code() int {
	return o._statusCode
}

func (o *CollectDiagnosticsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/opdb/collectDiagnostics][%d] collectDiagnostics default  %+v", o._statusCode, o.Payload)
}

func (o *CollectDiagnosticsDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/opdb/collectDiagnostics][%d] collectDiagnostics default  %+v", o._statusCode, o.Payload)
}

func (o *CollectDiagnosticsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CollectDiagnosticsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}