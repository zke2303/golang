package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	name := "zap"
	logger.Info("failed to fetch", zap.String("hello", name))
}
