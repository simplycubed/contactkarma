// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// InitGoogleContactSourceHandlerFunc turns a function with the right signature into a init google contact source handler
type InitGoogleContactSourceHandlerFunc func(InitGoogleContactSourceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn InitGoogleContactSourceHandlerFunc) Handle(params InitGoogleContactSourceParams) middleware.Responder {
	return fn(params)
}

// InitGoogleContactSourceHandler interface for that can handle valid init google contact source params
type InitGoogleContactSourceHandler interface {
	Handle(InitGoogleContactSourceParams) middleware.Responder
}

// NewInitGoogleContactSource creates a new http.Handler for the init google contact source operation
func NewInitGoogleContactSource(ctx *middleware.Context, handler InitGoogleContactSourceHandler) *InitGoogleContactSource {
	return &InitGoogleContactSource{Context: ctx, Handler: handler}
}

/* InitGoogleContactSource swagger:route GET /contacts/sources/google/init initGoogleContactSource

returns redirect url

*/
type InitGoogleContactSource struct {
	Context *middleware.Context
	Handler InitGoogleContactSourceHandler
}

func (o *InitGoogleContactSource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewInitGoogleContactSourceParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
