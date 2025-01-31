// Code generated by go-swagger; DO NOT EDIT.

package adminoperations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/transcom/mymove/pkg/gen/adminapi/adminoperations/access_codes"
	"github.com/transcom/mymove/pkg/gen/adminapi/adminoperations/admin_users"
	"github.com/transcom/mymove/pkg/gen/adminapi/adminoperations/electronic_order"
	"github.com/transcom/mymove/pkg/gen/adminapi/adminoperations/move"
	"github.com/transcom/mymove/pkg/gen/adminapi/adminoperations/notification"
	"github.com/transcom/mymove/pkg/gen/adminapi/adminoperations/office"
	"github.com/transcom/mymove/pkg/gen/adminapi/adminoperations/office_users"
	"github.com/transcom/mymove/pkg/gen/adminapi/adminoperations/organization"
	"github.com/transcom/mymove/pkg/gen/adminapi/adminoperations/transportation_service_provider_performances"
	"github.com/transcom/mymove/pkg/gen/adminapi/adminoperations/upload"
	"github.com/transcom/mymove/pkg/gen/adminapi/adminoperations/users"
	"github.com/transcom/mymove/pkg/gen/adminapi/adminoperations/webhook_subscriptions"
)

// NewMymoveAPI creates a new Mymove instance
func NewMymoveAPI(spec *loads.Document) *MymoveAPI {
	return &MymoveAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		AdminUsersCreateAdminUserHandler: admin_users.CreateAdminUserHandlerFunc(func(params admin_users.CreateAdminUserParams) middleware.Responder {
			return middleware.NotImplemented("operation admin_users.CreateAdminUser has not yet been implemented")
		}),
		OfficeUsersCreateOfficeUserHandler: office_users.CreateOfficeUserHandlerFunc(func(params office_users.CreateOfficeUserParams) middleware.Responder {
			return middleware.NotImplemented("operation office_users.CreateOfficeUser has not yet been implemented")
		}),
		WebhookSubscriptionsCreateWebhookSubscriptionHandler: webhook_subscriptions.CreateWebhookSubscriptionHandlerFunc(func(params webhook_subscriptions.CreateWebhookSubscriptionParams) middleware.Responder {
			return middleware.NotImplemented("operation webhook_subscriptions.CreateWebhookSubscription has not yet been implemented")
		}),
		AdminUsersGetAdminUserHandler: admin_users.GetAdminUserHandlerFunc(func(params admin_users.GetAdminUserParams) middleware.Responder {
			return middleware.NotImplemented("operation admin_users.GetAdminUser has not yet been implemented")
		}),
		ElectronicOrderGetElectronicOrdersTotalsHandler: electronic_order.GetElectronicOrdersTotalsHandlerFunc(func(params electronic_order.GetElectronicOrdersTotalsParams) middleware.Responder {
			return middleware.NotImplemented("operation electronic_order.GetElectronicOrdersTotals has not yet been implemented")
		}),
		MoveGetMoveHandler: move.GetMoveHandlerFunc(func(params move.GetMoveParams) middleware.Responder {
			return middleware.NotImplemented("operation move.GetMove has not yet been implemented")
		}),
		OfficeUsersGetOfficeUserHandler: office_users.GetOfficeUserHandlerFunc(func(params office_users.GetOfficeUserParams) middleware.Responder {
			return middleware.NotImplemented("operation office_users.GetOfficeUser has not yet been implemented")
		}),
		TransportationServiceProviderPerformancesGetTSPPHandler: transportation_service_provider_performances.GetTSPPHandlerFunc(func(params transportation_service_provider_performances.GetTSPPParams) middleware.Responder {
			return middleware.NotImplemented("operation transportation_service_provider_performances.GetTSPP has not yet been implemented")
		}),
		UploadGetUploadHandler: upload.GetUploadHandlerFunc(func(params upload.GetUploadParams) middleware.Responder {
			return middleware.NotImplemented("operation upload.GetUpload has not yet been implemented")
		}),
		UsersGetUserHandler: users.GetUserHandlerFunc(func(params users.GetUserParams) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUser has not yet been implemented")
		}),
		WebhookSubscriptionsGetWebhookSubscriptionHandler: webhook_subscriptions.GetWebhookSubscriptionHandlerFunc(func(params webhook_subscriptions.GetWebhookSubscriptionParams) middleware.Responder {
			return middleware.NotImplemented("operation webhook_subscriptions.GetWebhookSubscription has not yet been implemented")
		}),
		AccessCodesIndexAccessCodesHandler: access_codes.IndexAccessCodesHandlerFunc(func(params access_codes.IndexAccessCodesParams) middleware.Responder {
			return middleware.NotImplemented("operation access_codes.IndexAccessCodes has not yet been implemented")
		}),
		AdminUsersIndexAdminUsersHandler: admin_users.IndexAdminUsersHandlerFunc(func(params admin_users.IndexAdminUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation admin_users.IndexAdminUsers has not yet been implemented")
		}),
		ElectronicOrderIndexElectronicOrdersHandler: electronic_order.IndexElectronicOrdersHandlerFunc(func(params electronic_order.IndexElectronicOrdersParams) middleware.Responder {
			return middleware.NotImplemented("operation electronic_order.IndexElectronicOrders has not yet been implemented")
		}),
		MoveIndexMovesHandler: move.IndexMovesHandlerFunc(func(params move.IndexMovesParams) middleware.Responder {
			return middleware.NotImplemented("operation move.IndexMoves has not yet been implemented")
		}),
		NotificationIndexNotificationsHandler: notification.IndexNotificationsHandlerFunc(func(params notification.IndexNotificationsParams) middleware.Responder {
			return middleware.NotImplemented("operation notification.IndexNotifications has not yet been implemented")
		}),
		OfficeUsersIndexOfficeUsersHandler: office_users.IndexOfficeUsersHandlerFunc(func(params office_users.IndexOfficeUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation office_users.IndexOfficeUsers has not yet been implemented")
		}),
		OfficeIndexOfficesHandler: office.IndexOfficesHandlerFunc(func(params office.IndexOfficesParams) middleware.Responder {
			return middleware.NotImplemented("operation office.IndexOffices has not yet been implemented")
		}),
		OrganizationIndexOrganizationsHandler: organization.IndexOrganizationsHandlerFunc(func(params organization.IndexOrganizationsParams) middleware.Responder {
			return middleware.NotImplemented("operation organization.IndexOrganizations has not yet been implemented")
		}),
		TransportationServiceProviderPerformancesIndexTSPPsHandler: transportation_service_provider_performances.IndexTSPPsHandlerFunc(func(params transportation_service_provider_performances.IndexTSPPsParams) middleware.Responder {
			return middleware.NotImplemented("operation transportation_service_provider_performances.IndexTSPPs has not yet been implemented")
		}),
		UsersIndexUsersHandler: users.IndexUsersHandlerFunc(func(params users.IndexUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation users.IndexUsers has not yet been implemented")
		}),
		WebhookSubscriptionsIndexWebhookSubscriptionsHandler: webhook_subscriptions.IndexWebhookSubscriptionsHandlerFunc(func(params webhook_subscriptions.IndexWebhookSubscriptionsParams) middleware.Responder {
			return middleware.NotImplemented("operation webhook_subscriptions.IndexWebhookSubscriptions has not yet been implemented")
		}),
		AdminUsersUpdateAdminUserHandler: admin_users.UpdateAdminUserHandlerFunc(func(params admin_users.UpdateAdminUserParams) middleware.Responder {
			return middleware.NotImplemented("operation admin_users.UpdateAdminUser has not yet been implemented")
		}),
		MoveUpdateMoveHandler: move.UpdateMoveHandlerFunc(func(params move.UpdateMoveParams) middleware.Responder {
			return middleware.NotImplemented("operation move.UpdateMove has not yet been implemented")
		}),
		OfficeUsersUpdateOfficeUserHandler: office_users.UpdateOfficeUserHandlerFunc(func(params office_users.UpdateOfficeUserParams) middleware.Responder {
			return middleware.NotImplemented("operation office_users.UpdateOfficeUser has not yet been implemented")
		}),
		UsersUpdateUserHandler: users.UpdateUserHandlerFunc(func(params users.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation users.UpdateUser has not yet been implemented")
		}),
		WebhookSubscriptionsUpdateWebhookSubscriptionHandler: webhook_subscriptions.UpdateWebhookSubscriptionHandlerFunc(func(params webhook_subscriptions.UpdateWebhookSubscriptionParams) middleware.Responder {
			return middleware.NotImplemented("operation webhook_subscriptions.UpdateWebhookSubscription has not yet been implemented")
		}),
	}
}

