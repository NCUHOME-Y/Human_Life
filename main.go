package main

import (
	"Hack/initDatabase"
	"Hack/router"
	"Hack/zapLog"
	"go.uber.org/zap"
)

// var contextTimeOut context.Context
//var cancel context.CancelFunc

// IsAmt will Check the struct and return true or false

func main() {
	zapLog.InitLogger()
	initDatabase.InitDB()
	router.InitRouter()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Info("sync err", zap.Error(err))
		}
	}(zapLog.Logger)
	zapLog.SugarLogger.Info("write a err")

}
