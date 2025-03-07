// Code generated by codegen; DO NOT EDIT.

package elasticache

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EngineVersions() *schema.Table {
	return &schema.Table{
		Name:      "aws_elasticache_engine_versions",
		Resolver:  fetchElasticacheEngineVersions,
		Multiplex: client.ServiceAccountRegionMultiplexer("elasticache"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
				Description: `The AWS Account ID of the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "region",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
				Description: `The AWS Region of the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "cache_engine_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheEngineDescription"),
			},
			{
				Name:     "cache_engine_version_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheEngineVersionDescription"),
			},
			{
				Name:     "cache_parameter_group_family",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheParameterGroupFamily"),
			},
			{
				Name:     "engine",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Engine"),
			},
			{
				Name:     "engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EngineVersion"),
			},
		},
	}
}
