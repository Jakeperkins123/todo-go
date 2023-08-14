package logger

import "go.uber.org/zap"

func InitLogger() *zap.Logger {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
    return log
}
