package serviceparamvaluelookups

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/route"
	"github.com/transcom/mymove/pkg/testingsuite"
)

const defaultDistance = 1234

type ServiceParamValueLookupsSuite struct {
	testingsuite.PopTestSuite
	logger  Logger
	planner route.Planner
}

func (suite *ServiceParamValueLookupsSuite) SetupTest() {
	suite.DB().TruncateAll()
}

func TestServiceParamValueLookupsSuite(t *testing.T) {
	ts := &ServiceParamValueLookupsSuite{
		PopTestSuite: testingsuite.NewPopTestSuite(testingsuite.CurrentPackage()),
		logger:       zap.NewNop(), // Use a no-op logger during testing
		planner:      route.NewTestingPlanner(defaultDistance),
	}

	suite.Run(t, ts)
	ts.PopTestSuite.TearDown()
}
