package main

import (
	"log/slog"
	"os"
	"runtime"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}).WithAttrs([]slog.Attr{slog.String("app_version", "1.0.0")}))
	slog.SetDefault(logger)

	slog.Info("Golang", slog.String("version", runtime.Version()), slog.Group("Os Info",
		slog.String("os", runtime.GOOS),
		slog.String("arch", runtime.GOARCH),
		slog.String("compiler", runtime.Compiler),
		slog.Int("CPU", runtime.NumCPU()),
	))
	slog.Error("Wow, Golang")
	slog.Warn("Upa, Golang")
	slog.Debug("Debugging Golang")
}
