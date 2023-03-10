// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteContactSourceHandlerFunc turns a function with the right signature into a delete contact source handler
type DeleteContactSourceHandlerFunc func(DeleteContactSourceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteContactSourceHandlerFunc) Handle(params DeleteContactSourceParams) middleware.Responder {
	return fn(params)
}

// DeleteContactSourceHandler interface for that can handle valid delete contact source params
type DeleteContactSourceHandler interface {
	Handle(DeleteContactSourceParams) middleware.Responder
}

// NewDeleteContactSource creates a new http.Handler for the delete contact source operation
func NewDeleteContactSource(ctx *middleware.Context, handler DeleteContactSourceHandler) *DeleteContactSource {
	return &DeleteContactSource{Context: ctx, Handler: handler}
}

/* DeleteContactSource swagger:route DELETE /contacts/sources/{source_id} deleteContactSource

delete contact source by id

*/
type DeleteContactSource struct {
	Context *middleware.Context
	Handler DeleteContactSourceHandler
}

func (o *DeleteContactSource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteContactSourceParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteContactSourceBody delete contact source body
//
// swagger:model DeleteContactSourceBody
type DeleteContactSourceBody struct {

	// flag to indicate whether contacts should be removed from unified as well
	RemoveFromUnified bool `json:"remove_from_unified,omitempty"`
}

// Validate validates this delete contact source body
func (o *DeleteContactSourceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete contact source body based on context it is used
func (o *DeleteContactSourceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteContactSourceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteContactSourceBody) UnmarshalBinary(b []byte) error {
	var res DeleteContactSourceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
