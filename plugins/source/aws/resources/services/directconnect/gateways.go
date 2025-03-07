// Code generated by codegen; DO NOT EDIT.

package directconnect

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Gateways() *schema.Table {
	return &schema.Table{
		Name:        "aws_directconnect_gateways",
		Description: "https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DirectConnectGateway.html",
		Resolver:    fetchDirectconnectGateways,
		Multiplex:   client.ServiceAccountRegionMultiplexer("directconnect"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveGatewayARN(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectConnectGatewayId"),
			},
			{
				Name:     "amazon_side_asn",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AmazonSideAsn"),
			},
			{
				Name:     "direct_connect_gateway_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectConnectGatewayName"),
			},
			{
				Name:     "direct_connect_gateway_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectConnectGatewayState"),
			},
			{
				Name:     "owner_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerAccount"),
			},
			{
				Name:     "state_change_error",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StateChangeError"),
			},
		},

		Relations: []*schema.Table{
			GatewayAssociations(),
			GatewayAttachments(),
		},
	}
}
