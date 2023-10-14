package tracer

import (
	"notifsys/internal/config"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var Trace trace.Tracer

func New() {
	cfg := config.Get().APP

	Trace = otel.Tracer(cfg.Name)
	return
}
