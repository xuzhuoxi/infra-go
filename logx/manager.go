package logx

import (
	"errors"
	"fmt"
	"sync"
)

var (
	// DefaultLoggerManager is the package-wide default named logger registry.
	// DefaultLoggerManager 是包级默认的命名 Logger 注册表。
	DefaultLoggerManager = NewLoggerManager()
	// DefaultLoggerName is the default logger name used by LoggerManager.
	// DefaultLoggerName 是 LoggerManager 使用的默认 Logger 名称。
	DefaultLoggerName = "global"
)

// NewLoggerManager creates a new empty ILoggerManager.
// NewLoggerManager 创建一个新的空 ILoggerManager。
func NewLoggerManager() ILoggerManager {
	return &LoggerManager{}
}

// ILoggerManager manages a registry of named ILogger instances.
// ILoggerManager 管理命名 ILogger 实例的注册表。
type ILoggerManager interface {
	// GenLogger creates a logger from CfgLog and registers it under name.
	// GenLogger 根据 CfgLog 创建 Logger 并注册到指定名称。
	GenLogger(name string, config CfgLog) error
	// RegisterLogger registers an existing logger under name.
	// RegisterLogger 将已有 Logger 注册到指定名称。
	RegisterLogger(name string, logger ILogger) error
	// UnregisterLogger removes and returns the logger registered under name.
	// UnregisterLogger 取消指定名称的 Logger 注册并返回该实例。
	UnregisterLogger(name string) (logger ILogger, err error)
	// FindLogger returns the logger registered under name, or nil if not found.
	// FindLogger 查找指定名称的 Logger；未找到时返回 nil。
	FindLogger(name string) ILogger
	// Clear removes all registered loggers.
	// Clear 清除全部已注册的 Logger 信息。
	Clear()
	// SetDefault sets the default logger name used by GetDefaultLogger.
	// SetDefault 设置默认 Logger 名称。
	SetDefault(logName string) bool
	// GetDefaultLogger returns the logger registered as default.
	// GetDefaultLogger 返回默认 Logger。
	GetDefaultLogger() ILogger
	ILoggerManagerActions
}

// ILoggerManagerActions defines named dispatch methods for leveled logging.
// ILoggerManagerActions 定义按名称分发的分级日志方法。
type ILoggerManagerActions interface {
	// Print logs via the named logger without level prefix.
	// Print 通过指定名称的 Logger 普通记录，忽略前缀。
	Print(name string, v ...interface{})
	// Printf logs a formatted message via the named logger without level prefix.
	// Printf 通过指定名称的 Logger 普通格式化记录，忽略前缀。
	Printf(name string, format string, v ...interface{})
	// Println logs a line via the named logger without level prefix.
	// Println 通过指定名称的 Logger 普通记录一行，忽略前缀。
	Println(name string, v ...interface{})

	// Log logs at the given level via the named logger.
	// Log 通过指定名称的 Logger 按等级记录日志。
	Log(name string, level LogLevel, v ...interface{})
	// Logf logs a formatted message at the given level via the named logger.
	// Logf 通过指定名称的 Logger 按等级记录格式化日志。
	Logf(name string, level LogLevel, format string, v ...interface{})
	// Logln logs a line at the given level via the named logger.
	// Logln 通过指定名称的 Logger 按等级记录一行日志。
	Logln(name string, level LogLevel, v ...interface{})

	// Trace logs at trace level via the named logger.
	// Trace 通过指定名称的 Logger 记录跟踪级别日志。
	Trace(name string, v ...interface{})
	// Tracef logs a formatted trace message via the named logger.
	// Tracef 通过指定名称的 Logger 记录格式化跟踪级别日志。
	Tracef(name string, format string, v ...interface{})
	// Traceln logs a trace line via the named logger.
	// Traceln 通过指定名称的 Logger 记录一行跟踪级别日志。
	Traceln(name string, v ...interface{})
	// Debug logs at debug level via the named logger.
	// Debug 通过指定名称的 Logger 记录调试级别日志。
	Debug(name string, v ...interface{})
	// Debugf logs a formatted debug message via the named logger.
	// Debugf 通过指定名称的 Logger 记录格式化调试级别日志。
	Debugf(name string, format string, v ...interface{})
	// Debugln logs a debug line via the named logger.
	// Debugln 通过指定名称的 Logger 记录一行调试级别日志。
	Debugln(name string, v ...interface{})
	// Info logs at info level via the named logger.
	// Info 通过指定名称的 Logger 记录信息级别日志。
	Info(name string, v ...interface{})
	// Infof logs a formatted info message via the named logger.
	// Infof 通过指定名称的 Logger 记录格式化信息级别日志。
	Infof(name string, format string, v ...interface{})
	// Infoln logs an info line via the named logger.
	// Infoln 通过指定名称的 Logger 记录一行信息级别日志。
	Infoln(name string, v ...interface{})
	// Warn logs at warn level via the named logger.
	// Warn 通过指定名称的 Logger 记录警告级别日志。
	Warn(name string, v ...interface{})
	// Warnf logs a formatted warn message via the named logger.
	// Warnf 通过指定名称的 Logger 记录格式化警告级别日志。
	Warnf(name string, format string, v ...interface{})
	// Warnln logs a warn line via the named logger.
	// Warnln 通过指定名称的 Logger 记录一行警告级别日志。
	Warnln(name string, v ...interface{})
	// Error logs at error level via the named logger.
	// Error 通过指定名称的 Logger 记录错误级别日志。
	Error(name string, v ...interface{})
	// Errorf logs a formatted error message via the named logger.
	// Errorf 通过指定名称的 Logger 记录格式化错误级别日志。
	Errorf(name string, format string, v ...interface{})
	// Errorln logs an error line via the named logger.
	// Errorln 通过指定名称的 Logger 记录一行错误级别日志。
	Errorln(name string, v ...interface{})
	// Fatal logs at fatal level via the named logger.
	// Fatal 通过指定名称的 Logger 记录致命级别日志。
	Fatal(name string, v ...interface{})
	// Fatalf logs a formatted fatal message via the named logger.
	// Fatalf 通过指定名称的 Logger 记录格式化致命级别日志。
	Fatalf(name string, format string, v ...interface{})
	// Fatalln logs a fatal line via the named logger.
	// Fatalln 通过指定名称的 Logger 记录一行致命级别日志。
	Fatalln(name string, v ...interface{})
}

