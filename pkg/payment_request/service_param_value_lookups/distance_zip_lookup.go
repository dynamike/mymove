package serviceparamvaluelookups

import (
	"database/sql"
	"strconv"

	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
	"github.com/transcom/mymove/pkg/unit"
)

// DistanceZipLookup contains zip lookup
type DistanceZipLookup struct {
	PickupAddress      models.Address
	DestinationAddress models.Address
}

func (r DistanceZipLookup) lookup(keyData *ServiceItemParamKeyData) (string, error) {
	planner := keyData.planner
	db := keyData.db

	// Get the MTOServiceItem and associated MTOShipment and addresses
	mtoServiceItemID := keyData.MTOServiceItemID
	var mtoServiceItem models.MTOServiceItem
	err := db.
		Eager("MTOShipment", "MTOShipment.PickupAddress", "MTOShipment.DestinationAddress").
		Find(&mtoServiceItem, mtoServiceItemID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return "", services.NewNotFoundError(mtoServiceItemID, "looking for MTOServiceItemID")
		default:
			return "", err
		}
	}

	// Make sure there's an MTOShipment since that's nullable
	mtoShipmentID := mtoServiceItem.MTOShipmentID
	if mtoShipmentID == nil {
		return "", services.NewNotFoundError(uuid.Nil, "looking for MTOShipmentID")
	}

	mtoShipment := mtoServiceItem.MTOShipment
	if mtoShipment.Distance != nil {
		return strconv.Itoa(mtoShipment.Distance.Int()), nil
	}

	// Now calculate the distance between zips
	pickupZip := r.PickupAddress.PostalCode
	destinationZip := r.DestinationAddress.PostalCode
	distanceMiles, err := distanceZip(planner, pickupZip, destinationZip)
	if err != nil {
		return "", err
	}

	miles := unit.Miles(distanceMiles)
	mtoShipment.Distance = &miles
	err = db.Save(&mtoShipment)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(distanceMiles), nil
}
