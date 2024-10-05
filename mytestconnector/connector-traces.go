package mytestconnector

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type tracesConnector struct {
	component.StartFunc
	component.ShutdownFunc

	logger *zap.Logger
	config *Config

	tracesConsumer consumer.Traces
}

func newTracesConnector(
	set connector.Settings,
	config component.Config,
	traces consumer.Traces,
) (*tracesConnector, error) {
	return &tracesConnector{
		logger:         set.Logger,
		config:         config.(*Config),
		tracesConsumer: traces,
	}, nil
}

func (c *tracesConnector) ConsumeTraces(ctx context.Context, traces ptrace.Traces) error {
	for i := 0; i < traces.ResourceSpans().Len(); i++ {
		c.tracesConsumer.ConsumeTraces(ctx, traces)
	}

	return nil
}

func (*tracesConnector) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}
