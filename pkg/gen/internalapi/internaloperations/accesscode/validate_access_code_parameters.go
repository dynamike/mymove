// Code generated by go-swagger; DO NOT EDIT.

package accesscode

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewValidateAccessCodeParams creates a new ValidateAccessCodeParams object
//
// There are no default values defined in the spec.
func NewValidateAccessCodeParams() ValidateAccessCodeParams {

	return ValidateAccessCodeParams{}
}

// ValidateAccessCodeParams contains all the bound params for the validate access code operation
// typically these are obtained from a http.Request
//
// swagger:parameters validateAccessCode
type ValidateAccessCodeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*the code the access code represents and verifies if in use
	  Pattern: ^(HHG|PPM)-[A-Z0-9]{6}$
	  In: query
	*/
	Code *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewValidateAccessCodeParams() beforehand.
func (o *ValidateAccessCodeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCode, qhkCode, _ := qs.GetOK("code")
	if err := o.bindCode(qCode, qhkCode, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCode binds and validates parameter Code from query.
func (o *ValidateAccessCodeParams) bindCode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Code = &raw

	if err := o.validateCode(formats); err != nil {
		return err
	}

	return nil
}

// validateCode carries on validations for parameter Code
func (o *ValidateAccessCodeParams) validateCode(formats strfmt.Registry) error {

	if err := validate.Pattern("code", "query", *o.Code, `^(HHG|PPM)-[A-Z0-9]{6}$`); err != nil {
		return err
	}

	return nil
}
