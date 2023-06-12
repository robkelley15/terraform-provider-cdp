// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

// DescribeDatalakeReader is a Reader for the DescribeDatalake structure.
type DescribeDatalakeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeDatalakeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeDatalakeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeDatalakeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeDatalakeOK creates a DescribeDatalakeOK with default headers values
func NewDescribeDatalakeOK() *DescribeDatalakeOK {
	return &DescribeDatalakeOK{}
}

/*
DescribeDatalakeOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DescribeDatalakeOK struct {
	Payload *models.DescribeDatalakeResponse
}

// IsSuccess returns true when this describe datalake o k response has a 2xx status code
func (o *DescribeDatalakeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe datalake o k response has a 3xx status code
func (o *DescribeDatalakeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe datalake o k response has a 4xx status code
func (o *DescribeDatalakeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe datalake o k response has a 5xx status code
func (o *DescribeDatalakeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe datalake o k response a status code equal to that given
func (o *DescribeDatalakeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe datalake o k response
func (o *DescribeDatalakeOK) Code() int {
	return 200
}

func (o *DescribeDatalakeOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datalake/describeDatalake][%d] describeDatalakeOK  %+v", 200, o.Payload)
}

func (o *DescribeDatalakeOK) String() string {
	return fmt.Sprintf("[POST /api/v1/datalake/describeDatalake][%d] describeDatalakeOK  %+v", 200, o.Payload)
}

func (o *DescribeDatalakeOK) GetPayload() *models.DescribeDatalakeResponse {
	return o.Payload
}

func (o *DescribeDatalakeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeDatalakeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeDatalakeDefault creates a DescribeDatalakeDefault with default headers values
func NewDescribeDatalakeDefault(code int) *DescribeDatalakeDefault {
	return &DescribeDatalakeDefault{
		_statusCode: code,
	}
}

/*
DescribeDatalakeDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DescribeDatalakeDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this describe datalake default response has a 2xx status code
func (o *DescribeDatalakeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe datalake default response has a 3xx status code
func (o *DescribeDatalakeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe datalake default response has a 4xx status code
func (o *DescribeDatalakeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe datalake default response has a 5xx status code
func (o *DescribeDatalakeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe datalake default response a status code equal to that given
func (o *DescribeDatalakeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe datalake default response
func (o *DescribeDatalakeDefault) Code() int {
	return o._statusCode
}

func (o *DescribeDatalakeDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datalake/describeDatalake][%d] describeDatalake default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeDatalakeDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/datalake/describeDatalake][%d] describeDatalake default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeDatalakeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeDatalakeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}