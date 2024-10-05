package mytestconnector

import (
	"context"

	"github.com/practical-otel/collector-extension-template/mytestconnector/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
)

// NewFactory creates a factory for the spanmetrics connector.
func NewFactory() connector.Factory {
	return connector.NewFactory(
		metadata.Type,
		createDefaultConfig,
		connector.WithTracesToTraces(createTracesToTraces, metadata.TracesToTracesStability),
		connector.WithLogsToLogs(createLogsToLogs, metadata.LogsToLogsStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

func createTracesToTraces(
	_ context.Context,
	set connector.Settings,
	cfg component.Config,
	traces consumer.Traces,
) (connector.Traces, error) {
	return newTracesConnector(set, cfg, traces)
}

func createLogsToLogs(
	_ context.Context,
	set connector.Settings,
	cfg component.Config,
	logs consumer.Logs,
) (connector.Logs, error) {
	return newLogsConnector(set, cfg, logs)
}
