// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewDeleteContactNoteParams creates a new DeleteContactNoteParams object
//
// There are no default values defined in the spec.
func NewDeleteContactNoteParams() DeleteContactNoteParams {

	return DeleteContactNoteParams{}
}

// DeleteContactNoteParams contains all the bound params for the delete contact note operation
// typically these are obtained from a http.Request
//
// swagger:parameters delete-contact-note
type DeleteContactNoteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	XApigatewayAPIUserinfo string
	/*note id
	  Required: true
	  In: path
	*/
	NoteID string
	/*contact id
	  Required: true
	  In: path
	*/
	UnifiedID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteContactNoteParams() beforehand.
func (o *DeleteContactNoteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindXApigatewayAPIUserinfo(r.Header[http.CanonicalHeaderKey("X-Apigateway-Api-Userinfo")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rNoteID, rhkNoteID, _ := route.Params.GetOK("note_id")
	if err := o.bindNoteID(rNoteID, rhkNoteID, route.Formats); err != nil {
		res = append(res, err)
	}

	rUnifiedID, rhkUnifiedID, _ := route.Params.GetOK("unified_id")
	if err := o.bindUnifiedID(rUnifiedID, rhkUnifiedID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXApigatewayAPIUserinfo binds and validates parameter XApigatewayAPIUserinfo from header.
func (o *DeleteContactNoteParams) bindXApigatewayAPIUserinfo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("X-Apigateway-Api-Userinfo", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("X-Apigateway-Api-Userinfo", "header", raw); err != nil {
		return err
	}
	o.XApigatewayAPIUserinfo = raw

	return nil
}

// bindNoteID binds and validates parameter NoteID from path.
func (o *DeleteContactNoteParams) bindNoteID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.NoteID = raw

	return nil
}

// bindUnifiedID binds and validates parameter UnifiedID from path.
func (o *DeleteContactNoteParams) bindUnifiedID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.UnifiedID = raw

	return nil
}