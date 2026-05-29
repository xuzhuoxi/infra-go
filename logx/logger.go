package logx

import (
	"log"

	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/mathx"
	"github.com/xuzhuoxi/infra-go/osxu"
)

// LogType identifies a logging output sink (console or file variant).
// LogType 标识日志输出目标类型（控制台或各类文件输出）。
type LogType uint8

// LogLevel identifies the severity of a log message.
// LogLevel 标识日志消息的严重程度。
type LogLevel uint8

const (
	// LevelAll allows all log levels to be recorded.
	// LevelAll 允许记录所有日志等级。
	LevelAll LogLevel = iota
	// LevelTrace is the most verbose trace level.
	// LevelTrace 最详细的跟踪等级。
	LevelTrace
	// LevelDebug is the debug diagnostic level.
	// LevelDebug 调试诊断等级。
	LevelDebug
	// LevelInfo is the informational level.
	// LevelInfo 信息提示等级。
	LevelInfo
	// LevelWarn is the warning level.
	// LevelWarn 警告等级。
	LevelWarn
	// LevelError is the error level.
	// LevelError 错误等级。
	LevelError
	// LevelFatal is the fatal level; typically followed by process exit.
	// LevelFatal 致命错误等级；通常随后退出进程。
	LevelFatal
	// LevelOff disables all log output.
	// LevelOff 关闭全部日志输出。
	LevelOff
)

const (
	// TypeConsole is console (stderr) output.
	// TypeConsole 命令行（标准错误）输出。
	TypeConsole LogType = iota
	// TypeRollingFile is file output split by size.
	// TypeRollingFile 基于体积分割的文件日志。
	TypeRollingFile
	// TypeDailyFile is file output split by calendar date.
	// TypeDailyFile 基于日期分割的文件日志。
	TypeDailyFile
	// TypeDailyRollingFile is file output split by both date and size.
	// TypeDailyRollingFile 基于日期及体积分割的文件日志。
	TypeDailyRollingFile
)

// DefaultLogger returns the package-wide default ILogger instance.
// DefaultLogger 返回包级默认 ILogger 实例。
func DefaultLogger() ILogger {
	return defaultLogger
}

// CfgLog is a simplified log configuration for JSON/YAML unmarshaling.
// CfgLog 是便于 JSON/YAML 反序列化的简化日志配置。
type CfgLog struct {
	LogType  LogType  `json:"type" yaml:"type"`
	LogLevel LogLevel `json:"level" yaml:"level"`
	LogPath  string   `json:"path" yaml:"path"`
	LogSize  string   `json:"size" yaml:"size"`
}

// GetAbsLogPath returns the absolute log file path relative to the running directory.
// GetAbsLogPath 返回相对于运行目录的绝对日志文件路径。
func (o CfgLog) GetAbsLogPath() string {
	return filex.Combine(osxu.GetRunningDir(), o.LogPath)
}

// MaxSize parses LogSize into a mathx.SizeUnit.
// MaxSize 将 LogSize 解析为 mathx.SizeUnit。
func (o CfgLog) MaxSize() mathx.SizeUnit {
	return mathx.ParseSize(o.LogSize)
}

// ToLogConfig converts CfgLog into a full LogConfig.
// ToLogConfig 将 CfgLog 转换为完整的 LogConfig。
func (o CfgLog) ToLogConfig() LogConfig {
	return LogConfig{Type: o.LogType, Level: o.LogLevel,
		FilePath: o.GetAbsLogPath(),
		MaxSize:  o.MaxSize()}
}

// LogConfig holds detailed settings for a single log output sink.
// LogConfig 保存单个日志输出目标的详细配置。
type LogConfig struct {
	Type        LogType        `json:"type" yaml:"type"`
	Level       LogLevel       `json:"level" yaml:"level"`
	Flag        int            `json:"flag,omitempty" yaml:"flag,omitempty"`
	FileDir     string         `json:"dir" yaml:"dir"`
	FileName    string         `json:"name" yaml:"name"`
	FileExtName string         `json:"ext" yaml:"ext"`   // Must include leading ".". 要带"."。
	FilePath    string         `json:"path" yaml:"path"` // When set, overrides FileDir, FileName, and FileExtName. 设置后会覆盖 FileDir、FileName、FileExtName 属性。
	MaxSize     mathx.SizeUnit `json:"size,omitempty" yaml:"size,omitempty"`
}

