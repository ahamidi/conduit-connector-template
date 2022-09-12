package main

import (
	sdk "github.com/conduitio/conduit-connector-sdk"

	connectorName "github.com/ahamidi/conduit-connector-template"
)

func main() {
	sdk.Serve(connectorName.Connector)
}
