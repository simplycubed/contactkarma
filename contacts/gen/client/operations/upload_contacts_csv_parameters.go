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

// NewUploadContactsCsvParams creates a new UploadContactsCsvParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadContactsCsvParams() *UploadContactsCsvParams {
	return &UploadContactsCsvParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadContactsCsvParamsWithTimeout creates a new UploadContactsCsvParams object
// with the ability to set a timeout on a request.
func NewUploadContactsCsvParamsWithTimeout(timeout time.Duration) *UploadContactsCsvParams {
	return &UploadContactsCsvParams{
		timeout: timeout,
	}
}

// NewUploadContactsCsvParamsWithContext creates a new UploadContactsCsvParams object
// with the ability to set a context for a request.
func NewUploadContactsCsvParamsWithContext(ctx context.Context) *UploadContactsCsvParams {
	return &UploadContactsCsvParams{
		Context: ctx,
	}
}

// NewUploadContactsCsvParamsWithHTTPClient creates a new UploadContactsCsvParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadContactsCsvParamsWithHTTPClient(client *http.Client) *UploadContactsCsvParams {
	return &UploadContactsCsvParams{
		HTTPClient: client,
	}
}

/* UploadContactsCsvParams contains all the parameters to send to the API endpoint
   for the upload contacts csv operation.

   Typically these are written to a http.Request.
*/
type UploadContactsCsvParams struct {

	// XApigatewayAPIUserinfo.
	XApigatewayAPIUserinfo string

	/* File.

	   The uploaded file data
	*/
	File runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload contacts csv params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadContactsCsvParams) WithDefaults() *UploadContactsCsvParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload contacts csv params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadContactsCsvParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload contacts csv params
func (o *UploadContactsCsvParams) WithTimeout(timeout time.Duration) *UploadContactsCsvParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload contacts csv params
func (o *UploadContactsCsvParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload contacts csv params
func (o *UploadContactsCsvParams) WithContext(ctx context.Context) *UploadContactsCsvParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload contacts csv params
func (o *UploadContactsCsvParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload contacts csv params
func (o *UploadContactsCsvParams) WithHTTPClient(client *http.Client) *UploadContactsCsvParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload contacts csv params
func (o *UploadContactsCsvParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXApigatewayAPIUserinfo adds the xApigatewayAPIUserinfo to the upload contacts csv params
func (o *UploadContactsCsvParams) WithXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) *UploadContactsCsvParams {
	o.SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo)
	return o
}

// SetXApigatewayAPIUserinfo adds the xApigatewayApiUserinfo to the upload contacts csv params
func (o *UploadContactsCsvParams) SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) {
	o.XApigatewayAPIUserinfo = xApigatewayAPIUserinfo
}

// WithFile adds the file to the upload contacts csv params
func (o *UploadContactsCsvParams) WithFile(file runtime.NamedReadCloser) *UploadContactsCsvParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the upload contacts csv params
func (o *UploadContactsCsvParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WriteToRequest writes these params to a swagger request
func (o *UploadContactsCsvParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Apigateway-Api-Userinfo
	if err := r.SetHeaderParam("X-Apigateway-Api-Userinfo", o.XApigatewayAPIUserinfo); err != nil {
		return err
	}
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
