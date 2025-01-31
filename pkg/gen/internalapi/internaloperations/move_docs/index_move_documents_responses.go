// Code generated by go-swagger; DO NOT EDIT.

package move_docs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/internalmessages"
)

// IndexMoveDocumentsOKCode is the HTTP code returned for type IndexMoveDocumentsOK
const IndexMoveDocumentsOKCode int = 200

/*IndexMoveDocumentsOK returns list of move douments

swagger:response indexMoveDocumentsOK
*/
type IndexMoveDocumentsOK struct {

	/*
	  In: Body
	*/
	Payload internalmessages.MoveDocuments `json:"body,omitempty"`
}

// NewIndexMoveDocumentsOK creates IndexMoveDocumentsOK with default headers values
func NewIndexMoveDocumentsOK() *IndexMoveDocumentsOK {

	return &IndexMoveDocumentsOK{}
}

// WithPayload adds the payload to the index move documents o k response
func (o *IndexMoveDocumentsOK) WithPayload(payload internalmessages.MoveDocuments) *IndexMoveDocumentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the index move documents o k response
func (o *IndexMoveDocumentsOK) SetPayload(payload internalmessages.MoveDocuments) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IndexMoveDocumentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = internalmessages.MoveDocuments{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// IndexMoveDocumentsBadRequestCode is the HTTP code returned for type IndexMoveDocumentsBadRequest
const IndexMoveDocumentsBadRequestCode int = 400

/*IndexMoveDocumentsBadRequest invalid request

swagger:response indexMoveDocumentsBadRequest
*/
type IndexMoveDocumentsBadRequest struct {
}

// NewIndexMoveDocumentsBadRequest creates IndexMoveDocumentsBadRequest with default headers values
func NewIndexMoveDocumentsBadRequest() *IndexMoveDocumentsBadRequest {

	return &IndexMoveDocumentsBadRequest{}
}

// WriteResponse to the client
func (o *IndexMoveDocumentsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// IndexMoveDocumentsUnauthorizedCode is the HTTP code returned for type IndexMoveDocumentsUnauthorized
const IndexMoveDocumentsUnauthorizedCode int = 401

/*IndexMoveDocumentsUnauthorized request requires user authentication

swagger:response indexMoveDocumentsUnauthorized
*/
type IndexMoveDocumentsUnauthorized struct {
}

// NewIndexMoveDocumentsUnauthorized creates IndexMoveDocumentsUnauthorized with default headers values
func NewIndexMoveDocumentsUnauthorized() *IndexMoveDocumentsUnauthorized {

	return &IndexMoveDocumentsUnauthorized{}
}

// WriteResponse to the client
func (o *IndexMoveDocumentsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// IndexMoveDocumentsForbiddenCode is the HTTP code returned for type IndexMoveDocumentsForbidden
const IndexMoveDocumentsForbiddenCode int = 403

/*IndexMoveDocumentsForbidden user is not authorized

swagger:response indexMoveDocumentsForbidden
*/
type IndexMoveDocumentsForbidden struct {
}

// NewIndexMoveDocumentsForbidden creates IndexMoveDocumentsForbidden with default headers values
func NewIndexMoveDocumentsForbidden() *IndexMoveDocumentsForbidden {

	return &IndexMoveDocumentsForbidden{}
}

// WriteResponse to the client
func (o *IndexMoveDocumentsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
