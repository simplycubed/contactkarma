// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostContactNoteHandlerFunc turns a function with the right signature into a post contact note handler
type PostContactNoteHandlerFunc func(PostContactNoteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostContactNoteHandlerFunc) Handle(params PostContactNoteParams) middleware.Responder {
	return fn(params)
}

// PostContactNoteHandler interface for that can handle valid post contact note params
type PostContactNoteHandler interface {
	Handle(PostContactNoteParams) middleware.Responder
}

// NewPostContactNote creates a new http.Handler for the post contact note operation
func NewPostContactNote(ctx *middleware.Context, handler PostContactNoteHandler) *PostContactNote {
	return &PostContactNote{Context: ctx, Handler: handler}
}

/* PostContactNote swagger:route POST /unified/{unified_id}/notes postContactNote

post contact note

*/
type PostContactNote struct {
	Context *middleware.Context
	Handler PostContactNoteHandler
}

func (o *PostContactNote) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostContactNoteParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}