// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

// UpdateUserContactOKCode is the HTTP code returned for type UpdateUserContactOK
const UpdateUserContactOKCode int = 200

/*UpdateUserContactOK Updated

swagger:response updateUserContactOK
*/
type UpdateUserContactOK struct {

	/*
	  In: Body
	*/
	Payload *models.Unified `json:"body,omitempty"`
}

// NewUpdateUserContactOK creates UpdateUserContactOK with default headers values
func NewUpdateUserContactOK() *UpdateUserContactOK {

	return &UpdateUserContactOK{}
}

// WithPayload adds the payload to the update user contact o k response
func (o *UpdateUserContactOK) WithPayload(payload *models.Unified) *UpdateUserContactOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user contact o k response
func (o *UpdateUserContactOK) SetPayload(payload *models.Unified) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserContactOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateUserContactBadRequestCode is the HTTP code returned for type UpdateUserContactBadRequest
const UpdateUserContactBadRequestCode int = 400

/*UpdateUserContactBadRequest Error

swagger:response updateUserContactBadRequest
*/
type UpdateUserContactBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewUpdateUserContactBadRequest creates UpdateUserContactBadRequest with default headers values
func NewUpdateUserContactBadRequest() *UpdateUserContactBadRequest {

	return &UpdateUserContactBadRequest{}
}

// WithPayload adds the payload to the update user contact bad request response
func (o *UpdateUserContactBadRequest) WithPayload(payload *models.ErrorResponse) *UpdateUserContactBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user contact bad request response
func (o *UpdateUserContactBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserContactBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateUserContactUnauthorizedCode is the HTTP code returned for type UpdateUserContactUnauthorized
const UpdateUserContactUnauthorizedCode int = 401

/*UpdateUserContactUnauthorized Error

swagger:response updateUserContactUnauthorized
*/
type UpdateUserContactUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewUpdateUserContactUnauthorized creates UpdateUserContactUnauthorized with default headers values
func NewUpdateUserContactUnauthorized() *UpdateUserContactUnauthorized {

	return &UpdateUserContactUnauthorized{}
}

// WithPayload adds the payload to the update user contact unauthorized response
func (o *UpdateUserContactUnauthorized) WithPayload(payload *models.ErrorResponse) *UpdateUserContactUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user contact unauthorized response
func (o *UpdateUserContactUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserContactUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateUserContactNotFoundCode is the HTTP code returned for type UpdateUserContactNotFound
const UpdateUserContactNotFoundCode int = 404

/*UpdateUserContactNotFound Error

swagger:response updateUserContactNotFound
*/
type UpdateUserContactNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewUpdateUserContactNotFound creates UpdateUserContactNotFound with default headers values
func NewUpdateUserContactNotFound() *UpdateUserContactNotFound {

	return &UpdateUserContactNotFound{}
}

// WithPayload adds the payload to the update user contact not found response
func (o *UpdateUserContactNotFound) WithPayload(payload *models.ErrorResponse) *UpdateUserContactNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user contact not found response
func (o *UpdateUserContactNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserContactNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateUserContactInternalServerErrorCode is the HTTP code returned for type UpdateUserContactInternalServerError
const UpdateUserContactInternalServerErrorCode int = 500

/*UpdateUserContactInternalServerError Internal Server Error

swagger:response updateUserContactInternalServerError
*/
type UpdateUserContactInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewUpdateUserContactInternalServerError creates UpdateUserContactInternalServerError with default headers values
func NewUpdateUserContactInternalServerError() *UpdateUserContactInternalServerError {

	return &UpdateUserContactInternalServerError{}
}

// WithPayload adds the payload to the update user contact internal server error response
func (o *UpdateUserContactInternalServerError) WithPayload(payload *models.ErrorResponse) *UpdateUserContactInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user contact internal server error response
func (o *UpdateUserContactInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserContactInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
