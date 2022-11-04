package security

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createContactsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSecurityContactsClient(ctrl)
	s := services.Services{
		Security: services.SecurityClient{
			Contacts: mockClient,
		},
	}

	data := security.Contact{}
	require.Nil(t, faker.FakeObject(&data))

	result := security.NewContactListPage(security.ContactList{Value: &[]security.Contact{data}}, func(ctx context.Context, result security.ContactList) (security.ContactList, error) {
		return security.ContactList{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestSecurityContacts(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureSecurityContactsGenerator{}), createContactsMock, azure_client.TestOptions{})
}
