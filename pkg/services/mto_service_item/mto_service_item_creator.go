package mtoserviceitem

import (
	"fmt"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/services/query"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
)

type createMTOServiceItemQueryBuilder interface {
	FetchOne(model interface{}, filters []services.QueryFilter) error
	CreateOne(model interface{}) (*validate.Errors, error)
	Transaction(fn func(tx *pop.Connection) error) error
}

type mtoServiceItemCreator struct {
	builder          createMTOServiceItemQueryBuilder
	createNewBuilder func(db *pop.Connection) createMTOServiceItemQueryBuilder
}

// CreateMTOServiceItem creates an MTO Service Item
func (o *mtoServiceItemCreator) CreateMTOServiceItem(serviceItem *models.MTOServiceItem) (*models.MTOServiceItem, *validate.Errors, error) {
	var verrs *validate.Errors
	var err error

	var moveTaskOrder models.MoveTaskOrder
	moveTaskOrderID := serviceItem.MoveTaskOrderID
	queryFilters := []services.QueryFilter{
		query.NewQueryFilter("id", "=", moveTaskOrderID),
	}
	// check if MTO exists
	err = o.builder.FetchOne(&moveTaskOrder, queryFilters)
	if err != nil {
		return nil, nil, services.NewNotFoundError(moveTaskOrderID, fmt.Sprintf("MoveTaskOrderID: %s", err))
	}

	// check if shipment exists
	var mtoShipment models.MTOShipment
	var mtoShipmentID uuid.UUID
	if serviceItem.MTOShipmentID != nil {
		mtoShipmentID = *serviceItem.MTOShipmentID
	}
	queryFilters = []services.QueryFilter{
		query.NewQueryFilter("id", "=", mtoShipmentID),
	}
	err = o.builder.FetchOne(&mtoShipment, queryFilters)
	if err != nil {
		return nil, nil, services.NewNotFoundError(mtoShipmentID, fmt.Sprintf("MTOShipmentID: %s", err))
	}

	// find the re service code id
	var reService models.ReService
	reServiceCode := serviceItem.ReService.Code
	queryFilters = []services.QueryFilter{
		query.NewQueryFilter("code", "=", reServiceCode),
	}
	err = o.builder.FetchOne(&reService, queryFilters)
	if err != nil {
		return nil, nil, services.NewNotFoundError(uuid.Nil, fmt.Sprintf("failed to find service item code: %s; %s", reServiceCode, err))
	}

	// set re service for service item
	serviceItem.ReServiceID = reService.ID

	if serviceItem.ReService.Code == models.ReServiceCodeDOSHUT || serviceItem.ReService.Code == models.ReServiceCodeDDSHUT {
		if mtoShipment.PrimeEstimatedWeight == nil {
			return nil, nil, services.NewInvalidInputError(mtoShipmentID, nil, nil,
				fmt.Sprintf("MTOShipment with id: %s is missing the estimated weight required for this service item", mtoShipmentID))
		}
	}

	// create new items in a transaction in case of failure
	o.builder.Transaction(func(tx *pop.Connection) error {
		// create new builder to use tx
		txBuilder := o.createNewBuilder(tx)

		// create service item
		verrs, err = txBuilder.CreateOne(serviceItem)
		if verrs != nil || err != nil {
			return fmt.Errorf("%#v %e", verrs, err)
		}

		// create dimensions if any
		for index := range serviceItem.Dimensions {
			createDimension := &serviceItem.Dimensions[index]
			createDimension.MTOServiceItemID = serviceItem.ID
			verrs, err = txBuilder.CreateOne(createDimension)
			if verrs != nil || err != nil {
				return fmt.Errorf("%#v %e", verrs, err)
			}
		}

		// create customer contacts if any
		for index := range serviceItem.CustomerContacts {
			createCustContacts := &serviceItem.CustomerContacts[index]
			createCustContacts.MTOServiceItemID = serviceItem.ID
			verrs, err = txBuilder.CreateOne(createCustContacts)
			if verrs != nil || err != nil {
				return fmt.Errorf("%#v %e", verrs, err)
			}
		}

		return nil
	})

	if verrs != nil || err != nil {
		return nil, verrs, err
	}

	return serviceItem, nil, nil
}

// NewMTOServiceItemCreator returns a new MTO service item creator
func NewMTOServiceItemCreator(builder createMTOServiceItemQueryBuilder) services.MTOServiceItemCreator {
	// used inside a transaction and mocking
	createNewBuilder := func(db *pop.Connection) createMTOServiceItemQueryBuilder {
		return query.NewQueryBuilder(db)
	}

	return &mtoServiceItemCreator{builder: builder, createNewBuilder: createNewBuilder}
}
