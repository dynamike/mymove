package primeapi

import (
	"fmt"

	"github.com/gobuffalo/validate"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"

	paymentrequestop "github.com/transcom/mymove/pkg/gen/primeapi/primeoperations/payment_requests"
	"github.com/transcom/mymove/pkg/gen/primemessages"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/handlers/primeapi/internal/payloads"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
)

// CreatePaymentRequestHandler is the handler for creating payment requests
type CreatePaymentRequestHandler struct {
	handlers.HandlerContext
	services.PaymentRequestCreator
}

// Handle creates the payment request
func (h CreatePaymentRequestHandler) Handle(params paymentrequestop.CreatePaymentRequestParams) middleware.Responder {
	// TODO: authorization to create payment request

	logger := h.LoggerFromRequest(params.HTTPRequest)

	payload := params.Body
	if payload == nil {
		errPayload := payloads.ClientError(handlers.SQLErrMessage, "Invalid payment request: params Body is nil", h.GetTraceID())
		logger.Error("Invalid payment request: params Body is nil", zap.Any("payload", errPayload))
		return paymentrequestop.NewCreatePaymentRequestBadRequest().WithPayload(errPayload)
	}

	logger.Info("primeapi.CreatePaymentRequestHandler info", zap.String("pointOfContact", params.Body.PointOfContact))

	moveTaskOrderIDString := payload.MoveTaskOrderID.String()
	mtoID, err := uuid.FromString(moveTaskOrderIDString)
	if err != nil {
		logger.Error("Invalid payment request: params MoveTaskOrderID cannot be converted to a UUID",
			zap.String("MoveTaskOrderID", moveTaskOrderIDString), zap.Error(err))
		// create a custom verrs for returning a 422
		verrs :=
			&validate.Errors{Errors: map[string][]string{
				"move_task_order_id": {"id cannot be converted to UUID"},
			},
			}
		errPayload := payloads.ValidationError(err.Error(), h.GetTraceID(), verrs)
		return paymentrequestop.NewCreatePaymentRequestUnprocessableEntity().WithPayload(errPayload)
	}

	isFinal := false
	if payload.IsFinal != nil {
		isFinal = *payload.IsFinal
	}

	paymentRequest := models.PaymentRequest{
		IsFinal:         isFinal,
		MoveTaskOrderID: mtoID,
	}

	// Build up the paymentRequest.PaymentServiceItems using the incoming payload to offload Swagger data coming
	// in from the API. These paymentRequest.PaymentServiceItems will be used as a temp holder to process the incoming API data
	verrs := validate.NewErrors()
	paymentRequest.PaymentServiceItems, verrs, err = h.buildPaymentServiceItems(payload)
	if err != nil || verrs.HasAny() {

		logger.Error("could not build service items", zap.Error(err))
		// TODO: do not bail out before creating the payment request, we need the failed record
		//       we should create the failed record and store it as failed with a rejection
		errPayload := payloads.ValidationError(err.Error(), h.GetTraceID(), verrs)
		return paymentrequestop.NewCreatePaymentRequestUnprocessableEntity().WithPayload(errPayload)
	}

	createdPaymentRequest, err := h.PaymentRequestCreator.CreatePaymentRequest(&paymentRequest)
	if err != nil {
		logger.Error("Error creating payment request", zap.Error(err))
		if typedErr, ok := err.(services.InvalidCreateInputError); ok {
			verrs := typedErr.ValidationErrors
			detail := err.Error()
			payload := payloads.ValidationError(detail, h.GetTraceID(), verrs)

			logger.Error("Payment Request",
				zap.Any("payload", payload))
			return paymentrequestop.NewCreatePaymentRequestUnprocessableEntity().WithPayload(payload)
		}

		if _, ok := err.(services.NotFoundError); ok {
			payload := payloads.ClientError(handlers.NotFoundMessage, err.Error(), h.GetTraceID())

			logger.Error("Payment Request",
				zap.Any("payload", payload))
			return paymentrequestop.NewCreatePaymentRequestNotFound().WithPayload(payload)
		}
		if _, ok := err.(*services.BadDataError); ok {
			payload := payloads.ClientError(handlers.SQLErrMessage, err.Error(), h.GetTraceID())

			logger.Error("Payment Request",
				zap.Any("payload", payload))
			return paymentrequestop.NewCreatePaymentRequestBadRequest().WithPayload(payload)
		}
		logger.Error("Payment Request",
			zap.Any("payload", payload))
		return paymentrequestop.NewCreatePaymentRequestInternalServerError()
	}

	returnPayload := payloads.PaymentRequest(createdPaymentRequest)
	logger.Info("Successful payment request creation for mto ID", zap.String("moveID", moveTaskOrderIDString))
	return paymentrequestop.NewCreatePaymentRequestCreated().WithPayload(returnPayload)
}

func (h CreatePaymentRequestHandler) buildPaymentServiceItems(payload *primemessages.CreatePaymentRequestPayload) (models.PaymentServiceItems, *validate.Errors, error) {
	var paymentServiceItems models.PaymentServiceItems
	verrs := validate.NewErrors()
	for _, payloadServiceItem := range payload.ServiceItems {
		mtoServiceItemID, err := uuid.FromString(payloadServiceItem.ID.String())
		if err != nil {
			// create a custom verrs for returning a 422
			verrs = &validate.Errors{Errors: map[string][]string{
				"payment_service_item_id": {"id cannot be converted to UUID"},
			},
			}
			return nil, verrs, fmt.Errorf("could not convert service item ID [%v] to UUID: %w", payloadServiceItem.ID, err)
		}

		paymentServiceItem := models.PaymentServiceItem{
			// The rest of the model will be filled in when the payment request is created
			MTOServiceItemID: mtoServiceItemID,
		}

		paymentServiceItem.PaymentServiceItemParams = h.buildPaymentServiceItemParams(payloadServiceItem)

		paymentServiceItems = append(paymentServiceItems, paymentServiceItem)
	}

	return paymentServiceItems, verrs, nil
}

func (h CreatePaymentRequestHandler) buildPaymentServiceItemParams(payloadMTOServiceItem *primemessages.ServiceItem) models.PaymentServiceItemParams {
	var paymentServiceItemParams models.PaymentServiceItemParams

	for _, payloadServiceItemParam := range payloadMTOServiceItem.Params {
		paymentServiceItemParam := models.PaymentServiceItemParam{
			// ID and PaymentServiceItemID to be filled in when payment request is created
			IncomingKey: payloadServiceItemParam.Key,
			Value:       payloadServiceItemParam.Value,
		}

		paymentServiceItemParams = append(paymentServiceItemParams, paymentServiceItemParam)
	}

	return paymentServiceItemParams
}