type _LogInfo struct {
	Name   string
	Logger ILogger
}

// LoggerManager is the default ILoggerManager implementation.
// LoggerManager 是 ILoggerManager 的默认实现。
type LoggerManager struct {
	defaultLog string
	loggers    []*_LogInfo
	lock       sync.RWMutex
}

// SetDefault sets the default logger name used by GetDefaultLogger.
// SetDefault 设置默认 Logger 名称。
func (o *LoggerManager) SetDefault(logName string) bool {
	o.lock.Lock()
	defer o.lock.Unlock()
	index := o.findIndex(logName)
	if -1 == index {
		return false
	}
	o.defaultLog = logName
	return true
}

// GetDefaultLogger returns the logger registered as default.
// GetDefaultLogger 返回默认 Logger。
func (o *LoggerManager) GetDefaultLogger() ILogger {
	o.lock.RLock()
	defer o.lock.RUnlock()
	return o.FindLogger(o.defaultLog)
}

// GenLogger creates a logger from CfgLog and registers it under name.
// GenLogger 根据 CfgLog 创建 Logger 并注册到指定名称。
func (o *LoggerManager) GenLogger(name string, config CfgLog) error {
	o.lock.Lock()
	defer o.lock.Unlock()
	index := o.findIndex(name)
	if -1 != index {
		return errors.New(fmt.Sprintf("Logger[name=%s] is exists!", name))
	}
	logger := NewLogger()
	logger.SetConfig(config.ToLogConfig())
	o.loggers = append(o.loggers, &_LogInfo{Name: name, Logger: logger})
	return nil
}

// RegisterLogger registers an existing logger under name.
// RegisterLogger 将已有 Logger 注册到指定名称。
func (o *LoggerManager) RegisterLogger(name string, logger ILogger) error {
	o.lock.Lock()
	defer o.lock.Unlock()
	index := o.findIndex(name)
	if -1 != index {
		return errors.New(fmt.Sprintf("Logger[name=%s] is exists!", name))
	}
	o.loggers = append(o.loggers, &_LogInfo{Name: name, Logger: logger})
	return nil
}

// UnregisterLogger removes and returns the logger registered under name.
// UnregisterLogger 取消指定名称的 Logger 注册并返回该实例。
func (o *LoggerManager) UnregisterLogger(name string) (logger ILogger, err error) {
	o.lock.Lock()
	defer o.lock.Unlock()
	index := o.findIndex(name)
	if -1 == index {
		return nil, errors.New(fmt.Sprintf("No Logger[name=%s] exists!", name))
	}
	rs := o.loggers[index].Logger
	o.loggers = append(o.loggers[:index], o.loggers[index+1:]...)
	return rs, nil
}

// FindLogger returns the logger registered under name, or nil if not found.
// FindLogger 查找指定名称的 Logger；未找到时返回 nil。
func (o *LoggerManager) FindLogger(name string) ILogger {
	o.lock.RLock()
	defer o.lock.RUnlock()
	index := o.findIndex(name)
	if -1 == index {
		return nil
	}
	return o.loggers[index].Logger
}

// Clear removes all registered loggers.
// Clear 清除全部已注册的 Logger 信息。
func (o *LoggerManager) Clear() {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.loggers = nil
}

// Print logs via the named logger without level prefix.
// Print 通过指定名称的 Logger 普通记录，忽略前缀。
func (o *LoggerManager) Print(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Print(v...)
}

