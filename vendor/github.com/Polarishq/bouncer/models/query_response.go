package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// QueryResponse query response
// swagger:model QueryResponse
type QueryResponse struct {

	// events
	Events []*Event `json:"events"`

	// fields
	Fields []string `json:"fields"`

	// logic info
	LogicInfo *QueryLogicInfoMetaData `json:"logic_info,omitempty"`

	// meta data
	MetaData *ListMetaData `json:"meta_data,omitempty"`
}

// Validate validates this query response
func (m *QueryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFields(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLogicInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMetaData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResponse) validateEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {

		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {

			if err := m.Events[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *QueryResponse) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	return nil
}

func (m *QueryResponse) validateLogicInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.LogicInfo) { // not required
		return nil
	}

	if m.LogicInfo != nil {

		if err := m.LogicInfo.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *QueryResponse) validateMetaData(formats strfmt.Registry) error {

	if swag.IsZero(m.MetaData) { // not required
		return nil
	}

	if m.MetaData != nil {

		if err := m.MetaData.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}