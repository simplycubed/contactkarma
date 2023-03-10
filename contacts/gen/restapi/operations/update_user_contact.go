// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateUserContactHandlerFunc turns a function with the right signature into a update user contact handler
type UpdateUserContactHandlerFunc func(UpdateUserContactParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateUserContactHandlerFunc) Handle(params UpdateUserContactParams) middleware.Responder {
	return fn(params)
}

// UpdateUserContactHandler interface for that can handle valid update user contact params
type UpdateUserContactHandler interface {
	Handle(UpdateUserContactParams) middleware.Responder
}

// NewUpdateUserContact creates a new http.Handler for the update user contact operation
func NewUpdateUserContact(ctx *middleware.Context, handler UpdateUserContactHandler) *UpdateUserContact {
	return &UpdateUserContact{Context: ctx, Handler: handler}
}

/* UpdateUserContact swagger:route PATCH /unified/{unified_id} updateUserContact

update user's contact by id

*/
type UpdateUserContact struct {
	Context *middleware.Context
	Handler UpdateUserContactHandler
}

func (o *UpdateUserContact) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateUserContactParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
