package ediinvoice

import (
	"strings"
	"testing"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/db/sequence"
	edisegment "github.com/transcom/mymove/pkg/edi/segment"
	"github.com/transcom/mymove/pkg/testingsuite"
)

type InvoiceSuite struct {
	testingsuite.PopTestSuite
	logger       Logger
	Viper        *viper.Viper
	icnSequencer sequence.Sequencer
}

func TestInvoiceSuite(t *testing.T) {
	// Use a no-op logger during testing
	logger := zap.NewNop()

	flag := pflag.CommandLine
	// Flag to update the test EDI
	// Borrowed from https://about.sourcegraph.com/go/advanced-testing-in-go
	flag.Bool("update", false, "update .golden files")
	// Flag to toggle Invoice usage indicator from P>T (Production>Test)
	flag.Bool("send-prod-invoice", false, "Send Production Invoice")

	v := viper.New()
	v.BindPFlags(flag)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	hs := &InvoiceSuite{
		PopTestSuite: testingsuite.NewPopTestSuite(testingsuite.CurrentPackage()),
		logger:       logger,
		Viper:        v,
	}

	hs.icnSequencer = sequence.NewDatabaseSequencer(hs.DB(), ICNSequenceName)

	suite.Run(t, hs)
	hs.PopTestSuite.TearDown()
}

func (suite *InvoiceSuite) TestEDIString() {
	suite.T().Run("full EDI string is expected", func(t *testing.T) {
		date := edisegment.G62{
			DateQualifier: 10,
			Date:          "20200909",
		}
		invoice := Invoice858C{
			ISA: edisegment.ISA{
				AuthorizationInformationQualifier: "00",
				AuthorizationInformation:          "0084182369",
				SecurityInformationQualifier:      "00",
				SecurityInformation:               "0000000000",
				InterchangeSenderIDQualifier:      "ZZ",
				InterchangeSenderID:               "MYMOVE         ",
				InterchangeReceiverIDQualifier:    "12",
				InterchangeReceiverID:             "8004171844     ",
				InterchangeDate:                   "060102",
				InterchangeTime:                   "1504",
				InterchangeControlStandards:       "U",
				InterchangeControlVersionNumber:   "00401",
				InterchangeControlNumber:          9999,
				AcknowledgementRequested:          0,
				UsageIndicator:                    "T",
				ComponentElementSeparator:         "|",
			},
			GS: edisegment.GS{
				FunctionalIdentifierCode: "SI",
				ApplicationSendersCode:   "MYMOVE",
				ApplicationReceiversCode: "8004171844",
				Date:                     "190903",
				Time:                     "1617",
				GroupControlNumber:       1,
				ResponsibleAgencyCode:    "X",
				Version:                  "004010",
			},
			Header: []edisegment.Segment{
				&date,
			},
			ST: edisegment.ST{
				TransactionSetIdentifierCode: "858",
				TransactionSetControlNumber:  "ABCDE",
			},
			SE: edisegment.SE{
				NumberOfIncludedSegments:    12345,
				TransactionSetControlNumber: "ABCDE",
			},
			GE: edisegment.GE{
				NumberOfTransactionSetsIncluded: 1,
				GroupControlNumber:              1234567,
			},
			IEA: edisegment.IEA{
				NumberOfIncludedFunctionalGroups: 1,
				InterchangeControlNumber:         9999,
			},
		}

		ediString, err := invoice.EDIString(suite.logger)
		suite.NoError(err)
		suite.Equal(`ISA*00*0084182369*00*0000000000*ZZ*MYMOVE         *12*8004171844     *060102*1504*U*00401*000009999*0*T*|
GS*SI*MYMOVE*8004171844*190903*1617*1*X*004010
ST*858*ABCDE
G62*10*20200909**
SE*12345*ABCDE
GE*1*1234567
IEA*1*000009999
`, ediString)
	})
}

func (suite *InvoiceSuite) TestValidate() {
	suite.T().Run("everything validates successfully", func(t *testing.T) {
		invoice := Invoice858C{
			ISA: edisegment.ISA{
				AuthorizationInformationQualifier: "00",
				AuthorizationInformation:          "0084182369",
				SecurityInformationQualifier:      "00",
				SecurityInformation:               "0000000000",
				InterchangeSenderIDQualifier:      "ZZ",
				InterchangeSenderID:               "MYMOVE         ",
				InterchangeReceiverIDQualifier:    "12",
				InterchangeReceiverID:             "8004171844     ",
				InterchangeDate:                   "060102",
				InterchangeTime:                   "1504",
				InterchangeControlStandards:       "U",
				InterchangeControlVersionNumber:   "00401",
				InterchangeControlNumber:          9999,
				AcknowledgementRequested:          0,
				UsageIndicator:                    "T",
				ComponentElementSeparator:         "|",
			},
			GS: edisegment.GS{
				FunctionalIdentifierCode: "SI",
				ApplicationSendersCode:   "MYMOVE",
				ApplicationReceiversCode: "8004171844",
				Date:                     "190903",
				Time:                     "1617",
				GroupControlNumber:       1,
				ResponsibleAgencyCode:    "X",
				Version:                  "004010",
			},
			ST: edisegment.ST{
				TransactionSetIdentifierCode: "858",
				TransactionSetControlNumber:  "ABCDE",
			},
			SE: edisegment.SE{
				NumberOfIncludedSegments:    12345,
				TransactionSetControlNumber: "ABCDE",
			},
			GE: edisegment.GE{
				NumberOfTransactionSetsIncluded: 1,
				GroupControlNumber:              1234567,
			},
			IEA: edisegment.IEA{
				NumberOfIncludedFunctionalGroups: 1,
				InterchangeControlNumber:         9999,
			},
		}

		err := invoice.Validate()
		suite.NoError(err, "Failed to get invoice 858C as EDI string")
	})
}
