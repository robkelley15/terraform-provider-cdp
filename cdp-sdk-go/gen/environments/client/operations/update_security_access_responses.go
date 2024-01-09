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

// UpdateSecurityAccessReader is a Reader for the UpdateSecurityAccess structure.
type UpdateSecurityAccessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSecurityAccessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSecurityAccessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateSecurityAccessDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSecurityAccessOK creates a UpdateSecurityAccessOK with default headers values
func NewUpdateSecurityAccessOK() *UpdateSecurityAccessOK {
	return &UpdateSecurityAccessOK{}
}

/*
UpdateSecurityAccessOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type UpdateSecurityAccessOK struct {
	Payload *models.UpdateSecurityAccessResponse
}

// IsSuccess returns true when this update security access o k response has a 2xx status code
func (o *UpdateSecurityAccessOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update security access o k response has a 3xx status code
func (o *UpdateSecurityAccessOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update security access o k response has a 4xx status code
func (o *UpdateSecurityAccessOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update security access o k response has a 5xx status code
func (o *UpdateSecurityAccessOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update security access o k response a status code equal to that given
func (o *UpdateSecurityAccessOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update security access o k response
func (o *UpdateSecurityAccessOK) Code() int {
	return 200
}

func (o *UpdateSecurityAccessOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/updateSecurityAccess][%d] updateSecurityAccessOK  %+v", 200, o.Payload)
}

func (o *UpdateSecurityAccessOK) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/updateSecurityAccess][%d] updateSecurityAccessOK  %+v", 200, o.Payload)
}

func (o *UpdateSecurityAccessOK) GetPayload() *models.UpdateSecurityAccessResponse {
	return o.Payload
}

func (o *UpdateSecurityAccessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateSecurityAccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecurityAccessDefault creates a UpdateSecurityAccessDefault with default headers values
func NewUpdateSecurityAccessDefault(code int) *UpdateSecurityAccessDefault {
	return &UpdateSecurityAccessDefault{
		_statusCode: code,
	}
}

/*
UpdateSecurityAccessDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type UpdateSecurityAccessDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update security access default response has a 2xx status code
func (o *UpdateSecurityAccessDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update security access default response has a 3xx status code
func (o *UpdateSecurityAccessDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update security access default response has a 4xx status code
func (o *UpdateSecurityAccessDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update security access default response has a 5xx status code
func (o *UpdateSecurityAccessDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update security access default response a status code equal to that given
func (o *UpdateSecurityAccessDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update security access default response
func (o *UpdateSecurityAccessDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSecurityAccessDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/updateSecurityAccess][%d] updateSecurityAccess default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSecurityAccessDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/updateSecurityAccess][%d] updateSecurityAccess default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSecurityAccessDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSecurityAccessDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}