package main

import "go.uber.org/zap"

func main() {

	logger := zap.NewExample()
	logger.Debug("This is a bug")
	logger.Info("This is a Info")

}
