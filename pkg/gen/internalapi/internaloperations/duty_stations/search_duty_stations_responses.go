// Code generated by go-swagger; DO NOT EDIT.

package duty_stations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/internalmessages"
)

// SearchDutyStationsOKCode is the HTTP code returned for type SearchDutyStationsOK
const SearchDutyStationsOKCode int = 200

/*SearchDutyStationsOK the instance of the duty station

swagger:response searchDutyStationsOK
*/
type SearchDutyStationsOK struct {

	/*
	  In: Body
	*/
	Payload internalmessages.DutyStationsPayload `json:"body,omitempty"`
}

// NewSearchDutyStationsOK creates SearchDutyStationsOK with default headers values
func NewSearchDutyStationsOK() *SearchDutyStationsOK {

	return &SearchDutyStationsOK{}
}

// WithPayload adds the payload to the search duty stations o k response
func (o *SearchDutyStationsOK) WithPayload(payload internalmessages.DutyStationsPayload) *SearchDutyStationsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search duty stations o k response
func (o *SearchDutyStationsOK) SetPayload(payload internalmessages.DutyStationsPayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchDutyStationsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = internalmessages.DutyStationsPayload{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// SearchDutyStationsBadRequestCode is the HTTP code returned for type SearchDutyStationsBadRequest
const SearchDutyStationsBadRequestCode int = 400

/*SearchDutyStationsBadRequest invalid request

swagger:response searchDutyStationsBadRequest
*/
type SearchDutyStationsBadRequest struct {
}

// NewSearchDutyStationsBadRequest creates SearchDutyStationsBadRequest with default headers values
func NewSearchDutyStationsBadRequest() *SearchDutyStationsBadRequest {

	return &SearchDutyStationsBadRequest{}
}

// WriteResponse to the client
func (o *SearchDutyStationsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// SearchDutyStationsUnauthorizedCode is the HTTP code returned for type SearchDutyStationsUnauthorized
const SearchDutyStationsUnauthorizedCode int = 401

/*SearchDutyStationsUnauthorized request requires user authentication

swagger:response searchDutyStationsUnauthorized
*/
type SearchDutyStationsUnauthorized struct {
}

// NewSearchDutyStationsUnauthorized creates SearchDutyStationsUnauthorized with default headers values
func NewSearchDutyStationsUnauthorized() *SearchDutyStationsUnauthorized {

	return &SearchDutyStationsUnauthorized{}
}

// WriteResponse to the client
func (o *SearchDutyStationsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SearchDutyStationsForbiddenCode is the HTTP code returned for type SearchDutyStationsForbidden
const SearchDutyStationsForbiddenCode int = 403

/*SearchDutyStationsForbidden user is not authorized

swagger:response searchDutyStationsForbidden
*/
type SearchDutyStationsForbidden struct {
}

// NewSearchDutyStationsForbidden creates SearchDutyStationsForbidden with default headers values
func NewSearchDutyStationsForbidden() *SearchDutyStationsForbidden {

	return &SearchDutyStationsForbidden{}
}

// WriteResponse to the client
func (o *SearchDutyStationsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// SearchDutyStationsNotFoundCode is the HTTP code returned for type SearchDutyStationsNotFound
const SearchDutyStationsNotFoundCode int = 404

/*SearchDutyStationsNotFound matching duty station not found

swagger:response searchDutyStationsNotFound
*/
type SearchDutyStationsNotFound struct {
}

// NewSearchDutyStationsNotFound creates SearchDutyStationsNotFound with default headers values
func NewSearchDutyStationsNotFound() *SearchDutyStationsNotFound {

	return &SearchDutyStationsNotFound{}
}

// WriteResponse to the client
func (o *SearchDutyStationsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// SearchDutyStationsInternalServerErrorCode is the HTTP code returned for type SearchDutyStationsInternalServerError
const SearchDutyStationsInternalServerErrorCode int = 500

/*SearchDutyStationsInternalServerError internal server error

swagger:response searchDutyStationsInternalServerError
*/
type SearchDutyStationsInternalServerError struct {
}

// NewSearchDutyStationsInternalServerError creates SearchDutyStationsInternalServerError with default headers values
func NewSearchDutyStationsInternalServerError() *SearchDutyStationsInternalServerError {

	return &SearchDutyStationsInternalServerError{}
}

// WriteResponse to the client
func (o *SearchDutyStationsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
