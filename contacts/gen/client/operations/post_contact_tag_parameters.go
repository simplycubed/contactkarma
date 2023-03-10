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

	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

// NewPostContactTagParams creates a new PostContactTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostContactTagParams() *PostContactTagParams {
	return &PostContactTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostContactTagParamsWithTimeout creates a new PostContactTagParams object
// with the ability to set a timeout on a request.
func NewPostContactTagParamsWithTimeout(timeout time.Duration) *PostContactTagParams {
	return &PostContactTagParams{
		timeout: timeout,
	}
}

// NewPostContactTagParamsWithContext creates a new PostContactTagParams object
// with the ability to set a context for a request.
func NewPostContactTagParamsWithContext(ctx context.Context) *PostContactTagParams {
	return &PostContactTagParams{
		Context: ctx,
	}
}

// NewPostContactTagParamsWithHTTPClient creates a new PostContactTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostContactTagParamsWithHTTPClient(client *http.Client) *PostContactTagParams {
	return &PostContactTagParams{
		HTTPClient: client,
	}
}

/* PostContactTagParams contains all the parameters to send to the API endpoint
   for the post contact tag operation.

   Typically these are written to a http.Request.
*/
type PostContactTagParams struct {

	// XApigatewayAPIUserinfo.
	XApigatewayAPIUserinfo string

	// Body.
	Body *models.Tag

	/* UnifiedID.

	   contact id
	*/
	UnifiedID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post contact tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostContactTagParams) WithDefaults() *PostContactTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post contact tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostContactTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post contact tag params
func (o *PostContactTagParams) WithTimeout(timeout time.Duration) *PostContactTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post contact tag params
func (o *PostContactTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post contact tag params
func (o *PostContactTagParams) WithContext(ctx context.Context) *PostContactTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post contact tag params
func (o *PostContactTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post contact tag params
func (o *PostContactTagParams) WithHTTPClient(client *http.Client) *PostContactTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post contact tag params
func (o *PostContactTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXApigatewayAPIUserinfo adds the xApigatewayAPIUserinfo to the post contact tag params
func (o *PostContactTagParams) WithXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) *PostContactTagParams {
	o.SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo)
	return o
}

// SetXApigatewayAPIUserinfo adds the xApigatewayApiUserinfo to the post contact tag params
func (o *PostContactTagParams) SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) {
	o.XApigatewayAPIUserinfo = xApigatewayAPIUserinfo
}

// WithBody adds the body to the post contact tag params
func (o *PostContactTagParams) WithBody(body *models.Tag) *PostContactTagParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post contact tag params
func (o *PostContactTagParams) SetBody(body *models.Tag) {
	o.Body = body
}

// WithUnifiedID adds the unifiedID to the post contact tag params
func (o *PostContactTagParams) WithUnifiedID(unifiedID string) *PostContactTagParams {
	o.SetUnifiedID(unifiedID)
	return o
}

// SetUnifiedID adds the unifiedId to the post contact tag params
func (o *PostContactTagParams) SetUnifiedID(unifiedID string) {
	o.UnifiedID = unifiedID
}

// WriteToRequest writes these params to a swagger request
func (o *PostContactTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Apigateway-Api-Userinfo
	if err := r.SetHeaderParam("X-Apigateway-Api-Userinfo", o.XApigatewayAPIUserinfo); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
