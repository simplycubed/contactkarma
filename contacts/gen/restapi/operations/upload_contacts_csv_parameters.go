// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// UploadContactsCsvMaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var UploadContactsCsvMaxParseMemory int64 = 32 << 20

// NewUploadContactsCsvParams creates a new UploadContactsCsvParams object
//
// There are no default values defined in the spec.
func NewUploadContactsCsvParams() UploadContactsCsvParams {

	return UploadContactsCsvParams{}
}

// UploadContactsCsvParams contains all the bound params for the upload contacts csv operation
// typically these are obtained from a http.Request
//
// swagger:parameters upload-contacts-csv
type UploadContactsCsvParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	XApigatewayAPIUserinfo string
	/*The uploaded file data
	  Required: true
	  In: formData
	*/
	File io.ReadCloser
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUploadContactsCsvParams() beforehand.
func (o *UploadContactsCsvParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(UploadContactsCsvMaxParseMemory); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}

	if err := o.bindXApigatewayAPIUserinfo(r.Header[http.CanonicalHeaderKey("X-Apigateway-Api-Userinfo")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "file", err))
	} else if err := o.bindFile(file, fileHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.File = &runtime.File{Data: file, Header: fileHeader}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXApigatewayAPIUserinfo binds and validates parameter XApigatewayAPIUserinfo from header.
func (o *UploadContactsCsvParams) bindXApigatewayAPIUserinfo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("X-Apigateway-Api-Userinfo", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("X-Apigateway-Api-Userinfo", "header", raw); err != nil {
		return err
	}
	o.XApigatewayAPIUserinfo = raw

	return nil
}

// bindFile binds file parameter File.
//
// The only supported validations on files are MinLength and MaxLength
func (o *UploadContactsCsvParams) bindFile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}