package logger

import (
	"github.com/Brave-man/base/utils/logger"
	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {
	Log = logger.New()
}
