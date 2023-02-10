// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserNames user names
//
// swagger:model UserNames
type UserNames struct {

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// display name last first
	DisplayNameLastFirst string `json:"display_name_last_first,omitempty"`

	// family name
	FamilyName string `json:"family_name,omitempty"`

	// given name
	GivenName string `json:"given_name,omitempty"`

	// honorific prefix
	HonorificPrefix string `json:"honorific_prefix,omitempty"`

	// honorific suffix
	HonorificSuffix string `json:"honorific_suffix,omitempty"`

	// middle name
	MiddleName string `json:"middle_name,omitempty"`

	// phonetic family name
	PhoneticFamilyName string `json:"phonetic_family_name,omitempty"`

	// phonetic full name
	PhoneticFullName string `json:"phonetic_full_name,omitempty"`

	// phonetic given name
	PhoneticGivenName string `json:"phonetic_given_name,omitempty"`

	// phonetic honorific prefix
	PhoneticHonorificPrefix string `json:"phonetic_honorific_prefix,omitempty"`

	// phonetic honorific suffix
	PhoneticHonorificSuffix string `json:"phonetic_honorific_suffix,omitempty"`

	// phonetic middle name
	PhoneticMiddleName string `json:"phonetic_middle_name,omitempty"`

	// unstructured name
	UnstructuredName string `json:"unstructured_name,omitempty"`
}

// Validate validates this user names
func (m *UserNames) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user names based on context it is used
func (m *UserNames) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserNames) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserNames) UnmarshalBinary(b []byte) error {
	var res UserNames
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
