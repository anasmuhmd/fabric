package microsoft_test

import (
	"context"
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zclconf/go-cty/cty"

	"github.com/blackstork-io/fabric/internal/microsoft"
	client_mocks "github.com/blackstork-io/fabric/mocks/internalpkg/microsoft"
	"github.com/blackstork-io/fabric/pkg/diagnostics/diagtest"
	"github.com/blackstork-io/fabric/plugin"
	"github.com/blackstork-io/fabric/plugin/dataspec"
	"github.com/blackstork-io/fabric/plugin/plugindata"
	"github.com/blackstork-io/fabric/plugin/plugintest"
)

type MicrosoftGraphDataSourceTestSuite struct {
	suite.Suite
	plugin *plugin.Schema
	schema *plugin.DataSource
	cli    *client_mocks.MicrosoftGraphClient
}

func TestMicrosoftGraphDataSourceTestSuite(t *testing.T) {
	suite.Run(t, &MicrosoftGraphDataSourceTestSuite{})
}

func (s *MicrosoftGraphDataSourceTestSuite) SetupSuite() {
	s.plugin = microsoft.Plugin("1.0.0", nil, nil, (func(ctx context.Context, apiVersion string, cfg *dataspec.Block) (client microsoft.MicrosoftGraphClient, err error) {
		return s.cli, nil
	}))
	s.schema = s.plugin.DataSources["microsoft_graph"]
}

func (s *MicrosoftGraphDataSourceTestSuite) SetupTest() {
	s.cli = &client_mocks.MicrosoftGraphClient{}
}

func (s *MicrosoftGraphDataSourceTestSuite) TearDownTest() {
	s.cli.AssertExpectations(s.T())
}

func (s *MicrosoftGraphDataSourceTestSuite) TestSchema() {
	s.Require().NotNil(s.plugin)
	s.Require().NotNil(s.schema)
	s.NotNil(s.schema.Args)
	s.NotNil(s.schema.DataFunc)
	s.NotNil(s.schema.Config)
}

func (s *MicrosoftGraphDataSourceTestSuite) TestBasic() {
	expectedData := map[string]interface{}{
		"value": []interface{}{map[string]interface{}{
			"severity":    "High",
			"displayName": "Incident 1",
		}},
	}
	s.cli.On("QueryGraph", mock.Anything, "/security/incidents", url.Values{"$top": []string{"10"}}).Return(expectedData, nil)
	ctx := context.Background()
	result, diags := s.schema.DataFunc(ctx, &plugin.RetrieveDataParams{
		Config: plugintest.NewTestDecoder(s.T(), s.schema.Config).
			SetAttr("client_id", cty.StringVal("cid")).
			SetAttr("tenant_id", cty.StringVal("tid")).
			SetAttr("client_secret", cty.StringVal("csecret")).
			Decode(),
		Args: plugintest.NewTestDecoder(s.T(), s.schema.Args).
			SetAttr("endpoint", cty.StringVal("/security/incidents")).
			SetAttr("api_version", cty.StringVal("v1")).
			SetAttr("query_params", cty.MapVal(map[string]cty.Value{"$top": cty.StringVal("10")})).
			Decode(),
	})
	s.Nil(diags)
	parsedData, err := plugindata.ParseAny(expectedData)
	s.Nil(err)
	s.Equal(parsedData, result.AsPluginData())
}

func (s *MicrosoftGraphDataSourceTestSuite) TestClientError() {
	s.cli.On("QueryGraph", mock.Anything, "/security/incidents", url.Values{"$top": []string{"10"}}).Return(nil, errors.New("microsoft graph client returned status code: 400"))
	ctx := context.Background()
	result, diags := s.schema.DataFunc(ctx, &plugin.RetrieveDataParams{
		Config: plugintest.NewTestDecoder(s.T(), s.schema.Config).
			SetAttr("client_id", cty.StringVal("cid")).
			SetAttr("tenant_id", cty.StringVal("tid")).
			SetAttr("client_secret", cty.StringVal("csecret")).
			Decode(),
		Args: plugintest.NewTestDecoder(s.T(), s.schema.Args).
			SetAttr("endpoint", cty.StringVal("/security/incidents")).
			SetAttr("api_version", cty.StringVal("v1")).
			SetAttr("query_params", cty.MapVal(map[string]cty.Value{"$top": cty.StringVal("10")})).
			Decode(),
	})
	s.Nil(result)
	diagtest.Asserts{{
		diagtest.IsError,
		diagtest.DetailContains("microsoft graph client returned status code: 400"),
	}}.AssertMatch(s.T(), diags, nil)
}

func (s *MicrosoftGraphDataSourceTestSuite) TestMissingArgs() {
	plugintest.NewTestDecoder(
		s.T(),
		s.schema.Args,
	).Decode([]diagtest.Assert{
		diagtest.IsError,
		diagtest.SummaryEquals("Missing required attribute"),
		diagtest.DetailContains("endpoint"),
	})
}

func (s *MicrosoftGraphDataSourceTestSuite) TestMissingConfig() {
	plugintest.NewTestDecoder(
		s.T(),
		s.schema.Config,
	).Decode([]diagtest.Assert{
		diagtest.IsError,
		diagtest.SummaryEquals("Missing required attribute"),
		diagtest.DetailContains("client_id"),
	}, []diagtest.Assert{
		diagtest.IsError,
		diagtest.SummaryEquals("Missing required attribute"),
		diagtest.DetailContains("tenant_id"),
	})
}

func (s *MicrosoftGraphDataSourceTestSuite) TestMissingCredentials() {
	expectedData := map[string]interface{}{
		"value": []interface{}{map[string]interface{}{
			"severity":    "High",
			"displayName": "Incident 1",
		}},
	}
	s.cli.On("QueryGraph", mock.Anything, "/security/incidents", url.Values{"$top": []string{"10"}}).Return(expectedData, nil)
	ctx := context.Background()
	result, diags := s.schema.DataFunc(ctx, &plugin.RetrieveDataParams{
		Config: plugintest.NewTestDecoder(s.T(), s.schema.Config).
			SetAttr("client_id", cty.StringVal("cid")).
			SetAttr("tenant_id", cty.StringVal("tid")).
			Decode(),
		Args: plugintest.NewTestDecoder(s.T(), s.schema.Args).
			SetAttr("endpoint", cty.StringVal("/security/incidents")).
			SetAttr("api_version", cty.StringVal("v1")).
			SetAttr("query_params", cty.MapVal(map[string]cty.Value{"$top": cty.StringVal("10")})).
			Decode(),
	})
	s.Nil(diags)
	parsedData, err := plugindata.ParseAny(expectedData)
	s.Nil(err)
	s.Equal(parsedData, result.AsPluginData())
}
