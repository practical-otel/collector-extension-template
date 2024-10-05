package mytestconnector

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

type logsConnector struct {
	component.StartFunc
	component.ShutdownFunc

	logger *zap.Logger
	config *Config

	logsConsumer consumer.Logs
}

func newLogsConnector(
	set connector.Settings,
	config component.Config,
	logs consumer.Logs,
) (*logsConnector, error) {
	return &logsConnector{
		logger:       set.Logger,
		config:       config.(*Config),
		logsConsumer: logs,
	}, nil
}

func (c *logsConnector) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	for i := 0; i < logs.ResourceLogs().Len(); i++ {
		c.logsConsumer.ConsumeLogs(ctx, logs)
	}

	return nil
}

func (*logsConnector) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}
