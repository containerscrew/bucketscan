package logger

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"github.com/fatih/color"
	"golang.org/x/exp/slog"
)

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	// By default, all messages are white color.
	// You can customize each message inside the switch
	var msg = color.WhiteString(r.Message)

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
		msg = color.MagentaString(r.Message)
	case slog.LevelInfo:
		level = color.BlueString(level)
		msg = color.BlueString(r.Message)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	//b, err := json.MarshalIndent(fields, "", "  ")
	b, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[15:05:05]")
	//msg := color.CyanString(r.Message)

	h.l.Println(timeStr, level, msg, color.WhiteString(string(b)))

	return nil
}

func NewPrettyHandler(out io.Writer, opts PrettyHandlerOptions) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewTextHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}
