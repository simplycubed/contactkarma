// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteContactTagHandlerFunc turns a function with the right signature into a delete contact tag handler
type DeleteContactTagHandlerFunc func(DeleteContactTagParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteContactTagHandlerFunc) Handle(params DeleteContactTagParams) middleware.Responder {
	return fn(params)
}

// DeleteContactTagHandler interface for that can handle valid delete contact tag params
type DeleteContactTagHandler interface {
	Handle(DeleteContactTagParams) middleware.Responder
}

// NewDeleteContactTag creates a new http.Handler for the delete contact tag operation
func NewDeleteContactTag(ctx *middleware.Context, handler DeleteContactTagHandler) *DeleteContactTag {
	return &DeleteContactTag{Context: ctx, Handler: handler}
}

/* DeleteContactTag swagger:route DELETE /unified/{unified_id}/tags/{tag_id} deleteContactTag

delete contact tag

*/
type DeleteContactTag struct {
	Context *middleware.Context
	Handler DeleteContactTagHandler
}

func (o *DeleteContactTag) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteContactTagParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}