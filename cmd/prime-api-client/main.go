package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/transcom/mymove/pkg/cli"
)

const (
	// CertPathFlag is the path to the certificate to use for TLS
	CertPathFlag string = "certpath"
	// KeyPathFlag is the path to the key to use for TLS
	KeyPathFlag string = "keypath"
	// HostnameFlag is the hostname to connect to
	HostnameFlag string = "hostname"
	// PortFlag is the port to connect to
	PortFlag string = "port"
	// InsecureFlag indicates that TLS verification and validation can be skipped
	InsecureFlag string = "insecure"
)

// initRootFlags initializes flags relating to the prime api
func initRootFlags(flag *pflag.FlagSet) {
	cli.InitCACFlags(flag)
	cli.InitVerboseFlags(flag)

	flag.String(CertPathFlag, "./config/tls/devlocal-mtls.cer", "Path to the public cert")
	flag.String(KeyPathFlag, "./config/tls/devlocal-mtls.key", "Path to the private key")
	flag.String(HostnameFlag, cli.HTTPPrimeServerNameLocal, "The hostname to connect to")
	flag.Int(PortFlag, cli.MutualTLSPort, "The port to connect to")
	flag.Bool(InsecureFlag, false, "Skip TLS verification and validation")
}

// CheckRootConfig checks the validity of the prime api flags
func CheckRootConfig(v *viper.Viper) error {
	err := cli.CheckCAC(v)
	if err != nil {
		return err
	}

	err = cli.CheckVerbose(v)
	if err != nil {
		return err
	}

	if (v.GetString(CertPathFlag) != "" && v.GetString(KeyPathFlag) == "") || (v.GetString(CertPathFlag) == "" && v.GetString(KeyPathFlag) != "") {
		return fmt.Errorf("Both TLS certificate and key paths must be provided")
	}

	return nil
}

