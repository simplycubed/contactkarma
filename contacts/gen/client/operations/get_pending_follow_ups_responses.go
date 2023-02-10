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

// GetPendingFollowUpsReader is a Reader for the GetPendingFollowUps structure.
type GetPendingFollowUpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPendingFollowUpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPendingFollowUpsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPendingFollowUpsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPendingFollowUpsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPendingFollowUpsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPendingFollowUpsOK creates a GetPendingFollowUpsOK with default headers values
func NewGetPendingFollowUpsOK() *GetPendingFollowUpsOK {
	return &GetPendingFollowUpsOK{}
}

/* GetPendingFollowUpsOK describes a response with status code 200, with default header values.

Retrieved
*/
type GetPendingFollowUpsOK struct {
	Payload []*models.Unified
}

func (o *GetPendingFollowUpsOK) Error() string {
	return fmt.Sprintf("[GET /pending][%d] getPendingFollowUpsOK  %+v", 200, o.Payload)
}
func (o *GetPendingFollowUpsOK) GetPayload() []*models.Unified {
	return o.Payload
}

func (o *GetPendingFollowUpsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPendingFollowUpsBadRequest creates a GetPendingFollowUpsBadRequest with default headers values
func NewGetPendingFollowUpsBadRequest() *GetPendingFollowUpsBadRequest {
	return &GetPendingFollowUpsBadRequest{}
}

/* GetPendingFollowUpsBadRequest describes a response with status code 400, with default header values.

Error
*/
type GetPendingFollowUpsBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetPendingFollowUpsBadRequest) Error() string {
	return fmt.Sprintf("[GET /pending][%d] getPendingFollowUpsBadRequest  %+v", 400, o.Payload)
}
func (o *GetPendingFollowUpsBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetPendingFollowUpsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPendingFollowUpsUnauthorized creates a GetPendingFollowUpsUnauthorized with default headers values
func NewGetPendingFollowUpsUnauthorized() *GetPendingFollowUpsUnauthorized {
	return &GetPendingFollowUpsUnauthorized{}
}

/* GetPendingFollowUpsUnauthorized describes a response with status code 401, with default header values.

Error
*/
type GetPendingFollowUpsUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GetPendingFollowUpsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pending][%d] getPendingFollowUpsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetPendingFollowUpsUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetPendingFollowUpsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPendingFollowUpsInternalServerError creates a GetPendingFollowUpsInternalServerError with default headers values
func NewGetPendingFollowUpsInternalServerError() *GetPendingFollowUpsInternalServerError {
	return &GetPendingFollowUpsInternalServerError{}
}

/* GetPendingFollowUpsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetPendingFollowUpsInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetPendingFollowUpsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pending][%d] getPendingFollowUpsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetPendingFollowUpsInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetPendingFollowUpsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
