// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

// PatchContactNoteOKCode is the HTTP code returned for type PatchContactNoteOK
const PatchContactNoteOKCode int = 200

/*PatchContactNoteOK Updated

swagger:response patchContactNoteOK
*/
type PatchContactNoteOK struct {

	/*
	  In: Body
	*/
	Payload *models.Note `json:"body,omitempty"`
}

// NewPatchContactNoteOK creates PatchContactNoteOK with default headers values
func NewPatchContactNoteOK() *PatchContactNoteOK {

	return &PatchContactNoteOK{}
}

// WithPayload adds the payload to the patch contact note o k response
func (o *PatchContactNoteOK) WithPayload(payload *models.Note) *PatchContactNoteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch contact note o k response
func (o *PatchContactNoteOK) SetPayload(payload *models.Note) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchContactNoteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchContactNoteBadRequestCode is the HTTP code returned for type PatchContactNoteBadRequest
const PatchContactNoteBadRequestCode int = 400

/*PatchContactNoteBadRequest Error

swagger:response patchContactNoteBadRequest
*/
type PatchContactNoteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPatchContactNoteBadRequest creates PatchContactNoteBadRequest with default headers values
func NewPatchContactNoteBadRequest() *PatchContactNoteBadRequest {

	return &PatchContactNoteBadRequest{}
}

// WithPayload adds the payload to the patch contact note bad request response
func (o *PatchContactNoteBadRequest) WithPayload(payload *models.ErrorResponse) *PatchContactNoteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch contact note bad request response
func (o *PatchContactNoteBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchContactNoteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchContactNoteUnauthorizedCode is the HTTP code returned for type PatchContactNoteUnauthorized
const PatchContactNoteUnauthorizedCode int = 401

/*PatchContactNoteUnauthorized Error

swagger:response patchContactNoteUnauthorized
*/
type PatchContactNoteUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPatchContactNoteUnauthorized creates PatchContactNoteUnauthorized with default headers values
func NewPatchContactNoteUnauthorized() *PatchContactNoteUnauthorized {

	return &PatchContactNoteUnauthorized{}
}

// WithPayload adds the payload to the patch contact note unauthorized response
func (o *PatchContactNoteUnauthorized) WithPayload(payload *models.ErrorResponse) *PatchContactNoteUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch contact note unauthorized response
func (o *PatchContactNoteUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchContactNoteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchContactNoteNotFoundCode is the HTTP code returned for type PatchContactNoteNotFound
const PatchContactNoteNotFoundCode int = 404

/*PatchContactNoteNotFound Error

swagger:response patchContactNoteNotFound
*/
type PatchContactNoteNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPatchContactNoteNotFound creates PatchContactNoteNotFound with default headers values
func NewPatchContactNoteNotFound() *PatchContactNoteNotFound {

	return &PatchContactNoteNotFound{}
}

// WithPayload adds the payload to the patch contact note not found response
func (o *PatchContactNoteNotFound) WithPayload(payload *models.ErrorResponse) *PatchContactNoteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch contact note not found response
func (o *PatchContactNoteNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchContactNoteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchContactNoteInternalServerErrorCode is the HTTP code returned for type PatchContactNoteInternalServerError
const PatchContactNoteInternalServerErrorCode int = 500

/*PatchContactNoteInternalServerError Internal Server Error

swagger:response patchContactNoteInternalServerError
*/
type PatchContactNoteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPatchContactNoteInternalServerError creates PatchContactNoteInternalServerError with default headers values
func NewPatchContactNoteInternalServerError() *PatchContactNoteInternalServerError {

	return &PatchContactNoteInternalServerError{}
}

// WithPayload adds the payload to the patch contact note internal server error response
func (o *PatchContactNoteInternalServerError) WithPayload(payload *models.ErrorResponse) *PatchContactNoteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch contact note internal server error response
func (o *PatchContactNoteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchContactNoteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
