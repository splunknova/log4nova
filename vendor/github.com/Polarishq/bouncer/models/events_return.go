package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// EventsReturn events return
// swagger:model EventsReturn
type EventsReturn struct {

	// Number of bytes received
	Bytes int64 `json:"bytes,omitempty"`

	// Number of events received
	Count int64 `json:"count,omitempty"`
}

// Validate validates this events return
func (m *EventsReturn) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}