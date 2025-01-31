// Code generated by go-swagger; DO NOT EDIT.

package payment_request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/gen/primemessages"
)

// CreateUploadReader is a Reader for the CreateUpload structure.
type CreateUploadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUploadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUploadCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUploadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateUploadUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUploadForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUploadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateUploadUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUploadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUploadCreated creates a CreateUploadCreated with default headers values
func NewCreateUploadCreated() *CreateUploadCreated {
	return &CreateUploadCreated{}
}

/* CreateUploadCreated describes a response with status code 201, with default header values.

Successfully created upload of digital file.
*/
type CreateUploadCreated struct {
	Payload *primemessages.Upload
}

func (o *CreateUploadCreated) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadCreated  %+v", 201, o.Payload)
}
func (o *CreateUploadCreated) GetPayload() *primemessages.Upload {
	return o.Payload
}

func (o *CreateUploadCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.Upload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadBadRequest creates a CreateUploadBadRequest with default headers values
func NewCreateUploadBadRequest() *CreateUploadBadRequest {
	return &CreateUploadBadRequest{}
}

/* CreateUploadBadRequest describes a response with status code 400, with default header values.

The request payload is invalid.
*/
type CreateUploadBadRequest struct {
	Payload *primemessages.ClientError
}

func (o *CreateUploadBadRequest) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadBadRequest  %+v", 400, o.Payload)
}
func (o *CreateUploadBadRequest) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreateUploadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadUnauthorized creates a CreateUploadUnauthorized with default headers values
func NewCreateUploadUnauthorized() *CreateUploadUnauthorized {
	return &CreateUploadUnauthorized{}
}

/* CreateUploadUnauthorized describes a response with status code 401, with default header values.

The request was denied.
*/
type CreateUploadUnauthorized struct {
	Payload *primemessages.ClientError
}

func (o *CreateUploadUnauthorized) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateUploadUnauthorized) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreateUploadUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadForbidden creates a CreateUploadForbidden with default headers values
func NewCreateUploadForbidden() *CreateUploadForbidden {
	return &CreateUploadForbidden{}
}

/* CreateUploadForbidden describes a response with status code 403, with default header values.

The request was denied.
*/
type CreateUploadForbidden struct {
	Payload *primemessages.ClientError
}

func (o *CreateUploadForbidden) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadForbidden  %+v", 403, o.Payload)
}
func (o *CreateUploadForbidden) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreateUploadForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadNotFound creates a CreateUploadNotFound with default headers values
func NewCreateUploadNotFound() *CreateUploadNotFound {
	return &CreateUploadNotFound{}
}

/* CreateUploadNotFound describes a response with status code 404, with default header values.

The requested resource wasn't found.
*/
type CreateUploadNotFound struct {
	Payload *primemessages.ClientError
}

func (o *CreateUploadNotFound) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadNotFound  %+v", 404, o.Payload)
}
func (o *CreateUploadNotFound) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreateUploadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadUnprocessableEntity creates a CreateUploadUnprocessableEntity with default headers values
func NewCreateUploadUnprocessableEntity() *CreateUploadUnprocessableEntity {
	return &CreateUploadUnprocessableEntity{}
}

/* CreateUploadUnprocessableEntity describes a response with status code 422, with default header values.

The payload was unprocessable.
*/
type CreateUploadUnprocessableEntity struct {
	Payload *primemessages.ValidationError
}

func (o *CreateUploadUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateUploadUnprocessableEntity) GetPayload() *primemessages.ValidationError {
	return o.Payload
}

func (o *CreateUploadUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadInternalServerError creates a CreateUploadInternalServerError with default headers values
func NewCreateUploadInternalServerError() *CreateUploadInternalServerError {
	return &CreateUploadInternalServerError{}
}

/* CreateUploadInternalServerError describes a response with status code 500, with default header values.

A server error occurred.
*/
type CreateUploadInternalServerError struct {
	Payload *primemessages.Error
}

func (o *CreateUploadInternalServerError) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateUploadInternalServerError) GetPayload() *primemessages.Error {
	return o.Payload
}

func (o *CreateUploadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
