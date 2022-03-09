package connector

import (
	"context"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

type Destination struct {
	sdk.UnimplementedDestination

	config DestinationConfig
}

func NewDestination() sdk.Destination {
	return &Destination{}
}

func (d *Destination) Configure(ctx context.Context, cfg map[string]string) error {
	sdk.Logger(ctx).Info().Msg("Configuring a Destination connector...")
	config, err := ParseDestinationConfig(cfg)
	if err != nil {
		return err
	}
	d.config = config
	return nil
}

func (d *Destination) Open(ctx context.Context) error {
	return nil
}

func (d *Destination) Write(ctx context.Context, record sdk.Record) error {
	return nil
}

func (d *Destination) Flush(context.Context) error {
	return nil
}

func (d *Destination) Teardown(ctx context.Context) error {
	return nil
}
