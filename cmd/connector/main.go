package main

import (
	sdk "github.com/conduitio/conduit-connector-sdk"

	connector "github.com/ahamidi/conduit-connector-template"
)

func main() {
	sdk.Serve(connector.Specification, connector.NewSource, connector.NewDestination)
}