/*MymoveAPI The API for move.mil admin actions. */
type MymoveAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// AdminUsersCreateAdminUserHandler sets the operation handler for the create admin user operation
	AdminUsersCreateAdminUserHandler admin_users.CreateAdminUserHandler
	// OfficeUsersCreateOfficeUserHandler sets the operation handler for the create office user operation
	OfficeUsersCreateOfficeUserHandler office_users.CreateOfficeUserHandler
	// WebhookSubscriptionsCreateWebhookSubscriptionHandler sets the operation handler for the create webhook subscription operation
	WebhookSubscriptionsCreateWebhookSubscriptionHandler webhook_subscriptions.CreateWebhookSubscriptionHandler
	// AdminUsersGetAdminUserHandler sets the operation handler for the get admin user operation
	AdminUsersGetAdminUserHandler admin_users.GetAdminUserHandler
	// ElectronicOrderGetElectronicOrdersTotalsHandler sets the operation handler for the get electronic orders totals operation
	ElectronicOrderGetElectronicOrdersTotalsHandler electronic_order.GetElectronicOrdersTotalsHandler
	// MoveGetMoveHandler sets the operation handler for the get move operation
	MoveGetMoveHandler move.GetMoveHandler
	// OfficeUsersGetOfficeUserHandler sets the operation handler for the get office user operation
	OfficeUsersGetOfficeUserHandler office_users.GetOfficeUserHandler
	// TransportationServiceProviderPerformancesGetTSPPHandler sets the operation handler for the get t s p p operation
	TransportationServiceProviderPerformancesGetTSPPHandler transportation_service_provider_performances.GetTSPPHandler
	// UploadGetUploadHandler sets the operation handler for the get upload operation
	UploadGetUploadHandler upload.GetUploadHandler
	// UsersGetUserHandler sets the operation handler for the get user operation
	UsersGetUserHandler users.GetUserHandler
	// WebhookSubscriptionsGetWebhookSubscriptionHandler sets the operation handler for the get webhook subscription operation
	WebhookSubscriptionsGetWebhookSubscriptionHandler webhook_subscriptions.GetWebhookSubscriptionHandler
	// AccessCodesIndexAccessCodesHandler sets the operation handler for the index access codes operation
	AccessCodesIndexAccessCodesHandler access_codes.IndexAccessCodesHandler
	// AdminUsersIndexAdminUsersHandler sets the operation handler for the index admin users operation
	AdminUsersIndexAdminUsersHandler admin_users.IndexAdminUsersHandler
	// ElectronicOrderIndexElectronicOrdersHandler sets the operation handler for the index electronic orders operation
	ElectronicOrderIndexElectronicOrdersHandler electronic_order.IndexElectronicOrdersHandler
	// MoveIndexMovesHandler sets the operation handler for the index moves operation
	MoveIndexMovesHandler move.IndexMovesHandler
	// NotificationIndexNotificationsHandler sets the operation handler for the index notifications operation
	NotificationIndexNotificationsHandler notification.IndexNotificationsHandler
	// OfficeUsersIndexOfficeUsersHandler sets the operation handler for the index office users operation
	OfficeUsersIndexOfficeUsersHandler office_users.IndexOfficeUsersHandler
	// OfficeIndexOfficesHandler sets the operation handler for the index offices operation
	OfficeIndexOfficesHandler office.IndexOfficesHandler
	// OrganizationIndexOrganizationsHandler sets the operation handler for the index organizations operation
	OrganizationIndexOrganizationsHandler organization.IndexOrganizationsHandler
	// TransportationServiceProviderPerformancesIndexTSPPsHandler sets the operation handler for the index t s p ps operation
	TransportationServiceProviderPerformancesIndexTSPPsHandler transportation_service_provider_performances.IndexTSPPsHandler
	// UsersIndexUsersHandler sets the operation handler for the index users operation
	UsersIndexUsersHandler users.IndexUsersHandler
	// WebhookSubscriptionsIndexWebhookSubscriptionsHandler sets the operation handler for the index webhook subscriptions operation
	WebhookSubscriptionsIndexWebhookSubscriptionsHandler webhook_subscriptions.IndexWebhookSubscriptionsHandler
	// AdminUsersUpdateAdminUserHandler sets the operation handler for the update admin user operation
	AdminUsersUpdateAdminUserHandler admin_users.UpdateAdminUserHandler
	// MoveUpdateMoveHandler sets the operation handler for the update move operation
	MoveUpdateMoveHandler move.UpdateMoveHandler
	// OfficeUsersUpdateOfficeUserHandler sets the operation handler for the update office user operation
	OfficeUsersUpdateOfficeUserHandler office_users.UpdateOfficeUserHandler
	// UsersUpdateUserHandler sets the operation handler for the update user operation
	UsersUpdateUserHandler users.UpdateUserHandler
	// WebhookSubscriptionsUpdateWebhookSubscriptionHandler sets the operation handler for the update webhook subscription operation
	WebhookSubscriptionsUpdateWebhookSubscriptionHandler webhook_subscriptions.UpdateWebhookSubscriptionHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *MymoveAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *MymoveAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *MymoveAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *MymoveAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *MymoveAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *MymoveAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *MymoveAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *MymoveAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *MymoveAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the MymoveAPI
