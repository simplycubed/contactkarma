// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Nickname nickname
//
// swagger:model Nickname
type Nickname struct {

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this nickname
func (m *Nickname) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nickname based on context it is used
func (m *Nickname) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Nickname) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Nickname) UnmarshalBinary(b []byte) error {
	var res Nickname
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}