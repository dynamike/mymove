package ghcapi

import (
	"fmt"
	"net/http/httptest"
	"time"

	moverouter "github.com/transcom/mymove/pkg/services/move"

	"github.com/transcom/mymove/pkg/models/roles"

	"github.com/go-openapi/strfmt"

	"github.com/gofrs/uuid"

	routemocks "github.com/transcom/mymove/pkg/route/mocks"

	"github.com/transcom/mymove/pkg/services"

	"github.com/gobuffalo/validate/v3"

	mtoserviceitem "github.com/transcom/mymove/pkg/services/mto_service_item"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"

	"github.com/transcom/mymove/pkg/etag"
	mtoshipmentops "github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/mto_shipment"
	shipmentops "github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/shipment"
	"github.com/transcom/mymove/pkg/gen/ghcmessages"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services/fetch"
	"github.com/transcom/mymove/pkg/services/mocks"
	mtoshipment "github.com/transcom/mymove/pkg/services/mto_shipment"

	"github.com/transcom/mymove/pkg/services/query"
	"github.com/transcom/mymove/pkg/testdatagen"
)

type listMTOShipmentsSubtestData struct {
	mtoAgent       models.MTOAgent
	mtoServiceItem models.MTOServiceItem
	shipments      models.MTOShipments
	params         mtoshipmentops.ListMTOShipmentsParams
}

func (suite *HandlerSuite) makeListMTOShipmentsSubtestData() (subtestData *listMTOShipmentsSubtestData) {
	subtestData = &listMTOShipmentsSubtestData{}

	mto := testdatagen.MakeDefaultMove(suite.DB())
	mtoShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move: mto,
		MTOShipment: models.MTOShipment{
			Status:           models.MTOShipmentStatusSubmitted,
			CounselorRemarks: handlers.FmtString("counselor remark"),
		},
	})
	subtestData.mtoAgent = testdatagen.MakeMTOAgent(suite.DB(), testdatagen.Assertions{
		MTOAgent: models.MTOAgent{
			MTOShipmentID: mtoShipment.ID,
		},
	})
	subtestData.mtoServiceItem = testdatagen.MakeMTOServiceItem(suite.DB(), testdatagen.Assertions{
		MTOServiceItem: models.MTOServiceItem{
			MTOShipmentID: &mtoShipment.ID,
		},
	})

	subtestData.shipments = models.MTOShipments{mtoShipment}
	requestUser := testdatagen.MakeStubbedUser(suite.DB())

	req := httptest.NewRequest("GET", fmt.Sprintf("/move_task_orders/%s/mto_shipments", mto.ID.String()), nil)
	req = suite.AuthenticateUserRequest(req, requestUser)

	subtestData.params = mtoshipmentops.ListMTOShipmentsParams{
		HTTPRequest:     req,
		MoveTaskOrderID: *handlers.FmtUUID(mtoShipment.MoveTaskOrderID),
	}

	return subtestData
}

func (suite *HandlerSuite) TestListMTOShipmentsHandler() {
	suite.Run("Successful list fetch - Integration Test", func() {
		subtestData := suite.makeListMTOShipmentsSubtestData()
		params := subtestData.params
		shipments := subtestData.shipments
		mtoAgent := subtestData.mtoAgent
		mtoServiceItem := subtestData.mtoServiceItem

		queryBuilder := query.NewQueryBuilder(suite.DB())
		listFetcher := fetch.NewListFetcher(queryBuilder)
		fetcher := fetch.NewFetcher(queryBuilder)
		handler := ListMTOShipmentsHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			listFetcher,
			fetcher,
		}

		response := handler.Handle(params)
		suite.IsType(&mtoshipmentops.ListMTOShipmentsOK{}, response)

		okResponse := response.(*mtoshipmentops.ListMTOShipmentsOK)
		suite.Len(okResponse.Payload, 1)
		suite.Equal(shipments[0].ID.String(), okResponse.Payload[0].ID.String())
		suite.Equal(*shipments[0].CounselorRemarks, *okResponse.Payload[0].CounselorRemarks)
		suite.Equal(mtoAgent.ID.String(), okResponse.Payload[0].MtoAgents[0].ID.String())
		suite.Equal(mtoServiceItem.ID.String(), okResponse.Payload[0].MtoServiceItems[0].ID.String())
	})

	suite.Run("Failure list fetch - Internal Server Error", func() {
		subtestData := suite.makeListMTOShipmentsSubtestData()
		params := subtestData.params

		mockListFetcher := mocks.ListFetcher{}
		mockFetcher := mocks.Fetcher{}
		handler := ListMTOShipmentsHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			&mockListFetcher,
			&mockFetcher,
		}

		internalServerErr := errors.New("ServerError")

		mockFetcher.On("FetchRecord",
			mock.Anything,
			mock.Anything,
		).Return(nil)

		mockListFetcher.On("FetchRecordList",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(internalServerErr)

		response := handler.Handle(params)
		suite.IsType(&mtoshipmentops.ListMTOShipmentsInternalServerError{}, response)
	})

	suite.Run("Failure list fetch - 404 Not Found - Move Task Order ID", func() {
		subtestData := suite.makeListMTOShipmentsSubtestData()
		params := subtestData.params

		mockListFetcher := mocks.ListFetcher{}
		mockFetcher := mocks.Fetcher{}
		handler := ListMTOShipmentsHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			&mockListFetcher,
			&mockFetcher,
		}

		notfound := errors.New("Not found error")

		mockFetcher.On("FetchRecord",
			mock.Anything,
			mock.Anything,
		).Return(notfound)

		response := handler.Handle(params)
		suite.IsType(&mtoshipmentops.ListMTOShipmentsNotFound{}, response)
	})
}