func (o *MymoveAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AdminUsersCreateAdminUserHandler == nil {
		unregistered = append(unregistered, "admin_users.CreateAdminUserHandler")
	}
	if o.OfficeUsersCreateOfficeUserHandler == nil {
		unregistered = append(unregistered, "office_users.CreateOfficeUserHandler")
	}
	if o.WebhookSubscriptionsCreateWebhookSubscriptionHandler == nil {
		unregistered = append(unregistered, "webhook_subscriptions.CreateWebhookSubscriptionHandler")
	}
	if o.AdminUsersGetAdminUserHandler == nil {
		unregistered = append(unregistered, "admin_users.GetAdminUserHandler")
	}
	if o.ElectronicOrderGetElectronicOrdersTotalsHandler == nil {
		unregistered = append(unregistered, "electronic_order.GetElectronicOrdersTotalsHandler")
	}
	if o.MoveGetMoveHandler == nil {
		unregistered = append(unregistered, "move.GetMoveHandler")
	}
	if o.OfficeUsersGetOfficeUserHandler == nil {
		unregistered = append(unregistered, "office_users.GetOfficeUserHandler")
	}
	if o.TransportationServiceProviderPerformancesGetTSPPHandler == nil {
		unregistered = append(unregistered, "transportation_service_provider_performances.GetTSPPHandler")
	}
	if o.UploadGetUploadHandler == nil {
		unregistered = append(unregistered, "upload.GetUploadHandler")
	}
	if o.UsersGetUserHandler == nil {
		unregistered = append(unregistered, "users.GetUserHandler")
	}
	if o.WebhookSubscriptionsGetWebhookSubscriptionHandler == nil {
		unregistered = append(unregistered, "webhook_subscriptions.GetWebhookSubscriptionHandler")
	}
	if o.AccessCodesIndexAccessCodesHandler == nil {
		unregistered = append(unregistered, "access_codes.IndexAccessCodesHandler")
	}
	if o.AdminUsersIndexAdminUsersHandler == nil {
		unregistered = append(unregistered, "admin_users.IndexAdminUsersHandler")
	}
	if o.ElectronicOrderIndexElectronicOrdersHandler == nil {
		unregistered = append(unregistered, "electronic_order.IndexElectronicOrdersHandler")
	}
	if o.MoveIndexMovesHandler == nil {
		unregistered = append(unregistered, "move.IndexMovesHandler")
	}
	if o.NotificationIndexNotificationsHandler == nil {
		unregistered = append(unregistered, "notification.IndexNotificationsHandler")
	}
	if o.OfficeUsersIndexOfficeUsersHandler == nil {
		unregistered = append(unregistered, "office_users.IndexOfficeUsersHandler")
	}
	if o.OfficeIndexOfficesHandler == nil {
		unregistered = append(unregistered, "office.IndexOfficesHandler")
	}
	if o.OrganizationIndexOrganizationsHandler == nil {
		unregistered = append(unregistered, "organization.IndexOrganizationsHandler")
	}
	if o.TransportationServiceProviderPerformancesIndexTSPPsHandler == nil {
		unregistered = append(unregistered, "transportation_service_provider_performances.IndexTSPPsHandler")
	}
	if o.UsersIndexUsersHandler == nil {
		unregistered = append(unregistered, "users.IndexUsersHandler")
	}
	if o.WebhookSubscriptionsIndexWebhookSubscriptionsHandler == nil {
		unregistered = append(unregistered, "webhook_subscriptions.IndexWebhookSubscriptionsHandler")
	}
	if o.AdminUsersUpdateAdminUserHandler == nil {
		unregistered = append(unregistered, "admin_users.UpdateAdminUserHandler")
	}
	if o.MoveUpdateMoveHandler == nil {
		unregistered = append(unregistered, "move.UpdateMoveHandler")
	}
	if o.OfficeUsersUpdateOfficeUserHandler == nil {
		unregistered = append(unregistered, "office_users.UpdateOfficeUserHandler")
	}
	if o.UsersUpdateUserHandler == nil {
		unregistered = append(unregistered, "users.UpdateUserHandler")
	}
	if o.WebhookSubscriptionsUpdateWebhookSubscriptionHandler == nil {
		unregistered = append(unregistered, "webhook_subscriptions.UpdateWebhookSubscriptionHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *MymoveAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *MymoveAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *MymoveAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *MymoveAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *MymoveAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *MymoveAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the mymove API
func (o *MymoveAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *MymoveAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/admin_users"] = admin_users.NewCreateAdminUser(o.context, o.AdminUsersCreateAdminUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/office_users"] = office_users.NewCreateOfficeUser(o.context, o.OfficeUsersCreateOfficeUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/webhook_subscriptions"] = webhook_subscriptions.NewCreateWebhookSubscription(o.context, o.WebhookSubscriptionsCreateWebhookSubscriptionHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/admin_users/{adminUserId}"] = admin_users.NewGetAdminUser(o.context, o.AdminUsersGetAdminUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/electronic_orders/totals"] = electronic_order.NewGetElectronicOrdersTotals(o.context, o.ElectronicOrderGetElectronicOrdersTotalsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/moves/{moveID}"] = move.NewGetMove(o.context, o.MoveGetMoveHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/office_users/{officeUserId}"] = office_users.NewGetOfficeUser(o.context, o.OfficeUsersGetOfficeUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/transportation_service_provider_performances/{tsppId}"] = transportation_service_provider_performances.NewGetTSPP(o.context, o.TransportationServiceProviderPerformancesGetTSPPHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/uploads/{uploadId}"] = upload.NewGetUpload(o.context, o.UploadGetUploadHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{userId}"] = users.NewGetUser(o.context, o.UsersGetUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/webhook_subscriptions/{webhookSubscriptionId}"] = webhook_subscriptions.NewGetWebhookSubscription(o.context, o.WebhookSubscriptionsGetWebhookSubscriptionHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/access_codes"] = access_codes.NewIndexAccessCodes(o.context, o.AccessCodesIndexAccessCodesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/admin_users"] = admin_users.NewIndexAdminUsers(o.context, o.AdminUsersIndexAdminUsersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/electronic_orders"] = electronic_order.NewIndexElectronicOrders(o.context, o.ElectronicOrderIndexElectronicOrdersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/moves"] = move.NewIndexMoves(o.context, o.MoveIndexMovesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/notifications"] = notification.NewIndexNotifications(o.context, o.NotificationIndexNotificationsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/office_users"] = office_users.NewIndexOfficeUsers(o.context, o.OfficeUsersIndexOfficeUsersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/offices"] = office.NewIndexOffices(o.context, o.OfficeIndexOfficesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/organizations"] = organization.NewIndexOrganizations(o.context, o.OrganizationIndexOrganizationsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/transportation_service_provider_performances"] = transportation_service_provider_performances.NewIndexTSPPs(o.context, o.TransportationServiceProviderPerformancesIndexTSPPsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users"] = users.NewIndexUsers(o.context, o.UsersIndexUsersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/webhook_subscriptions"] = webhook_subscriptions.NewIndexWebhookSubscriptions(o.context, o.WebhookSubscriptionsIndexWebhookSubscriptionsHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/admin_users/{adminUserId}"] = admin_users.NewUpdateAdminUser(o.context, o.AdminUsersUpdateAdminUserHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/moves/{moveID}"] = move.NewUpdateMove(o.context, o.MoveUpdateMoveHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/office_users/{officeUserId}"] = office_users.NewUpdateOfficeUser(o.context, o.OfficeUsersUpdateOfficeUserHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/users/{userId}"] = users.NewUpdateUser(o.context, o.UsersUpdateUserHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/webhook_subscriptions/{webhookSubscriptionId}"] = webhook_subscriptions.NewUpdateWebhookSubscription(o.context, o.WebhookSubscriptionsUpdateWebhookSubscriptionHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *MymoveAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *MymoveAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *MymoveAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *MymoveAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *MymoveAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
