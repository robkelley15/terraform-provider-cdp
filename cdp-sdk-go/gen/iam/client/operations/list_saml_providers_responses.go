// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// ListSamlProvidersReader is a Reader for the ListSamlProviders structure.
type ListSamlProvidersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSamlProvidersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSamlProvidersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListSamlProvidersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSamlProvidersOK creates a ListSamlProvidersOK with default headers values
func NewListSamlProvidersOK() *ListSamlProvidersOK {
	return &ListSamlProvidersOK{}
}

/*
ListSamlProvidersOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListSamlProvidersOK struct {
	Payload *models.ListSamlProvidersResponse
}

// IsSuccess returns true when this list saml providers o k response has a 2xx status code
func (o *ListSamlProvidersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list saml providers o k response has a 3xx status code
func (o *ListSamlProvidersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list saml providers o k response has a 4xx status code
func (o *ListSamlProvidersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list saml providers o k response has a 5xx status code
func (o *ListSamlProvidersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list saml providers o k response a status code equal to that given
func (o *ListSamlProvidersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list saml providers o k response
func (o *ListSamlProvidersOK) Code() int {
	return 200
}

func (o *ListSamlProvidersOK) Error() string {
	return fmt.Sprintf("[POST /iam/listSamlProviders][%d] listSamlProvidersOK  %+v", 200, o.Payload)
}

func (o *ListSamlProvidersOK) String() string {
	return fmt.Sprintf("[POST /iam/listSamlProviders][%d] listSamlProvidersOK  %+v", 200, o.Payload)
}

func (o *ListSamlProvidersOK) GetPayload() *models.ListSamlProvidersResponse {
	return o.Payload
}

func (o *ListSamlProvidersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListSamlProvidersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSamlProvidersDefault creates a ListSamlProvidersDefault with default headers values
func NewListSamlProvidersDefault(code int) *ListSamlProvidersDefault {
	return &ListSamlProvidersDefault{
		_statusCode: code,
	}
}

/*
ListSamlProvidersDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListSamlProvidersDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list saml providers default response has a 2xx status code
func (o *ListSamlProvidersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list saml providers default response has a 3xx status code
func (o *ListSamlProvidersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list saml providers default response has a 4xx status code
func (o *ListSamlProvidersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list saml providers default response has a 5xx status code
func (o *ListSamlProvidersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list saml providers default response a status code equal to that given
func (o *ListSamlProvidersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list saml providers default response
func (o *ListSamlProvidersDefault) Code() int {
	return o._statusCode
}

func (o *ListSamlProvidersDefault) Error() string {
	return fmt.Sprintf("[POST /iam/listSamlProviders][%d] listSamlProviders default  %+v", o._statusCode, o.Payload)
}

func (o *ListSamlProvidersDefault) String() string {
	return fmt.Sprintf("[POST /iam/listSamlProviders][%d] listSamlProviders default  %+v", o._statusCode, o.Payload)
}

func (o *ListSamlProvidersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListSamlProvidersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}