// Printf logs a formatted message via the named logger without level prefix.
// Printf 通过指定名称的 Logger 普通格式化记录，忽略前缀。
func (o *LoggerManager) Printf(name string, format string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Printf(format, v...)
}

// Println logs a line via the named logger without level prefix.
// Println 通过指定名称的 Logger 普通记录一行，忽略前缀。
func (o *LoggerManager) Println(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Println(v...)
}

// Log logs at the given level via the named logger.
// Log 通过指定名称的 Logger 按等级记录日志。
func (o *LoggerManager) Log(name string, level LogLevel, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Log(level, v...)
}

// Logf logs a formatted message at the given level via the named logger.
// Logf 通过指定名称的 Logger 按等级记录格式化日志。
func (o *LoggerManager) Logf(name string, level LogLevel, format string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Logf(level, format, v...)
}

// Logln logs a line at the given level via the named logger.
// Logln 通过指定名称的 Logger 按等级记录一行日志。
func (o *LoggerManager) Logln(name string, level LogLevel, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Logln(level, v...)
}

// Trace logs at trace level via the named logger.
// Trace 通过指定名称的 Logger 记录跟踪级别日志。
func (o *LoggerManager) Trace(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Trace(v...)
}

// Tracef logs a formatted trace message via the named logger.
// Tracef 通过指定名称的 Logger 记录格式化跟踪级别日志。
func (o *LoggerManager) Tracef(name string, format string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Tracef(format, v...)
}

// Traceln logs a trace line via the named logger.
// Traceln 通过指定名称的 Logger 记录一行跟踪级别日志。
func (o *LoggerManager) Traceln(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Traceln(v...)
}

// Debug logs at debug level via the named logger.
// Debug 通过指定名称的 Logger 记录调试级别日志。
func (o *LoggerManager) Debug(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Debug(v...)
}

// Debugf logs a formatted debug message via the named logger.
// Debugf 通过指定名称的 Logger 记录格式化调试级别日志。
func (o *LoggerManager) Debugf(name string, format string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Debugf(format, v...)
}

// Debugln logs a debug line via the named logger.
// Debugln 通过指定名称的 Logger 记录一行调试级别日志。
func (o *LoggerManager) Debugln(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Debugln(v...)
}

// Info logs at info level via the named logger.
// Info 通过指定名称的 Logger 记录信息级别日志。
func (o *LoggerManager) Info(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Info(v...)
}

// Infof logs a formatted info message via the named logger.
// Infof 通过指定名称的 Logger 记录格式化信息级别日志。
func (o *LoggerManager) Infof(name string, format string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Infof(format, v...)
}

// Infoln logs an info line via the named logger.
// Infoln 通过指定名称的 Logger 记录一行信息级别日志。
func (o *LoggerManager) Infoln(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Infoln(v...)
}

// Warn logs at warn level via the named logger.
// Warn 通过指定名称的 Logger 记录警告级别日志。
func (o *LoggerManager) Warn(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Warn(v...)
}

// Warnf logs a formatted warn message via the named logger.
// Warnf 通过指定名称的 Logger 记录格式化警告级别日志。
func (o *LoggerManager) Warnf(name string, format string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Warnf(format, v...)
}

// Warnln logs a warn line via the named logger.
// Warnln 通过指定名称的 Logger 记录一行警告级别日志。
func (o *LoggerManager) Warnln(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Warnln(v...)
}

// Error logs at error level via the named logger.
// Error 通过指定名称的 Logger 记录错误级别日志。
func (o *LoggerManager) Error(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Error(v...)
}

// Errorf logs a formatted error message via the named logger.
// Errorf 通过指定名称的 Logger 记录格式化错误级别日志。
func (o *LoggerManager) Errorf(name string, format string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Errorf(format, v...)
}

// Errorln logs an error line via the named logger.
// Errorln 通过指定名称的 Logger 记录一行错误级别日志。
func (o *LoggerManager) Errorln(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Errorln(v...)
}

// Fatal logs at fatal level via the named logger.
// Fatal 通过指定名称的 Logger 记录致命级别日志。
func (o *LoggerManager) Fatal(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Fatal(v...)
}

// Fatalf logs a formatted fatal message via the named logger.
// Fatalf 通过指定名称的 Logger 记录格式化致命级别日志。
func (o *LoggerManager) Fatalf(name string, format string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Fatalf(format, v...)
}

// Fatalln logs a fatal line via the named logger.
// Fatalln 通过指定名称的 Logger 记录一行致命级别日志。
func (o *LoggerManager) Fatalln(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Fatalln(v...)
}

func (o *LoggerManager) find(name string) ILogger {
	index := o.findIndex(name)
	if -1 == index {
		return nil
	}
	return o.loggers[index].Logger
}

func (o *LoggerManager) findIndex(name string) int {
	if len(o.loggers) == 0 || len(name) == 0 {
		return -1
	}
	for index := range o.loggers {
		if o.loggers[index].Name == name {
			return index
		}
	}
	return -1
}
