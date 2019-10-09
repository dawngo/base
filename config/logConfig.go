package config

import (
	"os"
	"path"

	"gopkg.in/natefinch/lumberjack.v2"
)

var projectDir, _ = os.Getwd()
var LogConfig = lumberjack.Logger{
	Filename:   path.Join(projectDir, "log", "app.log"), // 日志路径
	MaxSize:    128,                                     // 日志大小
	MaxAge:     60,                                      // 文件最多保存多少天
	MaxBackups: 31,                                      // 日志文件最多保存多少个备份
	LocalTime:  true,
	Compress:   false,
}
