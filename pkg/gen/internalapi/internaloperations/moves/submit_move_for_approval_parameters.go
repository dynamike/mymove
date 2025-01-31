// Code generated by go-swagger; DO NOT EDIT.

package moves

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/transcom/mymove/pkg/gen/internalmessages"
)

// NewSubmitMoveForApprovalParams creates a new SubmitMoveForApprovalParams object
//
// There are no default values defined in the spec.
func NewSubmitMoveForApprovalParams() SubmitMoveForApprovalParams {

	return SubmitMoveForApprovalParams{}
}

// SubmitMoveForApprovalParams contains all the bound params for the submit move for approval operation
// typically these are obtained from a http.Request
//
// swagger:parameters submitMoveForApproval
type SubmitMoveForApprovalParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*UUID of the move
	  Required: true
	  In: path
	*/
	MoveID strfmt.UUID
	/*
	  Required: true
	  In: body
	*/
	SubmitMoveForApprovalPayload *internalmessages.SubmitMoveForApprovalPayload
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSubmitMoveForApprovalParams() beforehand.
func (o *SubmitMoveForApprovalParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rMoveID, rhkMoveID, _ := route.Params.GetOK("moveId")
	if err := o.bindMoveID(rMoveID, rhkMoveID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body internalmessages.SubmitMoveForApprovalPayload
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("submitMoveForApprovalPayload", "body", ""))
			} else {
				res = append(res, errors.NewParseError("submitMoveForApprovalPayload", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.SubmitMoveForApprovalPayload = &body
			}
		}
	} else {
		res = append(res, errors.Required("submitMoveForApprovalPayload", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMoveID binds and validates parameter MoveID from path.
func (o *SubmitMoveForApprovalParams) bindMoveID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("moveId", "path", "strfmt.UUID", raw)
	}
	o.MoveID = *(value.(*strfmt.UUID))

	if err := o.validateMoveID(formats); err != nil {
		return err
	}

	return nil
}

// validateMoveID carries on validations for parameter MoveID
func (o *SubmitMoveForApprovalParams) validateMoveID(formats strfmt.Registry) error {

	if err := validate.FormatOf("moveId", "path", "uuid", o.MoveID.String(), formats); err != nil {
		return err
	}
	return nil
}
