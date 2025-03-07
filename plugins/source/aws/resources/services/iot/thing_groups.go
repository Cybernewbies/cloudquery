// Code generated by codegen; DO NOT EDIT.

package iot

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ThingGroups() *schema.Table {
	return &schema.Table{
		Name:      "aws_iot_thing_groups",
		Resolver:  fetchIotThingGroups,
		Multiplex: client.ServiceAccountRegionMultiplexer("iot"),
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
				Name:     "things_in_group",
				Type:     schema.TypeStringArray,
				Resolver: ResolveIotThingGroupThingsInGroup,
			},
			{
				Name:          "policies",
				Type:          schema.TypeStringArray,
				Resolver:      ResolveIotThingGroupPolicies,
				IgnoreInTests: true,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: ResolveIotThingGroupTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ThingGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "index_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IndexName"),
			},
			{
				Name:     "query_string",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("QueryString"),
			},
			{
				Name:     "query_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("QueryVersion"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "thing_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ThingGroupId"),
			},
			{
				Name:     "thing_group_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ThingGroupMetadata"),
			},
			{
				Name:     "thing_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ThingGroupName"),
			},
			{
				Name:     "thing_group_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ThingGroupProperties"),
			},
			{
				Name:     "version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Version"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},
	}
}
