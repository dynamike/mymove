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

	mtoserviceitem "github.com/transcom/mymove/pkg/gen/supportclient/mto_service_item"
)

func initUpdateMTOServiceItemStatusFlags(flag *pflag.FlagSet) {
	flag.String(FilenameFlag, "", "Name of the file being passed in")

	flag.SortFlags = false
}

func checkUpdateMTOServiceItemStatusConfig(v *viper.Viper, args []string, logger *log.Logger) error {
	err := CheckRootConfig(v)
	if err != nil {
		logger.Fatal(err)
	}

	if v.GetString(FilenameFlag) == "" && (len(args) < 1 || len(args) > 0 && !containsDash(args)) {
		logger.Fatal(errors.New("support-update-mto-service-item-status expects a file to be passed in"))
	}

	return nil
}

func updateMTOServiceItemStatus(cmd *cobra.Command, args []string) error {
	v := viper.New()

	//  Create the logger
	//  Remove the prefix and any datetime data
	logger := log.New(os.Stdout, "", log.LstdFlags)

	errParseFlags := ParseFlags(cmd, v, args)
	if errParseFlags != nil {
		return errParseFlags
	}

	// Check the config before talking to the CAC
	err := checkUpdateMTOServiceItemStatusConfig(v, args, logger)
	if err != nil {
		logger.Fatal(err)
	}

	// Decode json from file that was passed into MTO Service item
	filename := v.GetString(FilenameFlag)
	var updateServiceItemParams mtoserviceitem.UpdateMTOServiceItemStatusParams
	err = decodeJSONFileToPayload(filename, containsDash(args), &updateServiceItemParams)
	if err != nil {
		logger.Fatal(err)
	}
	updateServiceItemParams.SetTimeout(time.Second * 30)

	// Create the client and open the cacStore
	supportGateway, cacStore, errCreateClient := CreateSupportClient(v)
	if errCreateClient != nil {
		return errCreateClient
	}
	// Defer closing the store until after the API call has completed
	if cacStore != nil {
		defer cacStore.Close()
	}

	// Make the API Call
	resp, err := supportGateway.MtoServiceItem.UpdateMTOServiceItemStatus(&updateServiceItemParams)
	if err != nil {
		return handleGatewayError(err, logger)
	}

	// Get the successful response payload and convert to json for output
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
