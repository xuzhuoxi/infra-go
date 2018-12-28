package logx

import (
	"github.com/xuzhuoxi/util-go/mathx"
	"log"
)

type LogLevel uint8
type LogType uint8

const (
	LevelAll LogLevel = iota
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelOff
)

const (
	TypeConsole LogType = iota
	TypeRollingFile
	TypeDailyFile
	TypeDailyRollingFile
)

var (
	level2prefix          = make(map[LogLevel]string)
	defaultLogger ILogger = NewLogger()
)

func init() {
	level2prefix[LevelTrace] = "[Trace]"
	level2prefix[LevelDebug] = "[Debug]"
	level2prefix[LevelInfo] = "[Info]"
	level2prefix[LevelWarn] = "[Warn]"
	level2prefix[LevelError] = "[Error]"
	level2prefix[LevelFatal] = "[Fatal]"
	defaultLogger.SetConfig(LogConfig{Type: TypeConsole, Level: LevelAll})
}

type LogConfig struct {
	Type        LogType
	Level       LogLevel
	Flag        int
	FileDir     string
	FileName    string
	FileExtName string
	MaxSize     mathx.SizeUint
}

func NewLogger() ILogger {
	instance := &logger{prefix: "", defaultFlag: log.LstdFlags, infoMap: make(map[LogType]*logInfo)}
	return instance
}

type ILogger interface {
	//设置日志前缀
	SetPrefix(prefix string)
	//设置日志等级，只有大于等级设置等级的日志才会记录
	//重置日志等级为level,t为空时重置全部
	SetLevel(level LogLevel, t ...LogType)
	// SetFlags sets the output flags for the logger.
	//重置日志flag,t为空时重置全部
	SetFlags(flag int, t ...LogType)
	//配置Log,要求fileDir以"/"结尾
	SetConfig(cfg LogConfig)
	//移除配置
	RemoveConfig(t LogType)

	Log(level LogLevel, v ...interface{})
	Logf(level LogLevel, format string, v ...interface{})
	Logln(level LogLevel, v ...interface{})

	Trace(v ...interface{})
	Tracef(format string, v ...interface{})
	Traceln(v ...interface{})
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Debugln(v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Infoln(v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Warnln(v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
}

func SetPrefix(prefix string) {
	defaultLogger.SetPrefix(prefix)
}

func SetLevel(level LogLevel, t ...LogType) {
	defaultLogger.SetLevel(level, t...)
}

func SetFlags(flag int, t ...LogType) {
	defaultLogger.SetFlags(flag, t...)
}

func SetConfig(cfg LogConfig) {
	defaultLogger.SetConfig(cfg)
}

func RemoveConfig(t LogType) {
	defaultLogger.RemoveConfig(t)
}

func Trace(v ...interface{}) {
	defaultLogger.Trace(v...)
}

func Tracef(format string, v ...interface{}) {
	defaultLogger.Tracef(format, v...)
}

func Traceln(v ...interface{}) {
	defaultLogger.Traceln(v...)
}

func Debug(v ...interface{}) {
	defaultLogger.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	defaultLogger.Debugf(format, v...)
}

func Debugln(v ...interface{}) {
	defaultLogger.Debugln(v...)
}

func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

func Infoln(v ...interface{}) {
	defaultLogger.Infoln(v...)
}

func Warn(v ...interface{}) {
	defaultLogger.Warn(v...)
}

func Warnf(format string, v ...interface{}) {
	defaultLogger.Warnf(format, v...)
}

func Warnln(v ...interface{}) {
	defaultLogger.Warnln(v...)
}

func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

func Errorln(v ...interface{}) {
	defaultLogger.Errorln(v...)
}

func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	defaultLogger.Fatalln(v...)
}
