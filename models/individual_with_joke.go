// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IndividualWithJoke individual with joke
// swagger:model IndividualWithJoke
type IndividualWithJoke struct {
	Individual

	// joke
	Joke *Joke `json:"joke,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IndividualWithJoke) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Individual
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Individual = aO0

	// now for regular properties
	var propsIndividualWithJoke struct {
		Joke *Joke `json:"joke,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsIndividualWithJoke); err != nil {
		return err
	}
	m.Joke = propsIndividualWithJoke.Joke

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IndividualWithJoke) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.Individual)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsIndividualWithJoke struct {
		Joke *Joke `json:"joke,omitempty"`
	}
	propsIndividualWithJoke.Joke = m.Joke

	jsonDataPropsIndividualWithJoke, errIndividualWithJoke := swag.WriteJSON(propsIndividualWithJoke)
	if errIndividualWithJoke != nil {
		return nil, errIndividualWithJoke
	}
	_parts = append(_parts, jsonDataPropsIndividualWithJoke)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this individual with joke
func (m *IndividualWithJoke) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Individual
	if err := m.Individual.Validate(formats); err != nil {
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

func (m *IndividualWithJoke) validateJoke(formats strfmt.Registry) error {

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
func (m *IndividualWithJoke) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IndividualWithJoke) UnmarshalBinary(b []byte) error {
	var res IndividualWithJoke
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}