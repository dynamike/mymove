// Code generated by go-swagger; DO NOT EDIT.

package shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// DeleteShipmentNoContentCode is the HTTP code returned for type DeleteShipmentNoContent
const DeleteShipmentNoContentCode int = 204

/*DeleteShipmentNoContent Successfully soft deleted the shipment

swagger:response deleteShipmentNoContent
*/
type DeleteShipmentNoContent struct {
}

// NewDeleteShipmentNoContent creates DeleteShipmentNoContent with default headers values
func NewDeleteShipmentNoContent() *DeleteShipmentNoContent {

	return &DeleteShipmentNoContent{}
}

// WriteResponse to the client
func (o *DeleteShipmentNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteShipmentForbiddenCode is the HTTP code returned for type DeleteShipmentForbidden
const DeleteShipmentForbiddenCode int = 403

/*DeleteShipmentForbidden The request was denied

swagger:response deleteShipmentForbidden
*/
type DeleteShipmentForbidden struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Error `json:"body,omitempty"`
}

// NewDeleteShipmentForbidden creates DeleteShipmentForbidden with default headers values
func NewDeleteShipmentForbidden() *DeleteShipmentForbidden {

	return &DeleteShipmentForbidden{}
}

// WithPayload adds the payload to the delete shipment forbidden response
func (o *DeleteShipmentForbidden) WithPayload(payload *ghcmessages.Error) *DeleteShipmentForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete shipment forbidden response
func (o *DeleteShipmentForbidden) SetPayload(payload *ghcmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteShipmentForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteShipmentNotFoundCode is the HTTP code returned for type DeleteShipmentNotFound
const DeleteShipmentNotFoundCode int = 404

/*DeleteShipmentNotFound The requested resource wasn't found

swagger:response deleteShipmentNotFound
*/
type DeleteShipmentNotFound struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Error `json:"body,omitempty"`
}

// NewDeleteShipmentNotFound creates DeleteShipmentNotFound with default headers values
func NewDeleteShipmentNotFound() *DeleteShipmentNotFound {

	return &DeleteShipmentNotFound{}
}

// WithPayload adds the payload to the delete shipment not found response
func (o *DeleteShipmentNotFound) WithPayload(payload *ghcmessages.Error) *DeleteShipmentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete shipment not found response
func (o *DeleteShipmentNotFound) SetPayload(payload *ghcmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteShipmentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteShipmentInternalServerErrorCode is the HTTP code returned for type DeleteShipmentInternalServerError
const DeleteShipmentInternalServerErrorCode int = 500

/*DeleteShipmentInternalServerError A server error occurred

swagger:response deleteShipmentInternalServerError
*/
type DeleteShipmentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Error `json:"body,omitempty"`
}

// NewDeleteShipmentInternalServerError creates DeleteShipmentInternalServerError with default headers values
func NewDeleteShipmentInternalServerError() *DeleteShipmentInternalServerError {

	return &DeleteShipmentInternalServerError{}
}

// WithPayload adds the payload to the delete shipment internal server error response
func (o *DeleteShipmentInternalServerError) WithPayload(payload *ghcmessages.Error) *DeleteShipmentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete shipment internal server error response
func (o *DeleteShipmentInternalServerError) SetPayload(payload *ghcmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteShipmentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
