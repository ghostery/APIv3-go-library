// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateContact create contact
// swagger:model createContact

type CreateContact struct {

	// Values of the attributes to fill. The attributes must exist in you contact database
	Attributes interface{} `json:"attributes,omitempty"`

	// Email address of the user. Mandatory if `attributes.sms` is not passed
	Email strfmt.Email `json:"email,omitempty"`

	// Blacklist the contact for emails (emailBlacklisted = true)
	EmailBlacklisted bool `json:"emailBlacklisted,omitempty"`

	// Ids of the lists to add the contact to
	ListIds []int64 `json:"listIds"`

	// Blacklist the contact for SMS (smsBlacklisted = true)
	SmsBlacklisted bool `json:"smsBlacklisted,omitempty"`

	// Facilitate to update existing contact in same request (updateEnabled = true)
	UpdateEnabled *bool `json:"updateEnabled,omitempty"`
}

/* polymorph createContact attributes false */

/* polymorph createContact email false */

/* polymorph createContact emailBlacklisted false */

/* polymorph createContact listIds false */

/* polymorph createContact smsBlacklisted false */

/* polymorph createContact updateEnabled false */

// Validate validates this create contact
func (m *CreateContact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateListIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateContact) validateListIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ListIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateContact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateContact) UnmarshalBinary(b []byte) error {
	var res CreateContact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
