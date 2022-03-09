package connector

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

// Specification returns the connector's specification.
func Specification() sdk.Specification {
	return sdk.Specification{
		Name:        "<connector name>",
		Summary:     "<describe your connector>",
		Description: "<describe your connector in detail>",
		Version:     "v0.1.0",
		Author:      "<your name>",
		SourceParams: map[string]sdk.Parameter{
			GlobalConfigParam: {
				Default:     "localhost:10000",
				Required:    true,
				Description: "The URL of the server.",
			},
			SourceConfigParam: {
				Default:     "false",
				Required:    false,
				Description: "An optional configuration parameter.",
			},
		},
		DestinationParams: map[string]sdk.Parameter{
			GlobalConfigParam: {
				Default:     "localhost:10000",
				Required:    true,
				Description: "The URL of the server.",
			},
			DestinationConfigParam: {
				Default:     "false",
				Required:    false,
				Description: "An optional configuration parameter.",
			},
		},
	}
}