// GetFileInfos resolves file directory, base name, and extension from the config.
// GetFileInfos 从配置中解析文件目录、基础文件名与扩展名。
func (o LogConfig) GetFileInfos() (fileDir, filename, fileExt string) {
	if o.FilePath == "" {
		return o.FileDir, o.FileName, o.FileExtName
	} else {
		dir, name := filex.Split(o.FilePath)
		fn, fe, _ := filex.SplitFileName(name)
		return dir, fn, fe
	}
}

// SetColorFormat enables or disables ANSI color prefixes on console output.
// SetColorFormat 启用或禁用控制台输出的 ANSI 颜色前缀。
func SetColorFormat(enable bool) {
	enableColorFormat = enable
}

// NewLogger creates a new ILogger with no outputs configured yet.
// NewLogger 创建一个尚未配置输出目标的新 ILogger。
func NewLogger() ILogger {
	instance := &logger{prefix: "", defaultFlag: log.LstdFlags, infoMap: make(map[LogType]*logInfo)}
	return instance
}

// ILoggerSetter allows injecting an ILogger implementation.
// ILoggerSetter 允许注入 ILogger 实现。
type ILoggerSetter interface {
	// SetLogger assigns the logger instance.
	// SetLogger 设置 Logger 实例。
	SetLogger(logger ILogger)
}

// ILoggerGetter provides access to an ILogger instance.
// ILoggerGetter 提供对 ILogger 实例的访问。
type ILoggerGetter interface {
	// GetLogger returns the configured logger, or the package default when unset.
	// GetLogger 返回已配置的 Logger；未设置时返回包默认 Logger。
	GetLogger() ILogger
}

// ILoggerSupport combines setter and getter for embeddable logger access.
// ILoggerSupport 组合设置与获取能力，便于嵌入类型使用。
type ILoggerSupport interface {
	ILoggerSetter
	ILoggerGetter
}

// LoggerSupport is a base struct that holds an optional ILogger.
// LoggerSupport 是持有可选 ILogger 的基础结构体。
type LoggerSupport struct {
	logger ILogger
}

// SetLogger assigns the logger instance.
// SetLogger 设置 Logger 实例。
func (s *LoggerSupport) SetLogger(logger ILogger) {
	s.logger = logger
}

// GetLogger returns the configured logger, or DefaultLogger when unset.
// GetLogger 返回已配置的 Logger；未设置时返回 DefaultLogger。
func (s *LoggerSupport) GetLogger() ILogger {
	if nil == s.logger {
		return DefaultLogger()
	}
	return s.logger
}

// ILoggerSettings defines configuration methods for a logger.
// ILoggerSettings 定义 Logger 的配置方法。
type ILoggerSettings interface {
	// SetPrefix sets the log prefix prepended to level-tagged messages.
	// SetPrefix 设置日志前缀。
	SetPrefix(prefix string)
	// SetLevel sets the minimum log level; only messages at or above this level are recorded.
	// SetLevel 设置日志等级，只有大于等于该等级的日志才会记录。
	// When t is empty, resets the level for all log types.
	// t 为空时重置全部。
	SetLevel(level LogLevel, t ...LogType)
	// SetFlags sets the output flags for the logger.
	// SetFlags 重置日志 flag。
	// When t is empty, resets flags for all log types.
	// t 为空时重置全部。
	SetFlags(flag int, t ...LogType)
	// SetConfig configures the logger; fileDir must end with "/".
	// SetConfig 配置 Log，要求 fileDir 以 "/" 结尾。
	SetConfig(cfg LogConfig)
	// RemoveConfig removes the configuration for the given log type.
	// RemoveConfig 移除配置。
	RemoveConfig(t LogType)
}

