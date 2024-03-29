package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// defer sugar.Sync()
	// sugar.Infow("failed to fetch URL",
	// 	"url", "http://example.com",
	// 	"attempt", 3,
	// 	"backoff", time.Second,
	// )
	// sugar.Infof("failed to fetch URL: %s", "http://example.com")
	url := "http://www.google.com"

	// logger, _ := zap.NewProduction()
	// defer logger.Sync() // flushes buffer, if any
	// sugar := logger.Sugar()
	// sugar.Infow("failed to fetch URL",
	// 	// Structured context as loosely typed key-value pairs.
	// 	"url", url,
	// 	"attempt", 3,
	// 	"backoff", time.Second,
	// )
	// sugar.Infof("Failed to fetch URL: %s", url)

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
