// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UploadContactsCsvHandlerFunc turns a function with the right signature into a upload contacts csv handler
type UploadContactsCsvHandlerFunc func(UploadContactsCsvParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UploadContactsCsvHandlerFunc) Handle(params UploadContactsCsvParams) middleware.Responder {
	return fn(params)
}

// UploadContactsCsvHandler interface for that can handle valid upload contacts csv params
type UploadContactsCsvHandler interface {
	Handle(UploadContactsCsvParams) middleware.Responder
}

// NewUploadContactsCsv creates a new http.Handler for the upload contacts csv operation
func NewUploadContactsCsv(ctx *middleware.Context, handler UploadContactsCsvHandler) *UploadContactsCsv {
	return &UploadContactsCsv{Context: ctx, Handler: handler}
}

/* UploadContactsCsv swagger:route POST /contacts/upload-csv uploadContactsCsv

upload contacts csv

*/
type UploadContactsCsv struct {
	Context *middleware.Context
	Handler UploadContactsCsvHandler
}

func (o *UploadContactsCsv) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUploadContactsCsvParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
