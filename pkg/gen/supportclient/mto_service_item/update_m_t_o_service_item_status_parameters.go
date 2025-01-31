// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/gen/supportmessages"
)

// NewUpdateMTOServiceItemStatusParams creates a new UpdateMTOServiceItemStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMTOServiceItemStatusParams() *UpdateMTOServiceItemStatusParams {
	return &UpdateMTOServiceItemStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMTOServiceItemStatusParamsWithTimeout creates a new UpdateMTOServiceItemStatusParams object
// with the ability to set a timeout on a request.
func NewUpdateMTOServiceItemStatusParamsWithTimeout(timeout time.Duration) *UpdateMTOServiceItemStatusParams {
	return &UpdateMTOServiceItemStatusParams{
		timeout: timeout,
	}
}

// NewUpdateMTOServiceItemStatusParamsWithContext creates a new UpdateMTOServiceItemStatusParams object
// with the ability to set a context for a request.
func NewUpdateMTOServiceItemStatusParamsWithContext(ctx context.Context) *UpdateMTOServiceItemStatusParams {
	return &UpdateMTOServiceItemStatusParams{
		Context: ctx,
	}
}

// NewUpdateMTOServiceItemStatusParamsWithHTTPClient creates a new UpdateMTOServiceItemStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMTOServiceItemStatusParamsWithHTTPClient(client *http.Client) *UpdateMTOServiceItemStatusParams {
	return &UpdateMTOServiceItemStatusParams{
		HTTPClient: client,
	}
}

/* UpdateMTOServiceItemStatusParams contains all the parameters to send to the API endpoint
   for the update m t o service item status operation.

   Typically these are written to a http.Request.
*/
type UpdateMTOServiceItemStatusParams struct {

	/* IfMatch.

	   Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error.

	*/
	IfMatch string

	// Body.
	Body *supportmessages.UpdateMTOServiceItemStatus

	/* MtoServiceItemID.

	   UUID of mto service item to use.
	*/
	MtoServiceItemID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update m t o service item status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMTOServiceItemStatusParams) WithDefaults() *UpdateMTOServiceItemStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update m t o service item status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMTOServiceItemStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update m t o service item status params
func (o *UpdateMTOServiceItemStatusParams) WithTimeout(timeout time.Duration) *UpdateMTOServiceItemStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update m t o service item status params
func (o *UpdateMTOServiceItemStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update m t o service item status params
func (o *UpdateMTOServiceItemStatusParams) WithContext(ctx context.Context) *UpdateMTOServiceItemStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update m t o service item status params
func (o *UpdateMTOServiceItemStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update m t o service item status params
func (o *UpdateMTOServiceItemStatusParams) WithHTTPClient(client *http.Client) *UpdateMTOServiceItemStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update m t o service item status params
func (o *UpdateMTOServiceItemStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the update m t o service item status params
func (o *UpdateMTOServiceItemStatusParams) WithIfMatch(ifMatch string) *UpdateMTOServiceItemStatusParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the update m t o service item status params
func (o *UpdateMTOServiceItemStatusParams) SetIfMatch(ifMatch string) {
	o.IfMatch = ifMatch
}

// WithBody adds the body to the update m t o service item status params
func (o *UpdateMTOServiceItemStatusParams) WithBody(body *supportmessages.UpdateMTOServiceItemStatus) *UpdateMTOServiceItemStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update m t o service item status params
func (o *UpdateMTOServiceItemStatusParams) SetBody(body *supportmessages.UpdateMTOServiceItemStatus) {
	o.Body = body
}

// WithMtoServiceItemID adds the mtoServiceItemID to the update m t o service item status params
func (o *UpdateMTOServiceItemStatusParams) WithMtoServiceItemID(mtoServiceItemID string) *UpdateMTOServiceItemStatusParams {
	o.SetMtoServiceItemID(mtoServiceItemID)
	return o
}

// SetMtoServiceItemID adds the mtoServiceItemId to the update m t o service item status params
func (o *UpdateMTOServiceItemStatusParams) SetMtoServiceItemID(mtoServiceItemID string) {
	o.MtoServiceItemID = mtoServiceItemID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMTOServiceItemStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param If-Match
	if err := r.SetHeaderParam("If-Match", o.IfMatch); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param mtoServiceItemID
	if err := r.SetPathParam("mtoServiceItemID", o.MtoServiceItemID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