func (suite *HandlerSuite) TestDeleteShipmentHandler() {
	suite.Run("Returns a 403 when the office user is not a service counselor", func() {
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		uuid := uuid.Must(uuid.NewV4())
		deleter := &mocks.ShipmentDeleter{}

		deleter.AssertNumberOfCalls(suite.T(), "DeleteShipment", 0)

		req := httptest.NewRequest("DELETE", fmt.Sprintf("/shipments/%s", uuid.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := DeleteShipmentHandler{
			handlerContext,
			deleter,
		}
		deletionParams := shipmentops.DeleteShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(uuid),
		}

		response := handler.Handle(deletionParams)
		suite.IsType(&shipmentops.DeleteShipmentForbidden{}, response)
	})

	suite.Run("Returns 204 when all validations pass", func() {
		shipment := testdatagen.MakeDefaultMTOShipmentMinimal(suite.DB())
		officeUser := testdatagen.MakeServicesCounselorOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		deleter := &mocks.ShipmentDeleter{}

		deleter.On("DeleteShipment", shipment.ID).Return(shipment.MoveTaskOrderID, nil)

		req := httptest.NewRequest("DELETE", fmt.Sprintf("/shipments/%s", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := DeleteShipmentHandler{
			handlerContext,
			deleter,
		}
		deletionParams := shipmentops.DeleteShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
		}

		response := handler.Handle(deletionParams)

		suite.IsType(&shipmentops.DeleteShipmentNoContent{}, response)
	})

	suite.Run("Returns 404 when deleter returns NotFoundError", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		officeUser := testdatagen.MakeServicesCounselorOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		deleter := &mocks.ShipmentDeleter{}

		deleter.On("DeleteShipment", shipment.ID).Return(uuid.Nil, services.NotFoundError{})

		req := httptest.NewRequest("DELETE", fmt.Sprintf("/shipments/%s", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := DeleteShipmentHandler{
			handlerContext,
			deleter,
		}
		deletionParams := shipmentops.DeleteShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
		}

		response := handler.Handle(deletionParams)
		suite.IsType(&shipmentops.DeleteShipmentNotFound{}, response)
	})

	suite.Run("Returns 403 when deleter returns ForbiddenError", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		officeUser := testdatagen.MakeServicesCounselorOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		deleter := &mocks.ShipmentDeleter{}

		deleter.On("DeleteShipment", shipment.ID).Return(uuid.Nil, services.ForbiddenError{})

		req := httptest.NewRequest("DELETE", fmt.Sprintf("/shipments/%s", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := DeleteShipmentHandler{
			handlerContext,
			deleter,
		}
		deletionParams := shipmentops.DeleteShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
		}

		response := handler.Handle(deletionParams)
		suite.IsType(&shipmentops.DeleteShipmentForbidden{}, response)
	})
}

func (suite *HandlerSuite) TestApproveShipmentHandler() {
	suite.Run("Returns 200 when all validations pass", func() {
		move := testdatagen.MakeAvailableMove(suite.DB())
		shipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			Move: move,
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusSubmitted,
			},
		})
		// Populate the reServices table with codes needed by the
		// HHG_LONGHAUL_DOMESTIC shipment type
		reServiceCodes := []models.ReServiceCode{
			models.ReServiceCodeDLH,
			models.ReServiceCodeFSC,
			models.ReServiceCodeDOP,
			models.ReServiceCodeDDP,
			models.ReServiceCodeDPK,
			models.ReServiceCodeDUPK,
		}
		for _, serviceCode := range reServiceCodes {
			testdatagen.MakeReService(suite.DB(), testdatagen.Assertions{
				ReService: models.ReService{
					Code:      serviceCode,
					Name:      "test",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			})
		}

		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		builder := query.NewQueryBuilder(suite.DB())
		moveRouter := moverouter.NewMoveRouter(suite.DB(), suite.TestLogger())
		approver := mtoshipment.NewShipmentApprover(
			suite.DB(),
			mtoshipment.NewShipmentRouter(suite.DB()),
			mtoserviceitem.NewMTOServiceItemCreator(builder, moveRouter),
			&routemocks.Planner{},
		)

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/approve", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
		handlerContext.SetTraceID(uuid.Must(uuid.NewV4()))

		handler := ApproveShipmentHandler{
			handlerContext,
			approver,
		}

		approveParams := shipmentops.ApproveShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.ApproveShipmentOK{}, response)
		suite.HasWebhookNotification(shipment.ID, handlerContext.GetTraceID())
	})

	suite.Run("Returns a 403 when the office user is not a TOO", func() {
		officeUser := testdatagen.MakeServicesCounselorOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		uuid := uuid.Must(uuid.NewV4())
		approver := &mocks.ShipmentApprover{}

		approver.AssertNumberOfCalls(suite.T(), "ApproveShipment", 0)

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/approve", uuid.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := ApproveShipmentHandler{
			handlerContext,
			approver,
		}
		approveParams := shipmentops.ApproveShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(uuid),
			IfMatch:     etag.GenerateEtag(time.Now()),
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.ApproveShipmentForbidden{}, response)
	})

	suite.Run("Returns 404 when approver returns NotFoundError", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		approver := &mocks.ShipmentApprover{}

		approver.On("ApproveShipment", shipment.ID, eTag).Return(nil, services.NotFoundError{})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/approve", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := ApproveShipmentHandler{
			handlerContext,
			approver,
		}
		approveParams := shipmentops.ApproveShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.ApproveShipmentNotFound{}, response)
	})

	suite.Run("Returns 409 when approver returns Conflict Error", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		approver := &mocks.ShipmentApprover{}

		approver.On("ApproveShipment", shipment.ID, eTag).Return(nil, services.ConflictError{})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/approve", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := ApproveShipmentHandler{
			handlerContext,
			approver,
		}
		approveParams := shipmentops.ApproveShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.ApproveShipmentConflict{}, response)
	})

	suite.Run("Returns 412 when eTag does not match", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(time.Now())
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		approver := &mocks.ShipmentApprover{}

		approver.On("ApproveShipment", shipment.ID, eTag).Return(nil, services.PreconditionFailedError{})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/approve", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := ApproveShipmentHandler{
			handlerContext,
			approver,
		}
		approveParams := shipmentops.ApproveShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.ApproveShipmentPreconditionFailed{}, response)
	})

	suite.Run("Returns 422 when approver returns validation errors", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		approver := &mocks.ShipmentApprover{}

		approver.On("ApproveShipment", shipment.ID, eTag).Return(nil, services.InvalidInputError{ValidationErrors: &validate.Errors{}})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/approve", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := ApproveShipmentHandler{
			handlerContext,
			approver,
		}
		approveParams := shipmentops.ApproveShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.ApproveShipmentUnprocessableEntity{}, response)
	})

	suite.Run("Returns 500 when approver returns unexpected error", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		approver := &mocks.ShipmentApprover{}

		approver.On("ApproveShipment", shipment.ID, eTag).Return(nil, errors.New("UnexpectedError"))

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/approve", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := ApproveShipmentHandler{
			handlerContext,
			approver,
		}
		approveParams := shipmentops.ApproveShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.ApproveShipmentInternalServerError{}, response)
	})
}

