// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
)

// GetUpgradeVwVersionsReader is a Reader for the GetUpgradeVwVersions structure.
type GetUpgradeVwVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUpgradeVwVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUpgradeVwVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUpgradeVwVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUpgradeVwVersionsOK creates a GetUpgradeVwVersionsOK with default headers values
func NewGetUpgradeVwVersionsOK() *GetUpgradeVwVersionsOK {
	return &GetUpgradeVwVersionsOK{}
}

/*
GetUpgradeVwVersionsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type GetUpgradeVwVersionsOK struct {
	Payload *models.GetUpgradeVwVersionsResponse
}

// IsSuccess returns true when this get upgrade vw versions o k response has a 2xx status code
func (o *GetUpgradeVwVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get upgrade vw versions o k response has a 3xx status code
func (o *GetUpgradeVwVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upgrade vw versions o k response has a 4xx status code
func (o *GetUpgradeVwVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get upgrade vw versions o k response has a 5xx status code
func (o *GetUpgradeVwVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get upgrade vw versions o k response a status code equal to that given
func (o *GetUpgradeVwVersionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get upgrade vw versions o k response
func (o *GetUpgradeVwVersionsOK) Code() int {
	return 200
}

func (o *GetUpgradeVwVersionsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/getUpgradeVwVersions][%d] getUpgradeVwVersionsOK  %+v", 200, o.Payload)
}

func (o *GetUpgradeVwVersionsOK) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/getUpgradeVwVersions][%d] getUpgradeVwVersionsOK  %+v", 200, o.Payload)
}

func (o *GetUpgradeVwVersionsOK) GetPayload() *models.GetUpgradeVwVersionsResponse {
	return o.Payload
}

func (o *GetUpgradeVwVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUpgradeVwVersionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUpgradeVwVersionsDefault creates a GetUpgradeVwVersionsDefault with default headers values
func NewGetUpgradeVwVersionsDefault(code int) *GetUpgradeVwVersionsDefault {
	return &GetUpgradeVwVersionsDefault{
		_statusCode: code,
	}
}

/*
GetUpgradeVwVersionsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetUpgradeVwVersionsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get upgrade vw versions default response has a 2xx status code
func (o *GetUpgradeVwVersionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get upgrade vw versions default response has a 3xx status code
func (o *GetUpgradeVwVersionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get upgrade vw versions default response has a 4xx status code
func (o *GetUpgradeVwVersionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get upgrade vw versions default response has a 5xx status code
func (o *GetUpgradeVwVersionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get upgrade vw versions default response a status code equal to that given
func (o *GetUpgradeVwVersionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get upgrade vw versions default response
func (o *GetUpgradeVwVersionsDefault) Code() int {
	return o._statusCode
}

func (o *GetUpgradeVwVersionsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/getUpgradeVwVersions][%d] getUpgradeVwVersions default  %+v", o._statusCode, o.Payload)
}

func (o *GetUpgradeVwVersionsDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/getUpgradeVwVersions][%d] getUpgradeVwVersions default  %+v", o._statusCode, o.Payload)
}

func (o *GetUpgradeVwVersionsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUpgradeVwVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}