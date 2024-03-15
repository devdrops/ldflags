package main

import (
	"github.com/devdrops/ldflags/logger"

	"go.uber.org/zap"
)

func main() {
	logger.InitGlobalLogger()

	zap.L().Info("application up and running")
}
