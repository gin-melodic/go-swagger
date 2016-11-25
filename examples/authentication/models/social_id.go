package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// SocialID social id
// swagger:model social_id
type SocialID struct {

	// ssn
	// Required: true
	// Min Length: 11
	Ssn *string `json:"ssn"`
}

// Validate validates this social id
func (m *SocialID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSsn(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SocialID) validateSsn(formats strfmt.Registry) error {

	if err := validate.Required("ssn", "body", m.Ssn); err != nil {
		return err
	}

	if err := validate.MinLength("ssn", "body", string(*m.Ssn), 11); err != nil {
		return err
	}

	return nil
}
