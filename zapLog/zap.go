package zapLog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger

func InitLogger() {
	encoder := getEncoder()
	writerSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writerSyncer, zap.DebugLevel)
	Logger = zap.New(core, zap.AddCaller())
	SugarLogger = Logger.Sugar()
}
func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
}
func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./log.log")
	return zapcore.AddSync(file)
}
