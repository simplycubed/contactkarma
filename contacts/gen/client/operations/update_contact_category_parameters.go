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

// NewUpdateContactCategoryParams creates a new UpdateContactCategoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateContactCategoryParams() *UpdateContactCategoryParams {
	return &UpdateContactCategoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateContactCategoryParamsWithTimeout creates a new UpdateContactCategoryParams object
// with the ability to set a timeout on a request.
func NewUpdateContactCategoryParamsWithTimeout(timeout time.Duration) *UpdateContactCategoryParams {
	return &UpdateContactCategoryParams{
		timeout: timeout,
	}
}

// NewUpdateContactCategoryParamsWithContext creates a new UpdateContactCategoryParams object
// with the ability to set a context for a request.
func NewUpdateContactCategoryParamsWithContext(ctx context.Context) *UpdateContactCategoryParams {
	return &UpdateContactCategoryParams{
		Context: ctx,
	}
}

// NewUpdateContactCategoryParamsWithHTTPClient creates a new UpdateContactCategoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateContactCategoryParamsWithHTTPClient(client *http.Client) *UpdateContactCategoryParams {
	return &UpdateContactCategoryParams{
		HTTPClient: client,
	}
}

/* UpdateContactCategoryParams contains all the parameters to send to the API endpoint
   for the update contact category operation.

   Typically these are written to a http.Request.
*/
type UpdateContactCategoryParams struct {

	// XApigatewayAPIUserinfo.
	XApigatewayAPIUserinfo string

	// Body.
	Body *models.UpdateCategoryDto

	/* UnifiedID.

	   contact id
	*/
	UnifiedID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update contact category params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateContactCategoryParams) WithDefaults() *UpdateContactCategoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update contact category params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateContactCategoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update contact category params
func (o *UpdateContactCategoryParams) WithTimeout(timeout time.Duration) *UpdateContactCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update contact category params
func (o *UpdateContactCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update contact category params
func (o *UpdateContactCategoryParams) WithContext(ctx context.Context) *UpdateContactCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update contact category params
func (o *UpdateContactCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update contact category params
func (o *UpdateContactCategoryParams) WithHTTPClient(client *http.Client) *UpdateContactCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update contact category params
func (o *UpdateContactCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXApigatewayAPIUserinfo adds the xApigatewayAPIUserinfo to the update contact category params
func (o *UpdateContactCategoryParams) WithXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) *UpdateContactCategoryParams {
	o.SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo)
	return o
}

// SetXApigatewayAPIUserinfo adds the xApigatewayApiUserinfo to the update contact category params
func (o *UpdateContactCategoryParams) SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) {
	o.XApigatewayAPIUserinfo = xApigatewayAPIUserinfo
}

// WithBody adds the body to the update contact category params
func (o *UpdateContactCategoryParams) WithBody(body *models.UpdateCategoryDto) *UpdateContactCategoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update contact category params
func (o *UpdateContactCategoryParams) SetBody(body *models.UpdateCategoryDto) {
	o.Body = body
}

// WithUnifiedID adds the unifiedID to the update contact category params
func (o *UpdateContactCategoryParams) WithUnifiedID(unifiedID string) *UpdateContactCategoryParams {
	o.SetUnifiedID(unifiedID)
	return o
}

// SetUnifiedID adds the unifiedId to the update contact category params
func (o *UpdateContactCategoryParams) SetUnifiedID(unifiedID string) {
	o.UnifiedID = unifiedID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateContactCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
