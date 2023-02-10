// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchContactNoteHandlerFunc turns a function with the right signature into a patch contact note handler
type PatchContactNoteHandlerFunc func(PatchContactNoteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchContactNoteHandlerFunc) Handle(params PatchContactNoteParams) middleware.Responder {
	return fn(params)
}

// PatchContactNoteHandler interface for that can handle valid patch contact note params
type PatchContactNoteHandler interface {
	Handle(PatchContactNoteParams) middleware.Responder
}

// NewPatchContactNote creates a new http.Handler for the patch contact note operation
func NewPatchContactNote(ctx *middleware.Context, handler PatchContactNoteHandler) *PatchContactNote {
	return &PatchContactNote{Context: ctx, Handler: handler}
}

/* PatchContactNote swagger:route PATCH /unified/{unified_id}/notes/{note_id} patchContactNote

patch contact note

*/
type PatchContactNote struct {
	Context *middleware.Context
	Handler PatchContactNoteHandler
}

func (o *PatchContactNote) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPatchContactNoteParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}