// Package logx
// Create on 2023/6/30
// @author xuzhuoxi
package logx

import (
	"errors"
	"fmt"
	"sync"
)

var (
	DefaultLoggerManager = NewLoggerManager()
	DefaultLoggerName    = "global"
)

func NewLoggerManager() ILoggerManager {
	return &LoggerManager{}
}

type ILoggerManager interface {
	// GenLogger
	// 生成Logger并加入注册
	GenLogger(name string, config CfgLog) error
	// RegisterLogger
	// 注册一个Logger到管理器
	RegisterLogger(name string, logger ILogger) error
	// UnregisterLogger
	// 取消Logger的注册
	UnregisterLogger(name string) (logger ILogger, err error)
	// FindLogger
	// 查找Logger
	FindLogger(name string) ILogger
	// Clear
	// 清除全部已注册的Logger信息
	Clear()
	// SetDefault
	// 设置默认Logger
	SetDefault(logName string) bool
	// GetDefaultLogger
	// 对默认Logger
	GetDefaultLogger() ILogger
	ILoggerManagerActions
}

type ILoggerManagerActions interface {
	// Print 普通记录，忽略前缀
	Print(name string, v ...interface{})
	// Printf 普通记录，忽略前缀
	Printf(name string, format string, v ...interface{})
	// Println 普通记录，忽略前缀
	Println(name string, v ...interface{})

	Log(name string, level LogLevel, v ...interface{})
	Logf(name string, level LogLevel, format string, v ...interface{})
	Logln(name string, level LogLevel, v ...interface{})

	Trace(name string, v ...interface{})
	Tracef(name string, format string, v ...interface{})
	Traceln(name string, v ...interface{})
	Debug(name string, v ...interface{})
	Debugf(name string, format string, v ...interface{})
	Debugln(name string, v ...interface{})
	Info(name string, v ...interface{})
	Infof(name string, format string, v ...interface{})
	Infoln(name string, v ...interface{})
	Warn(name string, v ...interface{})
	Warnf(name string, format string, v ...interface{})
	Warnln(name string, v ...interface{})
	Error(name string, v ...interface{})
	Errorf(name string, format string, v ...interface{})
	Errorln(name string, v ...interface{})
	Fatal(name string, v ...interface{})
	Fatalf(name string, format string, v ...interface{})
	Fatalln(name string, v ...interface{})
}

type _LogInfo struct {
	Name   string
	Logger ILogger
}

type LoggerManager struct {
	defaultLog string
	loggers    []*_LogInfo
	lock       sync.RWMutex
}

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

func (o *LoggerManager) GetDefaultLogger() ILogger {
	o.lock.RLock()
	defer o.lock.RUnlock()
	return o.FindLogger(o.defaultLog)
}

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

func (o *LoggerManager) FindLogger(name string) ILogger {
	o.lock.RLock()
	defer o.lock.RUnlock()
	index := o.findIndex(name)
	if -1 == index {
		return nil
	}
	return o.loggers[index].Logger
}

func (o *LoggerManager) Clear() {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.loggers = nil
}

func (o *LoggerManager) Print(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Print(v...)
}

func (o *LoggerManager) Printf(name string, format string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Printf(format, v...)
}

func (o *LoggerManager) Println(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Println(v...)
}

func (o *LoggerManager) Log(name string, level LogLevel, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Log(level, v...)
}

func (o *LoggerManager) Logf(name string, level LogLevel, format string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Logf(level, format, v...)
}

func (o *LoggerManager) Logln(name string, level LogLevel, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Logln(level, v...)
}

func (o *LoggerManager) Trace(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Trace(v...)
}

func (o *LoggerManager) Tracef(name string, format string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Tracef(format, v...)
}

func (o *LoggerManager) Traceln(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Traceln(v...)
}

func (o *LoggerManager) Debug(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Debug(v...)
}

func (o *LoggerManager) Debugf(name string, format string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Debugf(format, v...)
}

func (o *LoggerManager) Debugln(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Debugln(v...)
}

func (o *LoggerManager) Info(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Info(v...)
}

func (o *LoggerManager) Infof(name string, format string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Infof(format, v...)
}

func (o *LoggerManager) Infoln(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Infoln(v...)
}

func (o *LoggerManager) Warn(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Warn(v...)
}

func (o *LoggerManager) Warnf(name string, format string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Warnf(format, v...)
}

func (o *LoggerManager) Warnln(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Warnln(v...)
}

func (o *LoggerManager) Error(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Error(v...)
}

func (o *LoggerManager) Errorf(name string, format string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Errorf(format, v...)
}

func (o *LoggerManager) Errorln(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Errorln(v...)
}

func (o *LoggerManager) Fatal(name string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Fatal(v...)
}

func (o *LoggerManager) Fatalf(name string, format string, v ...interface{}) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	log := o.find(name)
	if nil == log {
		return
	}
	log.Fatalf(format, v...)
}

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
