// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

// GetUnifiedContactsOKCode is the HTTP code returned for type GetUnifiedContactsOK
const GetUnifiedContactsOKCode int = 200

/*GetUnifiedContactsOK Retrieved

swagger:response getUnifiedContactsOK
*/
type GetUnifiedContactsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Unified `json:"body,omitempty"`
}

// NewGetUnifiedContactsOK creates GetUnifiedContactsOK with default headers values
func NewGetUnifiedContactsOK() *GetUnifiedContactsOK {

	return &GetUnifiedContactsOK{}
}

// WithPayload adds the payload to the get unified contacts o k response
func (o *GetUnifiedContactsOK) WithPayload(payload []*models.Unified) *GetUnifiedContactsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get unified contacts o k response
func (o *GetUnifiedContactsOK) SetPayload(payload []*models.Unified) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUnifiedContactsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetUnifiedContactsBadRequestCode is the HTTP code returned for type GetUnifiedContactsBadRequest
const GetUnifiedContactsBadRequestCode int = 400

/*GetUnifiedContactsBadRequest Error

swagger:response getUnifiedContactsBadRequest
*/
type GetUnifiedContactsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetUnifiedContactsBadRequest creates GetUnifiedContactsBadRequest with default headers values
func NewGetUnifiedContactsBadRequest() *GetUnifiedContactsBadRequest {

	return &GetUnifiedContactsBadRequest{}
}

// WithPayload adds the payload to the get unified contacts bad request response
func (o *GetUnifiedContactsBadRequest) WithPayload(payload *models.ErrorResponse) *GetUnifiedContactsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get unified contacts bad request response
func (o *GetUnifiedContactsBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUnifiedContactsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUnifiedContactsUnauthorizedCode is the HTTP code returned for type GetUnifiedContactsUnauthorized
const GetUnifiedContactsUnauthorizedCode int = 401

/*GetUnifiedContactsUnauthorized Error

swagger:response getUnifiedContactsUnauthorized
*/
type GetUnifiedContactsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetUnifiedContactsUnauthorized creates GetUnifiedContactsUnauthorized with default headers values
func NewGetUnifiedContactsUnauthorized() *GetUnifiedContactsUnauthorized {

	return &GetUnifiedContactsUnauthorized{}
}

// WithPayload adds the payload to the get unified contacts unauthorized response
func (o *GetUnifiedContactsUnauthorized) WithPayload(payload *models.ErrorResponse) *GetUnifiedContactsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get unified contacts unauthorized response
func (o *GetUnifiedContactsUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUnifiedContactsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUnifiedContactsInternalServerErrorCode is the HTTP code returned for type GetUnifiedContactsInternalServerError
const GetUnifiedContactsInternalServerErrorCode int = 500

/*GetUnifiedContactsInternalServerError Internal Server Error

swagger:response getUnifiedContactsInternalServerError
*/
type GetUnifiedContactsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetUnifiedContactsInternalServerError creates GetUnifiedContactsInternalServerError with default headers values
func NewGetUnifiedContactsInternalServerError() *GetUnifiedContactsInternalServerError {

	return &GetUnifiedContactsInternalServerError{}
}

// WithPayload adds the payload to the get unified contacts internal server error response
func (o *GetUnifiedContactsInternalServerError) WithPayload(payload *models.ErrorResponse) *GetUnifiedContactsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get unified contacts internal server error response
func (o *GetUnifiedContactsInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUnifiedContactsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
