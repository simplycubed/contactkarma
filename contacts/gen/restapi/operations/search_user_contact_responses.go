// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

// SearchUserContactOKCode is the HTTP code returned for type SearchUserContactOK
const SearchUserContactOKCode int = 200

/*SearchUserContactOK Created

swagger:response searchUserContactOK
*/
type SearchUserContactOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Unified `json:"body,omitempty"`
}

// NewSearchUserContactOK creates SearchUserContactOK with default headers values
func NewSearchUserContactOK() *SearchUserContactOK {

	return &SearchUserContactOK{}
}

// WithPayload adds the payload to the search user contact o k response
func (o *SearchUserContactOK) WithPayload(payload []*models.Unified) *SearchUserContactOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search user contact o k response
func (o *SearchUserContactOK) SetPayload(payload []*models.Unified) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchUserContactOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Unified, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// SearchUserContactBadRequestCode is the HTTP code returned for type SearchUserContactBadRequest
const SearchUserContactBadRequestCode int = 400

/*SearchUserContactBadRequest Error

swagger:response searchUserContactBadRequest
*/
type SearchUserContactBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSearchUserContactBadRequest creates SearchUserContactBadRequest with default headers values
func NewSearchUserContactBadRequest() *SearchUserContactBadRequest {

	return &SearchUserContactBadRequest{}
}

// WithPayload adds the payload to the search user contact bad request response
func (o *SearchUserContactBadRequest) WithPayload(payload *models.ErrorResponse) *SearchUserContactBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search user contact bad request response
func (o *SearchUserContactBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchUserContactBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SearchUserContactUnauthorizedCode is the HTTP code returned for type SearchUserContactUnauthorized
const SearchUserContactUnauthorizedCode int = 401

/*SearchUserContactUnauthorized Error

swagger:response searchUserContactUnauthorized
*/
type SearchUserContactUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSearchUserContactUnauthorized creates SearchUserContactUnauthorized with default headers values
func NewSearchUserContactUnauthorized() *SearchUserContactUnauthorized {

	return &SearchUserContactUnauthorized{}
}

// WithPayload adds the payload to the search user contact unauthorized response
func (o *SearchUserContactUnauthorized) WithPayload(payload *models.ErrorResponse) *SearchUserContactUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search user contact unauthorized response
func (o *SearchUserContactUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchUserContactUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SearchUserContactInternalServerErrorCode is the HTTP code returned for type SearchUserContactInternalServerError
const SearchUserContactInternalServerErrorCode int = 500

/*SearchUserContactInternalServerError Internal Server Error

swagger:response searchUserContactInternalServerError
*/
type SearchUserContactInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSearchUserContactInternalServerError creates SearchUserContactInternalServerError with default headers values
func NewSearchUserContactInternalServerError() *SearchUserContactInternalServerError {

	return &SearchUserContactInternalServerError{}
}

// WithPayload adds the payload to the search user contact internal server error response
func (o *SearchUserContactInternalServerError) WithPayload(payload *models.ErrorResponse) *SearchUserContactInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search user contact internal server error response
func (o *SearchUserContactInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchUserContactInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