func (suite *HandlerSuite) TestRequestShipmentDiversionHandler() {
	suite.Run("Returns 200 when all validations pass", func() {
		move := testdatagen.MakeAvailableMove(suite.DB())
		shipment := testdatagen.MakeMTOShipmentMinimal(suite.DB(), testdatagen.Assertions{
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusApproved,
			},
			Move: move,
		})

		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		requester := mtoshipment.NewShipmentDiversionRequester(
			suite.DB(),
			mtoshipment.NewShipmentRouter(suite.DB()),
		)

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-diversion", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
		handlerContext.SetTraceID(uuid.Must(uuid.NewV4()))

		handler := RequestShipmentDiversionHandler{
			handlerContext,
			requester,
		}

		approveParams := shipmentops.RequestShipmentDiversionParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.RequestShipmentDiversionOK{}, response)
		suite.HasWebhookNotification(shipment.ID, handlerContext.GetTraceID())
	})

	suite.Run("Returns a 403 when the office user is not a TOO", func() {
		officeUser := testdatagen.MakeServicesCounselorOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		uuid := uuid.Must(uuid.NewV4())
		requester := &mocks.ShipmentDiversionRequester{}

		requester.AssertNumberOfCalls(suite.T(), "RequestShipmentDiversion", 0)

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-diversion", uuid.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentDiversionHandler{
			handlerContext,
			requester,
		}
		approveParams := shipmentops.RequestShipmentDiversionParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(uuid),
			IfMatch:     etag.GenerateEtag(time.Now()),
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.RequestShipmentDiversionForbidden{}, response)
	})

	suite.Run("Returns 404 when requester returns NotFoundError", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		requester := &mocks.ShipmentDiversionRequester{}

		requester.On("RequestShipmentDiversion", shipment.ID, eTag).Return(nil, services.NotFoundError{})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-diversion", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentDiversionHandler{
			handlerContext,
			requester,
		}
		approveParams := shipmentops.RequestShipmentDiversionParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.RequestShipmentDiversionNotFound{}, response)
	})

	suite.Run("Returns 409 when requester returns Conflict Error", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		requester := &mocks.ShipmentDiversionRequester{}

		requester.On("RequestShipmentDiversion", shipment.ID, eTag).Return(nil, mtoshipment.ConflictStatusError{})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-diversion", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentDiversionHandler{
			handlerContext,
			requester,
		}
		approveParams := shipmentops.RequestShipmentDiversionParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.RequestShipmentDiversionConflict{}, response)
	})

	suite.Run("Returns 412 when eTag does not match", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(time.Now())
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		requester := &mocks.ShipmentDiversionRequester{}

		requester.On("RequestShipmentDiversion", shipment.ID, eTag).Return(nil, services.PreconditionFailedError{})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-diversion", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentDiversionHandler{
			handlerContext,
			requester,
		}
		approveParams := shipmentops.RequestShipmentDiversionParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.RequestShipmentDiversionPreconditionFailed{}, response)
	})

	suite.Run("Returns 422 when requester returns validation errors", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		requester := &mocks.ShipmentDiversionRequester{}

		requester.On("RequestShipmentDiversion", shipment.ID, eTag).Return(nil, services.InvalidInputError{ValidationErrors: &validate.Errors{}})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-diversion", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentDiversionHandler{
			handlerContext,
			requester,
		}
		approveParams := shipmentops.RequestShipmentDiversionParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.RequestShipmentDiversionUnprocessableEntity{}, response)
	})

	suite.Run("Returns 500 when requester returns unexpected error", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		requester := &mocks.ShipmentDiversionRequester{}

		requester.On("RequestShipmentDiversion", shipment.ID, eTag).Return(nil, errors.New("UnexpectedError"))

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-diversion", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentDiversionHandler{
			handlerContext,
			requester,
		}
		approveParams := shipmentops.RequestShipmentDiversionParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.RequestShipmentDiversionInternalServerError{}, response)
	})
}

func (suite *HandlerSuite) TestApproveShipmentDiversionHandler() {
	suite.Run("Returns 200 when all validations pass", func() {
		move := testdatagen.MakeAvailableMove(suite.DB())
		shipment := testdatagen.MakeMTOShipmentMinimal(suite.DB(), testdatagen.Assertions{
			MTOShipment: models.MTOShipment{
				Status:    models.MTOShipmentStatusSubmitted,
				Diversion: true,
			},
			Move: move,
		})

		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		approver := mtoshipment.NewShipmentDiversionApprover(
			suite.DB(),
			mtoshipment.NewShipmentRouter(suite.DB()),
		)

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/approve-diversion", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
		handlerContext.SetTraceID(uuid.Must(uuid.NewV4()))

		handler := ApproveShipmentDiversionHandler{
			handlerContext,
			approver,
		}

		approveParams := shipmentops.ApproveShipmentDiversionParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.ApproveShipmentDiversionOK{}, response)
		suite.HasWebhookNotification(shipment.ID, handlerContext.GetTraceID())
	})

	suite.Run("Returns a 403 when the office user is not a TOO", func() {
		officeUser := testdatagen.MakeServicesCounselorOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		uuid := uuid.Must(uuid.NewV4())
		approver := &mocks.ShipmentDiversionApprover{}

		approver.AssertNumberOfCalls(suite.T(), "ApproveShipmentDiversion", 0)

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/approve-diversion", uuid.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := ApproveShipmentDiversionHandler{
			handlerContext,
			approver,
		}
		approveParams := shipmentops.ApproveShipmentDiversionParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(uuid),
			IfMatch:     etag.GenerateEtag(time.Now()),
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.ApproveShipmentDiversionForbidden{}, response)
	})

	suite.Run("Returns 404 when approver returns NotFoundError", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		approver := &mocks.ShipmentDiversionApprover{}

		approver.On("ApproveShipmentDiversion", shipment.ID, eTag).Return(nil, services.NotFoundError{})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/approve-diversion", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := ApproveShipmentDiversionHandler{
			handlerContext,
			approver,
		}
		approveParams := shipmentops.ApproveShipmentDiversionParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.ApproveShipmentDiversionNotFound{}, response)
	})

	suite.Run("Returns 409 when approver returns Conflict Error", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		approver := &mocks.ShipmentDiversionApprover{}

		approver.On("ApproveShipmentDiversion", shipment.ID, eTag).Return(nil, mtoshipment.ConflictStatusError{})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/approve-diversion", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := ApproveShipmentDiversionHandler{
			handlerContext,
			approver,
		}
		approveParams := shipmentops.ApproveShipmentDiversionParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.ApproveShipmentDiversionConflict{}, response)
	})

	suite.Run("Returns 412 when eTag does not match", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(time.Now())
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		approver := &mocks.ShipmentDiversionApprover{}

		approver.On("ApproveShipmentDiversion", shipment.ID, eTag).Return(nil, services.PreconditionFailedError{})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/approve-diversion", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := ApproveShipmentDiversionHandler{
			handlerContext,
			approver,
		}
		approveParams := shipmentops.ApproveShipmentDiversionParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.ApproveShipmentDiversionPreconditionFailed{}, response)
	})

	suite.Run("Returns 422 when approver returns validation errors", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		approver := &mocks.ShipmentDiversionApprover{}

		approver.On("ApproveShipmentDiversion", shipment.ID, eTag).Return(nil, services.InvalidInputError{ValidationErrors: &validate.Errors{}})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/approve-diversion", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := ApproveShipmentDiversionHandler{
			handlerContext,
			approver,
		}
		approveParams := shipmentops.ApproveShipmentDiversionParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.ApproveShipmentDiversionUnprocessableEntity{}, response)
	})

	suite.Run("Returns 500 when approver returns unexpected error", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		approver := &mocks.ShipmentDiversionApprover{}

		approver.On("ApproveShipmentDiversion", shipment.ID, eTag).Return(nil, errors.New("UnexpectedError"))

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/approve-diversion", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := ApproveShipmentDiversionHandler{
			handlerContext,
			approver,
		}
		approveParams := shipmentops.ApproveShipmentDiversionParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.ApproveShipmentDiversionInternalServerError{}, response)
	})
}

