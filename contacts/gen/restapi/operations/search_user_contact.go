// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SearchUserContactHandlerFunc turns a function with the right signature into a search user contact handler
type SearchUserContactHandlerFunc func(SearchUserContactParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SearchUserContactHandlerFunc) Handle(params SearchUserContactParams) middleware.Responder {
	return fn(params)
}

// SearchUserContactHandler interface for that can handle valid search user contact params
type SearchUserContactHandler interface {
	Handle(SearchUserContactParams) middleware.Responder
}

// NewSearchUserContact creates a new http.Handler for the search user contact operation
func NewSearchUserContact(ctx *middleware.Context, handler SearchUserContactHandler) *SearchUserContact {
	return &SearchUserContact{Context: ctx, Handler: handler}
}

/* SearchUserContact swagger:route POST /search searchUserContact

search user's contact

*/
type SearchUserContact struct {
	Context *middleware.Context
	Handler SearchUserContactHandler
}

func (o *SearchUserContact) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSearchUserContactParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
