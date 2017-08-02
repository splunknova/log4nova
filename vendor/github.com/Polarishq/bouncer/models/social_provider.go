package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// SocialProvider social provider
// swagger:model SocialProvider
type SocialProvider struct {

	// endpoint to authorize against
	// Required: true
	AuthorizeURL *string `json:"authorize_url"`

	// displayable name for provider
	// Required: true
	DisplayName *string `json:"display_name"`

	// unique provider id
	// Required: true
	ID *string `json:"id"`

	// url for provider image to display
	// Required: true
	ImgURL *string `json:"img_url"`
}

// Validate validates this social provider
func (m *SocialProvider) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorizeURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImgURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SocialProvider) validateAuthorizeURL(formats strfmt.Registry) error {

	if err := validate.Required("authorize_url", "body", m.AuthorizeURL); err != nil {
		return err
	}

	return nil
}

func (m *SocialProvider) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("display_name", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *SocialProvider) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SocialProvider) validateImgURL(formats strfmt.Registry) error {

	if err := validate.Required("img_url", "body", m.ImgURL); err != nil {
		return err
	}

	return nil
}