func (suite *HandlerSuite) TestRejectShipmentHandler() {
	reason := "reason"

	suite.Run("Returns 200 when all validations pass", func() {
		move := testdatagen.MakeAvailableMove(suite.DB())
		shipment := testdatagen.MakeMTOShipmentMinimal(suite.DB(), testdatagen.Assertions{
			Move: move,
		})

		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		rejecter := mtoshipment.NewShipmentRejecter(
			suite.DB(),
			mtoshipment.NewShipmentRouter(suite.DB()),
		)

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/reject", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
		handlerContext.SetTraceID(uuid.Must(uuid.NewV4()))

		handler := RejectShipmentHandler{
			handlerContext,
			rejecter,
		}

		params := shipmentops.RejectShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
			Body: &ghcmessages.RejectShipment{
				RejectionReason: &reason,
			},
		}

		suite.NoError(params.Body.Validate(strfmt.Default))

		response := handler.Handle(params)
		suite.IsType(&shipmentops.RejectShipmentOK{}, response)
		suite.HasWebhookNotification(shipment.ID, handlerContext.GetTraceID())
	})

	suite.Run("Returns a 403 when the office user is not a TOO", func() {
		officeUser := testdatagen.MakeServicesCounselorOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		uuid := uuid.Must(uuid.NewV4())
		rejecter := &mocks.ShipmentRejecter{}

		rejecter.AssertNumberOfCalls(suite.T(), "RejectShipment", 0)

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/reject", uuid.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RejectShipmentHandler{
			handlerContext,
			rejecter,
		}
		params := shipmentops.RejectShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(uuid),
			IfMatch:     etag.GenerateEtag(time.Now()),
			Body: &ghcmessages.RejectShipment{
				RejectionReason: &reason,
			},
		}

		suite.NoError(params.Body.Validate(strfmt.Default))

		response := handler.Handle(params)
		suite.IsType(&shipmentops.RejectShipmentForbidden{}, response)
	})

	suite.Run("Returns 404 when rejecter returns NotFoundError", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		rejecter := &mocks.ShipmentRejecter{}

		rejecter.On("RejectShipment", shipment.ID, eTag, &reason).Return(nil, services.NotFoundError{})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/reject", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RejectShipmentHandler{
			handlerContext,
			rejecter,
		}
		params := shipmentops.RejectShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
			Body: &ghcmessages.RejectShipment{
				RejectionReason: &reason,
			},
		}

		suite.NoError(params.Body.Validate(strfmt.Default))

		response := handler.Handle(params)
		suite.IsType(&shipmentops.RejectShipmentNotFound{}, response)
	})

	suite.Run("Returns 409 when rejecter returns Conflict Error", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		rejecter := &mocks.ShipmentRejecter{}

		rejecter.On("RejectShipment", shipment.ID, eTag, &reason).Return(nil, mtoshipment.ConflictStatusError{})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/reject", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RejectShipmentHandler{
			handlerContext,
			rejecter,
		}
		params := shipmentops.RejectShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
			Body: &ghcmessages.RejectShipment{
				RejectionReason: &reason,
			},
		}

		suite.NoError(params.Body.Validate(strfmt.Default))

		response := handler.Handle(params)
		suite.IsType(&shipmentops.RejectShipmentConflict{}, response)
	})

	suite.Run("Returns 412 when eTag does not match", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(time.Now())
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		rejecter := &mocks.ShipmentRejecter{}

		rejecter.On("RejectShipment", shipment.ID, eTag, &reason).Return(nil, services.PreconditionFailedError{})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/reject", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RejectShipmentHandler{
			handlerContext,
			rejecter,
		}
		params := shipmentops.RejectShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
			Body: &ghcmessages.RejectShipment{
				RejectionReason: &reason,
			},
		}

		suite.NoError(params.Body.Validate(strfmt.Default))

		response := handler.Handle(params)
		suite.IsType(&shipmentops.RejectShipmentPreconditionFailed{}, response)
	})

	suite.Run("Returns 422 when rejecter returns validation errors", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		rejecter := &mocks.ShipmentRejecter{}

		rejecter.On("RejectShipment", shipment.ID, eTag, &reason).Return(nil, services.InvalidInputError{ValidationErrors: &validate.Errors{}})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/reject", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RejectShipmentHandler{
			handlerContext,
			rejecter,
		}
		params := shipmentops.RejectShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
			Body: &ghcmessages.RejectShipment{
				RejectionReason: &reason,
			},
		}

		suite.NoError(params.Body.Validate(strfmt.Default))

		response := handler.Handle(params)
		suite.IsType(&shipmentops.RejectShipmentUnprocessableEntity{}, response)
	})

	suite.Run("Returns 500 when rejecter returns unexpected error", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		rejecter := &mocks.ShipmentRejecter{}

		rejecter.On("RejectShipment", shipment.ID, eTag, &reason).Return(nil, errors.New("UnexpectedError"))

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/reject", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RejectShipmentHandler{
			handlerContext,
			rejecter,
		}
		params := shipmentops.RejectShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
			Body: &ghcmessages.RejectShipment{
				RejectionReason: &reason,
			},
		}

		suite.NoError(params.Body.Validate(strfmt.Default))

		response := handler.Handle(params)
		suite.IsType(&shipmentops.RejectShipmentInternalServerError{}, response)
	})

	suite.Run("Requires rejection reason in Body of request", func() {
		move := testdatagen.MakeAvailableMove(suite.DB())
		shipment := testdatagen.MakeMTOShipmentMinimal(suite.DB(), testdatagen.Assertions{
			Move: move,
		})
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		rejecter := mtoshipment.NewShipmentRejecter(
			suite.DB(),
			mtoshipment.NewShipmentRouter(suite.DB()),
		)

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/reject", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RejectShipmentHandler{
			handlerContext,
			rejecter,
		}
		params := shipmentops.RejectShipmentParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
			Body:        &ghcmessages.RejectShipment{},
		}

		suite.Error(params.Body.Validate(strfmt.Default))

		response := handler.Handle(params)
		suite.IsType(&shipmentops.RejectShipmentUnprocessableEntity{}, response)
	})
}

