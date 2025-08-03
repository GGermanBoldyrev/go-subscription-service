package logger

import (
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func Init() {
	raw, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	log = raw.Sugar()
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	log.Fatalf(template, args...)
}
