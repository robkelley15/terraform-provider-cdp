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

// StartFreeIpaVerticalScalingReader is a Reader for the StartFreeIpaVerticalScaling structure.
type StartFreeIpaVerticalScalingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartFreeIpaVerticalScalingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartFreeIpaVerticalScalingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartFreeIpaVerticalScalingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartFreeIpaVerticalScalingOK creates a StartFreeIpaVerticalScalingOK with default headers values
func NewStartFreeIpaVerticalScalingOK() *StartFreeIpaVerticalScalingOK {
	return &StartFreeIpaVerticalScalingOK{}
}

/*
StartFreeIpaVerticalScalingOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type StartFreeIpaVerticalScalingOK struct {
	Payload *models.StartFreeIpaVerticalScalingResponse
}

// IsSuccess returns true when this start free ipa vertical scaling o k response has a 2xx status code
func (o *StartFreeIpaVerticalScalingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start free ipa vertical scaling o k response has a 3xx status code
func (o *StartFreeIpaVerticalScalingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start free ipa vertical scaling o k response has a 4xx status code
func (o *StartFreeIpaVerticalScalingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start free ipa vertical scaling o k response has a 5xx status code
func (o *StartFreeIpaVerticalScalingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start free ipa vertical scaling o k response a status code equal to that given
func (o *StartFreeIpaVerticalScalingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the start free ipa vertical scaling o k response
func (o *StartFreeIpaVerticalScalingOK) Code() int {
	return 200
}

func (o *StartFreeIpaVerticalScalingOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/startFreeIpaVerticalScaling][%d] startFreeIpaVerticalScalingOK  %+v", 200, o.Payload)
}

func (o *StartFreeIpaVerticalScalingOK) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/startFreeIpaVerticalScaling][%d] startFreeIpaVerticalScalingOK  %+v", 200, o.Payload)
}

func (o *StartFreeIpaVerticalScalingOK) GetPayload() *models.StartFreeIpaVerticalScalingResponse {
	return o.Payload
}

func (o *StartFreeIpaVerticalScalingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StartFreeIpaVerticalScalingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartFreeIpaVerticalScalingDefault creates a StartFreeIpaVerticalScalingDefault with default headers values
func NewStartFreeIpaVerticalScalingDefault(code int) *StartFreeIpaVerticalScalingDefault {
	return &StartFreeIpaVerticalScalingDefault{
		_statusCode: code,
	}
}

/*
StartFreeIpaVerticalScalingDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type StartFreeIpaVerticalScalingDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this start free ipa vertical scaling default response has a 2xx status code
func (o *StartFreeIpaVerticalScalingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this start free ipa vertical scaling default response has a 3xx status code
func (o *StartFreeIpaVerticalScalingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this start free ipa vertical scaling default response has a 4xx status code
func (o *StartFreeIpaVerticalScalingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this start free ipa vertical scaling default response has a 5xx status code
func (o *StartFreeIpaVerticalScalingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this start free ipa vertical scaling default response a status code equal to that given
func (o *StartFreeIpaVerticalScalingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the start free ipa vertical scaling default response
func (o *StartFreeIpaVerticalScalingDefault) Code() int {
	return o._statusCode
}

func (o *StartFreeIpaVerticalScalingDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/startFreeIpaVerticalScaling][%d] startFreeIpaVerticalScaling default  %+v", o._statusCode, o.Payload)
}

func (o *StartFreeIpaVerticalScalingDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/startFreeIpaVerticalScaling][%d] startFreeIpaVerticalScaling default  %+v", o._statusCode, o.Payload)
}

func (o *StartFreeIpaVerticalScalingDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *StartFreeIpaVerticalScalingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}