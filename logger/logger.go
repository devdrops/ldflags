package logger

import (
	"github.com/devdrops/ldflags/appinfo"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitGlobalLogger() {
	logger, _ := zap.NewProduction(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	defer logger.Sync()

	logger = logger.With(
		zap.String("app_name", appinfo.AppName),
		zap.String("app_version", appinfo.AppVersion),
		zap.String("app_build_at", appinfo.AppBuildAt),
	)

	zap.ReplaceGlobals(logger)
}
