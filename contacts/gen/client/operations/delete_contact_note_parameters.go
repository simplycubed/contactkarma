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

// NewDeleteContactNoteParams creates a new DeleteContactNoteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteContactNoteParams() *DeleteContactNoteParams {
	return &DeleteContactNoteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteContactNoteParamsWithTimeout creates a new DeleteContactNoteParams object
// with the ability to set a timeout on a request.
func NewDeleteContactNoteParamsWithTimeout(timeout time.Duration) *DeleteContactNoteParams {
	return &DeleteContactNoteParams{
		timeout: timeout,
	}
}

// NewDeleteContactNoteParamsWithContext creates a new DeleteContactNoteParams object
// with the ability to set a context for a request.
func NewDeleteContactNoteParamsWithContext(ctx context.Context) *DeleteContactNoteParams {
	return &DeleteContactNoteParams{
		Context: ctx,
	}
}

// NewDeleteContactNoteParamsWithHTTPClient creates a new DeleteContactNoteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteContactNoteParamsWithHTTPClient(client *http.Client) *DeleteContactNoteParams {
	return &DeleteContactNoteParams{
		HTTPClient: client,
	}
}

/* DeleteContactNoteParams contains all the parameters to send to the API endpoint
   for the delete contact note operation.

   Typically these are written to a http.Request.
*/
type DeleteContactNoteParams struct {

	// XApigatewayAPIUserinfo.
	XApigatewayAPIUserinfo string

	/* NoteID.

	   note id
	*/
	NoteID string

	/* UnifiedID.

	   contact id
	*/
	UnifiedID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete contact note params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContactNoteParams) WithDefaults() *DeleteContactNoteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete contact note params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContactNoteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete contact note params
func (o *DeleteContactNoteParams) WithTimeout(timeout time.Duration) *DeleteContactNoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete contact note params
func (o *DeleteContactNoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete contact note params
func (o *DeleteContactNoteParams) WithContext(ctx context.Context) *DeleteContactNoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete contact note params
func (o *DeleteContactNoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete contact note params
func (o *DeleteContactNoteParams) WithHTTPClient(client *http.Client) *DeleteContactNoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete contact note params
func (o *DeleteContactNoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXApigatewayAPIUserinfo adds the xApigatewayAPIUserinfo to the delete contact note params
func (o *DeleteContactNoteParams) WithXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) *DeleteContactNoteParams {
	o.SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo)
	return o
}

// SetXApigatewayAPIUserinfo adds the xApigatewayApiUserinfo to the delete contact note params
func (o *DeleteContactNoteParams) SetXApigatewayAPIUserinfo(xApigatewayAPIUserinfo string) {
	o.XApigatewayAPIUserinfo = xApigatewayAPIUserinfo
}

// WithNoteID adds the noteID to the delete contact note params
func (o *DeleteContactNoteParams) WithNoteID(noteID string) *DeleteContactNoteParams {
	o.SetNoteID(noteID)
	return o
}

// SetNoteID adds the noteId to the delete contact note params
func (o *DeleteContactNoteParams) SetNoteID(noteID string) {
	o.NoteID = noteID
}

// WithUnifiedID adds the unifiedID to the delete contact note params
func (o *DeleteContactNoteParams) WithUnifiedID(unifiedID string) *DeleteContactNoteParams {
	o.SetUnifiedID(unifiedID)
	return o
}

// SetUnifiedID adds the unifiedId to the delete contact note params
func (o *DeleteContactNoteParams) SetUnifiedID(unifiedID string) {
	o.UnifiedID = unifiedID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteContactNoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Apigateway-Api-Userinfo
	if err := r.SetHeaderParam("X-Apigateway-Api-Userinfo", o.XApigatewayAPIUserinfo); err != nil {
		return err
	}

	// path param note_id
	if err := r.SetPathParam("note_id", o.NoteID); err != nil {
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