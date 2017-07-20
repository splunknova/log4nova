package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// QueryLogicInfoMetaData query logic info meta data
// swagger:model QueryLogicInfoMetaData
type QueryLogicInfoMetaData struct {

	// Shows the user the backend logic to handle the search conditions; like expired search, cache and new search mismatch
	InfoMessage string `json:"info_message,omitempty"`
}

// Validate validates this query logic info meta data
func (m *QueryLogicInfoMetaData) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
