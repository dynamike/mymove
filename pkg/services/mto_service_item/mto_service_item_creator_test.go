package mtoserviceitem

import (
	"errors"
	"testing"

	"github.com/gobuffalo/pop"

	"github.com/transcom/mymove/pkg/services"

	"github.com/gobuffalo/validate"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

type testMTOServiceItemQueryBuilder struct {
	fakeCreateOne   func(model interface{}) (*validate.Errors, error)
	fakeFetchOne    func(model interface{}, filters []services.QueryFilter) error
	fakeTransaction func(func(tx *pop.Connection) error) error
}

func (t *testMTOServiceItemQueryBuilder) CreateOne(model interface{}) (*validate.Errors, error) {
	return t.fakeCreateOne(model)
}

func (t *testMTOServiceItemQueryBuilder) FetchOne(model interface{}, filters []services.QueryFilter) error {
	return t.fakeFetchOne(model, filters)
}

func (t *testMTOServiceItemQueryBuilder) Transaction(fn func(tx *pop.Connection) error) error {
	return t.fakeTransaction(fn)
}

func (suite *MTOServiceItemServiceSuite) TestCreateMTOServiceItem() {
	moveTaskOrder := testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{})
	dimension := testdatagen.MakeMTOServiceItemDimension(suite.DB(), testdatagen.Assertions{})
	serviceItem := models.MTOServiceItem{
		MoveTaskOrderID: moveTaskOrder.ID,
		MoveTaskOrder:   moveTaskOrder,
		Dimensions: models.MTOServiceItemDimensions{
			dimension,
		},
	}

	// Happy path
	suite.T().Run("If the user is created successfully it should be returned", func(t *testing.T) {
		fakeCreateOne := func(model interface{}) (*validate.Errors, error) {
			return nil, nil
		}
		fakeFetchOne := func(model interface{}, filters []services.QueryFilter) error {
			return nil
		}
		fakeTx := func(fn func(tx *pop.Connection) error) error {
			return fn(&pop.Connection{})
		}

		builder := &testMTOServiceItemQueryBuilder{
			fakeCreateOne:   fakeCreateOne,
			fakeFetchOne:    fakeFetchOne,
			fakeTransaction: fakeTx,
		}

		fakeCreateNewBuilder := func(db *pop.Connection) createMTOServiceItemQueryBuilder {
			return builder
		}

		creator := mtoServiceItemCreator{
			builder:          builder,
			createNewBuilder: fakeCreateNewBuilder,
		}
		createdServiceItem, verrs, err := creator.CreateMTOServiceItem(&serviceItem)

		suite.NoError(err)
		suite.Nil(verrs)
		suite.NotNil(createdServiceItem)
		suite.NotEmpty(createdServiceItem.Dimensions)
	})

	// Bad data which could be IDs that doesn't exist (MoveTaskOrderID or REServiceID)
	suite.T().Run("If error when trying to create, the create should fail", func(t *testing.T) {
		expectedError := "Can't create service item for some reason"
		verrs := validate.NewErrors()
		verrs.Add("test", expectedError)
		fakeCreateOne := func(model interface{}) (*validate.Errors, error) {
			return verrs, errors.New(expectedError)
		}
		fakeFetchOne := func(model interface{}, filters []services.QueryFilter) error {
			return nil
		}

		fakeTx := func(fn func(tx *pop.Connection) error) error {
			return fn(&pop.Connection{})
		}

		builder := &testMTOServiceItemQueryBuilder{
			fakeCreateOne:   fakeCreateOne,
			fakeFetchOne:    fakeFetchOne,
			fakeTransaction: fakeTx,
		}

		fakeCreateNewBuilder := func(db *pop.Connection) createMTOServiceItemQueryBuilder {
			return builder
		}

		creator := mtoServiceItemCreator{
			builder:          builder,
			createNewBuilder: fakeCreateNewBuilder,
		}

		createdServiceItem, verrs, _ := creator.CreateMTOServiceItem(&serviceItem)
		suite.Error(verrs)
		suite.Nil(createdServiceItem)
	})

	// If the service item we're trying to create is shuttle service and there is no estimated weight, it fails.
	suite.T().Run("If we try to create a shuttle service without the estimated weight it fails", func(t *testing.T) {
		shipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{})

		serviceItemNoWeight := models.MTOServiceItem{
			MoveTaskOrderID: moveTaskOrder.ID,
			MoveTaskOrder:   moveTaskOrder,
			MTOShipment:     shipment,
			MTOShipmentID:   &shipment.ID,
			ReService: models.ReService{
				Code: models.ReServiceCodeDDSHUT,
			},
		}

		fakeCreateOne := func(model interface{}) (*validate.Errors, error) {
			return nil, nil
		}
		fakeFetchOne := func(model interface{}, filters []services.QueryFilter) error {
			return nil
		}

		builder := &testMTOServiceItemQueryBuilder{
			fakeCreateOne: fakeCreateOne,
			fakeFetchOne:  fakeFetchOne,
		}

		creator := NewMTOServiceItemCreator(builder)
		createdServiceItem, _, err := creator.CreateMTOServiceItem(&serviceItemNoWeight)
		suite.Error(err)
		suite.Nil(createdServiceItem)
	})
}
