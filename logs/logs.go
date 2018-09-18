package logs

import (
	"github.com/wuqifei/server_lib/logs2"
	"github.com/wuqifei/server_lib/logs_plugin"
)

// 新建日志
func NewLog(filepath string) {
	m := &logs_plugin.MultiFileLogWriter{}
	m.Separate = []string{"emergency", "critical", "error", "warning", "info", "debug"}
	m.FullLogWriter = &logs_plugin.FileLogWriter{
		Daily:      true,
		MaxDays:    7,
		Rotate:     true,
		RotatePerm: "0440",
		Level:      logs2.LogLevelDebug,
		Perm:       "0660",
		Filename:   filepath,
	}
	logs2.DefaultLogger().Register("mutifile", m, logs_plugin.NewFilesWriter())
	logs2.DefaultLogger().SetDefaultLevel(logs2.LogLevelDebug)
	logs2.DefaultLogger().Async(3, 100)

}

// 紧急
func Emergency(f interface{}, v ...interface{}) {
	logs2.Emergency(f, v...)
}

// 严格的
func Critical(f interface{}, v ...interface{}) {
	logs2.Critical(f, v...)
}

// 错误的
func Error(f interface{}, v ...interface{}) {
	logs2.Error(f, v...)
}

// 警告
func Warning(f interface{}, v ...interface{}) {
	logs2.Warning(f, v...)
}

// 一般信息
func Info(f interface{}, v ...interface{}) {
	logs2.Info(f, v...)
}

// 一般信息
func Debug(f interface{}, v ...interface{}) {
	logs2.Debug(f, v...)
}
