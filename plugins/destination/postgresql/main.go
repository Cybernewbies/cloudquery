package main

import (
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/client"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/serve"
)

const (
	sentryDSN = "https://19d1257d36854a51b17c06614e76dc2d@o1396617.ingest.sentry.io/4503896817336320"
)

var (
	Version = "Development"
)

func main() {
	p := plugins.NewDestinationPlugin("postgresql", Version, client.New)
	serve.Destination(p, serve.WithDestinationSentryDSN(sentryDSN))
}
