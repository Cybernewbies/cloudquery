// Auto generated code - DO NOT EDIT.

package network

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
)

func TestNetworkPublicIPAddresses(t *testing.T) {
	client.MockTestHelper(t, PublicIPAddresses(), createPublicIPAddressesMock)
}

func createPublicIPAddressesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkPublicIPAddressesClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			PublicIPAddresses: mockClient,
		},
	}

	data := network.PublicIPAddress{}
	require.Nil(t, faker.FakeObject(&data))

	result := network.NewPublicIPAddressListResultPage(network.PublicIPAddressListResult{Value: &[]network.PublicIPAddress{data}}, func(ctx context.Context, result network.PublicIPAddressListResult) (network.PublicIPAddressListResult, error) {
		return network.PublicIPAddressListResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any()).Return(result, nil)
	return s
}
