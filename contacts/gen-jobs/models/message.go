// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Message Message
//
// swagger:model Message
type Message struct {

	// attributes
	Attributes interface{} `json:"attributes,omitempty"`

	// data
	// Format: byte
	Data strfmt.Base64 `json:"data,omitempty"`

	// delivery attempt
	DeliveryAttempt int64 `json:"deliveryAttempt,omitempty"`

	// message Id
	// Required: true
	MessageID *string `json:"messageId"`

	// ordering key
	OrderingKey string `json:"orderingKey,omitempty"`

	// publish time
	// Format: date-time
	PublishTime strfmt.DateTime `json:"publishTime,omitempty"`
}

// Validate validates this message
func (m *Message) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Message) validateMessageID(formats strfmt.Registry) error {

	if err := validate.Required("messageId", "body", m.MessageID); err != nil {
		return err
	}

	return nil
}

func (m *Message) validatePublishTime(formats strfmt.Registry) error {
	if swag.IsZero(m.PublishTime) { // not required
		return nil
	}

	if err := validate.FormatOf("publishTime", "body", "date-time", m.PublishTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this message based on context it is used
func (m *Message) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Message) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Message) UnmarshalBinary(b []byte) error {
	var res Message
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}