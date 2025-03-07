// Code generated by codegen; DO NOT EDIT.

package services

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func Apps() *schema.Table {
	return &schema.Table{
		Name:        "heroku_apps",
		Description: "https://devcenter.heroku.com/articles/platform-api-reference#app-attributes",
		Resolver:    fetchApps,
		Columns: []schema.Column{
			{
				Name:     "acm",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Acm"),
			},
			{
				Name:     "archived_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ArchivedAt"),
			},
			{
				Name:     "build_stack",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BuildStack"),
			},
			{
				Name:     "buildpack_provided_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BuildpackProvidedDescription"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "git_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GitURL"),
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
				Name:     "internal_routing",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("InternalRouting"),
			},
			{
				Name:     "maintenance",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Maintenance"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "organization",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Organization"),
			},
			{
				Name:     "owner",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Owner"),
			},
			{
				Name:     "region",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "released_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ReleasedAt"),
			},
			{
				Name:     "repo_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("RepoSize"),
			},
			{
				Name:     "slug_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SlugSize"),
			},
			{
				Name:     "space",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Space"),
			},
			{
				Name:     "stack",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Stack"),
			},
			{
				Name:     "team",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Team"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "web_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WebURL"),
			},
		},
	}
}

func fetchApps(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.AppList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- v
	}
	return nil
}
