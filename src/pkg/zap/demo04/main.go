package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewDevelopment(
		zap.Fields(
			zap.String("log_name", "test_log"),
			zap.String("log_author", "zhang"),
		),
	)
	defer logger.Sync()
	logger.Info("test fields output")
}
