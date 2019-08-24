// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Individual individual
// swagger:model Individual
type Individual struct {

	// region
	Region string `json:"Region,omitempty"`

	// gender
	// Enum: [Male Female Transgender N/A]
	Gender string `json:"gender,omitempty"`

	// joke
	Joke *Joke `json:"joke,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// surname
	Surname string `json:"surname,omitempty"`
}

// Validate validates this individual
func (m *Individual) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJoke(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var individualTypeGenderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Male","Female","Transgender","N/A"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		individualTypeGenderPropEnum = append(individualTypeGenderPropEnum, v)
	}
}

const (

	// IndividualGenderMale captures enum value "Male"
	IndividualGenderMale string = "Male"

	// IndividualGenderFemale captures enum value "Female"
	IndividualGenderFemale string = "Female"

	// IndividualGenderTransgender captures enum value "Transgender"
	IndividualGenderTransgender string = "Transgender"

	// IndividualGenderNA captures enum value "N/A"
	IndividualGenderNA string = "N/A"
)

// prop value enum
func (m *Individual) validateGenderEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, individualTypeGenderPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Individual) validateGender(formats strfmt.Registry) error {

	if swag.IsZero(m.Gender) { // not required
		return nil
	}

	// value enum
	if err := m.validateGenderEnum("gender", "body", m.Gender); err != nil {
		return err
	}

	return nil
}

func (m *Individual) validateJoke(formats strfmt.Registry) error {

	if swag.IsZero(m.Joke) { // not required
		return nil
	}

	if m.Joke != nil {
		if err := m.Joke.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("joke")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Individual) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Individual) UnmarshalBinary(b []byte) error {
	var res Individual
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
