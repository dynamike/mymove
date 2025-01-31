// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/gen/primemessages"
)

// CreateMTOShipmentReader is a Reader for the CreateMTOShipment structure.
type CreateMTOShipmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMTOShipmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMTOShipmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateMTOShipmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateMTOShipmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateMTOShipmentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateMTOShipmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateMTOShipmentOK creates a CreateMTOShipmentOK with default headers values
func NewCreateMTOShipmentOK() *CreateMTOShipmentOK {
	return &CreateMTOShipmentOK{}
}

/* CreateMTOShipmentOK describes a response with status code 200, with default header values.

Successfully created a MTO shipment.
*/
type CreateMTOShipmentOK struct {
	Payload *primemessages.MTOShipment
}

func (o *CreateMTOShipmentOK) Error() string {
	return fmt.Sprintf("[POST /mto-shipments][%d] createMTOShipmentOK  %+v", 200, o.Payload)
}
func (o *CreateMTOShipmentOK) GetPayload() *primemessages.MTOShipment {
	return o.Payload
}

func (o *CreateMTOShipmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.MTOShipment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMTOShipmentBadRequest creates a CreateMTOShipmentBadRequest with default headers values
func NewCreateMTOShipmentBadRequest() *CreateMTOShipmentBadRequest {
	return &CreateMTOShipmentBadRequest{}
}

/* CreateMTOShipmentBadRequest describes a response with status code 400, with default header values.

The request payload is invalid.
*/
type CreateMTOShipmentBadRequest struct {
	Payload *primemessages.ClientError
}

func (o *CreateMTOShipmentBadRequest) Error() string {
	return fmt.Sprintf("[POST /mto-shipments][%d] createMTOShipmentBadRequest  %+v", 400, o.Payload)
}
func (o *CreateMTOShipmentBadRequest) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreateMTOShipmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMTOShipmentNotFound creates a CreateMTOShipmentNotFound with default headers values
func NewCreateMTOShipmentNotFound() *CreateMTOShipmentNotFound {
	return &CreateMTOShipmentNotFound{}
}

/* CreateMTOShipmentNotFound describes a response with status code 404, with default header values.

The requested resource wasn't found.
*/
type CreateMTOShipmentNotFound struct {
	Payload *primemessages.ClientError
}

func (o *CreateMTOShipmentNotFound) Error() string {
	return fmt.Sprintf("[POST /mto-shipments][%d] createMTOShipmentNotFound  %+v", 404, o.Payload)
}
func (o *CreateMTOShipmentNotFound) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreateMTOShipmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMTOShipmentUnprocessableEntity creates a CreateMTOShipmentUnprocessableEntity with default headers values
func NewCreateMTOShipmentUnprocessableEntity() *CreateMTOShipmentUnprocessableEntity {
	return &CreateMTOShipmentUnprocessableEntity{}
}

/* CreateMTOShipmentUnprocessableEntity describes a response with status code 422, with default header values.

The payload was unprocessable.
*/
type CreateMTOShipmentUnprocessableEntity struct {
	Payload *primemessages.ValidationError
}

func (o *CreateMTOShipmentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /mto-shipments][%d] createMTOShipmentUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateMTOShipmentUnprocessableEntity) GetPayload() *primemessages.ValidationError {
	return o.Payload
}

func (o *CreateMTOShipmentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMTOShipmentInternalServerError creates a CreateMTOShipmentInternalServerError with default headers values
func NewCreateMTOShipmentInternalServerError() *CreateMTOShipmentInternalServerError {
	return &CreateMTOShipmentInternalServerError{}
}

/* CreateMTOShipmentInternalServerError describes a response with status code 500, with default header values.

A server error occurred.
*/
type CreateMTOShipmentInternalServerError struct {
	Payload *primemessages.Error
}

func (o *CreateMTOShipmentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /mto-shipments][%d] createMTOShipmentInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateMTOShipmentInternalServerError) GetPayload() *primemessages.Error {
	return o.Payload
}

func (o *CreateMTOShipmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
