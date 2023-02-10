// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContactSourceDeleted ContactSourceDeleted
//
// swagger:model ContactSourceDeleted
type ContactSourceDeleted struct {

	// contact source Id
	ContactSourceID string `json:"contactSourceId,omitempty"`

	// remove contacts from unified
	RemoveContactsFromUnified bool `json:"removeContactsFromUnified,omitempty"`

	// source
	Source string `json:"source,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`
}

// Validate validates this contact source deleted
func (m *ContactSourceDeleted) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this contact source deleted based on context it is used
func (m *ContactSourceDeleted) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContactSourceDeleted) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactSourceDeleted) UnmarshalBinary(b []byte) error {
	var res ContactSourceDeleted
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
