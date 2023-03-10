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

// NewGetContactNotesParams creates a new GetContactNotesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContactNotesParams() *GetContactNotesParams {
	return &GetContactNotesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContactNotesParamsWithTimeout creates a new GetContactNotesParams object
// with the ability to set a timeout on a request.
func NewGetContactNotesParamsWithTimeout(timeout time.Duration) *GetContactNotesParams {
	return &GetContactNotesParams{
		timeout: timeout,
	}
}

// NewGetContactNotesParamsWithContext creates a new GetContactNotesParams object
// with the ability to set a context for a request.
func NewGetContactNotesParamsWithContext(ctx context.Context) *GetContactNotesParams {
	return &GetContactNotesParams{
		Context: ctx,
	}
}

// NewGetContactNotesParamsWithHTTPClient creates a new GetContactNotesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContactNotesParamsWithHTTPClient(client *http.Client) *GetContactNotesParams {
	return &GetContactNotesParams{
		HTTPClient: client,
	}
}

/* GetContactNotesParams contains all the parameters to send to the API endpoint
   for the get contact notes operation.

   Typically these are written to a http.Request.
*/
type GetContactNotesParams struct {

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

// WithDefaults hydrates default values in the get contact notes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContactNotesParams) WithDefaults() *GetContactNotesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get contact notes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContactNotesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get contact notes params
func (o *GetContactNotesParams) WithTimeout(timeout time.Duration) *GetContactNotesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contact notes params
func (o *GetContactNotesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contact notes params
func (o *GetContactNotesParams) WithContext(ctx context.Context) *GetContactNotesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contact notes params
func (o *GetContactNotesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contact notes params
func (o *GetContactNotesParams) WithHTTPClient(client *http.Client) *GetContactNotesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contact notes params
func (o *GetContactNotesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXApigatewayAPIUserinfo adds the xApigatewayAPIUserinfo to the get contact notes params
func (o *GetContactNotesParams) WithXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) *GetContactNotesParams {
	o.SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo)
	return o
}

// SetXApigatewayAPIUserinfo adds the xApigatewayApiUserinfo to the get contact notes params
func (o *GetContactNotesParams) SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) {
	o.XApigatewayAPIUserinfo = xApigatewayAPIUserinfo
}

// WithLastDocumentID adds the lastDocumentID to the get contact notes params
func (o *GetContactNotesParams) WithLastDocumentID(lastDocumentID *string) *GetContactNotesParams {
	o.SetLastDocumentID(lastDocumentID)
	return o
}

// SetLastDocumentID adds the lastDocumentId to the get contact notes params
func (o *GetContactNotesParams) SetLastDocumentID(lastDocumentID *string) {
	o.LastDocumentID = lastDocumentID
}

// WithLimit adds the limit to the get contact notes params
func (o *GetContactNotesParams) WithLimit(limit *int64) *GetContactNotesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get contact notes params
func (o *GetContactNotesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithUnifiedID adds the unifiedID to the get contact notes params
func (o *GetContactNotesParams) WithUnifiedID(unifiedID string) *GetContactNotesParams {
	o.SetUnifiedID(unifiedID)
	return o
}

// SetUnifiedID adds the unifiedId to the get contact notes params
func (o *GetContactNotesParams) SetUnifiedID(unifiedID string) {
	o.UnifiedID = unifiedID
}

// WriteToRequest writes these params to a swagger request
func (o *GetContactNotesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
