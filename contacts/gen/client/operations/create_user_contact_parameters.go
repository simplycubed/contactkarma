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

// NewCreateUserContactParams creates a new CreateUserContactParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUserContactParams() *CreateUserContactParams {
	return &CreateUserContactParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserContactParamsWithTimeout creates a new CreateUserContactParams object
// with the ability to set a timeout on a request.
func NewCreateUserContactParamsWithTimeout(timeout time.Duration) *CreateUserContactParams {
	return &CreateUserContactParams{
		timeout: timeout,
	}
}

// NewCreateUserContactParamsWithContext creates a new CreateUserContactParams object
// with the ability to set a context for a request.
func NewCreateUserContactParamsWithContext(ctx context.Context) *CreateUserContactParams {
	return &CreateUserContactParams{
		Context: ctx,
	}
}

// NewCreateUserContactParamsWithHTTPClient creates a new CreateUserContactParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUserContactParamsWithHTTPClient(client *http.Client) *CreateUserContactParams {
	return &CreateUserContactParams{
		HTTPClient: client,
	}
}

/* CreateUserContactParams contains all the parameters to send to the API endpoint
   for the create user contact operation.

   Typically these are written to a http.Request.
*/
type CreateUserContactParams struct {

	// XApigatewayAPIUserinfo.
	XApigatewayAPIUserinfo string

	// Body.
	Body *models.CreateContactDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create user contact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserContactParams) WithDefaults() *CreateUserContactParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create user contact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserContactParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create user contact params
func (o *CreateUserContactParams) WithTimeout(timeout time.Duration) *CreateUserContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create user contact params
func (o *CreateUserContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create user contact params
func (o *CreateUserContactParams) WithContext(ctx context.Context) *CreateUserContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create user contact params
func (o *CreateUserContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create user contact params
func (o *CreateUserContactParams) WithHTTPClient(client *http.Client) *CreateUserContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create user contact params
func (o *CreateUserContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXApigatewayAPIUserinfo adds the xApigatewayAPIUserinfo to the create user contact params
func (o *CreateUserContactParams) WithXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) *CreateUserContactParams {
	o.SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo)
	return o
}

// SetXApigatewayAPIUserinfo adds the xApigatewayApiUserinfo to the create user contact params
func (o *CreateUserContactParams) SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) {
	o.XApigatewayAPIUserinfo = xApigatewayAPIUserinfo
}

// WithBody adds the body to the create user contact params
func (o *CreateUserContactParams) WithBody(body *models.CreateContactDto) *CreateUserContactParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create user contact params
func (o *CreateUserContactParams) SetBody(body *models.CreateContactDto) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
