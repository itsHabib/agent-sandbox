package main

import (
	"flag"
	"log/slog"
	"os"
)

func main() {
	debug := flag.Bool("debug", false, "enable debug logging")
	flag.Parse()

	level := slog.LevelInfo
	if *debug {
		level = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)

	slog.Debug("debug logging enabled")
	slog.Info("agent-sandbox")
}
