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

// NewGetLinkSuggestionsParams creates a new GetLinkSuggestionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLinkSuggestionsParams() *GetLinkSuggestionsParams {
	return &GetLinkSuggestionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLinkSuggestionsParamsWithTimeout creates a new GetLinkSuggestionsParams object
// with the ability to set a timeout on a request.
func NewGetLinkSuggestionsParamsWithTimeout(timeout time.Duration) *GetLinkSuggestionsParams {
	return &GetLinkSuggestionsParams{
		timeout: timeout,
	}
}

// NewGetLinkSuggestionsParamsWithContext creates a new GetLinkSuggestionsParams object
// with the ability to set a context for a request.
func NewGetLinkSuggestionsParamsWithContext(ctx context.Context) *GetLinkSuggestionsParams {
	return &GetLinkSuggestionsParams{
		Context: ctx,
	}
}

// NewGetLinkSuggestionsParamsWithHTTPClient creates a new GetLinkSuggestionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLinkSuggestionsParamsWithHTTPClient(client *http.Client) *GetLinkSuggestionsParams {
	return &GetLinkSuggestionsParams{
		HTTPClient: client,
	}
}

/* GetLinkSuggestionsParams contains all the parameters to send to the API endpoint
   for the get link suggestions operation.

   Typically these are written to a http.Request.
*/
type GetLinkSuggestionsParams struct {

	// XApigatewayAPIUserinfo.
	XApigatewayAPIUserinfo string

	// LastDocumentID.
	LastDocumentID *string

	// Limit.
	Limit *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get link suggestions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLinkSuggestionsParams) WithDefaults() *GetLinkSuggestionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get link suggestions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLinkSuggestionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get link suggestions params
func (o *GetLinkSuggestionsParams) WithTimeout(timeout time.Duration) *GetLinkSuggestionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get link suggestions params
func (o *GetLinkSuggestionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get link suggestions params
func (o *GetLinkSuggestionsParams) WithContext(ctx context.Context) *GetLinkSuggestionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get link suggestions params
func (o *GetLinkSuggestionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get link suggestions params
func (o *GetLinkSuggestionsParams) WithHTTPClient(client *http.Client) *GetLinkSuggestionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get link suggestions params
func (o *GetLinkSuggestionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXApigatewayAPIUserinfo adds the xApigatewayAPIUserinfo to the get link suggestions params
func (o *GetLinkSuggestionsParams) WithXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) *GetLinkSuggestionsParams {
	o.SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo)
	return o
}

// SetXApigatewayAPIUserinfo adds the xApigatewayApiUserinfo to the get link suggestions params
func (o *GetLinkSuggestionsParams) SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) {
	o.XApigatewayAPIUserinfo = xApigatewayAPIUserinfo
}

// WithLastDocumentID adds the lastDocumentID to the get link suggestions params
func (o *GetLinkSuggestionsParams) WithLastDocumentID(lastDocumentID *string) *GetLinkSuggestionsParams {
	o.SetLastDocumentID(lastDocumentID)
	return o
}

// SetLastDocumentID adds the lastDocumentId to the get link suggestions params
func (o *GetLinkSuggestionsParams) SetLastDocumentID(lastDocumentID *string) {
	o.LastDocumentID = lastDocumentID
}

// WithLimit adds the limit to the get link suggestions params
func (o *GetLinkSuggestionsParams) WithLimit(limit *int64) *GetLinkSuggestionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get link suggestions params
func (o *GetLinkSuggestionsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WriteToRequest writes these params to a swagger request
func (o *GetLinkSuggestionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}