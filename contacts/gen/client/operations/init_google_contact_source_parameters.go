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

// NewInitGoogleContactSourceParams creates a new InitGoogleContactSourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitGoogleContactSourceParams() *InitGoogleContactSourceParams {
	return &InitGoogleContactSourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitGoogleContactSourceParamsWithTimeout creates a new InitGoogleContactSourceParams object
// with the ability to set a timeout on a request.
func NewInitGoogleContactSourceParamsWithTimeout(timeout time.Duration) *InitGoogleContactSourceParams {
	return &InitGoogleContactSourceParams{
		timeout: timeout,
	}
}

// NewInitGoogleContactSourceParamsWithContext creates a new InitGoogleContactSourceParams object
// with the ability to set a context for a request.
func NewInitGoogleContactSourceParamsWithContext(ctx context.Context) *InitGoogleContactSourceParams {
	return &InitGoogleContactSourceParams{
		Context: ctx,
	}
}

// NewInitGoogleContactSourceParamsWithHTTPClient creates a new InitGoogleContactSourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitGoogleContactSourceParamsWithHTTPClient(client *http.Client) *InitGoogleContactSourceParams {
	return &InitGoogleContactSourceParams{
		HTTPClient: client,
	}
}

/* InitGoogleContactSourceParams contains all the parameters to send to the API endpoint
   for the init google contact source operation.

   Typically these are written to a http.Request.
*/
type InitGoogleContactSourceParams struct {

	// XApigatewayAPIUserinfo.
	XApigatewayAPIUserinfo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the init google contact source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitGoogleContactSourceParams) WithDefaults() *InitGoogleContactSourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the init google contact source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitGoogleContactSourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the init google contact source params
func (o *InitGoogleContactSourceParams) WithTimeout(timeout time.Duration) *InitGoogleContactSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the init google contact source params
func (o *InitGoogleContactSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the init google contact source params
func (o *InitGoogleContactSourceParams) WithContext(ctx context.Context) *InitGoogleContactSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the init google contact source params
func (o *InitGoogleContactSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the init google contact source params
func (o *InitGoogleContactSourceParams) WithHTTPClient(client *http.Client) *InitGoogleContactSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the init google contact source params
func (o *InitGoogleContactSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXApigatewayAPIUserinfo adds the xApigatewayAPIUserinfo to the init google contact source params
func (o *InitGoogleContactSourceParams) WithXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) *InitGoogleContactSourceParams {
	o.SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo)
	return o
}

// SetXApigatewayAPIUserinfo adds the xApigatewayApiUserinfo to the init google contact source params
func (o *InitGoogleContactSourceParams) SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) {
	o.XApigatewayAPIUserinfo = xApigatewayAPIUserinfo
}

// WriteToRequest writes these params to a swagger request
func (o *InitGoogleContactSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Apigateway-Api-Userinfo
	if err := r.SetHeaderParam("X-Apigateway-Api-Userinfo", o.XApigatewayAPIUserinfo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
