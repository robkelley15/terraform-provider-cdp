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

// UpdateDatabaseReader is a Reader for the UpdateDatabase structure.
type UpdateDatabaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDatabaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDatabaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateDatabaseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDatabaseOK creates a UpdateDatabaseOK with default headers values
func NewUpdateDatabaseOK() *UpdateDatabaseOK {
	return &UpdateDatabaseOK{}
}

/*
UpdateDatabaseOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type UpdateDatabaseOK struct {
	Payload models.UpdateDatabaseResponse
}

// IsSuccess returns true when this update database o k response has a 2xx status code
func (o *UpdateDatabaseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update database o k response has a 3xx status code
func (o *UpdateDatabaseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update database o k response has a 4xx status code
func (o *UpdateDatabaseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update database o k response has a 5xx status code
func (o *UpdateDatabaseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update database o k response a status code equal to that given
func (o *UpdateDatabaseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update database o k response
func (o *UpdateDatabaseOK) Code() int {
	return 200
}

func (o *UpdateDatabaseOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/opdb/updateDatabase][%d] updateDatabaseOK  %+v", 200, o.Payload)
}

func (o *UpdateDatabaseOK) String() string {
	return fmt.Sprintf("[POST /api/v1/opdb/updateDatabase][%d] updateDatabaseOK  %+v", 200, o.Payload)
}

func (o *UpdateDatabaseOK) GetPayload() models.UpdateDatabaseResponse {
	return o.Payload
}

func (o *UpdateDatabaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatabaseDefault creates a UpdateDatabaseDefault with default headers values
func NewUpdateDatabaseDefault(code int) *UpdateDatabaseDefault {
	return &UpdateDatabaseDefault{
		_statusCode: code,
	}
}

/*
UpdateDatabaseDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type UpdateDatabaseDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update database default response has a 2xx status code
func (o *UpdateDatabaseDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update database default response has a 3xx status code
func (o *UpdateDatabaseDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update database default response has a 4xx status code
func (o *UpdateDatabaseDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update database default response has a 5xx status code
func (o *UpdateDatabaseDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update database default response a status code equal to that given
func (o *UpdateDatabaseDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update database default response
func (o *UpdateDatabaseDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDatabaseDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/opdb/updateDatabase][%d] updateDatabase default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDatabaseDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/opdb/updateDatabase][%d] updateDatabase default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDatabaseDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDatabaseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}