// ILoggerActions defines leveled and plain logging methods.
// ILoggerActions 定义分级与普通日志记录方法。
type ILoggerActions interface {
	// Print logs a message without level prefix.
	// Print 普通记录，忽略前缀。
	Print(v ...interface{})
	// Printf logs a formatted message without level prefix.
	// Printf 普通记录，忽略前缀。
	Printf(format string, v ...interface{})
	// Println logs a line without level prefix.
	// Println 普通记录，忽略前缀。
	Println(v ...interface{})

	// Log logs a message at the given level.
	// Log 按指定等级记录日志。
	Log(level LogLevel, v ...interface{})
	// Logf logs a formatted message at the given level.
	// Logf 按指定等级记录格式化日志。
	Logf(level LogLevel, format string, v ...interface{})
	// Logln logs a line at the given level.
	// Logln 按指定等级记录一行日志。
	Logln(level LogLevel, v ...interface{})

	// Trace logs at trace level.
	// Trace 记录跟踪级别日志。
	Trace(v ...interface{})
	// Tracef logs a formatted message at trace level.
	// Tracef 记录格式化跟踪级别日志。
	Tracef(format string, v ...interface{})
	// Traceln logs a line at trace level.
	// Traceln 记录一行跟踪级别日志。
	Traceln(v ...interface{})
	// Debug logs at debug level.
	// Debug 记录调试级别日志。
	Debug(v ...interface{})
	// Debugf logs a formatted message at debug level.
	// Debugf 记录格式化调试级别日志。
	Debugf(format string, v ...interface{})
	// Debugln logs a line at debug level.
	// Debugln 记录一行调试级别日志。
	Debugln(v ...interface{})
	// Info logs at info level.
	// Info 记录信息级别日志。
	Info(v ...interface{})
	// Infof logs a formatted message at info level.
	// Infof 记录格式化信息级别日志。
	Infof(format string, v ...interface{})
	// Infoln logs a line at info level.
	// Infoln 记录一行信息级别日志。
	Infoln(v ...interface{})
	// Warn logs at warn level.
	// Warn 记录警告级别日志。
	Warn(v ...interface{})
	// Warnf logs a formatted message at warn level.
	// Warnf 记录格式化警告级别日志。
	Warnf(format string, v ...interface{})
	// Warnln logs a line at warn level.
	// Warnln 记录一行警告级别日志。
	Warnln(v ...interface{})
	// Error logs at error level.
	// Error 记录错误级别日志。
	Error(v ...interface{})
	// Errorf logs a formatted message at error level.
	// Errorf 记录格式化错误级别日志。
	Errorf(format string, v ...interface{})
	// Errorln logs a line at error level.
	// Errorln 记录一行错误级别日志。
	Errorln(v ...interface{})
	// Fatal logs at fatal level.
	// Fatal 记录致命级别日志。
	Fatal(v ...interface{})
	// Fatalf logs a formatted message at fatal level.
	// Fatalf 记录格式化致命级别日志。
	Fatalf(format string, v ...interface{})
	// Fatalln logs a line at fatal level.
	// Fatalln 记录一行致命级别日志。
	Fatalln(v ...interface{})
}

// ILogger is the full logging interface combining settings and actions.
// ILogger 是组合配置与记录能力的完整日志接口。
type ILogger interface {
	ILoggerSettings
	ILoggerActions
}

// SetPrefix sets the prefix on the package default logger.
// SetPrefix 设置包默认 Logger 的前缀。
func SetPrefix(prefix string) {
	defaultLogger.SetPrefix(prefix)
}

// SetLevel sets the minimum level on the package default logger.
// SetLevel 设置包默认 Logger 的最低日志等级。
func SetLevel(level LogLevel, t ...LogType) {
	defaultLogger.SetLevel(level, t...)
}

// SetFlags sets output flags on the package default logger.
// SetFlags 设置包默认 Logger 的输出 flag。
func SetFlags(flag int, t ...LogType) {
	defaultLogger.SetFlags(flag, t...)
}

// SetConfig applies configuration to the package default logger.
// SetConfig 为包默认 Logger 应用配置。
func SetConfig(cfg LogConfig) {
	defaultLogger.SetConfig(cfg)
}

