// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

// CreateUserContactReader is a Reader for the CreateUserContact structure.
type CreateUserContactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserContactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUserContactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserContactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateUserContactUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUserContactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUserContactOK creates a CreateUserContactOK with default headers values
func NewCreateUserContactOK() *CreateUserContactOK {
	return &CreateUserContactOK{}
}

/* CreateUserContactOK describes a response with status code 200, with default header values.

Created
*/
type CreateUserContactOK struct {
	Payload *models.Unified
}

func (o *CreateUserContactOK) Error() string {
	return fmt.Sprintf("[POST /unified][%d] createUserContactOK  %+v", 200, o.Payload)
}
func (o *CreateUserContactOK) GetPayload() *models.Unified {
	return o.Payload
}

func (o *CreateUserContactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unified)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserContactBadRequest creates a CreateUserContactBadRequest with default headers values
func NewCreateUserContactBadRequest() *CreateUserContactBadRequest {
	return &CreateUserContactBadRequest{}
}

/* CreateUserContactBadRequest describes a response with status code 400, with default header values.

Error
*/
type CreateUserContactBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CreateUserContactBadRequest) Error() string {
	return fmt.Sprintf("[POST /unified][%d] createUserContactBadRequest  %+v", 400, o.Payload)
}
func (o *CreateUserContactBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserContactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserContactUnauthorized creates a CreateUserContactUnauthorized with default headers values
func NewCreateUserContactUnauthorized() *CreateUserContactUnauthorized {
	return &CreateUserContactUnauthorized{}
}

/* CreateUserContactUnauthorized describes a response with status code 401, with default header values.

Error
*/
type CreateUserContactUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CreateUserContactUnauthorized) Error() string {
	return fmt.Sprintf("[POST /unified][%d] createUserContactUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateUserContactUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserContactUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserContactInternalServerError creates a CreateUserContactInternalServerError with default headers values
func NewCreateUserContactInternalServerError() *CreateUserContactInternalServerError {
	return &CreateUserContactInternalServerError{}
}

/* CreateUserContactInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateUserContactInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CreateUserContactInternalServerError) Error() string {
	return fmt.Sprintf("[POST /unified][%d] createUserContactInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateUserContactInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserContactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
