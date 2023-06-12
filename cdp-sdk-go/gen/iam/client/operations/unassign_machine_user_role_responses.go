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

// UnassignMachineUserRoleReader is a Reader for the UnassignMachineUserRole structure.
type UnassignMachineUserRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnassignMachineUserRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnassignMachineUserRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnassignMachineUserRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnassignMachineUserRoleOK creates a UnassignMachineUserRoleOK with default headers values
func NewUnassignMachineUserRoleOK() *UnassignMachineUserRoleOK {
	return &UnassignMachineUserRoleOK{}
}

/*
UnassignMachineUserRoleOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type UnassignMachineUserRoleOK struct {
	Payload models.UnassignMachineUserRoleResponse
}

// IsSuccess returns true when this unassign machine user role o k response has a 2xx status code
func (o *UnassignMachineUserRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unassign machine user role o k response has a 3xx status code
func (o *UnassignMachineUserRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unassign machine user role o k response has a 4xx status code
func (o *UnassignMachineUserRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unassign machine user role o k response has a 5xx status code
func (o *UnassignMachineUserRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unassign machine user role o k response a status code equal to that given
func (o *UnassignMachineUserRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the unassign machine user role o k response
func (o *UnassignMachineUserRoleOK) Code() int {
	return 200
}

func (o *UnassignMachineUserRoleOK) Error() string {
	return fmt.Sprintf("[POST /iam/unassignMachineUserRole][%d] unassignMachineUserRoleOK  %+v", 200, o.Payload)
}

func (o *UnassignMachineUserRoleOK) String() string {
	return fmt.Sprintf("[POST /iam/unassignMachineUserRole][%d] unassignMachineUserRoleOK  %+v", 200, o.Payload)
}

func (o *UnassignMachineUserRoleOK) GetPayload() models.UnassignMachineUserRoleResponse {
	return o.Payload
}

func (o *UnassignMachineUserRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnassignMachineUserRoleDefault creates a UnassignMachineUserRoleDefault with default headers values
func NewUnassignMachineUserRoleDefault(code int) *UnassignMachineUserRoleDefault {
	return &UnassignMachineUserRoleDefault{
		_statusCode: code,
	}
}

/*
UnassignMachineUserRoleDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type UnassignMachineUserRoleDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this unassign machine user role default response has a 2xx status code
func (o *UnassignMachineUserRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unassign machine user role default response has a 3xx status code
func (o *UnassignMachineUserRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unassign machine user role default response has a 4xx status code
func (o *UnassignMachineUserRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unassign machine user role default response has a 5xx status code
func (o *UnassignMachineUserRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unassign machine user role default response a status code equal to that given
func (o *UnassignMachineUserRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the unassign machine user role default response
func (o *UnassignMachineUserRoleDefault) Code() int {
	return o._statusCode
}

func (o *UnassignMachineUserRoleDefault) Error() string {
	return fmt.Sprintf("[POST /iam/unassignMachineUserRole][%d] unassignMachineUserRole default  %+v", o._statusCode, o.Payload)
}

func (o *UnassignMachineUserRoleDefault) String() string {
	return fmt.Sprintf("[POST /iam/unassignMachineUserRole][%d] unassignMachineUserRole default  %+v", o._statusCode, o.Payload)
}

func (o *UnassignMachineUserRoleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UnassignMachineUserRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}