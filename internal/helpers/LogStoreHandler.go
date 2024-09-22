package helpers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"project/config"
	"sync"
)

var _ slog.Handler = (*LogStoreHandler)(nil) // verify it extends the slog interface

type LogStoreHandler struct {
	slogHandler slog.Handler
	buffer      *bytes.Buffer
	mutex       *sync.Mutex
}

func (handler *LogStoreHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return handler.slogHandler.Enabled(ctx, level)
}

func (handler *LogStoreHandler) Handle(ctx context.Context, record slog.Record) error {
	level := record.Level.String()
	message := record.Message
	attrs, err := handler.computeAttrs(ctx, record)

	if err != nil {
		return fmt.Errorf("error when calling computeAttrs: %w", err)
	}

	if level == slog.LevelDebug.String() {
		_, err := config.LogStore.DebugWithContext(message, attrs)
		return err
	}

	if level == slog.LevelInfo.String() {
		_, err := config.LogStore.InfoWithContext(message, attrs)
		return err
	}

	if level == slog.LevelWarn.String() {
		_, err := config.LogStore.WarnWithContext(message, attrs)
		return err
	}

	if level == slog.LevelError.String() {
		_, err := config.LogStore.ErrorWithContext(message, attrs)
		return err
	}

	_, err = config.LogStore.FatalWithContext(message, attrs)
	return err
}

func (handler *LogStoreHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &LogStoreHandler{
		slogHandler: handler.slogHandler.WithAttrs(attrs),
		buffer:      handler.buffer,
		mutex:       handler.mutex,
	}
}

func (handler *LogStoreHandler) WithGroup(name string) slog.Handler {
	return &LogStoreHandler{
		slogHandler: handler.slogHandler.WithGroup(name),
		buffer:      handler.buffer,
		mutex:       handler.mutex,
	}
}

func (handler *LogStoreHandler) computeAttrs(
	ctx context.Context,
	r slog.Record,
) (map[string]any, error) {
	handler.mutex.Lock()
	defer func() {
		handler.buffer.Reset()
		handler.mutex.Unlock()
	}()

	if err := handler.slogHandler.Handle(ctx, r); err != nil {
		return nil, fmt.Errorf("error when calling inner handler's Handle: %w", err)
	}

	var attrs map[string]any
	err := json.Unmarshal(handler.buffer.Bytes(), &attrs)

	if err != nil {
		return nil, fmt.Errorf("error when unmarshaling inner handler's Handle result: %w", err)
	}

	return attrs, nil
}