func (suite *HandlerSuite) TestRequestShipmentCancellationHandler() {
	suite.Run("Returns 200 when all validations pass", func() {
		move := testdatagen.MakeAvailableMove(suite.DB())
		shipment := testdatagen.MakeMTOShipmentMinimal(suite.DB(), testdatagen.Assertions{
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusApproved,
			},
			Move: move,
		})

		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		canceler := mtoshipment.NewShipmentCancellationRequester(
			suite.DB(),
			mtoshipment.NewShipmentRouter(suite.DB()),
		)

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-cancellation", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
		handlerContext.SetTraceID(uuid.Must(uuid.NewV4()))

		handler := RequestShipmentCancellationHandler{
			handlerContext,
			canceler,
		}

		approveParams := shipmentops.RequestShipmentCancellationParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.RequestShipmentCancellationOK{}, response)
		suite.HasWebhookNotification(shipment.ID, handlerContext.GetTraceID())
	})

	suite.Run("Returns a 403 when the office user is not a TOO", func() {
		officeUser := testdatagen.MakeServicesCounselorOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		uuid := uuid.Must(uuid.NewV4())
		canceler := &mocks.ShipmentCancellationRequester{}

		canceler.AssertNumberOfCalls(suite.T(), "RequestShipmentCancellation", 0)

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-cancellation", uuid.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentCancellationHandler{
			handlerContext,
			canceler,
		}
		approveParams := shipmentops.RequestShipmentCancellationParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(uuid),
			IfMatch:     etag.GenerateEtag(time.Now()),
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.RequestShipmentCancellationForbidden{}, response)
	})

	suite.Run("Returns 404 when canceler returns NotFoundError", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		canceler := &mocks.ShipmentCancellationRequester{}

		canceler.On("RequestShipmentCancellation", shipment.ID, eTag).Return(nil, services.NotFoundError{})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-cancellation", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentCancellationHandler{
			handlerContext,
			canceler,
		}
		approveParams := shipmentops.RequestShipmentCancellationParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.RequestShipmentCancellationNotFound{}, response)
	})

	suite.Run("Returns 409 when canceler returns Conflict Error", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		canceler := &mocks.ShipmentCancellationRequester{}

		canceler.On("RequestShipmentCancellation", shipment.ID, eTag).Return(nil, mtoshipment.ConflictStatusError{})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-cancellation", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentCancellationHandler{
			handlerContext,
			canceler,
		}
		approveParams := shipmentops.RequestShipmentCancellationParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.RequestShipmentCancellationConflict{}, response)
	})

	suite.Run("Returns 412 when eTag does not match", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(time.Now())
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		canceler := &mocks.ShipmentCancellationRequester{}

		canceler.On("RequestShipmentCancellation", shipment.ID, eTag).Return(nil, services.PreconditionFailedError{})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-cancellation", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentCancellationHandler{
			handlerContext,
			canceler,
		}
		approveParams := shipmentops.RequestShipmentCancellationParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.RequestShipmentCancellationPreconditionFailed{}, response)
	})

	suite.Run("Returns 422 when canceler returns validation errors", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		canceler := &mocks.ShipmentCancellationRequester{}

		canceler.On("RequestShipmentCancellation", shipment.ID, eTag).Return(nil, services.InvalidInputError{ValidationErrors: &validate.Errors{}})

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-cancellation", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentCancellationHandler{
			handlerContext,
			canceler,
		}
		approveParams := shipmentops.RequestShipmentCancellationParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.RequestShipmentCancellationUnprocessableEntity{}, response)
	})

	suite.Run("Returns 500 when canceler returns unexpected error", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		eTag := etag.GenerateEtag(shipment.UpdatedAt)
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		canceler := &mocks.ShipmentCancellationRequester{}

		canceler.On("RequestShipmentCancellation", shipment.ID, eTag).Return(nil, errors.New("UnexpectedError"))

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-cancellation", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentCancellationHandler{
			handlerContext,
			canceler,
		}
		approveParams := shipmentops.RequestShipmentCancellationParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
			IfMatch:     eTag,
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.RequestShipmentCancellationInternalServerError{}, response)
	})
}

func (suite *HandlerSuite) TestRequestShipmentReweighHandler() {
	suite.Run("Returns 200 when all validations pass", func() {
		move := testdatagen.MakeAvailableMove(suite.DB())
		shipment := testdatagen.MakeMTOShipmentMinimal(suite.DB(), testdatagen.Assertions{
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusApproved,
			},
			Move: move,
		})

		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		reweighRequester := mtoshipment.NewShipmentReweighRequester(
			suite.DB(),
		)

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-reweigh", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
		handlerContext.SetTraceID(uuid.Must(uuid.NewV4()))

		handler := RequestShipmentReweighHandler{
			handlerContext,
			reweighRequester,
		}

		approveParams := shipmentops.RequestShipmentReweighParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
		}

		response := handler.Handle(approveParams)
		okResponse := response.(*shipmentops.RequestShipmentReweighOK)
		payload := okResponse.Payload
		suite.IsType(&shipmentops.RequestShipmentReweighOK{}, response)
		suite.Equal(strfmt.UUID(shipment.ID.String()), payload.ShipmentID)
		suite.EqualValues(models.ReweighRequesterTOO, payload.RequestedBy)
		suite.WithinDuration(time.Now(), (time.Time)(payload.RequestedAt), 2*time.Second)
		suite.HasWebhookNotification(shipment.ID, handlerContext.GetTraceID())
	})

	suite.Run("Returns a 403 when the office user is not a TOO", func() {
		officeUser := testdatagen.MakeServicesCounselorOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		uuid := uuid.Must(uuid.NewV4())
		reweighRequester := &mocks.ShipmentReweighRequester{}

		reweighRequester.AssertNumberOfCalls(suite.T(), "RequestShipmentReweigh", 0)

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-reweigh", uuid.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentReweighHandler{
			handlerContext,
			reweighRequester,
		}
		approveParams := shipmentops.RequestShipmentReweighParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(uuid),
		}

		response := handler.Handle(approveParams)
		suite.IsType(&shipmentops.RequestShipmentReweighForbidden{}, response)
	})

	suite.Run("Returns 404 when reweighRequester returns NotFoundError", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		reweighRequester := &mocks.ShipmentReweighRequester{}

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-reweigh", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentReweighHandler{
			handlerContext,
			reweighRequester,
		}
		params := shipmentops.RequestShipmentReweighParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
		}
		reweighRequester.On("RequestShipmentReweigh", params.HTTPRequest.Context(), shipment.ID).Return(nil, services.NotFoundError{})

		response := handler.Handle(params)
		suite.IsType(&shipmentops.RequestShipmentReweighNotFound{}, response)
	})

	suite.Run("Returns 409 when reweighRequester returns Conflict Error", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		reweighRequester := &mocks.ShipmentReweighRequester{}

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-reweigh", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentReweighHandler{
			handlerContext,
			reweighRequester,
		}
		params := shipmentops.RequestShipmentReweighParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
		}

		reweighRequester.On("RequestShipmentReweigh", params.HTTPRequest.Context(), shipment.ID).Return(nil, services.ConflictError{})

		response := handler.Handle(params)
		suite.IsType(&shipmentops.RequestShipmentReweighConflict{}, response)
	})

	suite.Run("Returns 422 when reweighRequester returns validation errors", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		reweighRequester := &mocks.ShipmentReweighRequester{}

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-reweigh", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentReweighHandler{
			handlerContext,
			reweighRequester,
		}
		params := shipmentops.RequestShipmentReweighParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
		}
		reweighRequester.On("RequestShipmentReweigh", params.HTTPRequest.Context(), shipment.ID).Return(nil, services.InvalidInputError{ValidationErrors: &validate.Errors{}})

		response := handler.Handle(params)
		suite.IsType(&shipmentops.RequestShipmentReweighUnprocessableEntity{}, response)
	})

	suite.Run("Returns 500 when reweighRequester returns unexpected error", func() {
		shipment := testdatagen.MakeStubbedShipment(suite.DB())
		officeUser := testdatagen.MakeTOOOfficeUser(suite.DB(), testdatagen.Assertions{Stub: true})
		reweighRequester := &mocks.ShipmentReweighRequester{}

		req := httptest.NewRequest("POST", fmt.Sprintf("/shipments/%s/request-reweigh", shipment.ID.String()), nil)
		req = suite.AuthenticateOfficeRequest(req, officeUser)
		handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

		handler := RequestShipmentReweighHandler{
			handlerContext,
			reweighRequester,
		}
		params := shipmentops.RequestShipmentReweighParams{
			HTTPRequest: req,
			ShipmentID:  *handlers.FmtUUID(shipment.ID),
		}

		reweighRequester.On("RequestShipmentReweigh", params.HTTPRequest.Context(), shipment.ID).Return(nil, errors.New("UnexpectedError"))

		response := handler.Handle(params)
		suite.IsType(&shipmentops.RequestShipmentReweighInternalServerError{}, response)
	})
}

