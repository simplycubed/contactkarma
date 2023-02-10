// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

// DeleteContactNoteOKCode is the HTTP code returned for type DeleteContactNoteOK
const DeleteContactNoteOKCode int = 200

/*DeleteContactNoteOK Deleted

swagger:response deleteContactNoteOK
*/
type DeleteContactNoteOK struct {

	/*
	  In: Body
	*/
	Payload *models.Message `json:"body,omitempty"`
}

// NewDeleteContactNoteOK creates DeleteContactNoteOK with default headers values
func NewDeleteContactNoteOK() *DeleteContactNoteOK {

	return &DeleteContactNoteOK{}
}

// WithPayload adds the payload to the delete contact note o k response
func (o *DeleteContactNoteOK) WithPayload(payload *models.Message) *DeleteContactNoteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete contact note o k response
func (o *DeleteContactNoteOK) SetPayload(payload *models.Message) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteContactNoteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteContactNoteBadRequestCode is the HTTP code returned for type DeleteContactNoteBadRequest
const DeleteContactNoteBadRequestCode int = 400

/*DeleteContactNoteBadRequest Error

swagger:response deleteContactNoteBadRequest
*/
type DeleteContactNoteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteContactNoteBadRequest creates DeleteContactNoteBadRequest with default headers values
func NewDeleteContactNoteBadRequest() *DeleteContactNoteBadRequest {

	return &DeleteContactNoteBadRequest{}
}

// WithPayload adds the payload to the delete contact note bad request response
func (o *DeleteContactNoteBadRequest) WithPayload(payload *models.ErrorResponse) *DeleteContactNoteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete contact note bad request response
func (o *DeleteContactNoteBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteContactNoteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteContactNoteUnauthorizedCode is the HTTP code returned for type DeleteContactNoteUnauthorized
const DeleteContactNoteUnauthorizedCode int = 401

/*DeleteContactNoteUnauthorized Error

swagger:response deleteContactNoteUnauthorized
*/
type DeleteContactNoteUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteContactNoteUnauthorized creates DeleteContactNoteUnauthorized with default headers values
func NewDeleteContactNoteUnauthorized() *DeleteContactNoteUnauthorized {

	return &DeleteContactNoteUnauthorized{}
}

// WithPayload adds the payload to the delete contact note unauthorized response
func (o *DeleteContactNoteUnauthorized) WithPayload(payload *models.ErrorResponse) *DeleteContactNoteUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete contact note unauthorized response
func (o *DeleteContactNoteUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteContactNoteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteContactNoteInternalServerErrorCode is the HTTP code returned for type DeleteContactNoteInternalServerError
const DeleteContactNoteInternalServerErrorCode int = 500

/*DeleteContactNoteInternalServerError Internal Server Error

swagger:response deleteContactNoteInternalServerError
*/
type DeleteContactNoteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteContactNoteInternalServerError creates DeleteContactNoteInternalServerError with default headers values
func NewDeleteContactNoteInternalServerError() *DeleteContactNoteInternalServerError {

	return &DeleteContactNoteInternalServerError{}
}

// WithPayload adds the payload to the delete contact note internal server error response
func (o *DeleteContactNoteInternalServerError) WithPayload(payload *models.ErrorResponse) *DeleteContactNoteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete contact note internal server error response
func (o *DeleteContactNoteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteContactNoteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}