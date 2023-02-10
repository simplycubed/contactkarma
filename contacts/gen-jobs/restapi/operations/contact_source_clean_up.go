// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ContactSourceCleanUpHandlerFunc turns a function with the right signature into a contact source clean up handler
type ContactSourceCleanUpHandlerFunc func(ContactSourceCleanUpParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ContactSourceCleanUpHandlerFunc) Handle(params ContactSourceCleanUpParams) middleware.Responder {
	return fn(params)
}

// ContactSourceCleanUpHandler interface for that can handle valid contact source clean up params
type ContactSourceCleanUpHandler interface {
	Handle(ContactSourceCleanUpParams) middleware.Responder
}

// NewContactSourceCleanUp creates a new http.Handler for the contact source clean up operation
func NewContactSourceCleanUp(ctx *middleware.Context, handler ContactSourceCleanUpHandler) *ContactSourceCleanUp {
	return &ContactSourceCleanUp{Context: ctx, Handler: handler}
}

/* ContactSourceCleanUp swagger:route POST /contact-source-clean-up contactSourceCleanUp

clean up job after deleting a contact source

*/
type ContactSourceCleanUp struct {
	Context *middleware.Context
	Handler ContactSourceCleanUpHandler
}

func (o *ContactSourceCleanUp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewContactSourceCleanUpParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}