type createMTOShipmentSubtestData struct {
	builder *query.Builder
	params  mtoshipmentops.CreateMTOShipmentParams
}

func (suite *HandlerSuite) makeCreateMTOShipmentSubtestData() (subtestData *createMTOShipmentSubtestData) {
	subtestData = &createMTOShipmentSubtestData{}

	mto := testdatagen.MakeAvailableMove(suite.DB())
	pickupAddress := testdatagen.MakeDefaultAddress(suite.DB())
	destinationAddress := testdatagen.MakeDefaultAddress(suite.DB())
	mtoShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		Move:        mto,
		MTOShipment: models.MTOShipment{},
	})

	mtoShipment.MoveTaskOrderID = mto.ID

	subtestData.builder = query.NewQueryBuilder(suite.DB())

	req := httptest.NewRequest("POST", "/mto-shipments", nil)

	subtestData.params = mtoshipmentops.CreateMTOShipmentParams{
		HTTPRequest: req,
		Body: &ghcmessages.CreateMTOShipment{
			MoveTaskOrderID:     handlers.FmtUUID(mtoShipment.MoveTaskOrderID),
			Agents:              nil,
			CustomerRemarks:     handlers.FmtString("customer remark"),
			CounselorRemarks:    handlers.FmtString("counselor remark"),
			RequestedPickupDate: handlers.FmtDatePtr(mtoShipment.RequestedPickupDate),
		},
	}
	subtestData.params.Body.DestinationAddress.Address = ghcmessages.Address{
		City:           &destinationAddress.City,
		Country:        destinationAddress.Country,
		PostalCode:     &destinationAddress.PostalCode,
		State:          &destinationAddress.State,
		StreetAddress1: &destinationAddress.StreetAddress1,
		StreetAddress2: destinationAddress.StreetAddress2,
		StreetAddress3: destinationAddress.StreetAddress3,
	}
	subtestData.params.Body.PickupAddress.Address = ghcmessages.Address{
		City:           &pickupAddress.City,
		Country:        pickupAddress.Country,
		PostalCode:     &pickupAddress.PostalCode,
		State:          &pickupAddress.State,
		StreetAddress1: &pickupAddress.StreetAddress1,
		StreetAddress2: pickupAddress.StreetAddress2,
		StreetAddress3: pickupAddress.StreetAddress3,
	}

	return subtestData
}

