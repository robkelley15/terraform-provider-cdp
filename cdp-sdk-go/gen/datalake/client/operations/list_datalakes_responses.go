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

// ListDatalakesReader is a Reader for the ListDatalakes structure.
type ListDatalakesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDatalakesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDatalakesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListDatalakesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDatalakesOK creates a ListDatalakesOK with default headers values
func NewListDatalakesOK() *ListDatalakesOK {
	return &ListDatalakesOK{}
}

/*
ListDatalakesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListDatalakesOK struct {
	Payload *models.ListDatalakesResponse
}

// IsSuccess returns true when this list datalakes o k response has a 2xx status code
func (o *ListDatalakesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list datalakes o k response has a 3xx status code
func (o *ListDatalakesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list datalakes o k response has a 4xx status code
func (o *ListDatalakesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list datalakes o k response has a 5xx status code
func (o *ListDatalakesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list datalakes o k response a status code equal to that given
func (o *ListDatalakesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list datalakes o k response
func (o *ListDatalakesOK) Code() int {
	return 200
}

func (o *ListDatalakesOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datalake/listDatalakes][%d] listDatalakesOK  %+v", 200, o.Payload)
}

func (o *ListDatalakesOK) String() string {
	return fmt.Sprintf("[POST /api/v1/datalake/listDatalakes][%d] listDatalakesOK  %+v", 200, o.Payload)
}

func (o *ListDatalakesOK) GetPayload() *models.ListDatalakesResponse {
	return o.Payload
}

func (o *ListDatalakesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListDatalakesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDatalakesDefault creates a ListDatalakesDefault with default headers values
func NewListDatalakesDefault(code int) *ListDatalakesDefault {
	return &ListDatalakesDefault{
		_statusCode: code,
	}
}

/*
ListDatalakesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListDatalakesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list datalakes default response has a 2xx status code
func (o *ListDatalakesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list datalakes default response has a 3xx status code
func (o *ListDatalakesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list datalakes default response has a 4xx status code
func (o *ListDatalakesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list datalakes default response has a 5xx status code
func (o *ListDatalakesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list datalakes default response a status code equal to that given
func (o *ListDatalakesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list datalakes default response
func (o *ListDatalakesDefault) Code() int {
	return o._statusCode
}

func (o *ListDatalakesDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datalake/listDatalakes][%d] listDatalakes default  %+v", o._statusCode, o.Payload)
}

func (o *ListDatalakesDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/datalake/listDatalakes][%d] listDatalakes default  %+v", o._statusCode, o.Payload)
}

func (o *ListDatalakesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListDatalakesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}