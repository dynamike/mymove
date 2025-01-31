// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DPSAuthCookieURLPayload d p s auth cookie URL payload
//
// swagger:model DPSAuthCookieURLPayload
type DPSAuthCookieURLPayload struct {

	// cookie url
	// Format: uri
	CookieURL strfmt.URI `json:"cookie_url,omitempty"`
}

// Validate validates this d p s auth cookie URL payload
func (m *DPSAuthCookieURLPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCookieURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DPSAuthCookieURLPayload) validateCookieURL(formats strfmt.Registry) error {
	if swag.IsZero(m.CookieURL) { // not required
		return nil
	}

	if err := validate.FormatOf("cookie_url", "body", "uri", m.CookieURL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this d p s auth cookie URL payload based on context it is used
func (m *DPSAuthCookieURLPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DPSAuthCookieURLPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DPSAuthCookieURLPayload) UnmarshalBinary(b []byte) error {
	var res DPSAuthCookieURLPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
