package logx

import (
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/mathx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"log"
)

type LogType uint8
type LogLevel uint8

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
	// TypeConsole 命令行
	TypeConsole LogType = iota
	// TypeRollingFile 基于体积分割的文件日志
	TypeRollingFile
	// TypeDailyFile 基于日期分割的文件日志
	TypeDailyFile
	// TypeDailyRollingFile 基于日期及体积分割的文件日志
	TypeDailyRollingFile
)

func DefaultLogger() ILogger {
	return defaultLogger
}

type CfgLog struct {
	LogType  LogType  `json:"type" yaml:"type"`
	LogLevel LogLevel `json:"level" yaml:"level"`
	LogPath  string   `json:"path" yaml:"path"`
	LogSize  string   `json:"size" yaml:"size"`
}

func (o CfgLog) GetAbsLogPath() string {
	return filex.Combine(osxu.GetRunningDir(), o.LogPath)
}

func (o CfgLog) MaxSize() mathx.SizeUnit {
	return mathx.ParseSize(o.LogSize)
}

func (o CfgLog) ToLogConfig() LogConfig {
	return LogConfig{Type: o.LogType, Level: o.LogLevel,
		FilePath: o.GetAbsLogPath(),
		MaxSize:  o.MaxSize()}
}

type LogConfig struct {
	Type        LogType        `json:"type" yaml:"type"`
	Level       LogLevel       `json:"level" yaml:"level"`
	Flag        int            `json:"flag,omitempty" yaml:"flag,omitempty"`
	FileDir     string         `json:"dir" yaml:"dir"`
	FileName    string         `json:"name" yaml:"name"`
	FileExtName string         `json:"ext" yaml:"ext"`   // 要带"."
	FilePath    string         `json:"path" yaml:"path"` // 这里有了就会覆盖FileDir,FileName,FileExtName属性
	MaxSize     mathx.SizeUnit `json:"size,omitempty" yaml:"size,omitempty"`
}

func (o LogConfig) GetFileInfos() (fileDir, filename, fileExt string) {
	if o.FilePath == "" {
		return o.FileDir, o.FileName, o.FileExtName
	} else {
		dir, name := filex.Split(o.FilePath)
		fn, fe, _ := filex.SplitFileName(name)
		return dir, fn, fe
	}
}

func SetColorFormat(enable bool) {
	enableColorFormat = enable
}

func NewLogger() ILogger {
	instance := &logger{prefix: "", defaultFlag: log.LstdFlags, infoMap: make(map[LogType]*logInfo)}
	return instance
}

type ILoggerSetter interface {
	SetLogger(logger ILogger)
}

type ILoggerGetter interface {
	GetLogger() ILogger
}

type ILoggerSupport interface {
	ILoggerSetter
	ILoggerGetter
}

type LoggerSupport struct {
	logger ILogger
}

func (s *LoggerSupport) SetLogger(logger ILogger) {
	s.logger = logger
}

func (s *LoggerSupport) GetLogger() ILogger {
	if nil == s.logger {
		return DefaultLogger()
	}
	return s.logger
}

type ILogger interface {
	// SetPrefix 设置日志前缀
	SetPrefix(prefix string)
	// SetLevel
	// 设置日志等级，只有大于等级设置等级的日志才会记录
	// 重置日志等级为level,t为空时重置全部
	SetLevel(level LogLevel, t ...LogType)
	// SetFlags sets the output flags for the logger.
	// 重置日志flag,t为空时重置全部
	SetFlags(flag int, t ...LogType)
	// SetConfig
	// 配置Log,要求fileDir以"/"结尾
	SetConfig(cfg LogConfig)
	// RemoveConfig 移除配置
	RemoveConfig(t LogType)

	// Print 普通记录，忽略前缀
	Print(v ...interface{})
	// Printf 普通记录，忽略前缀
	Printf(format string, v ...interface{})
	// Println 普通记录，忽略前缀
	Println(v ...interface{})

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

func Print(v ...interface{}) {
	defaultLogger.Print(v...)
}

func Printf(format string, v ...interface{}) {
	defaultLogger.Printf(format, v...)
}

func Println(v ...interface{}) {
	defaultLogger.Println(v...)
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
