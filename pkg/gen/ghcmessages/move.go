// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Move move
//
// swagger:model Move
type Move struct {

	// available to prime at
	// Format: date-time
	AvailableToPrimeAt *strfmt.DateTime `json:"availableToPrimeAt,omitempty"`

	// contractor
	Contractor *Contractor `json:"contractor,omitempty"`

	// contractor Id
	// Format: uuid
	ContractorID *strfmt.UUID `json:"contractorId,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// e tag
	ETag string `json:"eTag,omitempty"`

	// id
	// Example: 1f2270c7-7166-40ae-981e-b200ebdf3054
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// locator
	// Example: 1K43AR
	Locator string `json:"locator,omitempty"`

	// orders
	Orders *Order `json:"orders,omitempty"`

	// orders Id
	// Example: c56a4180-65aa-42ec-a945-5fd21dec0538
	// Format: uuid
	OrdersID strfmt.UUID `json:"ordersId,omitempty"`

	// reference Id
	// Example: 1001-3456
	ReferenceID *string `json:"referenceId,omitempty"`

	// service counseling completed at
	// Format: date-time
	ServiceCounselingCompletedAt *strfmt.DateTime `json:"serviceCounselingCompletedAt,omitempty"`

	// status
	Status MoveStatus `json:"status,omitempty"`

	// submitted at
	// Format: date-time
	SubmittedAt *strfmt.DateTime `json:"submittedAt,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this move
func (m *Move) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableToPrimeAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContractor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContractorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrdersID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCounselingCompletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmittedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Move) validateAvailableToPrimeAt(formats strfmt.Registry) error {
	if swag.IsZero(m.AvailableToPrimeAt) { // not required
		return nil
	}

	if err := validate.FormatOf("availableToPrimeAt", "body", "date-time", m.AvailableToPrimeAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Move) validateContractor(formats strfmt.Registry) error {
	if swag.IsZero(m.Contractor) { // not required
		return nil
	}

	if m.Contractor != nil {
		if err := m.Contractor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contractor")
			}
			return err
		}
	}

	return nil
}

func (m *Move) validateContractorID(formats strfmt.Registry) error {
	if swag.IsZero(m.ContractorID) { // not required
		return nil
	}

	if err := validate.FormatOf("contractorId", "body", "uuid", m.ContractorID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Move) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Move) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Move) validateOrders(formats strfmt.Registry) error {
	if swag.IsZero(m.Orders) { // not required
		return nil
	}

	if m.Orders != nil {
		if err := m.Orders.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orders")
			}
			return err
		}
	}

	return nil
}

func (m *Move) validateOrdersID(formats strfmt.Registry) error {
	if swag.IsZero(m.OrdersID) { // not required
		return nil
	}

	if err := validate.FormatOf("ordersId", "body", "uuid", m.OrdersID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Move) validateServiceCounselingCompletedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceCounselingCompletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("serviceCounselingCompletedAt", "body", "date-time", m.ServiceCounselingCompletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Move) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *Move) validateSubmittedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.SubmittedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("submittedAt", "body", "date-time", m.SubmittedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Move) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this move based on the context it is used
func (m *Move) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContractor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Move) contextValidateContractor(ctx context.Context, formats strfmt.Registry) error {

	if m.Contractor != nil {
		if err := m.Contractor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contractor")
			}
			return err
		}
	}

	return nil
}

func (m *Move) contextValidateOrders(ctx context.Context, formats strfmt.Registry) error {

	if m.Orders != nil {
		if err := m.Orders.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orders")
			}
			return err
		}
	}

	return nil
}

func (m *Move) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Move) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Move) UnmarshalBinary(b []byte) error {
	var res Move
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
