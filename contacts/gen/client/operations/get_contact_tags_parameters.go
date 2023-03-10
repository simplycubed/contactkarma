// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetContactTagsParams creates a new GetContactTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContactTagsParams() *GetContactTagsParams {
	return &GetContactTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContactTagsParamsWithTimeout creates a new GetContactTagsParams object
// with the ability to set a timeout on a request.
func NewGetContactTagsParamsWithTimeout(timeout time.Duration) *GetContactTagsParams {
	return &GetContactTagsParams{
		timeout: timeout,
	}
}

// NewGetContactTagsParamsWithContext creates a new GetContactTagsParams object
// with the ability to set a context for a request.
func NewGetContactTagsParamsWithContext(ctx context.Context) *GetContactTagsParams {
	return &GetContactTagsParams{
		Context: ctx,
	}
}

// NewGetContactTagsParamsWithHTTPClient creates a new GetContactTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContactTagsParamsWithHTTPClient(client *http.Client) *GetContactTagsParams {
	return &GetContactTagsParams{
		HTTPClient: client,
	}
}

/* GetContactTagsParams contains all the parameters to send to the API endpoint
   for the get contact tags operation.

   Typically these are written to a http.Request.
*/
type GetContactTagsParams struct {

	// XApigatewayAPIUserinfo.
	XApigatewayAPIUserinfo string

	// LastDocumentID.
	LastDocumentID *string

	// Limit.
	Limit *int64

	/* UnifiedID.

	   contact id
	*/
	UnifiedID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get contact tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContactTagsParams) WithDefaults() *GetContactTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get contact tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContactTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get contact tags params
func (o *GetContactTagsParams) WithTimeout(timeout time.Duration) *GetContactTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contact tags params
func (o *GetContactTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contact tags params
func (o *GetContactTagsParams) WithContext(ctx context.Context) *GetContactTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contact tags params
func (o *GetContactTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contact tags params
func (o *GetContactTagsParams) WithHTTPClient(client *http.Client) *GetContactTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contact tags params
func (o *GetContactTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXApigatewayAPIUserinfo adds the xApigatewayAPIUserinfo to the get contact tags params
func (o *GetContactTagsParams) WithXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) *GetContactTagsParams {
	o.SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo)
	return o
}

// SetXApigatewayAPIUserinfo adds the xApigatewayApiUserinfo to the get contact tags params
func (o *GetContactTagsParams) SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) {
	o.XApigatewayAPIUserinfo = xApigatewayAPIUserinfo
}

// WithLastDocumentID adds the lastDocumentID to the get contact tags params
func (o *GetContactTagsParams) WithLastDocumentID(lastDocumentID *string) *GetContactTagsParams {
	o.SetLastDocumentID(lastDocumentID)
	return o
}

// SetLastDocumentID adds the lastDocumentId to the get contact tags params
func (o *GetContactTagsParams) SetLastDocumentID(lastDocumentID *string) {
	o.LastDocumentID = lastDocumentID
}

// WithLimit adds the limit to the get contact tags params
func (o *GetContactTagsParams) WithLimit(limit *int64) *GetContactTagsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get contact tags params
func (o *GetContactTagsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithUnifiedID adds the unifiedID to the get contact tags params
func (o *GetContactTagsParams) WithUnifiedID(unifiedID string) *GetContactTagsParams {
	o.SetUnifiedID(unifiedID)
	return o
}

// SetUnifiedID adds the unifiedId to the get contact tags params
func (o *GetContactTagsParams) SetUnifiedID(unifiedID string) {
	o.UnifiedID = unifiedID
}

// WriteToRequest writes these params to a swagger request
func (o *GetContactTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Apigateway-Api-Userinfo
	if err := r.SetHeaderParam("X-Apigateway-Api-Userinfo", o.XApigatewayAPIUserinfo); err != nil {
		return err
	}

	if o.LastDocumentID != nil {

		// query param last_document_id
		var qrLastDocumentID string

		if o.LastDocumentID != nil {
			qrLastDocumentID = *o.LastDocumentID
		}
		qLastDocumentID := qrLastDocumentID
		if qLastDocumentID != "" {

			if err := r.SetQueryParam("last_document_id", qLastDocumentID); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	// path param unified_id
	if err := r.SetPathParam("unified_id", o.UnifiedID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
