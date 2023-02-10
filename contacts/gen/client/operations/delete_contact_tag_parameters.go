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

// NewDeleteContactTagParams creates a new DeleteContactTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteContactTagParams() *DeleteContactTagParams {
	return &DeleteContactTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteContactTagParamsWithTimeout creates a new DeleteContactTagParams object
// with the ability to set a timeout on a request.
func NewDeleteContactTagParamsWithTimeout(timeout time.Duration) *DeleteContactTagParams {
	return &DeleteContactTagParams{
		timeout: timeout,
	}
}

// NewDeleteContactTagParamsWithContext creates a new DeleteContactTagParams object
// with the ability to set a context for a request.
func NewDeleteContactTagParamsWithContext(ctx context.Context) *DeleteContactTagParams {
	return &DeleteContactTagParams{
		Context: ctx,
	}
}

// NewDeleteContactTagParamsWithHTTPClient creates a new DeleteContactTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteContactTagParamsWithHTTPClient(client *http.Client) *DeleteContactTagParams {
	return &DeleteContactTagParams{
		HTTPClient: client,
	}
}

/* DeleteContactTagParams contains all the parameters to send to the API endpoint
   for the delete contact tag operation.

   Typically these are written to a http.Request.
*/
type DeleteContactTagParams struct {

	// XApigatewayAPIUserinfo.
	XApigatewayAPIUserinfo string

	/* TagID.

	   tag id
	*/
	TagID string

	/* UnifiedID.

	   contact id
	*/
	UnifiedID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete contact tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContactTagParams) WithDefaults() *DeleteContactTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete contact tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContactTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete contact tag params
func (o *DeleteContactTagParams) WithTimeout(timeout time.Duration) *DeleteContactTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete contact tag params
func (o *DeleteContactTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete contact tag params
func (o *DeleteContactTagParams) WithContext(ctx context.Context) *DeleteContactTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete contact tag params
func (o *DeleteContactTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete contact tag params
func (o *DeleteContactTagParams) WithHTTPClient(client *http.Client) *DeleteContactTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete contact tag params
func (o *DeleteContactTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXApigatewayAPIUserinfo adds the xApigatewayAPIUserinfo to the delete contact tag params
func (o *DeleteContactTagParams) WithXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) *DeleteContactTagParams {
	o.SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo)
	return o
}

// SetXApigatewayAPIUserinfo adds the xApigatewayApiUserinfo to the delete contact tag params
func (o *DeleteContactTagParams) SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) {
	o.XApigatewayAPIUserinfo = xApigatewayAPIUserinfo
}

// WithTagID adds the tagID to the delete contact tag params
func (o *DeleteContactTagParams) WithTagID(tagID string) *DeleteContactTagParams {
	o.SetTagID(tagID)
	return o
}

// SetTagID adds the tagId to the delete contact tag params
func (o *DeleteContactTagParams) SetTagID(tagID string) {
	o.TagID = tagID
}

// WithUnifiedID adds the unifiedID to the delete contact tag params
func (o *DeleteContactTagParams) WithUnifiedID(unifiedID string) *DeleteContactTagParams {
	o.SetUnifiedID(unifiedID)
	return o
}

// SetUnifiedID adds the unifiedId to the delete contact tag params
func (o *DeleteContactTagParams) SetUnifiedID(unifiedID string) {
	o.UnifiedID = unifiedID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteContactTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Apigateway-Api-Userinfo
	if err := r.SetHeaderParam("X-Apigateway-Api-Userinfo", o.XApigatewayAPIUserinfo); err != nil {
		return err
	}

	// path param tag_id
	if err := r.SetPathParam("tag_id", o.TagID); err != nil {
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