func main() {
	root := cobra.Command{
		Use:   "prime-api-client [flags]",
		Short: "Prime API client",
		Long:  "Prime API client",
	}
	initRootFlags(root.PersistentFlags())

	fetchMTOsCommand := &cobra.Command{
		Use:          "fetch-mtos",
		Short:        "fetch mtos",
		Long:         "fetch move task orders",
		RunE:         fetchMTOs,
		SilenceUsage: true,
	}
	initFetchMTOsFlags(fetchMTOsCommand.Flags())
	root.AddCommand(fetchMTOsCommand)

	createMTOCommand := &cobra.Command{
		Use:   "support-create-mto",
		Short: "Create a MoveTaskOrder",
		Long: `
  This command creates a MoveTaskOrder object.
  It requires the caller to pass in a file using the --filename param.

  Endpoint path: /move-task-orders
  The file should contain json as follows:
    {
      "body": <MoveTaskOrder>
    }
  Please see API documentation for full details on the MoveTaskOrder definition.`,
		RunE:         createMTO,
		SilenceUsage: true,
	}
	initCreateMTOFlags(createMTOCommand.Flags())
	root.AddCommand(createMTOCommand)

	updateMTOShipmentCommand := &cobra.Command{
		Use:   "update-mto-shipment",
		Short: "update mto shipment",
		Long: `
  This command updates an MTO shipment.
  It requires the caller to pass in a file using the --filename arg.
  The file should contain path parameters, headers and a body for the payload.

  Endpoint path: move-task-orders/{moveTaskOrderID}/mto-shipments/{mtoShipmentID}
  The file should contain json as follows:
  	{
      "moveTaskOrderID": <uuid string>,
      "mtoShipmentID": <uuid string>,
      "ifMatch": <eTag>,
      "body": <MTOShipment>
  	}
  Please see API documentation for full details on the endpoint definition.`,
		RunE:         updateMTOShipment,
		SilenceUsage: true,
	}
	initUpdateMTOShipmentFlags(updateMTOShipmentCommand.Flags())
	root.AddCommand(updateMTOShipmentCommand)

	updatePostCounselingInfo := &cobra.Command{
		Use:          "update-post-counseling-info",
		Short:        "update post counseling info",
		Long:         "update post counseling info such as discovering that customer has a PPM",
		RunE:         updatePostCounselingInfo,
		SilenceUsage: true,
	}
	initUpdatePostCounselingInfoFlags(updatePostCounselingInfo.Flags())
	root.AddCommand(updatePostCounselingInfo)

	createMTOServiceItemCommand := &cobra.Command{
		Use:   "create-mto-service-item",
		Short: "Create mto service item",
		Long: `
  This command creates an MTO service item on an MTO shipment.
  It requires the caller to pass in a file using the --filename arg.
  The file should contain path parameters and headers and a body for the payload.

  Endpoint path: /move-task-orders/{moveTaskOrderID}/mto-shipments/{mtoShipmentID}/mto-service-items
  The file should contain json as follows:
  	{
  	"moveTaskOrderID": <uuid string>,
  	"mtoShipmentID": <uuid string>,
  	"ifMatch": <eTag>,
  	"body": <MTOServiceItem>
  	}
  Please see API documentation for full details on the endpoint definition.`,
		RunE:         createMTOServiceItem,
		SilenceUsage: true,
	}
	initCreateMTOServiceItemFlags(createMTOServiceItemCommand.Flags())
	root.AddCommand(createMTOServiceItemCommand)

	makeAvailableToPrimeCommand := &cobra.Command{
		Use:   "support-make-mto-available-to-prime",
		Short: "Make mto available to prime",
		Long: `
  This command makes an MTO available for prime consumption.
  This is a support endpoint and is not available in production.
  It requires the caller to pass in a file using the --filename arg.
  The file should contain path parameters and headers.

  Endpoint path: /move-task-orders/{moveTaskOrderID}/status
  The file should contain json as follows:
  	{
  	"moveTaskOrderID": <uuid string>,
  	"ifMatch": <eTag>,
  	}
  Please see API documentation for full details on the endpoint definition.`,
		RunE:         updateMTOStatus,
		SilenceUsage: true,
	}
	initUpdateMTOStatusFlags(makeAvailableToPrimeCommand.Flags())
	root.AddCommand(makeAvailableToPrimeCommand)

	updatePaymentRequestStatusCommand := &cobra.Command{
		Use:   "support-update-payment-request-status",
		Short: "Update payment request status for prime",
		Long: `
  This command allows prime to update payment request status.
  This is a support endpoint and is not available in production.
  It requires the caller to pass in a file using the --filename arg.
  The file should contain path parameters and headers.

  Endpoint path: /payment-requests/{paymentRequestID}/status
  The file should contain json as follows:
    {
      "paymentRequestID": <uuid string>,
      "ifMatch": <etag>,
      "body" : <paymentRequestStatus>
    }
  Please see API documentation for full details on the endpoint definition.`,
		RunE:         updatePaymentRequestStatus,
		SilenceUsage: true,
	}
	initUpdatePaymentRequestStatusFlags(updatePaymentRequestStatusCommand.Flags())
	root.AddCommand(updatePaymentRequestStatusCommand)

	getMoveTaskOrder := &cobra.Command{
		Use:   "support-get-mto",
		Short: "Get an individual mto",
		Long: `
  This command gets a single move task order by ID
  This is a support endpoint and is not available in production.
  It requires the caller to pass in a file using the --filename arg.
  The file should contain path parameters and headers.

  Endpoint path: /move-task-orders/{moveTaskOrderID}
  The file should contain json as follows:
  	{
  	"moveTaskOrderID": <uuid string>,
  	}
  Please see API documentation for full details on the endpoint definition.`,
		RunE:         getMTO,
		SilenceUsage: true,
	}
	initGetMTOFlags(getMoveTaskOrder.Flags())
	root.AddCommand(getMoveTaskOrder)

	updateMTOServiceItemStatus := &cobra.Command{
		Use:   "support-update-mto-service-item-status",
		Short: "Update service item status",
		Long: `
  This command allows prime to update the MTO service item status.
  This is a support endpoint and is not available in production.
  It requires the caller to pass in a file using the --filename arg.
  The file should contain a body defining the request body.

  Endpoint path: service-items/{mtoServiceItemID}/status
    {
      "mtoServiceItemID": <uuid string>,
      "ifMatch": <etag>,
      "body": {
        "status": "APPROVED"
    }
  Please see API documentation for full details on the endpoint definition.`,
		RunE:         updateMTOServiceItemStatus,
		SilenceUsage: true,
	}
	initUpdateMTOServiceItemStatusFlags(updateMTOServiceItemStatus.Flags())
	root.AddCommand(updateMTOServiceItemStatus)

	createPaymentRequestCommand := &cobra.Command{
		Use:   "create-payment-request",
		Short: "Create payment request",
		Long: `
  This command gets a single move task order by ID
  It requires the caller to pass in a file using the --filename arg.
  The file should contain a body defining the PaymentRequest object.
  Endpoint path: /payment-requests
  The file should contain json as follows:
  	{
  	"body": <PaymentRequest>,
  	}
  Please see API documentation for full details on the endpoint definition.`,
		RunE:         createPaymentRequest,
		SilenceUsage: true,
	}
	initCreatePaymentRequestFlags(createPaymentRequestCommand.Flags())
	root.AddCommand(createPaymentRequestCommand)

	createPaymentRequestUploadCommand := &cobra.Command{
		Use:          "create-payment-request-upload",
		Short:        "Create payment request upload",
		Long:         "Create payment request upload for a payment request",
		RunE:         createPaymentRequestUpload,
		SilenceUsage: true,
	}
	initCreatePaymentRequestUploadFlags(createPaymentRequestUploadCommand.Flags())
	root.AddCommand(createPaymentRequestUploadCommand)

	patchMTOShipmentStatusCommand := &cobra.Command{
		Use:   "support-patch-mto-shipment-status",
		Short: "Update MTO shipment status for prime",
		Long: `
  This command allows prime to update the MTO shipment status.
  This is a support endpoint and is not available in production.
  It requires the caller to pass in a file using the --filename arg.
  The file should contain a body defining the request body.

  Endpoint path: /mto-shipments/{mtoShipmentID}/status
  The file should contain json as follows:
    {
      "mtoShipmentID": <uuid string>,
      "ifMatch": <etag>,
      "body": <MtoShipmentRequestStatus>,
    }
  Please see API documentation for full details on the endpoint definition.`,
		RunE:         patchMTOShipmentStatus,
		SilenceUsage: true,
	}
	initPatchMTOShipmentStatusFlags(patchMTOShipmentStatusCommand.Flags())
	root.AddCommand(patchMTOShipmentStatusCommand)

	completionCommand := &cobra.Command{
		Use:   "completion",
		Short: "Generates bash completion scripts",
		Long:  "To install completion scripts run:\n\nprime-api-client completion > /usr/local/etc/bash_completion.d/prime-api-client",
		RunE: func(cmd *cobra.Command, args []string) error {
			return root.GenBashCompletion(os.Stdout)
		},
	}
	root.AddCommand(completionCommand)

	if err := root.Execute(); err != nil {
		panic(err)
	}
}
