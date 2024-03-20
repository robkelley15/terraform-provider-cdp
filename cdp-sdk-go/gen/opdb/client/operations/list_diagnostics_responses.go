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

// ListDiagnosticsReader is a Reader for the ListDiagnostics structure.
type ListDiagnosticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDiagnosticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDiagnosticsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListDiagnosticsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDiagnosticsOK creates a ListDiagnosticsOK with default headers values
func NewListDiagnosticsOK() *ListDiagnosticsOK {
	return &ListDiagnosticsOK{}
}

/*
ListDiagnosticsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListDiagnosticsOK struct {
	Payload *models.ListDiagnosticsResponse
}

// IsSuccess returns true when this list diagnostics o k response has a 2xx status code
func (o *ListDiagnosticsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list diagnostics o k response has a 3xx status code
func (o *ListDiagnosticsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list diagnostics o k response has a 4xx status code
func (o *ListDiagnosticsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list diagnostics o k response has a 5xx status code
func (o *ListDiagnosticsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list diagnostics o k response a status code equal to that given
func (o *ListDiagnosticsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list diagnostics o k response
func (o *ListDiagnosticsOK) Code() int {
	return 200
}

func (o *ListDiagnosticsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/opdb/listDiagnostics][%d] listDiagnosticsOK  %+v", 200, o.Payload)
}

func (o *ListDiagnosticsOK) String() string {
	return fmt.Sprintf("[POST /api/v1/opdb/listDiagnostics][%d] listDiagnosticsOK  %+v", 200, o.Payload)
}

func (o *ListDiagnosticsOK) GetPayload() *models.ListDiagnosticsResponse {
	return o.Payload
}

func (o *ListDiagnosticsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListDiagnosticsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDiagnosticsDefault creates a ListDiagnosticsDefault with default headers values
func NewListDiagnosticsDefault(code int) *ListDiagnosticsDefault {
	return &ListDiagnosticsDefault{
		_statusCode: code,
	}
}

/*
ListDiagnosticsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListDiagnosticsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list diagnostics default response has a 2xx status code
func (o *ListDiagnosticsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list diagnostics default response has a 3xx status code
func (o *ListDiagnosticsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list diagnostics default response has a 4xx status code
func (o *ListDiagnosticsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list diagnostics default response has a 5xx status code
func (o *ListDiagnosticsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list diagnostics default response a status code equal to that given
func (o *ListDiagnosticsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list diagnostics default response
func (o *ListDiagnosticsDefault) Code() int {
	return o._statusCode
}

func (o *ListDiagnosticsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/opdb/listDiagnostics][%d] listDiagnostics default  %+v", o._statusCode, o.Payload)
}

func (o *ListDiagnosticsDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/opdb/listDiagnostics][%d] listDiagnostics default  %+v", o._statusCode, o.Payload)
}

func (o *ListDiagnosticsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListDiagnosticsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}