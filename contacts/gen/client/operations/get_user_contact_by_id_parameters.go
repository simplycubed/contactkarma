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
)

// NewGetUserContactByIDParams creates a new GetUserContactByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserContactByIDParams() *GetUserContactByIDParams {
	return &GetUserContactByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserContactByIDParamsWithTimeout creates a new GetUserContactByIDParams object
// with the ability to set a timeout on a request.
func NewGetUserContactByIDParamsWithTimeout(timeout time.Duration) *GetUserContactByIDParams {
	return &GetUserContactByIDParams{
		timeout: timeout,
	}
}

// NewGetUserContactByIDParamsWithContext creates a new GetUserContactByIDParams object
// with the ability to set a context for a request.
func NewGetUserContactByIDParamsWithContext(ctx context.Context) *GetUserContactByIDParams {
	return &GetUserContactByIDParams{
		Context: ctx,
	}
}

// NewGetUserContactByIDParamsWithHTTPClient creates a new GetUserContactByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserContactByIDParamsWithHTTPClient(client *http.Client) *GetUserContactByIDParams {
	return &GetUserContactByIDParams{
		HTTPClient: client,
	}
}

/* GetUserContactByIDParams contains all the parameters to send to the API endpoint
   for the get user contact by id operation.

   Typically these are written to a http.Request.
*/
type GetUserContactByIDParams struct {

	// XApigatewayAPIUserinfo.
	XApigatewayAPIUserinfo string

	/* UnifiedID.

	   contact id
	*/
	UnifiedID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user contact by id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserContactByIDParams) WithDefaults() *GetUserContactByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user contact by id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserContactByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user contact by id params
func (o *GetUserContactByIDParams) WithTimeout(timeout time.Duration) *GetUserContactByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user contact by id params
func (o *GetUserContactByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user contact by id params
func (o *GetUserContactByIDParams) WithContext(ctx context.Context) *GetUserContactByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user contact by id params
func (o *GetUserContactByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user contact by id params
func (o *GetUserContactByIDParams) WithHTTPClient(client *http.Client) *GetUserContactByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user contact by id params
func (o *GetUserContactByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXApigatewayAPIUserinfo adds the xApigatewayAPIUserinfo to the get user contact by id params
func (o *GetUserContactByIDParams) WithXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) *GetUserContactByIDParams {
	o.SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo)
	return o
}

// SetXApigatewayAPIUserinfo adds the xApigatewayApiUserinfo to the get user contact by id params
func (o *GetUserContactByIDParams) SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) {
	o.XApigatewayAPIUserinfo = xApigatewayAPIUserinfo
}

// WithUnifiedID adds the unifiedID to the get user contact by id params
func (o *GetUserContactByIDParams) WithUnifiedID(unifiedID string) *GetUserContactByIDParams {
	o.SetUnifiedID(unifiedID)
	return o
}

// SetUnifiedID adds the unifiedId to the get user contact by id params
func (o *GetUserContactByIDParams) SetUnifiedID(unifiedID string) {
	o.UnifiedID = unifiedID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserContactByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Apigateway-Api-Userinfo
	if err := r.SetHeaderParam("X-Apigateway-Api-Userinfo", o.XApigatewayAPIUserinfo); err != nil {
		return err
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
