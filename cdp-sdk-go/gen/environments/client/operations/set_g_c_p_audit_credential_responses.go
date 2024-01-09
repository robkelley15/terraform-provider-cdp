// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// SetGCPAuditCredentialReader is a Reader for the SetGCPAuditCredential structure.
type SetGCPAuditCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetGCPAuditCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetGCPAuditCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSetGCPAuditCredentialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetGCPAuditCredentialOK creates a SetGCPAuditCredentialOK with default headers values
func NewSetGCPAuditCredentialOK() *SetGCPAuditCredentialOK {
	return &SetGCPAuditCredentialOK{}
}

/*
SetGCPAuditCredentialOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type SetGCPAuditCredentialOK struct {
	Payload *models.SetGCPAuditCredentialResponse
}

// IsSuccess returns true when this set g c p audit credential o k response has a 2xx status code
func (o *SetGCPAuditCredentialOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set g c p audit credential o k response has a 3xx status code
func (o *SetGCPAuditCredentialOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set g c p audit credential o k response has a 4xx status code
func (o *SetGCPAuditCredentialOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set g c p audit credential o k response has a 5xx status code
func (o *SetGCPAuditCredentialOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set g c p audit credential o k response a status code equal to that given
func (o *SetGCPAuditCredentialOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set g c p audit credential o k response
func (o *SetGCPAuditCredentialOK) Code() int {
	return 200
}

func (o *SetGCPAuditCredentialOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/setGCPAuditCredential][%d] setGCPAuditCredentialOK  %+v", 200, o.Payload)
}

func (o *SetGCPAuditCredentialOK) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/setGCPAuditCredential][%d] setGCPAuditCredentialOK  %+v", 200, o.Payload)
}

func (o *SetGCPAuditCredentialOK) GetPayload() *models.SetGCPAuditCredentialResponse {
	return o.Payload
}

func (o *SetGCPAuditCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SetGCPAuditCredentialResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetGCPAuditCredentialDefault creates a SetGCPAuditCredentialDefault with default headers values
func NewSetGCPAuditCredentialDefault(code int) *SetGCPAuditCredentialDefault {
	return &SetGCPAuditCredentialDefault{
		_statusCode: code,
	}
}

/*
SetGCPAuditCredentialDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type SetGCPAuditCredentialDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this set g c p audit credential default response has a 2xx status code
func (o *SetGCPAuditCredentialDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this set g c p audit credential default response has a 3xx status code
func (o *SetGCPAuditCredentialDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this set g c p audit credential default response has a 4xx status code
func (o *SetGCPAuditCredentialDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this set g c p audit credential default response has a 5xx status code
func (o *SetGCPAuditCredentialDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this set g c p audit credential default response a status code equal to that given
func (o *SetGCPAuditCredentialDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the set g c p audit credential default response
func (o *SetGCPAuditCredentialDefault) Code() int {
	return o._statusCode
}

func (o *SetGCPAuditCredentialDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/setGCPAuditCredential][%d] setGCPAuditCredential default  %+v", o._statusCode, o.Payload)
}

func (o *SetGCPAuditCredentialDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/setGCPAuditCredential][%d] setGCPAuditCredential default  %+v", o._statusCode, o.Payload)
}

func (o *SetGCPAuditCredentialDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetGCPAuditCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}