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

// CreateSamlProviderReader is a Reader for the CreateSamlProvider structure.
type CreateSamlProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSamlProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSamlProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateSamlProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSamlProviderOK creates a CreateSamlProviderOK with default headers values
func NewCreateSamlProviderOK() *CreateSamlProviderOK {
	return &CreateSamlProviderOK{}
}

/*
CreateSamlProviderOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type CreateSamlProviderOK struct {
	Payload *models.CreateSamlProviderResponse
}

// IsSuccess returns true when this create saml provider o k response has a 2xx status code
func (o *CreateSamlProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create saml provider o k response has a 3xx status code
func (o *CreateSamlProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create saml provider o k response has a 4xx status code
func (o *CreateSamlProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create saml provider o k response has a 5xx status code
func (o *CreateSamlProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create saml provider o k response a status code equal to that given
func (o *CreateSamlProviderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create saml provider o k response
func (o *CreateSamlProviderOK) Code() int {
	return 200
}

func (o *CreateSamlProviderOK) Error() string {
	return fmt.Sprintf("[POST /iam/createSamlProvider][%d] createSamlProviderOK  %+v", 200, o.Payload)
}

func (o *CreateSamlProviderOK) String() string {
	return fmt.Sprintf("[POST /iam/createSamlProvider][%d] createSamlProviderOK  %+v", 200, o.Payload)
}

func (o *CreateSamlProviderOK) GetPayload() *models.CreateSamlProviderResponse {
	return o.Payload
}

func (o *CreateSamlProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateSamlProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSamlProviderDefault creates a CreateSamlProviderDefault with default headers values
func NewCreateSamlProviderDefault(code int) *CreateSamlProviderDefault {
	return &CreateSamlProviderDefault{
		_statusCode: code,
	}
}

/*
CreateSamlProviderDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type CreateSamlProviderDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create saml provider default response has a 2xx status code
func (o *CreateSamlProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create saml provider default response has a 3xx status code
func (o *CreateSamlProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create saml provider default response has a 4xx status code
func (o *CreateSamlProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create saml provider default response has a 5xx status code
func (o *CreateSamlProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create saml provider default response a status code equal to that given
func (o *CreateSamlProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create saml provider default response
func (o *CreateSamlProviderDefault) Code() int {
	return o._statusCode
}

func (o *CreateSamlProviderDefault) Error() string {
	return fmt.Sprintf("[POST /iam/createSamlProvider][%d] createSamlProvider default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSamlProviderDefault) String() string {
	return fmt.Sprintf("[POST /iam/createSamlProvider][%d] createSamlProvider default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSamlProviderDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSamlProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}