// RemoveConfig removes a log type configuration from the package default logger.
// RemoveConfig 从包默认 Logger 移除指定类型的配置。
func RemoveConfig(t LogType) {
	defaultLogger.RemoveConfig(t)
}

// Print logs via the package default logger without level prefix.
// Print 通过包默认 Logger 普通记录，忽略前缀。
func Print(v ...interface{}) {
	defaultLogger.Print(v...)
}

// Printf logs a formatted message via the package default logger without level prefix.
// Printf 通过包默认 Logger 普通格式化记录，忽略前缀。
func Printf(format string, v ...interface{}) {
	defaultLogger.Printf(format, v...)
}

// Println logs a line via the package default logger without level prefix.
// Println 通过包默认 Logger 普通记录一行，忽略前缀。
func Println(v ...interface{}) {
	defaultLogger.Println(v...)
}

// Trace logs at trace level on the package default logger.
// Trace 通过包默认 Logger 记录跟踪级别日志。
func Trace(v ...interface{}) {
	defaultLogger.Trace(v...)
}

// Tracef logs a formatted trace message on the package default logger.
// Tracef 通过包默认 Logger 记录格式化跟踪级别日志。
func Tracef(format string, v ...interface{}) {
	defaultLogger.Tracef(format, v...)
}

// Traceln logs a trace line on the package default logger.
// Traceln 通过包默认 Logger 记录一行跟踪级别日志。
func Traceln(v ...interface{}) {
	defaultLogger.Traceln(v...)
}

// Debug logs at debug level on the package default logger.
// Debug 通过包默认 Logger 记录调试级别日志。
func Debug(v ...interface{}) {
	defaultLogger.Debug(v...)
}

// Debugf logs a formatted debug message on the package default logger.
// Debugf 通过包默认 Logger 记录格式化调试级别日志。
func Debugf(format string, v ...interface{}) {
	defaultLogger.Debugf(format, v...)
}

// Debugln logs a debug line on the package default logger.
// Debugln 通过包默认 Logger 记录一行调试级别日志。
func Debugln(v ...interface{}) {
	defaultLogger.Debugln(v...)
}

// Info logs at info level on the package default logger.
// Info 通过包默认 Logger 记录信息级别日志。
func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

// Infof logs a formatted info message on the package default logger.
// Infof 通过包默认 Logger 记录格式化信息级别日志。
func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

// Infoln logs an info line on the package default logger.
// Infoln 通过包默认 Logger 记录一行信息级别日志。
func Infoln(v ...interface{}) {
	defaultLogger.Infoln(v...)
}

// Warn logs at warn level on the package default logger.
// Warn 通过包默认 Logger 记录警告级别日志。
func Warn(v ...interface{}) {
	defaultLogger.Warn(v...)
}

// Warnf logs a formatted warn message on the package default logger.
// Warnf 通过包默认 Logger 记录格式化警告级别日志。
func Warnf(format string, v ...interface{}) {
	defaultLogger.Warnf(format, v...)
}

// Warnln logs a warn line on the package default logger.
// Warnln 通过包默认 Logger 记录一行警告级别日志。
func Warnln(v ...interface{}) {
	defaultLogger.Warnln(v...)
}

// Error logs at error level on the package default logger.
// Error 通过包默认 Logger 记录错误级别日志。
func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

// Errorf logs a formatted error message on the package default logger.
// Errorf 通过包默认 Logger 记录格式化错误级别日志。
func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

// Errorln logs an error line on the package default logger.
// Errorln 通过包默认 Logger 记录一行错误级别日志。
func Errorln(v ...interface{}) {
	defaultLogger.Errorln(v...)
}

// Fatal logs at fatal level on the package default logger.
// Fatal 通过包默认 Logger 记录致命级别日志。
func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v...)
}

// Fatalf logs a formatted fatal message on the package default logger.
// Fatalf 通过包默认 Logger 记录格式化致命级别日志。
func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}

// Fatalln logs a fatal line on the package default logger.
// Fatalln 通过包默认 Logger 记录一行致命级别日志。
func Fatalln(v ...interface{}) {
	defaultLogger.Fatalln(v...)
}
