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

// PatchContactTagReader is a Reader for the PatchContactTag structure.
type PatchContactTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchContactTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchContactTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchContactTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchContactTagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchContactTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchContactTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchContactTagOK creates a PatchContactTagOK with default headers values
func NewPatchContactTagOK() *PatchContactTagOK {
	return &PatchContactTagOK{}
}

/* PatchContactTagOK describes a response with status code 200, with default header values.

Updated
*/
type PatchContactTagOK struct {
	Payload *models.Tag
}

func (o *PatchContactTagOK) Error() string {
	return fmt.Sprintf("[PATCH /unified/{unified_id}/tags/{tag_id}][%d] patchContactTagOK  %+v", 200, o.Payload)
}
func (o *PatchContactTagOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *PatchContactTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchContactTagBadRequest creates a PatchContactTagBadRequest with default headers values
func NewPatchContactTagBadRequest() *PatchContactTagBadRequest {
	return &PatchContactTagBadRequest{}
}

/* PatchContactTagBadRequest describes a response with status code 400, with default header values.

Error
*/
type PatchContactTagBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PatchContactTagBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /unified/{unified_id}/tags/{tag_id}][%d] patchContactTagBadRequest  %+v", 400, o.Payload)
}
func (o *PatchContactTagBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchContactTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchContactTagUnauthorized creates a PatchContactTagUnauthorized with default headers values
func NewPatchContactTagUnauthorized() *PatchContactTagUnauthorized {
	return &PatchContactTagUnauthorized{}
}

/* PatchContactTagUnauthorized describes a response with status code 401, with default header values.

Error
*/
type PatchContactTagUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *PatchContactTagUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /unified/{unified_id}/tags/{tag_id}][%d] patchContactTagUnauthorized  %+v", 401, o.Payload)
}
func (o *PatchContactTagUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchContactTagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchContactTagNotFound creates a PatchContactTagNotFound with default headers values
func NewPatchContactTagNotFound() *PatchContactTagNotFound {
	return &PatchContactTagNotFound{}
}

/* PatchContactTagNotFound describes a response with status code 404, with default header values.

Error
*/
type PatchContactTagNotFound struct {
	Payload *models.ErrorResponse
}

func (o *PatchContactTagNotFound) Error() string {
	return fmt.Sprintf("[PATCH /unified/{unified_id}/tags/{tag_id}][%d] patchContactTagNotFound  %+v", 404, o.Payload)
}
func (o *PatchContactTagNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchContactTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchContactTagInternalServerError creates a PatchContactTagInternalServerError with default headers values
func NewPatchContactTagInternalServerError() *PatchContactTagInternalServerError {
	return &PatchContactTagInternalServerError{}
}

/* PatchContactTagInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PatchContactTagInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PatchContactTagInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /unified/{unified_id}/tags/{tag_id}][%d] patchContactTagInternalServerError  %+v", 500, o.Payload)
}
func (o *PatchContactTagInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchContactTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}