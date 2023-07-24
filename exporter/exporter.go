package frostdbexporter

import (
	"context"
	"log"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
)

type frostdbExporter struct {
}

func (f *frostdbExporter) Start(ctx context.Context, host component.Host) error {
	// TODO
	return nil
}

func (f *frostdbExporter) Shutdown(ctx context.Context) error {
	// TODO
	return nil
}

func (f *frostdbExporter) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{
		MutatesData: false,
	}
}

func (f *frostdbExporter) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	log.Printf("log received")
	return nil
}
