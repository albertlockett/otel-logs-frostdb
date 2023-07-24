package frostdbexporter

import (
	"context"

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
	_ context.Context,
	params exporter.CreateSettings,
	baseCfg component.Config,
) (exporter.Logs, error) {
	exporter := frostdbExporter{}
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
