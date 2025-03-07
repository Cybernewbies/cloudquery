// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func serverAdmins() *schema.Table {
	return &schema.Table{
		Name:     "azure_sql_server_admins",
		Resolver: fetchSQLServerAdmins,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "sql_server_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "administrator_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AdministratorType"),
			},
			{
				Name:     "login",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Login"),
			},
			{
				Name:     "sid",
				Type:     schema.TypeUUID,
				Resolver: schema.PathResolver("Sid"),
			},
			{
				Name:     "tenant_id",
				Type:     schema.TypeUUID,
				Resolver: schema.PathResolver("TenantID"),
			},
			{
				Name:     "azure_ad_only_authentication",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AzureADOnlyAuthentication"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchSQLServerAdmins(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ServerAdmins

	server := parent.Item.(sql.Server)
	resourceDetails, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return err
	}
	response, err := svc.ListByServer(ctx, resourceDetails.ResourceGroup, *server.Name)

	if err != nil {
		return err
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}

	return nil
}
