// Code generated by codegen; DO NOT EDIT.

package ecr

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Repositories() *schema.Table {
	return &schema.Table{
		Name:      "aws_ecr_repositories",
		Resolver:  fetchEcrRepositories,
		Multiplex: client.ServiceAccountRegionMultiplexer("api.ecr"),
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
				Resolver: schema.PathResolver("RepositoryArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRepositoryTags,
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "encryption_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EncryptionConfiguration"),
			},
			{
				Name:     "image_scanning_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ImageScanningConfiguration"),
			},
			{
				Name:     "image_tag_mutability",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageTagMutability"),
			},
			{
				Name:     "registry_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegistryId"),
			},
			{
				Name:     "repository_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RepositoryName"),
			},
			{
				Name:     "repository_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RepositoryUri"),
			},
		},

		Relations: []*schema.Table{
			RepositoryImages(),
		},
	}
}
