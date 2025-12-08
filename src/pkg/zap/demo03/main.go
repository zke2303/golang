package main

import "go.uber.org/zap"

func main(){
	logger,  _ := zap.NewProduction()
	defer logger.Sync()
	logger.Debug("this is a bug")
	logger.Info("this is a Info")
}
