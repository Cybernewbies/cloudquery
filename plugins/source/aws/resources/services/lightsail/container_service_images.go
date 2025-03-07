// Code generated by codegen; DO NOT EDIT.

package lightsail

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ContainerServiceImages() *schema.Table {
	return &schema.Table{
		Name:      "aws_lightsail_container_service_images",
		Resolver:  fetchLightsailContainerServiceImages,
		Multiplex: client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Name:     "container_service_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "digest",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Digest"),
			},
			{
				Name:     "image",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Image"),
			},
		},
	}
}
