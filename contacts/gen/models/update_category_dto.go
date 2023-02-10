// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateCategoryDto UpdateCategoryDto
//
// swagger:model UpdateCategoryDto
type UpdateCategoryDto struct {

	// category
	Category ContactCategory `json:"category,omitempty"`
}

// Validate validates this update category dto
func (m *UpdateCategoryDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateCategoryDto) validateCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.Category) { // not required
		return nil
	}

	if err := m.Category.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("category")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("category")
		}
		return err
	}

	return nil
}

// ContextValidate validate this update category dto based on the context it is used
func (m *UpdateCategoryDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateCategoryDto) contextValidateCategory(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Category.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("category")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("category")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateCategoryDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateCategoryDto) UnmarshalBinary(b []byte) error {
	var res UpdateCategoryDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}