func (suite *HandlerSuite) TestCreateMTOShipmentHandler() {
	// Set the traceID so we can use it to find the webhook notification
	handlerContext := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())
	handlerContext.SetTraceID(uuid.Must(uuid.NewV4()))
	moveRouter := moverouter.NewMoveRouter(suite.DB(), suite.TestLogger())

	suite.Run("Successful POST - Integration Test", func() {
		subtestData := suite.makeCreateMTOShipmentSubtestData()
		builder := subtestData.builder
		params := subtestData.params

		fetcher := fetch.NewFetcher(builder)
		creator := mtoshipment.NewMTOShipmentCreator(suite.DB(), builder, fetcher, moveRouter)
		handler := CreateMTOShipmentHandler{
			handlerContext,
			creator,
		}
		response := handler.Handle(params)
		okResponse := response.(*mtoshipmentops.CreateMTOShipmentOK)
		createMTOShipmentPayload := okResponse.Payload
		suite.IsType(&mtoshipmentops.CreateMTOShipmentOK{}, response)

		suite.Require().Equal(ghcmessages.MTOShipmentStatusSUBMITTED, createMTOShipmentPayload.Status, "MTO Shipment should have been submitted")
		suite.Require().Equal(createMTOShipmentPayload.ShipmentType, ghcmessages.MTOShipmentTypeHHG, "MTO Shipment should be an HHG")
		suite.Equal(string("customer remark"), *createMTOShipmentPayload.CustomerRemarks)
		suite.Equal(string("counselor remark"), *createMTOShipmentPayload.CounselorRemarks)
	})

	suite.Run("POST failure - 500", func() {
		subtestData := suite.makeCreateMTOShipmentSubtestData()
		params := subtestData.params

		mockCreator := mocks.MTOShipmentCreator{}

		handler := CreateMTOShipmentHandler{
			handlerContext,
			&mockCreator,
		}

		err := errors.New("ServerError")

		mockCreator.On("CreateMTOShipment",
			mock.Anything,
			mock.Anything,
		).Return(nil, err)

		response := handler.Handle(params)

		suite.IsType(&mtoshipmentops.CreateMTOShipmentInternalServerError{}, response)
	})

	suite.Run("POST failure - 422 -- Bad agent IDs set on shipment", func() {
		subtestData := suite.makeCreateMTOShipmentSubtestData()
		builder := subtestData.builder
		params := subtestData.params

		fetcher := fetch.NewFetcher(builder)
		creator := mtoshipment.NewMTOShipmentCreator(suite.DB(), builder, fetcher, moveRouter)

		handler := CreateMTOShipmentHandler{
			handlerContext,
			creator,
		}

		badID := params.Body.MoveTaskOrderID
		agent := &ghcmessages.MTOAgent{
			ID:            *badID,
			MtoShipmentID: *badID,
			FirstName:     handlers.FmtString("Mary"),
		}

		paramsBadIDs := params
		paramsBadIDs.Body.Agents = ghcmessages.MTOAgents{agent}

		response := handler.Handle(paramsBadIDs)
		suite.IsType(&mtoshipmentops.CreateMTOShipmentUnprocessableEntity{}, response)
		typedResponse := response.(*mtoshipmentops.CreateMTOShipmentUnprocessableEntity)
		suite.NotEmpty(typedResponse.Payload.InvalidFields)
	})

	suite.Run("POST failure - 422 - invalid input, missing pickup address", func() {
		subtestData := suite.makeCreateMTOShipmentSubtestData()
		builder := subtestData.builder
		params := subtestData.params

		fetcher := fetch.NewFetcher(builder)
		creator := mtoshipment.NewMTOShipmentCreator(suite.DB(), builder, fetcher, moveRouter)

		handler := CreateMTOShipmentHandler{
			handlerContext,
			creator,
		}

		badParams := params
		badParams.Body.PickupAddress.Address.StreetAddress1 = nil

		suite.NoError(badParams.Body.Validate(strfmt.Default))

		response := handler.Handle(badParams)
		suite.IsType(&mtoshipmentops.CreateMTOShipmentUnprocessableEntity{}, response)
		typedResponse := response.(*mtoshipmentops.CreateMTOShipmentUnprocessableEntity)
		// CreateMTOShipment is returning services.NewInvalidInputError without any validation errors
		// so InvalidFields won't be added to the payload.
		suite.Empty(typedResponse.Payload.InvalidFields)
	})

	suite.Run("POST failure - 404 -- not found", func() {
		subtestData := suite.makeCreateMTOShipmentSubtestData()
		builder := subtestData.builder
		params := subtestData.params

		fetcher := fetch.NewFetcher(builder)
		creator := mtoshipment.NewMTOShipmentCreator(suite.DB(), builder, fetcher, moveRouter)

		handler := CreateMTOShipmentHandler{
			handlerContext,
			creator,
		}

		uuidString := "d874d002-5582-4a91-97d3-786e8f66c763"
		badParams := params
		badParams.Body.MoveTaskOrderID = handlers.FmtUUID(uuid.FromStringOrNil(uuidString))

		response := handler.Handle(badParams)
		suite.IsType(&mtoshipmentops.CreateMTOShipmentNotFound{}, response)
	})

	suite.Run("POST failure - 400 -- nil body", func() {
		subtestData := suite.makeCreateMTOShipmentSubtestData()
		builder := subtestData.builder

		fetcher := fetch.NewFetcher(builder)
		creator := mtoshipment.NewMTOShipmentCreator(suite.DB(), builder, fetcher, moveRouter)

		handler := CreateMTOShipmentHandler{
			handlerContext,
			creator,
		}

		req := httptest.NewRequest("POST", "/mto-shipments", nil)

		paramsNilBody := mtoshipmentops.CreateMTOShipmentParams{
			HTTPRequest: req,
		}
		response := handler.Handle(paramsNilBody)

		suite.IsType(&mtoshipmentops.CreateMTOShipmentBadRequest{}, response)
	})
}

func (suite *HandlerSuite) getUpdateShipmentParams(originalShipment models.MTOShipment) mtoshipmentops.UpdateMTOShipmentParams {
	servicesCounselor := testdatagen.MakeDefaultOfficeUser(suite.DB())
	servicesCounselor.User.Roles = append(servicesCounselor.User.Roles, roles.Role{
		RoleType: roles.RoleTypeServicesCounselor,
	})
	pickupAddress := testdatagen.MakeDefaultAddress(suite.DB())
	pickupAddress.StreetAddress1 = "123 Fake Test St NW"
	destinationAddress := testdatagen.MakeDefaultAddress(suite.DB())
	destinationAddress.StreetAddress1 = "54321 Test Fake Rd SE"
	customerRemarks := "help"
	counselorRemarks := "counselor approved"
	mtoAgent := testdatagen.MakeDefaultMTOAgent(suite.DB())
	agents := ghcmessages.MTOAgents{&ghcmessages.MTOAgent{
		FirstName: mtoAgent.FirstName,
		LastName:  mtoAgent.LastName,
		Email:     mtoAgent.Email,
		Phone:     mtoAgent.Phone,
		AgentType: string(mtoAgent.MTOAgentType),
	}}

	req := httptest.NewRequest("PATCH", fmt.Sprintf("/move_task_orders/%s/mto_shipments/%s", originalShipment.MoveTaskOrderID.String(), originalShipment.ID.String()), nil)
	req = suite.AuthenticateOfficeRequest(req, servicesCounselor)

	eTag := etag.GenerateEtag(originalShipment.UpdatedAt)

	payload := ghcmessages.UpdateShipment{
		RequestedPickupDate:   strfmt.Date(time.Now()),
		RequestedDeliveryDate: strfmt.Date(time.Now()),
		ShipmentType:          ghcmessages.MTOShipmentTypeHHG,
		CustomerRemarks:       &customerRemarks,
		CounselorRemarks:      &counselorRemarks,
		Agents:                agents,
	}
	payload.DestinationAddress.Address = ghcmessages.Address{
		City:           &destinationAddress.City,
		Country:        destinationAddress.Country,
		PostalCode:     &destinationAddress.PostalCode,
		State:          &destinationAddress.State,
		StreetAddress1: &destinationAddress.StreetAddress1,
		StreetAddress2: destinationAddress.StreetAddress2,
		StreetAddress3: destinationAddress.StreetAddress3,
	}
	payload.PickupAddress.Address = ghcmessages.Address{
		City:           &pickupAddress.City,
		Country:        pickupAddress.Country,
		PostalCode:     &pickupAddress.PostalCode,
		State:          &pickupAddress.State,
		StreetAddress1: &pickupAddress.StreetAddress1,
		StreetAddress2: pickupAddress.StreetAddress2,
		StreetAddress3: pickupAddress.StreetAddress3,
	}

	params := mtoshipmentops.UpdateMTOShipmentParams{
		HTTPRequest: req,
		ShipmentID:  *handlers.FmtUUID(originalShipment.ID),
		Body:        &payload,
		IfMatch:     eTag,
	}

	return params
}

