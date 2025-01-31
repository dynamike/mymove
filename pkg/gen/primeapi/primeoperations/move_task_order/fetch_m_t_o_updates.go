// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// FetchMTOUpdatesHandlerFunc turns a function with the right signature into a fetch m t o updates handler
type FetchMTOUpdatesHandlerFunc func(FetchMTOUpdatesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FetchMTOUpdatesHandlerFunc) Handle(params FetchMTOUpdatesParams) middleware.Responder {
	return fn(params)
}

// FetchMTOUpdatesHandler interface for that can handle valid fetch m t o updates params
type FetchMTOUpdatesHandler interface {
	Handle(FetchMTOUpdatesParams) middleware.Responder
}

// NewFetchMTOUpdates creates a new http.Handler for the fetch m t o updates operation
func NewFetchMTOUpdates(ctx *middleware.Context, handler FetchMTOUpdatesHandler) *FetchMTOUpdates {
	return &FetchMTOUpdates{Context: ctx, Handler: handler}
}

/* FetchMTOUpdates swagger:route GET /move-task-orders moveTaskOrder fetchMTOUpdates

fetchMTOUpdates

_[Deprecated: sunset on August 31, 2021]_ This endpoint is deprecated. Please use `listMoves`.

Gets all moves that have been reviewed and approved by the TOO. The `since` parameter can be used to filter this
list down to only the moves that have been updated since the provided timestamp. A move will be considered
updated if the `updatedAt` timestamp on the move is later than the provided date and time.

**WIP**: The original goal was to also look at the `updateAt` timestamps of the nested objects - such as the
shipments, service items, etc. This has not been implemented.

**WIP**: Include what causes moves to leave this list. Currently, once the `availableToPrimeAt` timestamp has
been set, that move will always appear in this list.


*/
type FetchMTOUpdates struct {
	Context *middleware.Context
	Handler FetchMTOUpdatesHandler
}

func (o *FetchMTOUpdates) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewFetchMTOUpdatesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
