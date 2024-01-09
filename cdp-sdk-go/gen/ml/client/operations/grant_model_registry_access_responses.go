// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
)

// GrantModelRegistryAccessReader is a Reader for the GrantModelRegistryAccess structure.
type GrantModelRegistryAccessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GrantModelRegistryAccessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGrantModelRegistryAccessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGrantModelRegistryAccessDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGrantModelRegistryAccessOK creates a GrantModelRegistryAccessOK with default headers values
func NewGrantModelRegistryAccessOK() *GrantModelRegistryAccessOK {
	return &GrantModelRegistryAccessOK{}
}

/*
GrantModelRegistryAccessOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type GrantModelRegistryAccessOK struct {
	Payload *models.GrantModelRegistryAccessResponse
}

// IsSuccess returns true when this grant model registry access o k response has a 2xx status code
func (o *GrantModelRegistryAccessOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this grant model registry access o k response has a 3xx status code
func (o *GrantModelRegistryAccessOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this grant model registry access o k response has a 4xx status code
func (o *GrantModelRegistryAccessOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this grant model registry access o k response has a 5xx status code
func (o *GrantModelRegistryAccessOK) IsServerError() bool {
	return false
}

// IsCode returns true when this grant model registry access o k response a status code equal to that given
func (o *GrantModelRegistryAccessOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the grant model registry access o k response
func (o *GrantModelRegistryAccessOK) Code() int {
	return 200
}

func (o *GrantModelRegistryAccessOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/grantModelRegistryAccess][%d] grantModelRegistryAccessOK  %+v", 200, o.Payload)
}

func (o *GrantModelRegistryAccessOK) String() string {
	return fmt.Sprintf("[POST /api/v1/ml/grantModelRegistryAccess][%d] grantModelRegistryAccessOK  %+v", 200, o.Payload)
}

func (o *GrantModelRegistryAccessOK) GetPayload() *models.GrantModelRegistryAccessResponse {
	return o.Payload
}

func (o *GrantModelRegistryAccessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrantModelRegistryAccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGrantModelRegistryAccessDefault creates a GrantModelRegistryAccessDefault with default headers values
func NewGrantModelRegistryAccessDefault(code int) *GrantModelRegistryAccessDefault {
	return &GrantModelRegistryAccessDefault{
		_statusCode: code,
	}
}

/*
GrantModelRegistryAccessDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GrantModelRegistryAccessDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this grant model registry access default response has a 2xx status code
func (o *GrantModelRegistryAccessDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this grant model registry access default response has a 3xx status code
func (o *GrantModelRegistryAccessDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this grant model registry access default response has a 4xx status code
func (o *GrantModelRegistryAccessDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this grant model registry access default response has a 5xx status code
func (o *GrantModelRegistryAccessDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this grant model registry access default response a status code equal to that given
func (o *GrantModelRegistryAccessDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the grant model registry access default response
func (o *GrantModelRegistryAccessDefault) Code() int {
	return o._statusCode
}

func (o *GrantModelRegistryAccessDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/grantModelRegistryAccess][%d] grantModelRegistryAccess default  %+v", o._statusCode, o.Payload)
}

func (o *GrantModelRegistryAccessDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/ml/grantModelRegistryAccess][%d] grantModelRegistryAccess default  %+v", o._statusCode, o.Payload)
}

func (o *GrantModelRegistryAccessDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GrantModelRegistryAccessDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}