package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	mtoShipment "github.com/transcom/mymove/pkg/gen/supportclient/mto_shipment"
)

func initPatchMTOShipmentStatusFlags(flag *pflag.FlagSet) {
	flag.String(FilenameFlag, "", "Name of the file being passed in")

	flag.SortFlags = false
}

func checkPatchMTOShipmentStatusConfig(v *viper.Viper, args []string, logger *log.Logger) error {
	err := CheckRootConfig(v)
	if err != nil {
		logger.Fatal(err)
	}

	if v.GetString(FilenameFlag) == "" && (len(args) < 1 || len(args) > 0 && !containsDash(args)) {
		logger.Fatal(errors.New("support-patch-mto-shipment-status expects a file to be passed in"))
	}

	return nil
}

func patchMTOShipmentStatus(cmd *cobra.Command, args []string) error {
	v := viper.New()

	//  Create the logger
	//  Remove the prefix and any datetime data
	logger := log.New(os.Stdout, "", log.LstdFlags)

	errParseFlags := ParseFlags(cmd, v, args)
	if errParseFlags != nil {
		return errParseFlags
	}

	// Check the config before talking to the CAC
	err := checkPatchMTOShipmentStatusConfig(v, args, logger)
	if err != nil {
		logger.Fatal(err)
	}

	// Decode json from file that was passed in
	filename := v.GetString(FilenameFlag)
	var patchMTOShipmentParams mtoShipment.PatchMTOShipmentStatusParams
	err = decodeJSONFileToPayload(filename, containsDash(args), &patchMTOShipmentParams)
	if err != nil {
		logger.Fatal(err)
	}
	patchMTOShipmentParams.SetTimeout(time.Second * 30)

	// Create the client and open the cacStore
	supportGateway, cacStore, errCreateClient := CreateSupportClient(v)
	if errCreateClient != nil {
		return errCreateClient
	}
	// Defer closing the store until after the API call has completed
	if cacStore != nil {
		defer cacStore.Close()
	}

	resp, err := supportGateway.MtoShipment.PatchMTOShipmentStatus(&patchMTOShipmentParams)
	if err != nil {
		return handleGatewayError(err, logger)
	}

	payload := resp.GetPayload()
	if payload != nil {
		payload, errJSONMarshall := json.Marshal(payload)
		if errJSONMarshall != nil {
			logger.Fatal(errJSONMarshall)
		}
		fmt.Println(string(payload))
	} else {
		logger.Fatal(resp.Error())
	}

	return nil
}
