package logger

import (
	"io"
	"log/slog"
	"os"
	"time"
)

type Options struct {
	Format     Format
	Level      slog.Level
	Output     io.Writer
	AddSource  bool
	TimeFormat string
}

type Format int

const (
	TextFormat Format = iota
	JSONFormat
)

func New(opts Options) *slog.Logger {
	if opts.Output == nil {
		opts.Output = os.Stdout
	}

	handlerOpts := &slog.HandlerOptions{
		Level:     opts.Level,
		AddSource: opts.AddSource,
	}

	if opts.TimeFormat != "" {
		handlerOpts.ReplaceAttr = func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					return slog.String(slog.TimeKey, t.Format(opts.TimeFormat))
				}
			}
			return a
		}
	}

	var handler slog.Handler
	switch opts.Format {
	case JSONFormat:
		handler = slog.NewJSONHandler(opts.Output, handlerOpts)
	default:
		handler = slog.NewTextHandler(opts.Output, handlerOpts)
	}

	return slog.New(handler)
}
