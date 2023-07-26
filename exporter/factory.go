package frostdbexporter

import (
	"context"

	"github.com/albertlockett/otel-logs-frostdb/query"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

const (
	typeStr = "frostdb"
)

func createDefaultConfig() component.Config {
	return &Config{}
}

func createLogsExporter(
	ctx context.Context,
	params exporter.CreateSettings,
	baseCfg component.Config,
) (exporter.Logs, error) {
	columnstore, err := createColumnStore(ctx)
	if err != nil {
		return nil, err
	}

	err = query.StartQueryServer()
	if err != nil {
		return nil, err
	}

	startQueryPoller(columnstore)

	exporter := frostdbExporter{
		columnstore,
	}

	return &exporter, nil
}

func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		typeStr,
		createDefaultConfig,
		exporter.WithLogs(
			createLogsExporter,
			component.StabilityLevelAlpha,
		),
	)
}