func (suite *HandlerSuite) TestUpdateShipmentHandler() {
	planner := &routemocks.Planner{}
	planner.On("TransitDistance",
		mock.Anything,
		mock.Anything,
	).Return(400, nil)
	moveRouter := moverouter.NewMoveRouter(suite.DB(), suite.TestLogger())

	suite.Run("Successful PATCH - Integration Test", func() {
		builder := query.NewQueryBuilder(suite.DB())
		fetcher := fetch.NewFetcher(builder)
		updater := mtoshipment.NewMTOShipmentUpdater(suite.DB(), builder, fetcher, planner, moveRouter)
		handler := UpdateShipmentHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			fetcher,
			updater,
		}

		oldShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusSubmitted,
			},
		})
		params := suite.getUpdateShipmentParams(oldShipment)

		// Run swagger validations
		suite.NoError(params.Body.Validate(strfmt.Default))

		response := handler.Handle(params)
		suite.IsType(&mtoshipmentops.UpdateMTOShipmentOK{}, response)

		updatedShipment := response.(*mtoshipmentops.UpdateMTOShipmentOK).Payload
		suite.Equal(oldShipment.ID.String(), updatedShipment.ID.String())
		suite.Equal(params.Body.CustomerRemarks, updatedShipment.CustomerRemarks)
		suite.Equal(params.Body.CounselorRemarks, updatedShipment.CounselorRemarks)
		suite.Equal(params.Body.PickupAddress.StreetAddress1, updatedShipment.PickupAddress.StreetAddress1)
		suite.Equal(params.Body.DestinationAddress.StreetAddress1, updatedShipment.DestinationAddress.StreetAddress1)
		suite.Equal(params.Body.RequestedPickupDate.String(), updatedShipment.RequestedPickupDate.String())
		suite.Equal(params.Body.Agents[0].FirstName, updatedShipment.MtoAgents[0].FirstName)
		suite.Equal(params.Body.Agents[0].LastName, updatedShipment.MtoAgents[0].LastName)
		suite.Equal(params.Body.Agents[0].Email, updatedShipment.MtoAgents[0].Email)
		suite.Equal(params.Body.Agents[0].Phone, updatedShipment.MtoAgents[0].Phone)
		suite.Equal(params.Body.Agents[0].AgentType, updatedShipment.MtoAgents[0].AgentType)
		suite.Equal(oldShipment.ID.String(), string(updatedShipment.MtoAgents[0].MtoShipmentID))
		suite.NotEmpty(updatedShipment.MtoAgents[0].ID)
		suite.Equal(params.Body.RequestedDeliveryDate.String(), updatedShipment.RequestedDeliveryDate.String())
	})

	suite.Run("PATCH failure - 400 -- nil body", func() {
		builder := query.NewQueryBuilder(suite.DB())
		fetcher := fetch.NewFetcher(builder)
		updater := mtoshipment.NewMTOShipmentUpdater(suite.DB(), builder, fetcher, planner, moveRouter)
		handler := UpdateShipmentHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			fetcher,
			updater,
		}

		oldShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusSubmitted,
			},
		})
		params := suite.getUpdateShipmentParams(oldShipment)
		params.Body = nil

		response := handler.Handle(params)

		suite.IsType(&mtoshipmentops.UpdateMTOShipmentUnprocessableEntity{}, response)
	})

	suite.Run("PATCH failure - 404 -- not found", func() {
		builder := query.NewQueryBuilder(suite.DB())
		fetcher := fetch.NewFetcher(builder)
		updater := mtoshipment.NewMTOShipmentUpdater(suite.DB(), builder, fetcher, planner, moveRouter)
		handler := UpdateShipmentHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			fetcher,
			updater,
		}

		uuidString := handlers.FmtUUID(uuid.FromStringOrNil("d874d002-5582-4a91-97d3-786e8f66c763"))
		oldShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusSubmitted,
			},
		})
		params := suite.getUpdateShipmentParams(oldShipment)
		params.ShipmentID = *uuidString

		// Run swagger validations
		suite.NoError(params.Body.Validate(strfmt.Default))

		response := handler.Handle(params)

		suite.IsType(&mtoshipmentops.UpdateMTOShipmentNotFound{}, response)
	})

	suite.Run("PATCH failure - 412 -- etag mismatch", func() {
		builder := query.NewQueryBuilder(suite.DB())
		fetcher := fetch.NewFetcher(builder)
		updater := mtoshipment.NewMTOShipmentUpdater(suite.DB(), builder, fetcher, planner, moveRouter)
		handler := UpdateShipmentHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			fetcher,
			updater,
		}

		oldShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusSubmitted,
			},
		})
		params := suite.getUpdateShipmentParams(oldShipment)
		params.IfMatch = "intentionally-bad-if-match-header-value"

		// Run swagger validations
		suite.NoError(params.Body.Validate(strfmt.Default))

		response := handler.Handle(params)

		suite.IsType(&mtoshipmentops.UpdateMTOShipmentPreconditionFailed{}, response)
	})

	suite.Run("PATCH failure - 412 -- shipment shouldn't be updatable", func() {
		builder := query.NewQueryBuilder(suite.DB())
		fetcher := fetch.NewFetcher(builder)
		updater := mtoshipment.NewMTOShipmentUpdater(suite.DB(), builder, fetcher, planner, moveRouter)
		handler := UpdateShipmentHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			fetcher,
			updater,
		}

		oldShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusDraft,
			},
		})

		params := suite.getUpdateShipmentParams(oldShipment)

		// Run swagger validations
		suite.NoError(params.Body.Validate(strfmt.Default))

		response := handler.Handle(params)

		suite.IsType(&mtoshipmentops.UpdateMTOShipmentPreconditionFailed{}, response)
	})

	suite.Run("PATCH failure - 500", func() {
		builder := query.NewQueryBuilder(suite.DB())
		mockUpdater := mocks.MTOShipmentUpdater{}
		fetcher := fetch.NewFetcher(builder)
		handler := UpdateShipmentHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			fetcher,
			&mockUpdater,
		}

		err := errors.New("ServerError")

		mockUpdater.On("UpdateMTOShipment",
			mock.Anything,
			mock.Anything,
		).Return(nil, err)
		mockUpdater.On("RetrieveMTOShipment",
			mock.Anything,
			mock.Anything,
		).Return(nil, err)
		mockUpdater.On("CheckIfMTOShipmentCanBeUpdated",
			mock.Anything,
			mock.Anything,
		).Return(nil, err)

		oldShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusSubmitted,
			},
		})
		params := suite.getUpdateShipmentParams(oldShipment)

		// Run swagger validations
		suite.NoError(params.Body.Validate(strfmt.Default))

		response := handler.Handle(params)

		suite.IsType(&mtoshipmentops.UpdateMTOShipmentInternalServerError{}, response)
	})

}
