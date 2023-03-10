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

// PostContactNoteReader is a Reader for the PostContactNote structure.
type PostContactNoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostContactNoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostContactNoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostContactNoteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostContactNoteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostContactNoteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostContactNoteOK creates a PostContactNoteOK with default headers values
func NewPostContactNoteOK() *PostContactNoteOK {
	return &PostContactNoteOK{}
}

/* PostContactNoteOK describes a response with status code 200, with default header values.

Created
*/
type PostContactNoteOK struct {
	Payload *models.Message
}

func (o *PostContactNoteOK) Error() string {
	return fmt.Sprintf("[POST /unified/{unified_id}/notes][%d] postContactNoteOK  %+v", 200, o.Payload)
}
func (o *PostContactNoteOK) GetPayload() *models.Message {
	return o.Payload
}

func (o *PostContactNoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContactNoteBadRequest creates a PostContactNoteBadRequest with default headers values
func NewPostContactNoteBadRequest() *PostContactNoteBadRequest {
	return &PostContactNoteBadRequest{}
}

/* PostContactNoteBadRequest describes a response with status code 400, with default header values.

Error
*/
type PostContactNoteBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PostContactNoteBadRequest) Error() string {
	return fmt.Sprintf("[POST /unified/{unified_id}/notes][%d] postContactNoteBadRequest  %+v", 400, o.Payload)
}
func (o *PostContactNoteBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostContactNoteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContactNoteUnauthorized creates a PostContactNoteUnauthorized with default headers values
func NewPostContactNoteUnauthorized() *PostContactNoteUnauthorized {
	return &PostContactNoteUnauthorized{}
}

/* PostContactNoteUnauthorized describes a response with status code 401, with default header values.

Error
*/
type PostContactNoteUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *PostContactNoteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /unified/{unified_id}/notes][%d] postContactNoteUnauthorized  %+v", 401, o.Payload)
}
func (o *PostContactNoteUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostContactNoteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContactNoteInternalServerError creates a PostContactNoteInternalServerError with default headers values
func NewPostContactNoteInternalServerError() *PostContactNoteInternalServerError {
	return &PostContactNoteInternalServerError{}
}

/* PostContactNoteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostContactNoteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PostContactNoteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /unified/{unified_id}/notes][%d] postContactNoteInternalServerError  %+v", 500, o.Payload)
}
func (o *PostContactNoteInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostContactNoteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
