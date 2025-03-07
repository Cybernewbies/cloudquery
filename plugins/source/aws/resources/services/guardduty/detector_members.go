// Code generated by codegen; DO NOT EDIT.

package guardduty

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DetectorMembers() *schema.Table {
	return &schema.Table{
		Name:      "aws_guardduty_detector_members",
		Resolver:  fetchGuarddutyDetectorMembers,
		Multiplex: client.ServiceAccountRegionMultiplexer("guardduty"),
		Columns: []schema.Column{
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "detector_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountId"),
			},
			{
				Name:     "email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Email"),
			},
			{
				Name:     "master_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MasterId"),
			},
			{
				Name:     "relationship_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RelationshipStatus"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "administrator_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AdministratorId"),
			},
			{
				Name:     "detector_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DetectorId"),
			},
			{
				Name:     "invited_at",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InvitedAt"),
			},
		},
	}
}
