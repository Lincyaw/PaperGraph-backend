package logger

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"time"
)

func New() *slog.Logger {
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,            // 添加日志的源文件和行号
		Level:     slog.LevelDebug, // 设置日志级别为 Info
	})

	logger := slog.New(h)

	return logger
}

var Global = New()

func Debug(msg string, args ...any) {
	if !Global.Enabled(context.Background(), slog.LevelDebug) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelDebug, msg, pcs[0])
	r.Add(args...)
	_ = Global.Handler().Handle(context.Background(), r)
}

func Info(msg string, args ...any) {
	if !Global.Enabled(context.Background(), slog.LevelInfo) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelInfo, msg, pcs[0])
	r.Add(args...)
	_ = Global.Handler().Handle(context.Background(), r)
}

func Warn(msg string, args ...any) {
	if !Global.Enabled(context.Background(), slog.LevelWarn) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelWarn, msg, pcs[0])
	r.Add(args...)

	_ = Global.Handler().Handle(context.Background(), r)
}

func Error(msg string, args ...any) {
	if !Global.Enabled(context.Background(), slog.LevelError) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelError, msg, pcs[0])
	r.Add(args...)

	_ = Global.Handler().Handle(context.Background(), r)
}

func With(args ...any) *slog.Logger {
	return Global.With(args...)
}
