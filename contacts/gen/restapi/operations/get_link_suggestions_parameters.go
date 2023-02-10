// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetLinkSuggestionsParams creates a new GetLinkSuggestionsParams object
//
// There are no default values defined in the spec.
func NewGetLinkSuggestionsParams() GetLinkSuggestionsParams {

	return GetLinkSuggestionsParams{}
}

// GetLinkSuggestionsParams contains all the bound params for the get link suggestions operation
// typically these are obtained from a http.Request
//
// swagger:parameters get-link-suggestions
type GetLinkSuggestionsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	XApigatewayAPIUserinfo string
	/*
	  In: query
	*/
	LastDocumentID *string
	/*
	  In: query
	*/
	Limit *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetLinkSuggestionsParams() beforehand.
func (o *GetLinkSuggestionsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXApigatewayAPIUserinfo(r.Header[http.CanonicalHeaderKey("X-Apigateway-Api-Userinfo")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	qLastDocumentID, qhkLastDocumentID, _ := qs.GetOK("last_document_id")
	if err := o.bindLastDocumentID(qLastDocumentID, qhkLastDocumentID, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXApigatewayAPIUserinfo binds and validates parameter XApigatewayAPIUserinfo from header.
func (o *GetLinkSuggestionsParams) bindXApigatewayAPIUserinfo(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindLastDocumentID binds and validates parameter LastDocumentID from query.
func (o *GetLinkSuggestionsParams) bindLastDocumentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.LastDocumentID = &raw

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetLinkSuggestionsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}
