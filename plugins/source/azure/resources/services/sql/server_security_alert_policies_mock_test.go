// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func createServerSecurityAlertPoliciesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLServerSecurityAlertPoliciesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ServerSecurityAlertPolicies: mockClient,
		},
	}

	data := sql.ServerSecurityAlertPolicy{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewLogicalServerSecurityAlertPolicyListResultPage(sql.LogicalServerSecurityAlertPolicyListResult{Value: &[]sql.ServerSecurityAlertPolicy{data}}, func(ctx context.Context, result sql.LogicalServerSecurityAlertPolicyListResult) (sql.LogicalServerSecurityAlertPolicyListResult, error) {
		return sql.LogicalServerSecurityAlertPolicyListResult{}, nil
	})

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
