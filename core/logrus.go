package core

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"wheatwaves/global"

	"github.com/sirupsen/logrus"
)

// 颜色常量
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type logFormatter struct{}

// 重写实现Formatter（entery *logurs.Entry）接口
func (t *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int // 不同level 不同颜色
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	case logrus.WarnLevel:
		levelColor = yellow
	default:
		levelColor = blue
	}
	var b *bytes.Buffer // 缓冲池
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	prefix := global.Config.Logger.Prefix
	// 自定义日期格式
	timestamp := entry.Time.Format("2024-06-10 00:00:00")
	if entry.HasCaller() {
		// 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// 自定义输出格式
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s】 \x1b[%dm[%s]\x1b[0m %s\n", prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	mLog := logrus.New()               // 新建实例
	mLog.SetOutput(os.Stdout)          // 设置输出类型
	mLog.SetReportCaller(true)         // 开启返回函数名和行号
	mLog.SetFormatter(&logFormatter{}) // 设置定义的formatter
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level) // 设置最低level
	return mLog
}

func InitDefaultLogger() {
	// 全局log
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logFormatter{})
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}
