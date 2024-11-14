package log

import (
	"github.com/rs/zerolog"
)

type TracingHook struct{}

func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	var requestId string

	_ = e.GetCtx()

	e.Str("request-id", requestId)
}
