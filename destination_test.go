package connector_test

import (
	"context"
	connector "github.com/ahamidi/conduit-connector-template"
	"strings"
	"testing"
)

func TestConfigureDestination_FailsWhenConfigEmpty(t *testing.T) {
	con := connector.Destination{}
	err := con.Configure(context.Background(), make(map[string]string))
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if strings.HasPrefix(err.Error(), "config is invalid:") {
		t.Errorf("expected error to be about missing config, got %v", err)
	}
}

func TestConfigureDestination_FailsWhenConfigInvalid(t *testing.T) {
	con := connector.Destination{}
	err := con.Configure(context.Background(), map[string]string{"foobar": "foobar"})
	if err != nil {
		t.Errorf("expected no error, got %v", err)

	}

	if strings.HasPrefix(err.Error(), "config is missing:") {
		t.Errorf("expected error to be about invalid config, got %v", err)
	}
}

func TestTeardown_NoOpen(t *testing.T) {
	con := connector.NewDestination()
	err := con.Teardown(context.Background())
	if err != nil {
		t.Errorf("expected no error, got %v", err)

	}
}
