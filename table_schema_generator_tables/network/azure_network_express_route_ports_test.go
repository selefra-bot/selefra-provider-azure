package network

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createExpressRoutePortsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkExpressRoutePortsClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			ExpressRoutePorts: mockClient,
		},
	}

	data := network.ExpressRoutePort{}
	require.Nil(t, faker.FakeObject(&data))

	result := network.NewExpressRoutePortListResultPage(network.ExpressRoutePortListResult{Value: &[]network.ExpressRoutePort{data}}, func(ctx context.Context, result network.ExpressRoutePortListResult) (network.ExpressRoutePortListResult, error) {
		return network.ExpressRoutePortListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestNetworkExpressRoutePorts(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureNetworkExpressRoutePortsGenerator{}), createExpressRoutePortsMock, azure_client.TestOptions{})
}
