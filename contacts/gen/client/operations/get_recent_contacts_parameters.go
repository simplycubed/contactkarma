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

// NewGetRecentContactsParams creates a new GetRecentContactsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRecentContactsParams() *GetRecentContactsParams {
	return &GetRecentContactsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecentContactsParamsWithTimeout creates a new GetRecentContactsParams object
// with the ability to set a timeout on a request.
func NewGetRecentContactsParamsWithTimeout(timeout time.Duration) *GetRecentContactsParams {
	return &GetRecentContactsParams{
		timeout: timeout,
	}
}

// NewGetRecentContactsParamsWithContext creates a new GetRecentContactsParams object
// with the ability to set a context for a request.
func NewGetRecentContactsParamsWithContext(ctx context.Context) *GetRecentContactsParams {
	return &GetRecentContactsParams{
		Context: ctx,
	}
}

// NewGetRecentContactsParamsWithHTTPClient creates a new GetRecentContactsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRecentContactsParamsWithHTTPClient(client *http.Client) *GetRecentContactsParams {
	return &GetRecentContactsParams{
		HTTPClient: client,
	}
}

/* GetRecentContactsParams contains all the parameters to send to the API endpoint
   for the get recent contacts operation.

   Typically these are written to a http.Request.
*/
type GetRecentContactsParams struct {

	// XApigatewayAPIUserinfo.
	XApigatewayAPIUserinfo string

	// LastDocumentID.
	LastDocumentID *string

	// LastDocumentLastContact.
	//
	// Format: date-time
	LastDocumentLastContact *strfmt.DateTime

	// Limit.
	Limit *int64

	// MaxDays.
	MaxDays *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get recent contacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRecentContactsParams) WithDefaults() *GetRecentContactsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get recent contacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRecentContactsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get recent contacts params
func (o *GetRecentContactsParams) WithTimeout(timeout time.Duration) *GetRecentContactsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recent contacts params
func (o *GetRecentContactsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recent contacts params
func (o *GetRecentContactsParams) WithContext(ctx context.Context) *GetRecentContactsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recent contacts params
func (o *GetRecentContactsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recent contacts params
func (o *GetRecentContactsParams) WithHTTPClient(client *http.Client) *GetRecentContactsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recent contacts params
func (o *GetRecentContactsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXApigatewayAPIUserinfo adds the xApigatewayAPIUserinfo to the get recent contacts params
func (o *GetRecentContactsParams) WithXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) *GetRecentContactsParams {
	o.SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo)
	return o
}

// SetXApigatewayAPIUserinfo adds the xApigatewayApiUserinfo to the get recent contacts params
func (o *GetRecentContactsParams) SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) {
	o.XApigatewayAPIUserinfo = xApigatewayAPIUserinfo
}

// WithLastDocumentID adds the lastDocumentID to the get recent contacts params
func (o *GetRecentContactsParams) WithLastDocumentID(lastDocumentID *string) *GetRecentContactsParams {
	o.SetLastDocumentID(lastDocumentID)
	return o
}

// SetLastDocumentID adds the lastDocumentId to the get recent contacts params
func (o *GetRecentContactsParams) SetLastDocumentID(lastDocumentID *string) {
	o.LastDocumentID = lastDocumentID
}

// WithLastDocumentLastContact adds the lastDocumentLastContact to the get recent contacts params
func (o *GetRecentContactsParams) WithLastDocumentLastContact(lastDocumentLastContact *strfmt.DateTime) *GetRecentContactsParams {
	o.SetLastDocumentLastContact(lastDocumentLastContact)
	return o
}

// SetLastDocumentLastContact adds the lastDocumentLastContact to the get recent contacts params
func (o *GetRecentContactsParams) SetLastDocumentLastContact(lastDocumentLastContact *strfmt.DateTime) {
	o.LastDocumentLastContact = lastDocumentLastContact
}

// WithLimit adds the limit to the get recent contacts params
func (o *GetRecentContactsParams) WithLimit(limit *int64) *GetRecentContactsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get recent contacts params
func (o *GetRecentContactsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaxDays adds the maxDays to the get recent contacts params
func (o *GetRecentContactsParams) WithMaxDays(maxDays *int64) *GetRecentContactsParams {
	o.SetMaxDays(maxDays)
	return o
}

// SetMaxDays adds the maxDays to the get recent contacts params
func (o *GetRecentContactsParams) SetMaxDays(maxDays *int64) {
	o.MaxDays = maxDays
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecentContactsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.LastDocumentLastContact != nil {

		// query param last_document_last_contact
		var qrLastDocumentLastContact strfmt.DateTime

		if o.LastDocumentLastContact != nil {
			qrLastDocumentLastContact = *o.LastDocumentLastContact
		}
		qLastDocumentLastContact := qrLastDocumentLastContact.String()
		if qLastDocumentLastContact != "" {

			if err := r.SetQueryParam("last_document_last_contact", qLastDocumentLastContact); err != nil {
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

	if o.MaxDays != nil {

		// query param max_days
		var qrMaxDays int64

		if o.MaxDays != nil {
			qrMaxDays = *o.MaxDays
		}
		qMaxDays := swag.FormatInt64(qrMaxDays)
		if qMaxDays != "" {

			if err := r.SetQueryParam("max_days", qMaxDays); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
