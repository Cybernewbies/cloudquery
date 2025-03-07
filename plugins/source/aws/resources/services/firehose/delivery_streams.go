// Code generated by codegen; DO NOT EDIT.

package firehose

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DeliveryStreams() *schema.Table {
	return &schema.Table{
		Name:                "aws_firehose_delivery_streams",
		Resolver:            fetchFirehoseDeliveryStreams,
		PreResourceResolver: getDeliveryStream,
		Multiplex:           client.ServiceAccountRegionMultiplexer("firehose"),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveFirehoseDeliveryStreamTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeliveryStreamARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "delivery_stream_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeliveryStreamName"),
			},
			{
				Name:     "delivery_stream_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeliveryStreamStatus"),
			},
			{
				Name:     "delivery_stream_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeliveryStreamType"),
			},
			{
				Name:     "destinations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Destinations"),
			},
			{
				Name:     "has_more_destinations",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HasMoreDestinations"),
			},
			{
				Name:     "version_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VersionId"),
			},
			{
				Name:     "create_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreateTimestamp"),
			},
			{
				Name:     "delivery_stream_encryption_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DeliveryStreamEncryptionConfiguration"),
			},
			{
				Name:     "failure_description",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FailureDescription"),
			},
			{
				Name:     "last_update_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdateTimestamp"),
			},
			{
				Name:     "source",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Source"),
			},
		},
	}
}
